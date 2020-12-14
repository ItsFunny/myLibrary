package com.charlie.crypt.opts.abs;

import com.charlie.crypt.EnumAsymmetricAlgorithm;
import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.opts.IAsymmetricOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:17
 */
public abstract class AbsAsymmetricOpts implements IAsymmetricOpts
{
    private EnumAsymmetricAlgorithm enumAsymmetricAlgorithm;
    private EnumBaseType baseType;

    public AbsAsymmetricOpts()
    {

    }
    @Override
    public EnumAsymmetricAlgorithm getDetailAsymmetricAlgorithm()
    {
        return this.enumAsymmetricAlgorithm;
    }

    @Override
    public EnumBaseType getBaseType()
    {
        return baseType;
    }

    public AbsAsymmetricOpts( EnumBaseType baseType,EnumAsymmetricAlgorithm enumAsymmetricAlgorithm)
    {
        this.enumAsymmetricAlgorithm = enumAsymmetricAlgorithm;
        this.baseType = baseType;
    }

    public EnumAsymmetricAlgorithm getEnumAsymmetricAlgorithm()
    {
        return enumAsymmetricAlgorithm;
    }

    public void setEnumAsymmetricAlgorithm(EnumAsymmetricAlgorithm enumAsymmetricAlgorithm)
    {
        this.enumAsymmetricAlgorithm = enumAsymmetricAlgorithm;
    }


    public void setBaseType(EnumBaseType baseType)
    {
        this.baseType = baseType;
    }
}
