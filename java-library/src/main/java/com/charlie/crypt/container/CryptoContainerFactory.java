package com.charlie.crypt.container;

import com.charlie.crypt.IAsymmetricCrypto;
import com.charlie.crypt.IHash;
import com.charlie.crypt.ISymmetricCrypto;
import com.charlie.crypt.container.CryptoContainer;
import com.charlie.crypt.factory.AsymmetricCryptoFactory;
import com.charlie.crypt.factory.HashFactory;
import com.charlie.crypt.factory.SymmetricCryptoFactory;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:30
 */
public class CryptoContainerFactory
{
    private CryptoContainer cryptoContainer;

    private CryptoContainerFactory(){
    }

    private static class CryptoContainerHolder
    {
        private static CryptoContainerFactory INSTANCE = new CryptoContainerFactory();
    }
    public static CryptoContainerFactory getInstance(){
        return CryptoContainerHolder.INSTANCE;
    }

    public CryptoContainer getCryptoContainer()
    {
        return cryptoContainer;
    }


    public static void defaultInit()throws Exception{
        IHash iHash = HashFactory.defaultHashChain();
        ISymmetricCrypto iSymmetricCrypto = SymmetricCryptoFactory.defaultSymmetricCryptoChain();
        IAsymmetricCrypto iAsymmetricCrypto = AsymmetricCryptoFactory.defaultAsymmetricCryptoChain();
        CryptoContainerFactory.getInstance().cryptoContainer=new CryptoContainer(iHash,iSymmetricCrypto,iAsymmetricCrypto);
        CryptoContainerFactory.getInstance().cryptoContainer.initOnce();
    }
}
