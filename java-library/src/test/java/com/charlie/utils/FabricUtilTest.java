package com.charlie.utils;

import org.apache.commons.codec.binary.Hex;
import org.hyperledger.fabric.protos.common.Common;
import org.junit.Test;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.Base64;

public class FabricUtilTest
{

    @Test
    public void getRWSetFromBlock() throws Exception
    {
        String path = "/Users/joker/fsdownload/a.block";
        byte[] bytes = FileUtils.fileToBytes(path);
        Common.Block block = Common.Block.parseFrom(bytes);
        System.out.println(block);
//        BlockInfo blockInfo = JSONUtil.jsonFileToObj(, Common.Block.class);
//        System.out.println(blockInfo);

    }

    @Test
    public void test111() throws Exception
    {
        byte[] bytes = GMUtil.parseSM2PrvK("FEYPhWXO8r+SuTGHH1sAHfUi7bWwZkbN5jtlDILQN3k=");
        String msg = "123";
        String pubKey = "205NGcj2F2/cB4YeOeMZUKURUJIOrLRoJOP/Yi88kHsAY6hcpmCz3mEzLLtNQJUTdv0mXaFVPDg0hk8j37i4/A==";
        byte[] pubBytes = Base64Utils.decode(pubKey);
        byte[] sign = GMUtil.sign(bytes, msg.getBytes());
        boolean b = GMUtil.verifyOwner(pubBytes, msg.getBytes(), sign);
        System.out.println(b);

    }

    class A
    {
        String name = "123";

        public String getName()
        {
            return name;
        }

        public void setName(String name)
        {
            this.name = name;
        }
    }

    @Test
    public void test222() throws Exception
    {
        String key = "f86781b9ffbb8f1a84cf74da5f9fe425";
        String appId = "88888";
        String reqJson = "123";
        String l = "1609899849872";
        String encode = hmacData(appId, key, l, reqJson);
        System.out.println(encode);
    }

    public static byte[] encodeHmacSHA256(byte[] data, byte[] key) throws Exception
    {
        // 还原密钥
        SecretKeySpec secretKey = new SecretKeySpec(key, "HmacSHA256");
        // 实例化 mac
        Mac mac = Mac.getInstance(secretKey.getAlgorithm());
        // 初始化 mac
        mac.init(secretKey);
        // 执行消息摘要
        byte[] bytes = mac.doFinal(data);
        System.out.println("======");
        System.out.println(Base64.getEncoder().encodeToString(bytes));
        return bytes;

    }

    public static String hmacData(String appId, String secret, String timestamp, String body) throws Exception
    {
//        StringBuilder sb = new StringBuilder(appId);
//        sb.append(timestamp).append(body);
//        System.out.println("hmac加密的原数据为:"+sb.toString());
//        byte[] signData = encodeHmacSHA256(sb.toString().getBytes(StandardCharsets.UTF_8), secret.getBytes(StandardCharsets.UTF_8));
//        System.out.println(HexUtil.bytes2HexString(signData));
//        return Base64.getEncoder().encodeToString(signData);

        StringBuilder sb = new StringBuilder(appId);
        sb.append(timestamp).append(body);
        System.out.println("=============");
        System.out.println("hmac加密的原数据为:"+sb.toString());
        byte[] signData = encodeHmacSHA256(secret.getBytes(StandardCharsets.UTF_8), sb.toString().getBytes(StandardCharsets.UTF_8));
        System.out.println("signDataHeader:"+Base64.getEncoder().encodeToString(signData));
        return Base64.getEncoder().encodeToString(signData);
    }

    @Test
    public void test11111() throws Exception
    {
        // 还原密钥
        SecretKeySpec secretKey = new SecretKeySpec("f86781b9ffbb8f1a84cf74da5f9fe425".getBytes(StandardCharsets.UTF_8), "HmacSHA256");
        // 实例化 mac
        Mac mac = Mac.getInstance(secretKey.getAlgorithm());
        // 初始化 mac
        mac.init(secretKey);
        // 执行消息摘要
        byte[] bytes = mac.doFinal("123".getBytes(StandardCharsets.UTF_8));
        System.out.println(Hex.encodeHexString(bytes));
        System.out.println(Base64.getEncoder().encodeToString(bytes));

    }

    @Test
    public void test2222()
    {
        String str = "313233";
        byte[] bytes = HexUtil.hexString2Bytes(str);
        System.out.println(new String(bytes));
    }
}