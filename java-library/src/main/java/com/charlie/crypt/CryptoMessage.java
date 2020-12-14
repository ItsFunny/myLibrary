package com.charlie.crypt;

import com.alibaba.fastjson.JSONObject;
import com.charlie.base.IBase;
import com.charlie.base.ITO;
import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.crypt.opts.ISymmetricOpts;
import com.charlie.exception.DeserializeException;
import com.charlie.exception.MessageNotCompleteException;
import com.charlie.exception.SerializeException;
import com.charlie.utils.JSONUtil;
import org.apache.commons.lang3.builder.Builder;

public class CryptoMessage implements Builder<CryptoMessage>, ITO<byte[]>
{
    String platformId;
    // 加密前hash
    byte[] hashBeforeEncrypt;
    // 对加密前密文hash的hash方法
//    EnumHashMethod hashMethod;
    // 对密文信息的加密方法,既对称加密方法
//    EnumSymmetryEncryptionType symmEncryptMethod;
    // 多个加密信封信息
    Envelope envelopInfo;

    private IHashOpts hashOpts;
    private ISymmetricOpts symmetricOpts;

    public CryptoMessage()
    {
    }

//    public CryptoMessage(EnumSymmetryEncryptionType encryptMethod, Envelope envelopInfo) {
//        setEncryptMethod(encryptMethod);
//        setEnvelopInfo(envelopInfo);
//    }

//    public CryptoMessage(String platformId, byte[] hashBeforeEncrypt, EnumHashMethod hashMethod, EnumSymmetryEncryptionType encryptMethod,
//                         Envelope envelopInfo) {
//        setPlatformId(platformId);
//        setHashBeforeEncrypt(hashBeforeEncrypt);
//        setHashMethod(hashMethod);
//        setEncryptMethod(encryptMethod);
//        setEnvelopInfo(envelopInfo);
//    }

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


    public String getPlatformId()
    {
        return platformId;
    }


    public void setPlatformId(String platformId)
    {
        this.platformId = platformId;
    }


    public byte[] getHashBeforeEncrypt()
    {
        return hashBeforeEncrypt;
    }


    public void setHashBeforeEncrypt(byte[] hashBeforeEncrypt)
    {
        this.hashBeforeEncrypt = hashBeforeEncrypt;
    }


//    public EnumHashMethod getHashMethod() {
//        return hashMethod;
//    }


//    public void setHashMethod(EnumHashMethod hashMethod) {
//        this.hashMethod = hashMethod;
//    }


//    public void setEncryptMethod(EnumSymmetryEncryptionType encryptMethod) {
//        if (encryptMethod == null) {
//            throw new MessageNotCompleteException("Required value [encryptMethod] should not be set to null.");
//        }
//        this.symmEncryptMethod = encryptMethod;
//    }


    public Envelope getEnvelopInfo()
    {
        return envelopInfo;
    }


//    public EnumSymmetryEncryptionType getSymmEncryptMethod()
//    {
//        return symmEncryptMethod;
//    }

//    public void setSymmEncryptMethod(EnumSymmetryEncryptionType symmEncryptMethod)
//    {
//        this.symmEncryptMethod = symmEncryptMethod;
//    }

    public void setEnvelopInfo(Envelope envelopInfo)
    {
        if (envelopInfo == null)
        {
            throw new MessageNotCompleteException("Required value [envelopInfo] should not be set to null.");
        }
        this.envelopInfo = envelopInfo;
    }


    public CryptoMessage shallowCopy()
    {
        CryptoMessage obj = new CryptoMessage();
        obj.platformId = this.platformId;
        obj.hashBeforeEncrypt = this.hashBeforeEncrypt;
        obj.hashOpts = this.hashOpts;
        obj.symmetricOpts = this.symmetricOpts;
        obj.envelopInfo = this.envelopInfo;
        return obj;
    }

    @Override
    public CryptoMessage build()
    {
        return null;
    }


    public String toJson()
    {
        final CryptoMessage message = this;
        CryptoMessageJSONObject res = new CryptoMessageJSONObject();
        String hashClazzName = message.getHashOpts().getClass().getName();
        String symmClazzName = message.getSymmetricOpts().getClass().getName();
        String asymmClazzName = message.getEnvelopInfo().getAsymmetricOpts().getClass().getName();
        String messageJson = JSONUtil.toFormattedJson(message);
        res.setMessageJson(messageJson);
        res.setHashClazzName(hashClazzName);
        res.setSymmClazzName(symmClazzName);
        res.setAsymmClazzName(asymmClazzName);
        return JSONUtil.toFormattedJson(res);
    }



    @Override
    public byte[] to()
    {
        return this.toJson().getBytes();
    }



}