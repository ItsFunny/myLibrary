package com.charlie.crypt;

public class CryptoMessageJSONObject
{
    private String hashClazzName;
    private String symmClazzName;
    private String asymmClazzName;
    private String messageJson;

    public String getMessageJson()
    {
        return messageJson;
    }

    public void setMessageJson(String messageJson)
    {
        this.messageJson = messageJson;
    }

    public static  CryptoMessageJSONObject toJson(CryptoMessage cryptoMessage){
        CryptoMessageJSONObject cryptoMessageJSONObject=new CryptoMessageJSONObject();

        return cryptoMessageJSONObject;
    }




    public String getHashClazzName()
    {
        return hashClazzName;
    }

    public void setHashClazzName(String hashClazzName)
    {
        this.hashClazzName = hashClazzName;
    }

    public String getSymmClazzName()
    {
        return symmClazzName;
    }

    public void setSymmClazzName(String symmClazzName)
    {
        this.symmClazzName = symmClazzName;
    }

    public String getAsymmClazzName()
    {
        return asymmClazzName;
    }

    public void setAsymmClazzName(String asymmClazzName)
    {
        this.asymmClazzName = asymmClazzName;
    }
}