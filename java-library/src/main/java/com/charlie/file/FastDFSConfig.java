package com.charlie.file;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-26 13:31
 */
@Data
public class FastDFSConfig
{
    private String httpServerUrl;
    private String trackerServers;
    private String connectTimeoutInSeconds;
    private String networkTimeoutInSeconds;
    private String charset;
    private String httpAntiStealToken;
    private String httpSecretKey;
    private String httpTrackerHttpPort;
}
