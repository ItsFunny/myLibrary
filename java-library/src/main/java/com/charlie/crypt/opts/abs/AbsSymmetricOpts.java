package com.charlie.crypt.opts.abs;

import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.EnumSymmetryEncryptionType;
import com.charlie.crypt.opts.ISymmetricOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:07
 */
public abstract  class AbsSymmetricOpts implements ISymmetricOpts
{
    protected EnumSymmetryEncryptionType encryptionType;
    public AbsSymmetricOpts(EnumSymmetryEncryptionType enumSymmetryEncryptionType){
        this.encryptionType=enumSymmetryEncryptionType;
    }
    @Override
    public EnumBaseType getBaseType()
    {
        return EnumBaseType.ENUM_SYMMETRIC_SHA;
    }

    @Override
    public EnumSymmetryEncryptionType getDetailSymmeticOpts()
    {
        return this.encryptionType;
    }
}
