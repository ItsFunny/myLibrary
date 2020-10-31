package com.charlie.bigfile;

import com.charlie.utils.FileUtils;
import lombok.Data;
import lombok.extern.slf4j.Slf4j;

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
@Slf4j
public abstract class AbstractBigFileProcessor implements FileChunkProcessor
{

    public AbstractBigFileProcessor(ChunkHandler chunkHandler)
    {
        this.chunkHandler = chunkHandler;
    }

    protected ChunkHandler chunkHandler;

    // 判断这个文件是否存在,不存在则创建
    // 同时判断文件是否已经存在了,因为存在全部合并完上传成功后的重复上传情况
    protected abstract FileInfo createIfDirectoryNotExist(FileChunkDecorator req);


    // chunk 信息


    @Override
    public ProcessResp process(FileChunkDecorator req) throws IOException
    {
        ProcessResp result = new ProcessResp();
        FileInfo fileInfo = this.createIfDirectoryNotExist(req);
        if (fileInfo.isUploaded())
        {
            log.debug("该文件:{}早已经上传成功,直接返回结果,并且更新碎片信息", req.getFileOriginName());
            ChunkInfo chunkInfo = new ChunkInfo();
            chunkInfo.setStatus(BigFileConstants.CHUNK_SUCCESS);
            chunkInfo.setChunkMd5(req.getChunkMd5());
            chunkInfo.setChunkSize(req.getChunkSize());
            chunkInfo.setUploadedSize(req.getChunkSize());
            this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            result.setMerged(true);
            return result;
        }
        log.debug("该文件尚未上传,查询碎片信息");
        ChunkInfo chunkInfo = this.chunkHandler.getChunkInfo(req);
        if (chunkInfo != null && (chunkInfo.getStatus() == BigFileConstants.CHUNK_SUCCESS || chunkInfo.getStatus() == BigFileConstants.CHUNK_ON_PROCESS))
        {
            log.debug("该processId=[{}],碎片chunk={}的碎皮已经上传成功或者是正在上传中,结束该次请求,状态为:{}", req.getProcessId(), req.getChunk(), chunkInfo.getStatus());
            return result;
        }

        File chunkFile = null;
        String directory = fileInfo.getFileStoreDirectory();
        if (!directory.endsWith(File.separator)) directory += File.separator;
        String fileOriginName = req.getFileOriginName();
        String chunkName = fileOriginName + "-" + req.getChunk();
        String chunkPath = directory + File.separator + chunkName;
        log.debug("chunkPath的存储路径为:" + chunkPath);
        chunkFile = new File(chunkPath);
        if (chunkInfo == null)
        {
            log.debug("该碎片{}不存在,创建碎片", chunkName);
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

        log.info("chunk={},写入到文件:{} 中", chunkName, chunkPath);
        InputStream inputStream = req.getInputStream();
        BufferedInputStream bufferInput = null;
        RandomAccessFile randomAccessFile = null;
        try
        {
            randomAccessFile = new RandomAccessFile(chunkFile, "rw");
            bufferInput = new BufferedInputStream(inputStream);
        } catch (FileNotFoundException e)
        {
            log.error("打开文件写入流失败:" + e.getMessage());
            throw e;
        }

        byte[] buffer = new byte[1024];
        try
        {
            int bytesRead = 0;
            int lastWrite = chunkInfo.getUploadedSize();
            log.debug("processId:{} 碎片:{},上次写入的位置为:{}", req.getProcessId(), req.getChunkMd5(), lastWrite);
            // 将输入写新定位到写入的位置
            if (lastWrite > 0)
            {
                bufferInput.skip(lastWrite);
            }
            while ((bytesRead = bufferInput.read(buffer)) != -1)
            {
                // 写入到文件中
                randomAccessFile.seek(lastWrite);
                randomAccessFile.write(buffer, 0, bytesRead);
                lastWrite += bytesRead;
                // 更新chunk
                chunkInfo.setUploadedSize(chunkInfo.getUploadedSize() + bytesRead);
                log.debug("跟新该chunk已经读取的字节数:{},百分比为:{}", chunkInfo.getUploadedSize(), (float) chunkInfo.getUploadedSize() / chunkInfo.getChunkSize());
                this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            }
            chunkInfo.setStatus(BigFileConstants.CHUNK_SUCCESS);
            this.chunkHandler.updateChunk(req.getProcessId(), chunkInfo);
            if (this.chunkHandler.increase(req.getProcessId(), req.getTotalChunk()))
            {
                log.debug("所有chunk都已经上传,需要开启合并");
                FileUtils.mergeFile(directory, directory + req.getFileOriginName());
                result.setMerged(true);
                log.debug("删除本次processId记录");
                this.chunkHandler.clearChunk(req.getProcessId());
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
        return result;
    }
}
