package cn.bidsun.blockchain.model;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-28 23:19
 */
public class RegisterEnrollReq
{
    private String userName;
    private String password;

    private String mspId;

    public String getMspId()
    {
        return mspId;
    }

    public void setMspId(String mspId)
    {
        this.mspId = mspId;
    }

    public String getPassword()
    {
        return password;
    }

    public void setPassword(String password)
    {
        this.password = password;
    }

    public String getUserName()
    {
        return userName;
    }

    public void setUserName(String userName)
    {
        this.userName = userName;
    }
}
