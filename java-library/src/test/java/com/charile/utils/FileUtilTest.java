package com.charile.utils;

import org.junit.Test;

import java.io.File;
import java.io.IOException;

import static org.junit.Assert.*;

public class FileUtilTest
{

    @Test
    public void getSuffix()
    {
        String path = "/a/b/c.jpg";
        String suffix = FileUtils.getSuffix(path);
        System.out.println(suffix);
    }


    @Test
    public void splitFiles() throws IOException
    {
        String originFile = "/Users/joker/Downloads/阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";
        String toDirectFilePath = "/Users/joker/Desktop/split";
        // 按100m 划分
        FileUtils.splitFiles(originFile, 100, toDirectFilePath, false);
    }

    @Test
    public void mergeFile() throws IOException
    {
        String directory = "/Users/joker/Desktop/split";
        String outputFile = "/Users/joker/Desktop/阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";
        FileUtils.mergeFile(directory, outputFile);
    }

    @Test
    public void testSplit() throws Exception
    {
        String file = "/Users/joker/Downloads/1.pdf";
        FileUtils.splitFiles("/Users/joker/Downloads/1.pdf", 80, "/Users/joker/Downloads/1", false);
    }

    @Test
    public void testtt() throws Exception
    {
        long a = 1024 * 1024 * 1024;
        System.out.println(a);
//        String file1 = "/Users/joker/Downloads/1/1.pdf-1";
//        String file2 = "/Users/joker/Downloads/11.pdf";
//        byte[] bytes = FileUtils.readFileToByteArray(new File(file1));
//        byte[] bytes2 = FileUtils.readFileToByteArray(new File(file2));
//        for (int i = 0; i < bytes.length; i++)
//        {
//            if (bytes[i] != bytes2[i])
//            {
//                System.out.println(bytes[i]);
//                System.out.println(bytes2[i]);
//                throw new RuntimeException("不匹配:");
//            }
//        }
    }

}