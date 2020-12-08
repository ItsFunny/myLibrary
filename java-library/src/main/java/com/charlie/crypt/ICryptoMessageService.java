package com.charlie.crypt;

import com.charlie.base.IInitOnce;
import com.charlie.exception.DecryptException;
import com.charlie.exception.EncryptException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 08:49
 */
// 加密
public interface ICryptoMessageService extends IInitOnce
{
    CryptoMessage encrypt(CryptoMessageBO cryptoMessageBO)throws EncryptException;

    CryptoMessageBO decrypt(CryptoMessage cryptoMessage)throws DecryptException;
}
