package com.charlie.crypt;


import com.charlie.exception.MessageNotCompleteException;


public class Envelope  {
    // 加密信封信息
    byte[] envelopeData;
    // 创建加密信封的加密方法,既非对称加解密
    EnumCertAlgorithm encryptMethod;
    // 生成加密信封使用时用到的加密公钥(方便用户找到匹配的私钥进行进行解密)(可选，根据实际业务)
    String encryptPublicKey;
    // 信封标识符 (可选,根据实际业务需要设置, 保函业务需要)
//    @Optional
    String envelopeIdentifier;
//    @Optional
    String extension;
//    @Optional
    String description;

    public Envelope() {
    }

    public Envelope(byte[] envelopeData, EnumCertAlgorithm encryptMethod, String encryptPublicKey) {
        setEnvelopeData(envelopeData);
        setEncryptMethod(encryptMethod);
        setEncryptPublicKey(encryptPublicKey);
    }

    public Envelope(byte[] envelopeData, EnumCertAlgorithm encryptMethod, String encryptPublicKey, String envelopeIdentifier, String extension,
                      String description) {
        setEnvelopeData(envelopeData);
        setEncryptMethod(encryptMethod);
        setEncryptPublicKey(encryptPublicKey);
        setEnvelopeIdentifier(envelopeIdentifier);
        setExtension(extension);
        setDescription(description);
    }

    public byte[] getEnvelopeData() {
        return envelopeData;
    }

    public void setEnvelopeData(byte[] envelopeData) {
        if (envelopeData == null) {
            throw new RuntimeException("Required value [envelopeData] should not be set to null.");
        }
        this.envelopeData = envelopeData;
    }

    public EnumCertAlgorithm getEncryptMethod() {
        return encryptMethod;
    }

    public void setEncryptMethod(EnumCertAlgorithm encryptMethod) {
        if (encryptMethod == null) {
            throw new RuntimeException("Required value [encryptMethod] should not be set to null.");
        }
        this.encryptMethod = encryptMethod;
    }

    public String getEncryptPublicKey() {
        return encryptPublicKey;
    }

    public void setEncryptPublicKey(String encryptPublicKey) {
        if (encryptPublicKey == null) {
            throw new RuntimeException("Required value [encryptPublicKey] should not be set to null.");
        }
        this.encryptPublicKey = encryptPublicKey;
    }

    public String getEnvelopeIdentifier() {
        return envelopeIdentifier;
    }

    public void setEnvelopeIdentifier(String envelopeIdentifier) {
        this.envelopeIdentifier = envelopeIdentifier;
    }

    public String getExtension() {
        return extension;
    }

    public void setExtension(String extension) {
        this.extension = extension;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }



    public Envelope shallowCopy() {
        Envelope obj = new Envelope();
        obj.envelopeData = this.envelopeData;
        obj.encryptMethod = this.encryptMethod;
        obj.encryptPublicKey = this.encryptPublicKey;
        obj.envelopeIdentifier = this.envelopeIdentifier;
        obj.extension = this.extension;
        obj.description = this.description;
        return obj;
    }



    
    public void validateComplete() throws MessageNotCompleteException
    {
        if (this.getEnvelopeData() == null)
            throw new MessageNotCompleteException("Required value [EnvelopeData] should not be set to null.");
        if (this.getEncryptMethod() == null)
            throw new MessageNotCompleteException("Required value [EncryptMethod] should not be set to null.");
        if (this.getEncryptPublicKey() == null)
            throw new MessageNotCompleteException("Required value [EncryptPublicKey] should not be set to null.");
    }

    /********************************************************************************
     * 以下为自定义函数
     ********************************************************************************/

}