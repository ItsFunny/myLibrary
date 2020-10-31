package com.charlie.bigfile;

/**
 * @author Charlie
 * @When
 * @Description handle chunk的信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 14:44
 */
public interface ChunkHandler
{
    // 获取碎片信息
    ChunkInfo getChunkInfo(FileChunkDecorator decorator);

    void updateChunk(String processId, ChunkInfo chunkInfo);

    // 已经上传的块数++ ,返回值bool 代表是否需要合并
    boolean increase(String processId, int maxCount);

    // clear chunk 删除所有的碎片缓存记录
    void clearChunk(String processId);

    ChunkInfo getPercentage(String processorId, String chunkMd5);


}
