package com.charlie.bigfile;

import java.io.InputStream;

/**
 * @author Charlie
 * @When
 * @Description 大文件chunk 的修饰器
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:27
 */
public interface FileChunkDecorator
{
    int getChunk();

    int getTotalChunk();

    String getChunkMd5();

    // 该chunk的大小
    int getChunkSize();


    String getProcessId();

    // 文件原始名称
    String getFileOriginName();


    InputStream getInputStream();
}
