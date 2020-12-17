package com.charlie.crypt.impl;

import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.IHashChain;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.exception.HashException;
import com.charlie.template.BaseTemplatePatternV3;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:21
 */
public abstract class AbsHasher extends BaseTemplatePatternV3<EnumBaseType> implements IHashChain
{
    private boolean inited;
    protected IHashChain next;

    @Override
    public byte[] hash(IHashOpts hashOpts, byte[] originData) throws HashException
    {
        if (this.validIsMine(hashOpts.getBaseType()))
        {
            return this.hash(hashOpts.getDetailHashOpts(),originData);
        } else if (null != this.next)
        {
            return this.next.hash(hashOpts, originData);
        } else
        {
            throw new HashException("找不到匹配的hash消费者");
        }
    }

    protected abstract byte[] hash(EnumHashMethod hashMethod, byte[] originData);

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
