package cn.bidsun.blockchain.model;

import java.security.PrivateKey;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-28 23:29
 */
public class EnrollResp
{
    private String signCertPem;
    private String privateKeyPem;

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
