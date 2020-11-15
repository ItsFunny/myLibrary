package cn.bidsun.blockchain.model;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-30 13:13
 */
public class UserRegisterReq
{
    private String userName;
    private String password;
    private String mspId;

    private String affiliation;


    public String getAffiliation()
    {
        return affiliation;
    }

    public void setAffiliation(String affiliation)
    {
        this.affiliation = affiliation;
    }

    public String getMspId()
    {
        return mspId;
    }

    public void setMspId(String mspId)
    {
        this.mspId = mspId;
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
}
