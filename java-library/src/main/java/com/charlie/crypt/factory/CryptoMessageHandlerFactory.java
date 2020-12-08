package com.charlie.crypt.factory;

import com.charlie.crypt.ICryptoMessageHandler;
import com.charlie.crypt.impl.DefaultCryptoMessageHandlerImpl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 14:02
 */
public class CryptoMessageHandlerFactory
{
    public static ICryptoMessageHandler defaultCryptoMessageHandler(){
        ICryptoMessageHandler  cryptoMessageCreator= DefaultCryptoMessageHandlerImpl.newInstance();
        return cryptoMessageCreator;
    }

}
