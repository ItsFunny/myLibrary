package com.charile.blockchain.util;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Properties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 15:47
 */
public class TLSUtils
{
    public static Properties loadTLSFile(String rootTLSCert, String hostName) throws IOException
    {
        Properties properties = new Properties();
//    # 其实只需要一个TLS根证书就可以了，比如TLS相关的秘钥等都是可选的
        properties.put("pemBytes", Files.readAllBytes(Paths.get(rootTLSCert)));
        properties.setProperty("sslProvider", "openSSL");
        properties.setProperty("negotiationType", "TLS");
        properties.setProperty("trustServerCertificate", "true");
        properties.setProperty("hostnameOverride", hostName);
        return properties;
    }


}
