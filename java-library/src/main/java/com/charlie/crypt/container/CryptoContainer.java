package com.charlie.crypt.container;

import com.charlie.base.AbstractInitOnce;
import com.charlie.crypt.*;
import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.crypt.opts.ISymmetricOpts;
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
public class CryptoContainer extends AbstractInitOnce implements IHash, ISymmetricCrypto,IAsymmetricCrypto
{
    private IHash hasher;
    private ISymmetricCrypto symmetricCrypto;
    private IAsymmetricCrypto asymmetricCrypto;

    public CryptoContainer(IHash hasher,ISymmetricCrypto symmetricCrypto,IAsymmetricCrypto asymmetricCrypto){
        this.hasher=hasher;
        this.symmetricCrypto=symmetricCrypto;
        this.asymmetricCrypto=asymmetricCrypto;
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

    @Override
    public String getPublicKey(EnumBaseType serializable)
    {
        return this.asymmetricCrypto.getPublicKey(serializable);
    }

    @Override
    public byte[] asymmEncrypt(IAsymmetricOpts asymmetricOpts, byte[] origin) throws EncryptException
    {
        return this.asymmetricCrypto.asymmEncrypt(asymmetricOpts,origin);
    }

    @Override
    public byte[] asymmDecrypt(IAsymmetricOpts asymmetricCrypto, byte[] encrypt) throws DecoderException
    {
        return this.asymmetricCrypto.asymmDecrypt(asymmetricCrypto,encrypt);
    }

    @Override
    public byte[] hash(IHashOpts hashOpts, byte[] originData) throws HashException
    {
        return this.hasher.hash(hashOpts,originData);
    }

    @Override
    public byte[] symmEncrypt(ISymmetricOpts symmetricOpts, byte[] origin) throws EncryptException
    {
        return this.symmetricCrypto.symmEncrypt(symmetricOpts,origin);
    }

    @Override
    public byte[] symmDecrypt(ISymmetricOpts symmetricOpts, byte[] encrypt) throws DecoderException
    {
        return this.symmetricCrypto.symmDecrypt(symmetricOpts,encrypt);
    }
}
