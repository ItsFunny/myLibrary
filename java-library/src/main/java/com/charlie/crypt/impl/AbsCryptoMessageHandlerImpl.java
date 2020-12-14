package com.charlie.crypt.impl;

import com.charlie.crypt.*;
import com.charlie.exception.DecryptException;
import com.charlie.utils.Base64Utils;
import com.charlie.utils.DebugUtil;
import com.charlie.utils.JSONUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:17
 */
public abstract class AbsCryptoMessageHandlerImpl implements ICryptoMessageHandler
{
    protected  final Logger logger = LoggerFactory.getLogger(AbsCryptoMessageHandlerImpl.class);
    private boolean inited;

//    protected IHash hash;
    protected IEnvelopeHandler  envelopeHandler;
//    protected ISymmetricCrypto symmetricCrypto;

    @Override
    public CryptoMessage encrypt(CryptoMessageBO req)
    {
        return this.doEncrypt(req);
    }

    protected abstract CryptoMessage doEncrypt(CryptoMessageBO messageBO);

    @Override
    public CryptoMessageBO decrypt(CryptoMessage cryptoMessage)throws DecryptException
    {
       return this.doDecrypt(cryptoMessage);
    }

    protected abstract CryptoMessageBO doDecrypt(CryptoMessage cryptoMessage);

    @Override
    public void initOnce() throws Exception
    {
        if (inited) return;
        this.init();
        this.inited = true;
    }

    protected abstract void init() throws Exception;
}
