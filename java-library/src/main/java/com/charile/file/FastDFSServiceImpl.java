package com.charile.file;

import org.csource.common.MyException;
import org.csource.fastdfs.ClientGlobal;
import org.csource.fastdfs.StorageServer;
import org.csource.fastdfs.TrackerClient;
import org.csource.fastdfs.TrackerServer;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-26 12:44
 */
public class FastDFSServiceImpl extends AbstractFIleStrategy
{
    private TrackerClient client;

    public FastDFSServiceImpl() {}

    private String httpServerUrl = "";

    public FastDFSServiceImpl(FastDFSConfig config) throws IOException, MyException
    {
        //初始化配置
        Properties p = new Properties();
        p.setProperty("fastdfs.tracker_servers", config.getTrackerServers());
        p.setProperty("fastdfs.connect_timeout_in_seconds", config.getConnectTimeoutInSeconds());
        p.setProperty("fastdfs.network_timeout_in_seconds", config.getNetworkTimeoutInSeconds());
        p.setProperty("fastdfs.charset", config.getCharset());
        p.setProperty("fastdfs.http_anti_steal_token", config.getHttpAntiStealToken());
        p.setProperty("fastdfs.http_secret_key", config.getHttpSecretKey());
        p.setProperty("fastdfs.http_tracker_http_port", config.getHttpTrackerHttpPort());
        ClientGlobal.initByProperties(p);
        httpServerUrl = config.getHttpServerUrl();

        TrackerClient trackerClient = new TrackerClient(ClientGlobal.g_tracker_group);
        TrackerServer trackerServer = trackerClient.getTrackerServer();
        if (trackerServer == null)
        {
            throw new IllegalStateException("getConnection return null");
        }

        StorageServer storageServer = trackerClient.getStoreStorage(trackerServer);
        if (storageServer == null)
        {
            throw new IllegalStateException("getStoreStorage return null");
        }
    }


    @Override
    protected BatchFileUploadWrapperResp doBatchUpload(BatchFileUploadWrapper wrapper)
    {
        return null;
    }

    /**
     * Upload upload resp.
     *
     * @param file        the file
     * @param storePath   the store path  不可以以 / 起始,并且当只有1个目录的时候,要删除后面的/
     * @param newFileName the new file name,新的名字必须包含后缀
     * @param key         the key
     * @return the upload resp: 包含绝对存储路径和映射路径
     * @throws IOException the io exception
     */
    @Override
    public UploadResponse upload(MultipartFile file, String storePath, String newFileName, String key) throws IOException
    {
        return null;
    }

    @Override
    public UploadResponse upload(InputStream inputStream, String storePath, String newFileName, String key) throws IOException
    {
        return null;
    }

    /**
     * Delete boolean.
     *
     * @param filePathName the file path name
     * @return the boolean
     */
    @Override
    public Boolean delete(String filePathName)
    {
        return null;
    }
}
