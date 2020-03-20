/**
 * @author joker
 * @date 创建时间：2018年8月8日 下午4:47:02
 */
package com.charile.file;

import java.io.File;
import java.io.IOException;
import java.io.InputStream;

import org.springframework.web.multipart.MultipartFile;

/**
 * 默认的文件上传,上传于本地
 *
 * @author joker
 * @date 创建时间：2018年8月8日 下午4:47:02
 */
public class DefaultFileService extends AbstractFIleStrategy
{

    @Override
    public UploadResponse upload(MultipartFile file, String storePath, String newFileName, String key)
    {
        UploadResponse result = new UploadResponse();
        String visitPrefix = getVisitPrefix(key);
        String storeBasePath = getStoreBasePath(key);
        File dirFile = new File(storeBasePath + File.separator + storePath);
        if (!dirFile.exists())
        {
            dirFile.mkdirs();
        }
        File newFiel = new File(dirFile.getAbsolutePath() + File.separator + newFileName);
        try
        {
            file.transferTo(newFiel);
            result.setStorePath(storeBasePath + File.separator + storePath + File.separator + newFileName);
            result.setMappingPath(visitPrefix + File.separator + storePath + File.separator + newFileName);
            return result;
        } catch (IllegalStateException | IOException e)
        {
            e.printStackTrace();
            return null;
        }
    }

    @Override
    public UploadResponse upload(InputStream inputStream, String storePath, String newFileName, String key) throws IOException
    {
        return null;
    }

    @Override
    public Boolean delete(String filePathName)
    {
        File file = null;
        if (filePathName.contains("."))
        {

            file = new File(filePathName);
        } else
        {
            file = new File(filePathName + File.separator);
        }
        if (!file.exists())
        {
            return true;
        }
        if (file.isDirectory())
        {
            File[] listFiles = file.listFiles();
            if (listFiles != null)
            {
                for (File file2 : listFiles)
                {
                    delete(file2.getAbsolutePath());
                }
            }
            return file.delete();
        } else
        {
            return file.delete();
        }
    }

}
