package com.charlie.crypt;

import com.charlie.exception.EncryptException;
import io.netty.handler.codec.DecoderException;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:13
 */
public interface ICryptoService
{
    byte[] encrypt(Serializable serializable, byte[] origin) throws EncryptException;
    byte[] decrypt(Serializable serializable,byte[] encrypt) throws DecoderException;
}
