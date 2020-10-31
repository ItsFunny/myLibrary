package com.charlie.bigfile;

import lombok.Data;

import java.io.InputStream;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:25
 */
@Data
public class BigFileUploadProcessInfo implements FileChunkDecorator
{
    // 第几块
    private int chunk;
    // 文件的总块数
    private int totalChunk;
    private int chunkSize;
    // chunk的大小
    private String chunkMd5;
    // 会话id
    private String processId;
    private String fileOriginName;
    private InputStream inputStream;
}
