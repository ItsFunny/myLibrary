package com.charlie.crypt.impl;

import com.charlie.crypt.EnumBaseType;
import com.charlie.crypt.EnumSymmetryEncryptionType;
import com.charlie.exception.DecryptException;
import com.charlie.exception.EncryptException;
import com.charlie.utils.AesCBCUtil;
import com.charlie.utils.AesCFBUtil;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 09:45
 */
public class DefaultAESSymmetricCryptoImpl extends AbsSymmetricCryptoImpl
{
    private DefaultAESSymmetricCryptoImpl()
    {

    }

    public static DefaultAESSymmetricCryptoImpl newInstance()
    {
        DefaultAESSymmetricCryptoImpl impl = new DefaultAESSymmetricCryptoImpl();
        return impl;
    }

    @Override
    public byte[] encrypt(EnumSymmetryEncryptionType type, byte[] origin) throws EncryptException
    {
        try
        {
            switch (type)
            {
                case AES_CFB:
                    return AesCFBUtil.encryptBytes(symmKey, origin);
                case AES_CBC:
                    return AesCBCUtil.encrypt(origin, symmKey);
                default:
                    throw new EncryptException("找不到匹配的aes算法");
            }
        } catch (Exception e)
        {
            throw new EncryptException(e);
        }
    }

    @Override
    public byte[] decrypt(EnumSymmetryEncryptionType type, byte[] encrypt) throws DecryptException
    {

        try
        {
            switch (type)
            {
                case AES_CFB:
                    AesCFBUtil.decryptBytes(symmKey, encrypt);
                case AES_CBC:
                    AesCBCUtil.decrypt(encrypt, symmKey);
                default:
                    throw new EncryptException("找不到匹配的aes算法");
            }
        } catch (Exception e)
        {
            throw new DecryptException(e);
        }
    }

    @Override
    protected void init()
    {
        // default
        this.symmKey = "SSV9LoEBaJI6d9xgdqdJdzSYXaUEgZxXWjF0frqegNU=";
    }

    @Override
    public Boolean validIsMine(EnumBaseType type)
    {
        return type.equals(EnumBaseType.ENUM_SYMMETRIC_AES);
    }
}
