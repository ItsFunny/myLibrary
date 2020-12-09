package com.charlie.crypt.container;

import com.charlie.base.AbstractInitOnce;
import com.charlie.crypt.IAsymmetricCrypto;
import com.charlie.crypt.IHash;
import com.charlie.crypt.ISymmetricCrypto;
import com.charlie.exception.ConfigException;
import com.charlie.exception.DecryptException;
import com.charlie.exception.EncryptException;
import com.charlie.exception.HashException;
import io.netty.handler.codec.DecoderException;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description 包含了所有的hash, symm和asymm方法
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:23
 */
public class CryptoContainer extends AbstractInitOnce
{
    private IHash hasher;
    private ISymmetricCrypto symmetricCrypto;
    private IAsymmetricCrypto asymmetricCrypto;

    CryptoContainer(){

    }

    public byte[] hash(Serializable serializable, byte[] originData) throws HashException
    {
        return hasher.hash(serializable, originData);
    }

    public byte[] asymmEncrypt(Serializable serializable, byte[] originData) throws EncryptException
    {
        return this.asymmetricCrypto.encrypt(serializable, originData);
    }

    public byte[] asymmDecrypt(Serializable serializable, byte[] encrypData) throws DecryptException
    {
        return this.asymmetricCrypto.decrypt(serializable, encrypData);
    }

    public byte[] symmEncrypt(Serializable serializable, byte[] originData) throws EncryptException
    {
        return this.symmetricCrypto.encrypt(serializable, originData);
    }

    public byte[] symmDecrypt(Serializable serializable, byte[] encrypData) throws DecryptException
    {
        return this.symmetricCrypto.decrypt(serializable, encrypData);
    }

    @Override
    protected void init() throws ConfigException
    {
        try
        {
            if (null!=this.hasher){
                this.hasher.initOnce();
            }
            if (null!=this.symmetricCrypto){
                this.symmetricCrypto.initOnce();
            }
            if (null!=this.asymmetricCrypto){
                this.asymmetricCrypto.initOnce();
            }
        }catch (Exception e){
            throw new ConfigException(e);
        }


    }
}
