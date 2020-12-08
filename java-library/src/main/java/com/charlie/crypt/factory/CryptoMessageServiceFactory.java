package com.charlie.crypt.factory;

import com.charlie.crypt.ICryptoMessageService;
import com.charlie.crypt.impl.DefaultCryptoMessageServiceImpl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 14:27
 */
public class CryptoMessageServiceFactory
{
    public static ICryptoMessageService defaultCryptoMessageService(){
        ICryptoMessageService cryptoMessageService= DefaultCryptoMessageServiceImpl.newInstance();
        return cryptoMessageService;
    }
}
