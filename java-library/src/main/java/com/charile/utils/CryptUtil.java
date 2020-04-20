package com.charile.utils;

import org.apache.commons.codec.binary.Hex;
import org.apache.commons.codec.digest.DigestUtils;
import org.apache.commons.io.FileUtils;

import java.io.*;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Base64;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2019-12-27 11:31
 */
public class CryptUtil
{
    /**
     * MD5加密
     *
     * @param password
     * @return
     * @author joker
     * @date 创建时间：2018年3月5日 下午7:26:17
     */
    public static String md5Encrypt(String password)
    {

        try
        {
            // 得到一个信息摘要器
            MessageDigest digest = MessageDigest.getInstance("md5");
            byte[] result = digest.digest(password.getBytes());
            StringBuffer buffer = new StringBuffer();
            // 把每一个byte 做一个与运算 0xff;
            for (byte b : result)
            {
                // 与运算
                int number = b & 0xff;// 加盐
                String str = Integer.toHexString(number);
                if (str.length() == 1)
                {
                    buffer.append("0");
                }
                buffer.append(str);
            }

            // 标准的md5加密后的结果
            return new String(buffer);
        } catch (NoSuchAlgorithmException e)
        {
            e.printStackTrace();
            return "";
        }
    }


    public static String Sha256(String str)
    {
        return getSHA256StrJava(str);
    }

    private static String getSHA256StrJava(String str)
    {

        MessageDigest messageDigest;
        String encodeStr = "";
        try
        {
            messageDigest = MessageDigest.getInstance("SHA-256");
            messageDigest.update(str.getBytes("UTF-8"));
            encodeStr = byte2Hex(messageDigest.digest());
        } catch (NoSuchAlgorithmException e)
        {
            e.printStackTrace();
            throw new RuntimeException("算法错误", e);
        } catch (UnsupportedEncodingException e)
        {
            e.printStackTrace();
            throw new RuntimeException("编码失败", e);
        }
        return encodeStr;
    }

    private static String byte2Hex(byte[] bytes)
    {
        StringBuffer stringBuffer = new StringBuffer();
        String temp = null;
        for (int i = 0; i < bytes.length; i++)
        {
            temp = Integer.toHexString(bytes[i] & 0xFF);
            if (temp.length() == 1)
            {
//1得到一位的进行补0操作
                stringBuffer.append("0");
            }
            stringBuffer.append(temp);
        }
        return stringBuffer.toString();
    }

    private final static String[] strHex = {"0", "1", "2", "3", "4", "5",
            "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"};

    public static String md5File(String path) throws IOException, NoSuchAlgorithmException
    {
//        String s = DigestUtils.md5Hex(new FileInputStream(path));
//        return s;

//        FileInputStream fileInputStream = new FileInputStream(new File(path));
        byte[] bytes = FileUtils.readFileToByteArray(new File(path));
        MessageDigest md5 = null;
        md5 = MessageDigest.getInstance("MD5");
        md5.update(bytes);
        byte[] m = md5.digest();//加密
        // BASE64Encoder encoder = new BASE64Encoder();
        //血坑！！！！！！！！！！！！！
        return Base64.getEncoder().encodeToString(m);
    }
}
