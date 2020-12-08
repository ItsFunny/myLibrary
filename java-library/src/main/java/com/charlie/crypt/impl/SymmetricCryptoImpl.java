package com.charlie.crypt.impl;

import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.ISymmetricCrypto;
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
public abstract class SymmetricCryptoImpl extends BaseTemplatePatternV3<Serializable> implements ISymmetricCrypto
{
    private boolean inited;
    protected String symmKey;

    protected  SymmetricCryptoImpl next;

    @Override
    public byte[] encrypt(Serializable serializable, byte[] origin) throws EncryptException
    {
        if (this.validIsMine(serializable))
        {
            return this.encrypt(origin);
        }else if (null!=this.next)
        {
            return this.next.encrypt(serializable,origin);
        }else{
            throw new EncryptException("找不到匹配的加密者");
        }
    }
    protected  abstract  byte[] encrypt(byte[] originData)throws EncryptException;

    @Override
    public byte[] decrypt(Serializable serializable, byte[] encrypt) throws DecoderException
    {
        if (this.validIsMine(serializable))
        {
            return this.decrypt(encrypt);
        }else if (null!=this.next)
        {
            return this.next.decrypt(serializable,encrypt);
        }else{
            throw new DecoderException("找不到匹配的解密者");
        }
    }

    protected  abstract  byte[] decrypt(byte[] encryptData)throws DecryptException;


    @Override
    public void initOnce() throws Exception
    {
        if (inited)return;
        this.init();
        this.inited=true;
    }

    protected abstract void init();

}
