package com.charile.bigfile;

import com.charile.utils.FileUtil;
import lombok.Data;
import lombok.extern.log4j.Log4j;
import lombok.extern.log4j.Log4j2;
import org.apache.commons.io.filefilter.FalseFileFilter;
import org.apache.commons.lang3.StringUtils;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import sun.java2d.SurfaceDataProxy;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.*;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.stream.Collectors;

import static org.junit.Assert.*;

@Data
@Log4j
public class AbstractBigFileProcessorTest
{

    private Logger logger = LoggerFactory.getLogger(AbstractBigFileProcessorTest.class);
    private MemoryChunkHandler memoryChunkHandler = new MemoryChunkHandler();

    @Test
    public void t()
    {

        logger.debug("1");
        logger.info("1233");
        logger.error("gg");
    }

    @Test
    public void process() throws IOException, InterruptedException
    {
        // 1. 先拆分文件  1.46G
        String originlName = "阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";


        // 碎片存储路径
        String chunkOutPutDirectroy = "/Users/joker/Desktop/split";
        // 合并的文件路径
        String mergeFile = "/Users/joker/Desktop/split";
        // 每个文件分100m
        String originFile = "/Users/joker/Downloads/阳光电影www.ygdy8.com.勇者斗恶龙：你的故事.BD.1080p.日语中英双字.mkv";
        logger.debug("1. 拆分文件");
        FileUtil.splitFiles(originFile, 100, chunkOutPutDirectroy, false);

        // 创建多线程,用于进度条展示
        logger.info("根据文件数量创建多线程,开始分片上传碎片");
        File directoryFile = new File(chunkOutPutDirectroy);
        File[] dirFiles = directoryFile.listFiles();
        List<File> files = new ArrayList<>();
        for (File file : dirFiles)
        {
            if (StringUtils.contains(file.getName(), originlName))
            {
                files.add(file);
            }
        }

        String processId = createProcessId();
        ArrayBlockingQueue<String> chunkMd5List = new ArrayBlockingQueue<>(files.size());
        CountDownLatch countDownLatch = new CountDownLatch(files.size());

        for (int i = 0; i < files.size(); i++)
        {
            final int index = i;
            new Thread(() ->
            {
                DefaultBigFileProcessor processor = new DefaultBigFileProcessor(memoryChunkHandler);
                processor.setStorePath(mergeFile);
                BigFileUploadProcessInfo info = new BigFileUploadProcessInfo();
                info.setChunk(index);
                String chunkMd5 = getChunkMd5();
                chunkMd5List.add(chunkMd5);
                countDownLatch.countDown();
                info.setChunkMd5(chunkMd5);
                info.setChunkSize((int) files.get(index).length());
                info.setFileOriginName(originlName);
                info.setProcessId(processId);
                info.setTotalChunk(files.size());
                FileInputStream fileInputStream = null;
                try
                {
                    fileInputStream = new FileInputStream(files.get(index));
                    info.setInputStream(fileInputStream);
                    countDownLatch.await();
                    processor.process(info);
                } catch (FileNotFoundException e)
                {
                    e.printStackTrace();
                } catch (IOException e)
                {
                    e.printStackTrace();
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            }).start();
        }

        countDownLatch.await();
        log.info("开始进度条查询");

        CountDownLatch stop = new CountDownLatch(chunkMd5List.size());
        List<String> idList = new ArrayList<>(chunkMd5List.size());
        List<String> collect = new ArrayList<>(chunkMd5List);
//        this.queryPercentage(stop, memoryChunkHandler, processId, collect);
        this.querySingle(memoryChunkHandler, processId, collect);

//        stop.await();


        System.out.println("程序结束");
    }

    private void querySingle(MemoryChunkHandler memoryChunkHandler, String processId, List<String> collect)
    {
        Random random = new Random();
        String md5 = collect.get(0);
        float percentage = 0.0f;
        while (percentage < 100)
        {
            ChunkInfo info = memoryChunkHandler.getPercentage(processId, md5);
            if (info != null)
            {
                int uploadedSize = info.getUploadedSize();
                int chunkSize = info.getChunkSize();
                float per = (float) uploadedSize / chunkSize;
                System.out.println(String.format("chunk={%s} 进度为{%s}, {%s}", md5, per * 100 + "%", StringUtils.repeat(">", (int) (per * 100))));
                percentage = per * 100;
                try
                {
                    Thread.sleep(random.nextInt(100));
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            } else
            {
                try
                {
                    Thread.sleep(10);
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            }
        }
        System.out.println("end");
    }

    private String getChunkMd5()
    {return UUID.randomUUID().toString();}

    private String createProcessId()
    {
        return "123456";
    }

    private void queryPercentage(CountDownLatch stop, MemoryChunkHandler memoryChunkHandler, String processId, List<String> chunkMd5List)
    {
        CountDownLatch countDownLatch = new CountDownLatch(chunkMd5List.size());
        for (int i = 0; i < chunkMd5List.size(); i++)
        {
            final int index = i;
            new Thread(() ->
            {
                String md5 = null;
                md5 = chunkMd5List.get(index);
                logger.debug("获取到md5:{},剩余:{}", md5, chunkMd5List.size());
                countDownLatch.countDown();
                try
                {
                    countDownLatch.await();
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }

                logger.debug("开始查询");
                float percentage = 0.0f;
                while (percentage <= 100)
                {
                    ChunkInfo info = memoryChunkHandler.getPercentage(processId, md5);
                    if (info != null)
                    {
                        int uploadedSize = info.getUploadedSize();
                        int chunkSize = info.getChunkSize();
                        float per = (float) uploadedSize / chunkSize;
                        System.out.println(String.format("chunk={%s} 进度为{%s}, {%s}", md5, per * 100 + "%", StringUtils.repeat(">", (int) (per * 100))));
                        percentage = per * 100;
                        try
                        {
                            Thread.sleep(new Random().nextInt(100));
                        } catch (InterruptedException e)
                        {
                            e.printStackTrace();
                        }
                    } else
                    {
                        try
                        {
                            Thread.sleep(10);
                        } catch (InterruptedException e)
                        {
                            e.printStackTrace();
                        }
                    }
                }
                stop.countDown();
            }).start();
        }


    }
}