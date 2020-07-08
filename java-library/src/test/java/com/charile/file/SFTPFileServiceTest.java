package com.charile.file;

import org.junit.Test;

import java.io.IOException;

public class SFTPFileServiceTest
{

    @Test
    public void upload() throws IOException
    {
//        Map<String, String> basePathMap = new HashMap<>();
//        basePathMap.put(IFileStrategy.IMG_TYPE, "/home/vlink");
//
//        Map<String, String> visitPrefixMap = new HashMap<>();
//        visitPrefixMap.put(IFileStrategy.IMG_TYPE, "imgs");
//
//        SFTPFileService sftpFileService = new SFTPFileService("118.178.85.47", "zengz", "gh123456!");
//        sftpFileService.setVisitPrefixMap(visitPrefixMap);
//        sftpFileService.setBasePathMap(basePathMap);
////        "39.100.60.58", 22, "vlink", "gh123456!"
//        File file = new File("/Users/joker/Desktop/a.jpg");
//        byte[] bytes = FileUtils.readFileToByteArray(file);
//        MultipartFile file1 = new MockMultipartFile("asd", bytes);
//
//        String fileName = UUID.randomUUID().toString() + ".jpg";
//        UploadResponse files = sftpFileService.upload(file1, "files", fileName, IFileStrategy.IMG_TYPE);
//        System.out.println(files);
//        System.out.println(upload);
    }

    @Test
    public void delete()
    {
    }
}