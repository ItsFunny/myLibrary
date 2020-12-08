package com.charlie.crypt;

import com.charlie.base.IInitOnce;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:47
 */
public interface IAsymmetricCrypto extends IInitOnce,ICryptoService
{
    String getPublicKey(Serializable serializable);
}
