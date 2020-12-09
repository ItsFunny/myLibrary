package com.charlie.crypt.impl;

import com.charlie.crypt.*;
import com.charlie.crypt.opts.ISymmetricOpts;
import com.charlie.exception.DecryptException;
import com.charlie.exception.EncryptException;
import com.charlie.template.BaseTemplatePatternV3;
import com.charlie.utils.AesCFBUtil;
import io.netty.handler.codec.DecoderException;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description CFB模式
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:25
 */
public abstract class AbsSymmetricCryptoImpl extends BaseTemplatePatternV3<EnumBaseType> implements ISymmetricCryptoChain
{
    private boolean inited;
    protected String symmKey;

    protected  ISymmetricCryptoChain next;

    @Override
    public byte[] symmEncrypt(ISymmetricOpts symmetricOpts, byte[] origin) throws EncryptException
    {
        if (this.validIsMine(symmetricOpts.getBaseType()))
        {
            return this.encrypt(symmetricOpts.getDetailSymmeticOpts(),origin);
        }else if (null!=this.next)
        {
            return this.next.symmEncrypt(symmetricOpts,origin);
        }else{
            throw new EncryptException("找不到匹配的加密者");
        }
    }
    protected  abstract  byte[] encrypt(EnumSymmetryEncryptionType detailType, byte[] originData)throws EncryptException;

    @Override
    public byte[] symmDecrypt(ISymmetricOpts symmetricOpts, byte[] encrypt) throws DecoderException
    {
        if (this.validIsMine(symmetricOpts.getBaseType()))
        {
            return this.decrypt(symmetricOpts.getDetailSymmeticOpts(),encrypt);
        }else if (null!=this.next)
        {
            return this.next.symmDecrypt(symmetricOpts,encrypt);
        }else{
            throw new DecoderException("找不到匹配的解密者");
        }
    }

    protected  abstract  byte[] decrypt(EnumSymmetryEncryptionType detailType,byte[] encryptData)throws DecryptException;

    @Override
    public ISymmetricCryptoChain getNext()
    {
        return this.next;
    }

    @Override
    public void setNext(ISymmetricCryptoChain next)
    {
        this.next=next;
    }

    @Override
    public void initOnce() throws Exception
    {
        if (inited)return;
        this.init();
        this.inited=true;
    }

    protected abstract void init();

}
