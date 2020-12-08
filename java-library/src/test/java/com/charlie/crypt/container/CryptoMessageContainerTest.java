package com.charlie.crypt.container;

import com.charlie.crypt.*;
import com.charlie.utils.JSONUtil;
import com.charlie.utils.RandomUtils;
import org.junit.Test;

import static org.junit.Assert.*;

public class CryptoMessageContainerTest
{

    public CryptoMessageBO newMessageBO(byte[] originData){
        CryptoMessageBO cryptoMessageBO=new CryptoMessageBO();
        cryptoMessageBO.setHashMethod(EnumHashMethod.MD5);
        cryptoMessageBO.setSymmEncryptMethod(EnumSymmetryEncryptionType.AES_CFB);
        EnvelopBO envelopBO = newEnvelope(originData);
        cryptoMessageBO.setEnvelope(envelopBO);
        return cryptoMessageBO;
    }

    public EnvelopBO newEnvelope(byte[] originData){
        EnvelopBO envelopBO=new EnvelopBO();
        envelopBO.setOriginData(originData);
        envelopBO.setCertAlgorithm(EnumCertAlgorithm.SM2_256);
        envelopBO.setDescription("description:"+ RandomUtils.randomString(5));
        envelopBO.setExtension("extension:"+RandomUtils.randomString(5));
        return envelopBO;
    }

    @Test
    public void testCreate()throws Exception{
        byte[] originData="123".getBytes();
        CryptoMessageContainer container=new CryptoMessageContainer();
        container.initOnce();
        CryptoMessageBO cryptoMessageBO = newMessageBO(originData);
        CryptoMessage cryptoMessage = container.buildCryptoMessage(cryptoMessageBO);
        System.out.println(JSONUtil.toFormattedJson(cryptoMessage));

        CryptoMessageBO decrypt = container.decrypt(cryptoMessage);
        System.out.println(decrypt);

    }

}