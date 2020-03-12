package com.charile.file;

import org.apache.commons.io.FileUtils;
import org.junit.Test;
import org.springframework.mock.web.MockMultipartFile;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;
import java.util.Random;
import java.util.UUID;

import static org.junit.Assert.*;

public class FTPFileServiceTest
{

    @Test
    public void upload() throws IOException
    {
//        FTPClientPool.FTPBean ftpBean = new FTPClientPool.FTPBean("39.100.60.58", 22, "vlink", "gh123456!");
//        ftpBean.setSftp(true);
//        FTPFileService ftpFileService = new FTPFileService(ftpBean);
//        ftpFileService.setHttpPortal("http://");
//        File file = new File("/Users/joker/Desktop/a.jpg");
//        byte[] bytes = new byte[0];
//        bytes = FileUtils.readFileToByteArray(file);
//        MultipartFile file1 = new MockMultipartFile("asd", bytes);
//        String ftpPath = "/home/shengtang/";
//        String newName = UUID.randomUUID().toString();
//        String upload = ftpFileService.upload(file1, ftpPath, newName, IFileStrategy.IMG_TYPE);
//        System.out.println(upload);

    }


    @Test
    public void changeEncode()
    {
    }

    @Test
    public void delete()
    {
    }
}