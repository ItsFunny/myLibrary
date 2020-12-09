package com.charlie.crypt;

import com.charlie.base.IInitOnce;
import com.charlie.crypt.opts.ISymmetricOpts;
import com.charlie.exception.EncryptException;
import io.netty.handler.codec.DecoderException;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description 对称加密
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:23
 */
public interface ISymmetricCrypto  extends IInitOnce
{
    byte[] symmEncrypt(ISymmetricOpts symmetricOpts, byte[] origin) throws EncryptException;
    byte[] symmDecrypt(ISymmetricOpts symmetricOpts,byte[] encrypt) throws DecoderException;
}
