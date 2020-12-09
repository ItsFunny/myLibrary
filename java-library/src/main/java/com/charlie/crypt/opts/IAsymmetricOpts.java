package com.charlie.crypt.opts;

import com.charlie.crypt.EnumAsymmetricAlgorithm;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:17
 */
public interface IAsymmetricOpts extends  IBaseOpts
{
    EnumAsymmetricAlgorithm getDetailAsymmetricAlgorithm();
}
