package com.charile.file;

import com.charile.exception.FTPConfigException;
import com.charile.exception.SFTPDataTransferException;
import com.charile.exception.SFTPLoginException;
import com.charile.utils.SFTPUtil;
import com.jcraft.jsch.JSchException;
import com.jcraft.jsch.SftpException;
import lombok.Data;
import lombok.extern.java.Log;
import lombok.extern.log4j.Log4j2;
import org.springframework.web.multipart.MultipartFile;

import java.io.*;

/**
 * The type Sftp file service.
 *
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间 ：2020-03-09 09:18
 */
@Data
public class SFTPFileService extends AbstractFIleStrategy
{
    private SFTPUtil sftpUtil;

    /**
     * Instantiates a new Sftp file service.
     *
     * @param ip       the ip
     * @param userName the user name
     * @param password the password
     */
    public SFTPFileService(String ip, String userName, String password)
    {
        this.sftpUtil = new SFTPUtil(userName, password, ip, 22);
    }

    //
    @Override
    public UploadResponse upload(MultipartFile file, String storePath, String newFileName, String key) throws IOException
    {
        // 变更 /qwe/ => qwe
        storePath = this.replaceStorepath(storePath);
        UploadResponse result = new UploadResponse();

        InputStream inputStream = file.getInputStream();
        return getUploadResponse(inputStream, storePath, newFileName, key, result);
    }

    @Override
    public UploadResponse upload(InputStream inputStream, String storePath, String newFileName, String key) throws IOException
    {
        // 变更 /qwe/ => qwe
        storePath = this.replaceStorepath(storePath);
        UploadResponse result = new UploadResponse();

        return getUploadResponse(inputStream, storePath, newFileName, key, result);
    }

    private String replaceStorepath(String storePath)
    {
        Character c = '/';
        if (c.equals(storePath.charAt(0)))
        {
            storePath = storePath.substring(1);
        }
        if (c.equals(storePath.charAt(storePath.length() - 1)))
        {
            storePath = storePath.substring(0, storePath.length() - 1);
        }
        return storePath;

    }

    private UploadResponse getUploadResponse(InputStream inputStream, String storePath, String newFileName, String key, UploadResponse result)
    {
        try
        {
            try
            {
                sftpUtil.login();
            } catch (JSchException e)
            {
                e.printStackTrace();
                throw new SFTPLoginException(e);
            }
            String storeBasePath = getStoreBasePath(key);
            String visitPrefix = getVisitPrefix(key);
            if (storeBasePath == null)
            {
                throw new FTPConfigException("配置错误:" + String.format("%s对应的找不到所属的信息", key));
            }
            try
            {
                sftpUtil.upload(storeBasePath + File.separator, storePath, newFileName, inputStream);
            } catch (SftpException e)
            {
                e.printStackTrace();
                throw new SFTPDataTransferException(e);
            }

            result.setStorePath(storeBasePath + File.separator + storePath + File.separator + newFileName);
            result.setMappingPath(visitPrefix + File.separator + storePath + File.separator + newFileName);
            return result;
        } finally
        {
            if (null != this.sftpUtil)
            {
                sftpUtil.logout();
            }

        }
    }

    @Override
    public Boolean delete(String filePathName)
    {
        return null;
    }
}
