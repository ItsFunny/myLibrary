package com.charile.bigfile;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:55
 */
public interface BigFileConstants
{
    // 传输失败
    byte CHUNK_FAIL = 1 << 0;
    // 正在传输中
    byte CHUNK_ON_PROCESS = 1 << 1;
    // 成功
    byte CHUNK_SUCCESS = 1 << 2;
}
