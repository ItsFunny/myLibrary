package com.charlie.crypt.factory;

import com.charlie.crypt.ISymmetricCrypto;
import com.charlie.crypt.impl.DefaultAESCFBSymmetricCryptoImpl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 11:08
 */
public class SymmetricCryptoFactory
{
    public static ISymmetricCrypto defaultSymmetricCryptoChain(){
        return DefaultAESCFBSymmetricCryptoImpl.newInstance();
    }
}
