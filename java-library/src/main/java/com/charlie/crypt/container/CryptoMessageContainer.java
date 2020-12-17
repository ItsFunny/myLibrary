package com.charlie.crypt.container;

import com.charlie.base.IInitOnce;
import com.charlie.crypt.*;
import com.charlie.crypt.factory.CryptoMessageServiceFactory;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 14:15
 */
public class CryptoMessageContainer implements IInitOnce
{
    private boolean inited;
    private ICryptoMessageService cryptoMessageService;

    public CryptoMessageContainer(){ this.cryptoMessageService= CryptoMessageServiceFactory.defaultCryptoMessageService(); }


    public CryptoMessage buildCryptoMessage(CryptoMessageBO cryptoMessageBO)
    {
        return this.cryptoMessageService.encrypt(cryptoMessageBO);
    }

    public CryptoMessageBO decrypt(CryptoMessage cryptoMessage){
        return this.cryptoMessageService.decrypt(cryptoMessage);
    }

    @Override
    public void initOnce() throws Exception
    {
        if (inited)return;
        this.cryptoMessageService.initOnce();
        this.inited=true;
    }
}
