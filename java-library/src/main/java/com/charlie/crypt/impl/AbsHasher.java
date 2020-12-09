package com.charlie.crypt.impl;

import com.charlie.crypt.IHash;
import com.charlie.crypt.IHashChain;
import com.charlie.exception.HashException;
import com.charlie.template.BaseTemplatePatternV3;
import org.apache.commons.collections.map.HashedMap;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:21
 */
public abstract class AbsHasher extends BaseTemplatePatternV3<Serializable> implements IHashChain
{
    private boolean inited;
    protected IHashChain next;

    @Override
    public byte[] hash(Serializable serializable, byte[] originData) throws HashException
    {
        if (this.validIsMine(serializable))
        {
            return this.hash(originData);
        } else if (null != this.next)
        {
            return this.next.hash(serializable, originData);
        } else
        {
            throw new HashException("找不到匹配的hash消费者");
        }
    }

    protected abstract byte[] hash(byte[] originData);

    @Override
    public void initOnce() throws Exception
    {
        if (inited) return;
        this.init();
        this.inited = true;
    }

    protected abstract void init();

    @Override
    public IHashChain getNextHash()
    {
        return this.next;
    }

    @Override
    public void setNextHash(IHashChain hash)
    {
        this.next=hash;
    }
}
