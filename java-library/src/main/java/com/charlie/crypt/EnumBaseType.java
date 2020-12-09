package com.charlie.crypt;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:01
 */
public enum EnumBaseType
{
    ENUM_ASYMMETRIC_ECDSA(1,"ecdsa非对称加密"),
    ENUM_ASYMMETRIC_SM2(2,"sm2非对称加密"),
    ENUM_ASYMMETRIC_RSA(3,"rsa非对称加密"),

    ENUM_SYMMETRIC_AES(10,"AES对称加密"),
    ENUM_SYMMETRIC_SM3(11,"SM3对称几秒"),


    ENUM_HASH_MD5(100,"MD5哈希算法"),
    ENUM_HASH_SHA(101,"SHA哈希算法"),

    ;
    private int value;
    private String description;

    EnumBaseType(int value, String description)
    {
        this.value = value;
        this.description = description;
    }

    public int getValue()
    {
        return value;
    }

    public void setValue(int value)
    {
        this.value = value;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }}
