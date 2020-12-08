package com.charlie.crypt.impl;

import com.charlie.constants.PemConstant;
import com.charlie.crypt.EnumCertAlgorithm;
import com.charlie.exception.EncryptException;
import com.charlie.model.CertInfo;
import com.charlie.utils.Base64Utils;
import com.charlie.utils.GMUtil;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:01
 */
public class DefaultSM2AsymmetricCryptoImpl extends  AsymmetricCryptoImpl
{
    public static DefaultSM2AsymmetricCryptoImpl newInstance(){
        DefaultSM2AsymmetricCryptoImpl defaultSM2AsymmetricCrypto=new DefaultSM2AsymmetricCryptoImpl();
        return defaultSM2AsymmetricCrypto;
    }
    @Override
    protected void init()throws Exception
    {
        this.prvKey="-----BEGIN PRIVATE KEY-----" +
                "V9JJ6IqFSMUVxugnDdlLNKCQesm4bPflogZAoCzotQQ=" +
                "-----END PRIVATE KEY-----";
        this.prvKeyBytes= Base64Utils.decode(GMUtil.replace(PemConstant.PRIVATEKEY,this.prvKey));
        String  cert="-----BEGIN CERTIFICATE-----" +
                "MIIB3DCCAYKgAwIBAgIUGDewX7bRwzq6iYcqTxlgiu9xvOEwCgYIKoEcz1UBg3Uw" +
                "XjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWpp" +
                "bmcxEzARBgNVBAoTCmJpZHN1bi5jb20xFjAUBgNVBAMTDWNhLmJpZHN1bi5jb20w" +
                "HhcNMjAxMjA1MDYxODU1WhcNMjEwMTE1MjIxODU1WjAhMQ8wDQYDVQQLEwZjbGll" +
                "bnQxDjAMBgNVBAMTBVh2N1BXMFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEMgrm" +
                "bof+ChEn2x0GX48fPq84Wm7E+kaqvdlc+cg0AShGS6GOXg/ALMngE7T0/CgM5l95" +
                "9LZlfex3QyNNFucbG6NbMFkwKwYDVR0jBCQwIoAgbWW7LFnLZRKhHRmYummNhD6D" +
                "2rYYkHAu9eV//waElIwwKgYIKgMEBQYHCAEEHnsiYXR0cnMiOnsiaXNFQklEU1VO" +
                "IjoidHJ1ZSJ9fTAKBggqgRzPVQGDdQNIADBFAiEApAXdtuQH7kA1UEBUhpIikiPo" +
                "L43jkazO4hRMDClaxR4CIADQjXyc2nyrNAtEh25J3kEuzMkkP1BZe2c41ZAL0E7I" +
                "-----END CERTIFICATE-----";
        CertInfo certInfo = GMUtil.parseSM2CertStr(cert);
        String pubKey = certInfo.getPubKey();
        this.pubKey=pubKey;
        this.pubKeyBytes=Base64Utils.decode(this.pubKey);
    }

    @Override
    protected byte[] encrypt(byte[] origin) throws EncryptException
    {
        return GMUtil.encrypt(this.pubKeyBytes,origin);
    }

    @Override
    protected byte[] decrypt(byte[] encrypt) throws EncryptException
    {
        return GMUtil.decrypt(this.prvKeyBytes,encrypt);
    }

    @Override
    public Boolean validIsMine(Serializable type)
    {
        return EnumCertAlgorithm.SM2_256.equals(type);
    }
}
