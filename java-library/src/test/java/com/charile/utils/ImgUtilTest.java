package com.charile.utils;

import com.jcraft.jsch.IO;
import lombok.Data;
import org.junit.Test;

import java.io.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-18 17:08
 */
public class ImgUtilTest
{

    @Test
    public void testCompress() throws IOException
    {
        String originFilePath = "/Users/joker/Desktop/aaa.jpg";
        String orgiinOutPutPath = "/Users/joker/Desktop/bbb.jpg";
        boolean b = ImgUtil.compressPic(originFilePath, orgiinOutPutPath);
        System.out.println(b);
    }

    @Test
    public void testPngComress() throws IOException
    {
        File in = new File("/Users/joker/Desktop/b.jpg");
        InputStream inputStream = new FileInputStream(in);
        ByteArrayOutputStream out = new ByteArrayOutputStream();
        ImgUtil.compressPng(inputStream, out);
    }
}
