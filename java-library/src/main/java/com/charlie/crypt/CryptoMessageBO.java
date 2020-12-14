package com.charlie.crypt;

import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.crypt.opts.ISymmetricOpts;

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
//    EnumHashMethod hashMethod;
    // 对密文信息的加密方法,对称加密
//    EnumSymmetryEncryptionType symmEncryptMethod;
    private EnvelopBO envelope;

    IHashOpts hashOpts;
    ISymmetricOpts symmetricOpts;
    IAsymmetricOpts asymmetricOpts;

    public IHashOpts getHashOpts()
    {
        return hashOpts;
    }

    public void setHashOpts(IHashOpts hashOpts)
    {
        this.hashOpts = hashOpts;
    }

    public ISymmetricOpts getSymmetricOpts()
    {
        return symmetricOpts;
    }

    public void setSymmetricOpts(ISymmetricOpts symmetricOpts)
    {
        this.symmetricOpts = symmetricOpts;
    }

    public IAsymmetricOpts getAsymmetricOpts()
    {
        return asymmetricOpts;
    }

    public void setAsymmetricOpts(IAsymmetricOpts asymmetricOpts)
    {
        this.asymmetricOpts = asymmetricOpts;
    }

//    public EnumSymmetryEncryptionType getSymmEncryptMethod()
//    {
//        return symmEncryptMethod;
//    }
//
//    public void setSymmEncryptMethod(EnumSymmetryEncryptionType symmEncryptMethod)
//    {
//        this.symmEncryptMethod = symmEncryptMethod;
//    }

    public String getPlatformId()
    {
        return platformId;
    }

    public void setPlatformId(String platformId)
    {
        this.platformId = platformId;
    }


//    public EnumHashMethod getHashMethod()
//    {
//        return hashMethod;
//    }
//
//    public void setHashMethod(EnumHashMethod hashMethod)
//    {
//        this.hashMethod = hashMethod;
//    }


    public EnvelopBO getEnvelope()
    {
        return envelope;
    }

    public void setEnvelope(EnvelopBO envelope)
    {
        this.envelope = envelope;
    }
}
