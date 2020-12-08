package com.charlie.crypt.factory;

import com.charlie.crypt.IAsymmetricCrypto;
import com.charlie.crypt.impl.DefaultSM2AsymmetricCryptoImpl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 13:57
 */
public class AsymmetricCryptoFactory
{
    public static IAsymmetricCrypto defaultAsymmetricCryptoChain(){
        return new DefaultSM2AsymmetricCryptoImpl();
    }
}
