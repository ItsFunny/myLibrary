package com.charile.utils;

import org.junit.Test;

import java.io.IOException;

import static org.junit.Assert.*;

public class FileUtilTest
{

    @Test
    public void getSuffix()
    {
        String path = "/a/b/c.jpg";
        String suffix = FileUtil.getSuffix(path);
        System.out.println(suffix);
    }


    @Test
    public void splitFiles() throws IOException
    {
        String originFile = "/Users/joker/Downloads/阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";
        String toDirectFilePath = "/Users/joker/Desktop/split";
        // 按100m 划分
        FileUtil.splitFiles(originFile, 100, toDirectFilePath, false);
    }

    @Test
    public void mergeFile() throws IOException
    {
        String directory = "/Users/joker/Desktop/split";
        String outputFile = "/Users/joker/Desktop/阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";
        FileUtil.mergeFile(directory, outputFile);
    }

}