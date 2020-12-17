package com.charlie.crypt.impl;

import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.EnumHashMethod;
import com.charlie.utils.MD5Utils;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:24
 */
public class MD5Hasher extends AbsHasher
{
    @Override
    protected byte[] hash(EnumHashMethod hashMethod,byte[] originData)
    {
        return MD5Utils.getMD5(originData).getBytes();
    }

    @Override
    protected void init()
    {
    }

    @Override
    public Boolean validIsMine(EnumBaseType type)
    {
        return EnumBaseType.ENUM_HASH_MD5.equals(type);
    }
}
