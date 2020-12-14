package com.charlie.crypt.opts.impl;

import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.opts.abs.AbsHashOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:05
 */
public class DefaultHashOpts extends AbsHashOpts
{
    public  DefaultHashOpts()
    {
        super();
    }

    public DefaultHashOpts(EnumHashMethod hashMethod)
    {
        super(hashMethod);
    }
}
