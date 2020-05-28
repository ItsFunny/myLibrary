package com.charile.bigfile;

import org.junit.Test;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

import static org.junit.Assert.*;

public class AbstractBigFileProcessorTest
{

    @Test
    public void process() throws IOException
    {
        DefaultBigFileProcessor processor = new DefaultBigFileProcessor(new MemoryChunkHandler());
        processor.setStorePath("/Users/joker/Desktop");
        BigFileUploadProcessInfo info = new BigFileUploadProcessInfo();
        info.setChunk(1);
        info.setChunkMd5("MD5");
        info.setChunkSize(9000);
        info.setFileOriginName("originName.flv");
        info.setProcessId("processId");
        FileInputStream fileInputStream = new FileInputStream(new File("/Users/joker/Desktop/未闻花名/公司/jav.flv"));
        info.setInputStream(fileInputStream);
        processor.process(info);
    }
}