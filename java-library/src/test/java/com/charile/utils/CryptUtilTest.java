package com.charile.utils;

import org.apache.commons.io.FileUtils;
import org.junit.Test;
import sun.misc.IOUtils;

import java.io.File;
import java.io.IOException;
import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.List;

import static org.junit.Assert.*;

public class CryptUtilTest
{

    @Test
    public void md5File() throws IOException, NoSuchAlgorithmException
    {
//        426d1ad84581dac49a5ed203a417825e
        String path = "/Users/joker/Desktop/a.jpg";
        String s = CryptUtil.md5File(path);
        System.out.println(s);
        String s1 = CryptUtil.Sha256(s);
        System.out.println();
        String s2 = "62cae8002fc1285cb3552602fb52f2badac7eda6b4047e0434c624887e9f0bd4";
        System.out.println(s1.equalsIgnoreCase(s2));
    }


    @Test
    public void sha256() throws Exception
    {
        String file1 = "/Users/joker/Downloads/1.pdf";
        String file2 = "/Users/joker/Downloads/2.pdf";
        String file3 = "/Users/joker/Downloads/3.pdf";
        byte[] bytes1 = FileUtils.readFileToByteArray(new File(file1));
        byte[] bytes2 = FileUtils.readFileToByteArray(new File(file2));
        byte[] bytes3 = FileUtils.readFileToByteArray(new File(file3));
        for (int i = 0; i < bytes3.length; i++)
        {
            if (bytes1[i] != bytes3[i])
            {
                byte[] normalBytes = new byte[20];
                for (int j = 0; j < 20; j++)
                {
                    // normalBytes.add(bytes1[i + j]);
                    normalBytes[j] = bytes1[i + j];
                }


                List<Byte> badBytes2 = new ArrayList<>();
                for (int j = 1; j < 20; j++)
                {
                    badBytes2.add(bytes2[i + j]);
                }

                List<Byte> badBytes3 = new ArrayList<>();
                for (int j = 1; j < 20; j++)
                {
                    badBytes3.add(bytes3[i + j]);
                }

                System.out.println(1);

            }


        }
        for (int i = 0; i < bytes1.length; i++)
        {
            if (bytes1[i] != bytes2[i])
            {
                System.out.println("file1 前一个元素为:" + bytes1[i - 1]);
                System.out.println("file2 的前一个元素为:" + bytes2[i - 1]);
                System.out.println("该下标所属的元素不合:" + i + "file1的元素为:" + bytes1[i] + " bytes2的值为:" + bytes2[i]);
                System.out.println("file1 下一个元素为:" + bytes1[i + 1]);
                System.out.println("file2 的下一个元素为:" + bytes2[i + 1]);

                throw new RuntimeException("该下标所属的元素不合:" + i);
            }

        }
        String sha1 = CryptUtil.sha256(bytes1);
        System.out.println(sha1);
        String sha2 = CryptUtil.sha256(bytes2);
        System.out.println(sha2);
    }

    @Test
    public void test2() throws Exception
    {
        File directory = new File("/Users/joker/Downloads/2/");
        File[] files = directory.listFiles();
        for (File file : files)
        {
            byte[] bytes = FileUtils.readFileToByteArray(file);
            System.out.println(bytes[0]);
            if (bytes[0] == -122 || bytes[1] == -122)
            {
                System.out.println(file.getAbsolutePath());
                throw new RuntimeException(file.getAbsolutePath());
            }
        }
//        FileUtil.mergeFile("/Users/joker/Downloads/2/", "/Users/joker/Downloads/5.pdf");
    }
}