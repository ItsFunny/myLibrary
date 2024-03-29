package com.charlie.crypt;

import com.charlie.base.IInitOnce;
import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.exception.EncryptException;
import io.netty.handler.codec.DecoderException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:47
 */
public interface IAsymmetricCrypto extends IInitOnce
{
    String getPublicKey(EnumBaseType serializable);

    byte[] asymmEncrypt(IAsymmetricOpts asymmetricOpts, byte[] origin) throws EncryptException;
    byte[] asymmDecrypt(IAsymmetricOpts asymmetricCrypto,byte[] encrypt) throws DecoderException;
}
