package com.charlie.crypt.impl;

import com.charlie.crypt.EnumAsymmetricAlgorithm;
import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.IAsymmetricCryptoChain;
import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.exception.EncryptException;
import com.charlie.template.BaseTemplatePatternV3;
import io.netty.handler.codec.DecoderException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:47
 */
public abstract class AbsAsymmetricCryptoImpl extends BaseTemplatePatternV3<EnumBaseType> implements IAsymmetricCryptoChain
{
    protected String pubKey;
    protected String prvKey;
    protected byte[] pubKeyBytes;
    protected byte[] prvKeyBytes;
    private boolean inited;

    protected IAsymmetricCryptoChain next;

    @Override
    public void initOnce() throws Exception
    {
        if (inited) return;
        this.init();
        inited = true;
    }

    protected abstract void init() throws Exception;

    @Override
    public byte[] asymmEncrypt(IAsymmetricOpts asymmetricOpts, byte[] origin) throws EncryptException
    {
        if (this.validIsMine(asymmetricOpts.getBaseType()))
        {
            return this.encrypt(asymmetricOpts.getDetailAsymmetricAlgorithm(),origin);
        } else if (this.next != null)
        {
            return this.next.asymmEncrypt(asymmetricOpts, origin);
        } else
        {
            throw new EncryptException("找不到匹配的非对称加密者");
        }
    }

    protected abstract byte[] encrypt(EnumAsymmetricAlgorithm detailAlgorithm,byte[] origin) throws EncryptException;

    @Override
    public byte[] asymmDecrypt(IAsymmetricOpts asymmetricOpts, byte[] encrypt) throws DecoderException
    {
        if (this.validIsMine(asymmetricOpts.getBaseType()))
        {
            return this.decrypt(asymmetricOpts.getDetailAsymmetricAlgorithm(),encrypt);
        } else if (this.next != null)
        {
            return this.next.asymmDecrypt(asymmetricOpts, encrypt);
        } else
        {
            throw new EncryptException("找不到匹配的非对称解密者");
        }
    }

    protected abstract byte[] decrypt(EnumAsymmetricAlgorithm detailAlgorithm, byte[] encrypt) throws EncryptException;


    @Override
    public IAsymmetricCryptoChain getNext()
    {
        return this.next;
    }

    @Override
    public void setNext(IAsymmetricCryptoChain next)
    {
        this.next = next;
    }

    @Override
    public String getPublicKey(EnumBaseType serializable)
    {
        if (this.validIsMine(serializable))
        {
            return this.pubKey;
        } else if (this.next != null)
        {
            return this.next.getPublicKey(serializable);
        } else
        {
            throw new EncryptException("找不到匹配的获取公钥者");
        }
    }


}
