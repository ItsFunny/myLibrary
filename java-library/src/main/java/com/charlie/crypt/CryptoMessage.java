package com.charlie.crypt;

import com.charlie.exception.MessageNotCompleteException;
import org.apache.commons.lang3.builder.Builder;

public class CryptoMessage  implements Builder<CryptoMessage>
{
    String platformId;
    // 加密前hash
    byte[] hashBeforeEncrypt;
    // 对加密前密文hash的hash方法
    EnumHashMethod hashMethod;
    // 对密文信息的加密方法,既对称加密方法
    EnumSymmetryEncryptionType symmEncryptMethod;
    // 多个加密信封信息
    Envelope envelopInfo;

    public CryptoMessage() {
    }

    public CryptoMessage(EnumSymmetryEncryptionType encryptMethod, Envelope envelopInfo) {
        setEncryptMethod(encryptMethod);
        setEnvelopInfo(envelopInfo);
    }

    public CryptoMessage(String platformId, byte[] hashBeforeEncrypt, EnumHashMethod hashMethod, EnumSymmetryEncryptionType encryptMethod,
                         Envelope envelopInfo) {
        setPlatformId(platformId);
        setHashBeforeEncrypt(hashBeforeEncrypt);
        setHashMethod(hashMethod);
        setEncryptMethod(encryptMethod);
        setEnvelopInfo(envelopInfo);
    }

  
    public String getPlatformId() {
        return platformId;
    }

  
    public void setPlatformId(String platformId) {
        this.platformId = platformId;
    }

  
    public byte[] getHashBeforeEncrypt() {
        return hashBeforeEncrypt;
    }

  
    public void setHashBeforeEncrypt(byte[] hashBeforeEncrypt) {
        this.hashBeforeEncrypt = hashBeforeEncrypt;
    }

  
    public EnumHashMethod getHashMethod() {
        return hashMethod;
    }

  
    public void setHashMethod(EnumHashMethod hashMethod) {
        this.hashMethod = hashMethod;
    }

  

  
    public void setEncryptMethod(EnumSymmetryEncryptionType encryptMethod) {
        if (encryptMethod == null) {
            throw new MessageNotCompleteException("Required value [encryptMethod] should not be set to null.");
        }
        this.symmEncryptMethod = encryptMethod;
    }

  
    public Envelope getEnvelopInfo() {
        return envelopInfo;
    }


    public EnumSymmetryEncryptionType getSymmEncryptMethod()
    {
        return symmEncryptMethod;
    }

    public void setSymmEncryptMethod(EnumSymmetryEncryptionType symmEncryptMethod)
    {
        this.symmEncryptMethod = symmEncryptMethod;
    }

    public void setEnvelopInfo(Envelope envelopInfo) {
        if (envelopInfo == null) {
            throw new MessageNotCompleteException("Required value [envelopInfo] should not be set to null.");
        }
        this.envelopInfo = envelopInfo;
    }

  


  

  
    public CryptoMessage shallowCopy() {
        CryptoMessage obj = new CryptoMessage();
        obj.platformId = this.platformId;
        obj.hashBeforeEncrypt = this.hashBeforeEncrypt;
        obj.hashMethod = this.hashMethod;
        obj.symmEncryptMethod = this.symmEncryptMethod;
        obj.envelopInfo = this.envelopInfo;
        return obj;
    }

    @Override
    public CryptoMessage build()
    {
        return null;
    }


    /********************************************************************************
     * 以下为自定义函数
     ********************************************************************************/

}