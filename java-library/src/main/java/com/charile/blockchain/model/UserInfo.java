package com.charile.blockchain.model;

import com.charile.base.IKeyImporter;
import lombok.Data;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;
import org.hyperledger.fabric.sdk.identity.X509Enrollment;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.PrivateKey;
import java.util.Set;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 15:20
 */
@Data
public class UserInfo implements User
{
    protected String name;
    protected String enrollSecret;
    protected String mspId;
    private Set<String> roles;

    private String account;
    private String affiliation;
    private Enrollment enrollment;

    public UserInfo(IKeyImporter keyImporter, String mspId, String name, String keyFile, String certFile)
    {
        this.name = name;
        this.mspId = mspId;
        try
        {
            this.enrollment = this.loadFromPemFile(keyImporter, keyFile, certFile);
        } catch (Exception e)
        {
            throw new RuntimeException(e);
        }
    }

    public UserInfo(IKeyImporter keyImporter, String mspId, String name, byte[] keyBytes, byte[] certBytes)
    {
        this.name = name;
        this.mspId = mspId;
        try
        {
            this.enrollment = new X509Enrollment(keyImporter.bytes2PrivateKey(keyBytes), new String(certBytes));
        } catch (Exception e)
        {
            throw new RuntimeException(e);
        }
    }



    private Enrollment loadFromPemFile(IKeyImporter keyImporter, String keyFile, String certFile) throws Exception
    {
        byte[] keyPem = Files.readAllBytes(Paths.get(keyFile));     //载入私钥PEM文本
        byte[] certPem = Files.readAllBytes(Paths.get(certFile));   //载入证书PEM文本
        PrivateKey privateKey = keyImporter.bytes2PrivateKey(keyPem);    //将PEM文本转换为私钥对象
        return new X509Enrollment(privateKey, new String(certPem));  //创建并返回X509Enrollment对象
    }

}
