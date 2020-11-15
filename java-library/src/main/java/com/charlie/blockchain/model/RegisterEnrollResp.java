package com.charlie.blockchain.model;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-28 23:19
 */
public class RegisterEnrollResp
{
    private String signCertPem;
    private String privateKeyPem;

    private String userName;
    private String password;


    public String getUserName()
    {
        return userName;
    }

    public void setUserName(String userName)
    {
        this.userName = userName;
    }

    public String getPassword()
    {
        return password;
    }

    public void setPassword(String password)
    {
        this.password = password;
    }

    public String getSignCertPem()
    {
        return signCertPem;
    }

    public void setSignCertPem(String signCertPem)
    {
        this.signCertPem = signCertPem;
    }

    public String getPrivateKeyPem()
    {
        return privateKeyPem;
    }

    public void setPrivateKeyPem(String privateKeyPem)
    {
        this.privateKeyPem = privateKeyPem;
    }
}
