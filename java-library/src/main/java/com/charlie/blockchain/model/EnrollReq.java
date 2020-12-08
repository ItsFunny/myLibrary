package com.charlie.blockchain.model;


import com.charlie.blockchain.IKeyPairGenerator;
import com.charlie.service.IValidator;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-28 23:29
 */
public class EnrollReq implements IValidator
{
    private String userName;
    private String password;
    private String profile;

    private IKeyPairGenerator keyPairGenerator;

    public String getProfile()
    {
        return profile;
    }

    public void setProfile(String profile)
    {
        this.profile = profile;
    }

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

    public IKeyPairGenerator getKeyPairGenerator()
    {
        return keyPairGenerator;
    }

    public void setKeyPairGenerator(IKeyPairGenerator keyPairGenerator)
    {
        this.keyPairGenerator = keyPairGenerator;
    }

    @Override
    public void valid()
    {

    }
}
