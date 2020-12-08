package com.charlie.crypt.impl;

import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.EnumSymmetryEncryptionType;
import com.charlie.exception.DecryptException;
import com.charlie.exception.EncryptException;
import com.charlie.utils.AesCFBUtil;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:45
 */
public class DefaultAESCFBSymmetricCryptoImpl extends SymmetricCryptoImpl
{
    private DefaultAESCFBSymmetricCryptoImpl(){

    }
    public static DefaultAESCFBSymmetricCryptoImpl newInstance(){
        DefaultAESCFBSymmetricCryptoImpl impl=new DefaultAESCFBSymmetricCryptoImpl();
        return impl;
    }

    @Override
    public byte[] encrypt(byte[] origin) throws EncryptException
    {
        try
        {
            return AesCFBUtil.encryptBytes(symmKey, origin);
        } catch (Exception e)
        {
            throw new EncryptException(e);
        }
    }

    @Override
    public byte[] decrypt(byte[] encrypt) throws DecryptException
    {
        try
        {
            return AesCFBUtil.decryptBytes(symmKey, encrypt);
        } catch (Exception e)
        {
            throw new DecryptException(e);
        }
    }

    @Override
    protected void init()
    {
        // default
        this.symmKey="SSV9LoEBaJI6d9xgdqdJdzSYXaUEgZxXWjF0frqegNU=";
    }

    @Override
    public Boolean validIsMine(Serializable type)
    {
        return type.equals(EnumSymmetryEncryptionType.AES_CFB);
    }
}
