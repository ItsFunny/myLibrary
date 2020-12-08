package com.charlie.crypt;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:15
 */
public class CryptoMessageBO
{
    String platformId;
    // 对加密前密文hash的hash方法
    EnumHashMethod hashMethod;
    // 对密文信息的加密方法
    EnumSymmetryEncryptionType symmEncryptMethod;


    private EnvelopBO envelope;

    public EnumSymmetryEncryptionType getSymmEncryptMethod()
    {
        return symmEncryptMethod;
    }

    public void setSymmEncryptMethod(EnumSymmetryEncryptionType symmEncryptMethod)
    {
        this.symmEncryptMethod = symmEncryptMethod;
    }

    public String getPlatformId()
    {
        return platformId;
    }

    public void setPlatformId(String platformId)
    {
        this.platformId = platformId;
    }


    public EnumHashMethod getHashMethod()
    {
        return hashMethod;
    }

    public void setHashMethod(EnumHashMethod hashMethod)
    {
        this.hashMethod = hashMethod;
    }


    public EnvelopBO getEnvelope()
    {
        return envelope;
    }

    public void setEnvelope(EnvelopBO envelope)
    {
        this.envelope = envelope;
    }
}
