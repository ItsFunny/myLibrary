package com.charlie.crypt.opts.abs;

import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.opts.IHashOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:00
 */
public abstract  class AbsHashOpts implements IHashOpts
{
    protected  EnumHashMethod hashMethod;

    public AbsHashOpts(EnumHashMethod hashMethod)
    {
        this.hashMethod = hashMethod;
    }

    @Override
    public EnumHashMethod getDetailHashOpts()
    {
        return hashMethod;
    }

    @Override
    public EnumBaseType getBaseType()
    {
        return  EnumBaseType.ENUM_HASH_MD5;
    }
}
