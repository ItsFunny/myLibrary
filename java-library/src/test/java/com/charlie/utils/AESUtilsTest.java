package com.charlie.utils;

import org.apache.commons.codec.binary.Base64;
import org.apache.commons.codec.binary.Hex;
import org.junit.Test;

public class AESUtilsTest
{

    @Test
    public void generateKey()
    {
    }

    // ecb 测试
    @Test
    public void encrypt() throws Exception
    {
        String text = "ebidsun";
        String password = "123";
        byte[] bytes = AESUtils.generateKey(password);
        String s1 = Base64.encodeBase64String(bytes);
        System.out.println(s1);
//        String s = HexUtil.bytes2HexString(bytes);
//        System.out.println(s);
        String s = Base64.encodeBase64String(AESUtils.encrypt(text, password));
        System.out.println(s);
    }


    @Test
    public void encryptByCFB() throws Exception
    {
        String text = "XJ521521521521";
        String password = "521";
        byte[] bytes = AESUtils.generateKey(password);
        String s1 = Hex.encodeHexString(bytes);
        System.out.println(s1);

        String iv = "1234567891234567";
        byte[] ivBytes = iv.getBytes();
        String s = Hex.encodeHexString(AESUtils.encryptByCFB(text.getBytes("UTF-8"), password, ivBytes));
        System.out.println(s);

    }

    @Test
    public void ecbEncryptByKeyAndIV()
    {
        byte[] keyBytes = Base64.decodeBase64("kb6/1Ol9gbFtdzhRFf44Pg==");
        byte[] ivBytes = Base64.decodeBase64("AAECAwQFBgcICQoLDA0ODw==");
        String str = "ebidusn";
        byte[] bytes = AESUtils.ecbEncryptByKey(str, keyBytes);
        System.out.println(Base64.encodeBase64String(bytes));
    }

    @Test
    public void encryptByCFBWithEbidsun() throws Exception
    {
        String text = "ebidsun";
        String iv = "AAECAwQFBgcICQoLDA0ODw==";
        byte[] ivBytes = Base64.decodeBase64(iv);
        byte[] keyBytes = Base64.decodeBase64("kb6/1Ol9gbFtdzhRFf44Pg==");
        String s = Base64.encodeBase64String(AESUtils.encryptByCFBWithBytes(text.getBytes("UTF-8"), keyBytes, ivBytes));
        System.out.println(s);
    }

    @Test
    public void ecbEncryptByKey()
    {
        String body1 = "{\"orderId\": \"BIAOXIN7793792099979100166300165\",\"payResult\": \"success\",\"timeStamp\": \"1579082454\"}";
        String password = "poiGslCuKcpiQoyx";
        byte[] encrypt = AESUtils.encrypt(body1, password);
        System.out.println(new String(encrypt));
    }

    @Test
    public void encryptByCFB1()
    {
    }

    @Test
    public void encryptByCFBWithBytes()
    {
    }
}