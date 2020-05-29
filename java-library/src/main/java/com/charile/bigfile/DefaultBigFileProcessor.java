package com.charile.bigfile;

import lombok.Data;
import lombok.extern.log4j.Log4j;
import lombok.extern.log4j.Log4j2;

import java.io.File;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 15:59
 */
@Data
public class DefaultBigFileProcessor extends AbstractBigFileProcessor
{
    // 碎片以及整个作品的存储路径
    private String storePath;

    public DefaultBigFileProcessor(ChunkHandler chunkHandler)
    {
        super(chunkHandler);
    }


    public void setStorePath(String storePath)
    {
        if (!storePath.endsWith(File.separator))
        {
            storePath += File.separator;
        }

        this.storePath = storePath;
    }

    @Override
    protected FileInfo createIfDirectoryNotExist(FileChunkDecorator req)
    {
        FileInfo result = new FileInfo();
        String processId = req.getProcessId();
        String fileOriginName = req.getFileOriginName();
        String directory = this.storePath + File.separator + processId;
        String filePath = directory + File.separator + fileOriginName;
        File file = new File(filePath);
        if (!file.exists())
        {
            // 创建文件夹
            if (!file.getParentFile().exists())
            {
                //如果目标文件所在的目录不存在，则创建父目录
                file.getParentFile().mkdirs();
            }
            result.setFileStoreDirectory(directory);
        } else
        {
            result.setUploaded(true);
        }
        return result;
    }

}
