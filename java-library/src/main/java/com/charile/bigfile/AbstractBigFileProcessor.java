package com.charile.bigfile;

import lombok.Data;
import lombok.extern.log4j.Log4j2;
import org.jpedal.io.RandomAccessFileBuffer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:33
 */
@Data
public abstract class AbstractBigFileProcessor implements FileChunkProcessor
{
    private static final Logger log = LoggerFactory.getLogger(AbstractBigFileProcessor.class);

    public AbstractBigFileProcessor(ChunkHandler chunkHandler)
    {
        this.chunkHandler = chunkHandler;
    }

    protected ChunkHandler chunkHandler;

    // 判断这个文件是否存在,不存在则创建
    // 同时判断文件是否已经存在了,因为存在全部合并完上传成功后的重复上传
    protected abstract FileInfo createIfDirectoryNotExist(FileChunkDecorator req);


    // chunk 信息


    @Override
    public void process(FileChunkDecorator req) throws IOException
    {
        FileInfo fileInfo = this.createIfDirectoryNotExist(req);
        if (fileInfo.isUploaded())
        {
            log.debug("该文件:{}早已经上传成功,直接返回结果", req.getFileOriginName());
            return;
        }
        log.debug("该文件尚未上传,查询碎片信息");
        ChunkInfo chunkInfo = this.chunkHandler.getChunkInfo(req);
        if (chunkInfo != null && (chunkInfo.getStatus() == BigFileConstants.CHUNK_SUCCESS || chunkInfo.getStatus() == BigFileConstants.CHUNK_ON_PROCESS))
        {
            log.debug("该processId=[{}],碎片chunk={}的碎皮已经上传成功或者是正在上传中,结束该次请求,状态为:{}", req.getProcessId(), req.getChunk(), chunkInfo.getStatus());
            return;
        }

        File chunkFile = null;
        String directory = fileInfo.getFileStoreDirectory();
        String fileOriginName = req.getFileOriginName();
        String chunkName = fileOriginName + "-" + req.getChunk();
        String chunkPath = directory + File.separator + chunkName;
        log.debug("chunk的名称为:" + chunkName);
        log.debug("chunkPath的存储路径为:" + chunkPath);
        chunkFile = new File(chunkPath);
        if (chunkInfo == null)
        {
            log.debug("碎片不存在,创建碎片");
            chunkInfo = new ChunkInfo();
            chunkInfo.setUploadedSize(0);
            chunkInfo.setStatus(BigFileConstants.CHUNK_ON_PROCESS);
            chunkInfo.setChunkSize(req.getChunkSize());
            chunkInfo.setChunkMd5(req.getChunkMd5());
            if (chunkFile.exists())
            {
                log.debug("缓存丢失,不管成功,失败与否,都删除");
                chunkFile.delete();
            }
        }

        log.debug("分段写入到数据中");
        InputStream inputStream = req.getInputStream();
        BufferedInputStream bufferInput = null;
        RandomAccessFile randomAccessFile = null;
        try
        {
            randomAccessFile = new RandomAccessFile(chunkFile, "rw");
            bufferInput = new BufferedInputStream(inputStream);
        } catch (FileNotFoundException e)
        {
            log.error("发开文件写入流失败:" + e.getMessage());
            throw e;
        }
        byte[] buffer = new byte[1024];
        try
        {
            int bytesRead = 0;
            long lastWrite = 0;
            while ((bytesRead = bufferInput.read(buffer)) != -1)
            {
                // 写入到文件中
                log.debug("读取到字节长度为:" + bytesRead);
                randomAccessFile.seek(lastWrite);
                randomAccessFile.write(buffer, 0, bytesRead);
                lastWrite += bytesRead;
                // 更新chunk
                chunkInfo.setUploadedSize(chunkInfo.getUploadedSize() + bytesRead);
                log.debug("跟新该chunk已经读取的字节数:{},百分比为:{}", chunkInfo.getUploadedSize(), chunkInfo.getUploadedSize() / chunkInfo.getChunkSize());
                this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            }
            chunkInfo.setStatus(BigFileConstants.CHUNK_SUCCESS);
            this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            if (this.chunkHandler.increase(req.getProcessId(), req.getTotalChunk()))
            {
                log.debug("所有chunk都已经上传,需要开启合并");
                // FIXME  合并数据
            }
        } catch (IOException e)
        {
            e.printStackTrace();
            log.error("发生io失败,需要将块状态更新为失败:" + e.getMessage());
            chunkInfo.setStatus(BigFileConstants.CHUNK_FAIL);
            this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            throw e;
        } finally
        {
            try
            {
                bufferInput.close();
                randomAccessFile.close();
            } catch (IOException e)
            {
                log.error("关闭bufferedInput失败:" + e.getMessage());
            }
        }
    }
}
