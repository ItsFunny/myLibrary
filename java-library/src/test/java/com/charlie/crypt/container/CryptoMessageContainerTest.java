package com.charlie.crypt.container;

import com.charlie.crypt.*;
import com.charlie.crypt.opts.impl.DefaultAsymmetricOpts;
import com.charlie.crypt.opts.impl.DefaultHashOpts;
import com.charlie.crypt.opts.impl.DefaultSymmetricOpts;
import com.charlie.utils.JSONUtil;
import com.charlie.utils.RandomUtils;
import org.junit.Test;


public class CryptoMessageContainerTest
{

    public CryptoMessageBO newMessageBO(byte[] originData)
    {
        CryptoMessageBO cryptoMessageBO = new CryptoMessageBO();
        cryptoMessageBO.setHashOpts(new DefaultHashOpts(EnumHashMethod.MD5));
        cryptoMessageBO.setSymmetricOpts(new DefaultSymmetricOpts(EnumSymmetryEncryptionType.AES_CFB));
        EnvelopBO envelopBO = newEnvelope(originData);
        cryptoMessageBO.setEnvelope(envelopBO);
        return cryptoMessageBO;
    }

    public EnvelopBO newEnvelope(byte[] originData)
    {
        EnvelopBO envelopBO = new EnvelopBO();
        envelopBO.setOriginData(originData);
        envelopBO.setAsymmetricOpts(new DefaultAsymmetricOpts(EnumBaseType.ENUM_ASYMMETRIC_SM2, EnumAsymmetricAlgorithm.SM2_256));
        envelopBO.setDescription("description:" + RandomUtils.randomString(5));
        envelopBO.setExtension("extension:" + RandomUtils.randomString(5));
        return envelopBO;
    }

    @Test
    public void testCreate() throws Exception
    {
        byte[] originData = "123".getBytes();
        CryptoMessageContainer container = new CryptoMessageContainer();
        container.initOnce();
        CryptoContainerFactory.defaultInit();
        CryptoMessageBO cryptoMessageBO = newMessageBO(originData);
        CryptoMessage cryptoMessage = container.buildCryptoMessage(cryptoMessageBO);
        String s=JSONUtil.toFormattedJson(cryptoMessage);
        cryptoMessage= JSONUtil.json2Obj(s, CryptoMessage.class);
        System.out.println(JSONUtil.toFormattedJson(cryptoMessage));

        CryptoMessageBO decrypt = container.decrypt(cryptoMessage);
        System.out.println(decrypt);

    }

}