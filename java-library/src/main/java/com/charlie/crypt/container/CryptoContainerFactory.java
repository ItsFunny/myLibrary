package com.charlie.crypt.container;

import com.charlie.crypt.container.CryptoContainer;

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

    private CryptoContainerFactory()
    {
        this.cryptoContainer = new CryptoContainer();
    }

    private static class CryptoContainerHolder
    {
        private static CryptoContainerFactory INSTANCE = new CryptoContainerFactory();
    }

    public static CryptoContainerFactory getInstance()
    {
        return CryptoContainerHolder.INSTANCE;
    }
    public  void initSlow() throws Exception
    {
        cryptoContainer.initOnce();
    }
}
