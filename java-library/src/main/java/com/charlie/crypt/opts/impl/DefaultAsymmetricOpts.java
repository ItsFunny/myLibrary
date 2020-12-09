package com.charlie.crypt.opts.impl;

import com.charlie.crypt.EnumAsymmetricAlgorithm;
import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.opts.abs.AbsAsymmetricOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:20
 */
public class DefaultAsymmetricOpts extends AbsAsymmetricOpts
{
    public DefaultAsymmetricOpts(EnumAsymmetricAlgorithm enumAsymmetricAlgorithm, EnumBaseType baseType)
    {
        super(enumAsymmetricAlgorithm, baseType);
    }
}
