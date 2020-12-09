package com.charlie.crypt.opts.impl;

import com.charlie.crypt.EnumSymmetryEncryptionType;
import com.charlie.crypt.opts.abs.AbsSymmetricOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:12
 */
public class DefaultSymmetricOpts extends AbsSymmetricOpts
{
    public DefaultSymmetricOpts(EnumSymmetryEncryptionType enumSymmetryEncryptionType)
    {
        super(enumSymmetryEncryptionType);
    }
}
