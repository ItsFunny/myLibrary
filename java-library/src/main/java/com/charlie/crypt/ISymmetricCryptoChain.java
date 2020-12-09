package com.charlie.crypt;

import com.charlie.base.IInitOnce;

/**
 * @author Charlie
 * @When
 * @Description 对称加密
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:23
 */
public interface ISymmetricCryptoChain extends ISymmetricCrypto
{
    ISymmetricCryptoChain getNext();
    void setNext(ISymmetricCryptoChain next);
}
