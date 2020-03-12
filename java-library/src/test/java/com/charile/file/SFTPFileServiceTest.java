package com.charile.file;

import org.apache.commons.io.FileUtils;
import org.junit.Test;
import org.springframework.mock.web.MockMultipartFile;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import static org.junit.Assert.*;

public class SFTPFileServiceTest
{

    @Test
    public void upload() throws IOException
    {
        Map<String, String> basePathMap = new HashMap<>();
        SFTPFileService sftpFileService = new SFTPFileService("39.100.60.58", "vlink", "gh123456!");
        basePathMap.put(IFileStrategy.IMG_TYPE, "/home/vlink");
        sftpFileService.setBasePathMap(basePathMap);
//        "39.100.60.58", 22, "vlink", "gh123456!"
        File file = new File("/Users/joker/Desktop/a.jpg");
        byte[] bytes = FileUtils.readFileToByteArray(file);
        MultipartFile file1 = new MockMultipartFile("asd", bytes);

        String fileName = UUID.randomUUID().toString() + ".jpg";
//        String upload = sftpFileService.upload(file1, "files", fileName, IFileStrategy.IMG_TYPE);
//        System.out.println(upload);
    }

    @Test
    public void delete()
    {
    }
}