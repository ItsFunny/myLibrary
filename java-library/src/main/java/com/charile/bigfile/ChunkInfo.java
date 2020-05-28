package com.charile.bigfile;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:44
 */
@Data
public class ChunkInfo
{
    // 状态
    private byte status;
    // 该块的大小
    private int chunkSize;
    // 已上传的字节
    private int uploadedSize;

    // 该chunk的md5
    private String chunkMd5;

}
