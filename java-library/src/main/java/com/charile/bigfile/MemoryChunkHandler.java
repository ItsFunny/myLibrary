package com.charile.bigfile;

import lombok.Data;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.locks.ReentrantLock;

/**
 * @author Charlie
 * @When
 * @Description 业务逻辑原因, 并不需要实时, 因此不需要考虑并发安全, 既updateChunkInfo并不需要lock
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-28 16:04
 */
@Data
public class MemoryChunkHandler implements ChunkHandler
{
    private Logger logger = LoggerFactory.getLogger(MemoryChunkHandler.class);
    // key 为processorId  ,value 的map 为 这个文件下的所有碎片信息  key 为文件的md5码, value 为chunk信息
    private final ConcurrentHashMap<String, ChunkHolder> chunkInfoMap = new ConcurrentHashMap<>();

    private final ConcurrentHashMap<String, AtomicInteger> chunkCountMap = new ConcurrentHashMap<>();

    private ReentrantLock lock;

    public MemoryChunkHandler()
    {
        this.lock = new ReentrantLock();
    }

    @Data
    private class ChunkHolder
    {
        private ConcurrentHashMap<String, ChunkInfo> chunkInfoMap;
    }


    @Override
    public ChunkInfo getChunkInfo(FileChunkDecorator decorator)
    {
        ChunkHolder chunkHolder = chunkInfoMap.get(decorator.getProcessId());
        if (chunkHolder == null) return null;
        // 不用管并发

        return chunkHolder.getChunkInfoMap() == null ? null : chunkHolder.getChunkInfoMap().get(decorator.getChunkMd5());
    }

    @Override
    public void updateChunk(String processId, ChunkInfo chunkInfo)
    {
        ChunkHolder chunkHolder = chunkInfoMap.get(processId);
        ConcurrentHashMap<String, ChunkInfo> cmap = null;
        if (chunkHolder == null)
        {

            try
            {
                this.lock.lock();
                chunkHolder = chunkInfoMap.get(processId);
                if (chunkHolder == null)
                {
                    chunkHolder = new ChunkHolder();
                    cmap = new ConcurrentHashMap<>();
                    chunkHolder.setChunkInfoMap(cmap);
                    chunkInfoMap.put(processId, chunkHolder);
                } else
                {
                    cmap = chunkHolder.getChunkInfoMap();
                    if (cmap == null)
                    {
                        cmap = new ConcurrentHashMap<>();
                    }
                }
            } finally
            {
                this.lock.unlock();
            }
        } else
        {
            cmap = chunkHolder.getChunkInfoMap();
            if (cmap == null)
            {
                try
                {
                    this.lock.lock();
                    cmap = chunkHolder.getChunkInfoMap();
                    if (cmap == null)
                    {
                        cmap = new ConcurrentHashMap<>();
                        chunkHolder.setChunkInfoMap(cmap);
                    }
                } finally
                {
                    this.lock.unlock();
                }
            }
        }
        cmap.put(chunkInfo.getChunkMd5(), chunkInfo);
    }

    @Override
    public boolean increase(String processId, int maxCount)
    {
        AtomicInteger uploadedCount = chunkCountMap.get(processId);

        int currentUploaded = 0;
        if (uploadedCount == null)
        {
            try
            {
                this.lock.lock();
                uploadedCount = chunkCountMap.get(processId);
                if (uploadedCount == null)
                {
                    uploadedCount = new AtomicInteger(1);
                    chunkCountMap.put(processId, uploadedCount);
                } else
                {
                    currentUploaded = uploadedCount.incrementAndGet();
                }

            } finally
            {
                this.lock.unlock();
            }
        } else
        {
            currentUploaded = uploadedCount.incrementAndGet();


        }
        logger.info("线程{}更新碎片数目,已上传碎片数目为:{},剩余:{}块", Thread.currentThread().getName(), currentUploaded, maxCount - currentUploaded);
        if (currentUploaded == maxCount)
        {
            return true;
        }

        return false;
    }

    @Override
    public void clearChunk(String processId)
    {
        this.chunkCountMap.remove(processId);
        this.chunkInfoMap.remove(processId);
    }

    @Override
    public ChunkInfo getPercentage(String processorId, String chunkMd5)
    {
        ChunkHolder chunkHolder = this.chunkInfoMap.get(processorId);
        if (null == chunkHolder) return null;
        ConcurrentHashMap<String, ChunkInfo> chunkInfoMap = chunkHolder.getChunkInfoMap();
        if (chunkInfoMap == null) return null;
        return chunkInfoMap.get(chunkMd5);
    }

}
