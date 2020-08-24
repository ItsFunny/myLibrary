package com.charile.cache;

import com.charile.cache.soft.AbstractCHMSoftReferenceCache;
import com.charile.cache.soft.SoftReferenceInfo;
import com.charile.utils.SFTPUtil;

import java.lang.ref.Reference;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-08-04 14:58
 */
public class SFTHCache extends AbstractCHMSoftReferenceCache<Integer, SFTPUtil.SFTPWrapper>
{

    private ClearStrategy<SFTPUtil.SFTPWrapper> SFTP_CLEAR_STRATEGY = (queue) ->
    {
        SoftReferenceInfo<Integer, SFTPUtil.SFTPWrapper> poll = (SoftReferenceInfo<Integer, SFTPUtil.SFTPWrapper>) queue.poll();
        while (null != poll)
        {
            Integer key = poll.getKey();
            this.dataMap.remove(key);
            SFTPUtil.SFTPWrapper sftpWrapper = poll.get();
            if (null != sftpWrapper) sftpWrapper.logOut();
        }
    };

    public SFTHCache(ObjectCreateStrategy<SFTPUtil.SFTPWrapper> createStrategy)
    {
        super(createStrategy);
        this.clearStrategy = SFTP_CLEAR_STRATEGY;
    }

    public SFTHCache(ClearStrategy<SFTPUtil.SFTPWrapper> clearStrategy, ObjectCreateStrategy<SFTPUtil.SFTPWrapper> createStrategy)
    {
        super(clearStrategy, createStrategy);
    }
}
