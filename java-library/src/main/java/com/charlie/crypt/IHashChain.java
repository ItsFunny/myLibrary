package com.charlie.crypt;

import com.charlie.base.IInitOnce;
import com.charlie.exception.HashException;

import java.io.Serializable;

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
