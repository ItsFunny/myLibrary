package com.charlie.crypt.impl;

import com.charlie.crypt.CryptoMessage;
import com.charlie.crypt.CryptoMessageBO;
import com.charlie.crypt.ICryptoMessageHandler;
import com.charlie.crypt.ICryptoMessageService;
import com.charlie.crypt.factory.CryptoMessageHandlerFactory;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 14:21
 */
public class DefaultCryptoMessageServiceImpl  implements ICryptoMessageService
{
    private ICryptoMessageHandler cryptoMessageHandler;
    private boolean inited;

    private DefaultCryptoMessageServiceImpl(){

    }
    public static DefaultCryptoMessageServiceImpl newInstance(){
        DefaultCryptoMessageServiceImpl defaultInstance=new DefaultCryptoMessageServiceImpl();
        defaultInstance.cryptoMessageHandler= CryptoMessageHandlerFactory.defaultCryptoMessageHandler();
        return defaultInstance;
    }

    @Override
    public CryptoMessage encrypt(CryptoMessageBO cryptoMessageBO)
    {
        return this.cryptoMessageHandler.encrypt(cryptoMessageBO);
    }

    @Override
    public CryptoMessageBO decrypt(CryptoMessage cryptoMessage)
    {
        return this.cryptoMessageHandler.decrypt(cryptoMessage);
    }


    @Override
    public void initOnce() throws Exception
    {
        if(inited)return;
        this.cryptoMessageHandler.initOnce();
        this.inited=true;
    }
}
