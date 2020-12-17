package com.charlie.crypt;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:20
 */
public interface IHashChain extends  IHash
{
    IHashChain getNextHash();
    void setNextHash(IHashChain hash);
}
