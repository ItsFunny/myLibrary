package com.charlie.crypt.factory;

import com.charlie.crypt.CryptoMessage;
import com.charlie.exception.DeserializeException;
import com.charlie.exception.SerializeException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-10 12:00
 */
public class CryptoMessageFactory implements IBaseByteFactory<CryptoMessage>
{
    @Override
    public byte[] to(CryptoMessage f) throws SerializeException
    {
        return new byte[0];
    }

    @Override
    public CryptoMessage from(byte[] bytes) throws DeserializeException
    {
        return null;
    }
}
