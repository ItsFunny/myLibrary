package com.charlie.crypt;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:58
 */
public interface IAsymmetricCryptoChain extends  IAsymmetricCrypto
{
    IAsymmetricCryptoChain getNext();
    void setNext(IAsymmetricCryptoChain next);
}
