package com.charlie.crypt.factory;

import com.alibaba.fastjson.JSONObject;
import com.alibaba.fastjson.TypeReference;
import com.alibaba.fastjson.parser.Feature;
import com.charlie.base.IBaseByteFactory;
import com.charlie.crypt.CryptoMessage;
import com.charlie.crypt.CryptoMessageBO;
import com.charlie.crypt.CryptoMessageJSONObject;
import com.charlie.crypt.IHash;
import com.charlie.crypt.opts.IAsymmetricOpts;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.exception.DeserializeException;
import com.charlie.exception.SerializeException;
import com.charlie.utils.JSONUtil;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.checkerframework.checker.units.qual.A;
import sun.util.resources.cldr.br.CalendarData_br_FR;

import com.alibaba.fastjson.TypeReference;

import java.util.List;

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
    private static final CryptoMessageFactory CRYPTO_MESSAGE_FACTORY = new CryptoMessageFactory();

    public static byte[] toBytes(CryptoMessage f) throws SerializeException
    {
        return CRYPTO_MESSAGE_FACTORY.to(f);
    }

    public static CryptoMessage fromBytes(byte[] bytes)
    {
        return CRYPTO_MESSAGE_FACTORY.from(bytes);
    }

    @Override
    public byte[] to(CryptoMessage f) throws SerializeException
    {
//        byte[] bytes = f.to();
//        String symmName = f.getSymmetricOpts().getClass().getName();
//        String hashName = f.getHashOpts().getClass().getName();
//        String asymmName = f.getEnvelopInfo().getAsymmetricOpts().getClass().getName();
//        CryptoMessageJSONObject cryptoMessageJSONObject = new CryptoMessageJSONObject();
//        cryptoMessageJSONObject.setMessageJson(new String(bytes));
//        cryptoMessageJSONObject.setSymmClazzName(symmName);
//        cryptoMessageJSONObject.setAsymmClazzName(asymmName);
//        cryptoMessageJSONObject.setHashClazzName(hashName);
//        return JSONUtil.toFormattedJson(cryptoMessageJSONObject).getBytes();
        return f.to();
    }

    @Override
    public CryptoMessage from(byte[] bytes) throws DeserializeException
    {
        CryptoMessage cryptoMessage = new CryptoMessage();
        return cryptoMessage.from(bytes);
//        String s=new String(bytes);
//        CryptoMessageJSONObject cryptoMessageJSONObject = JSONUtil.json2Obj(new String(bytes), CryptoMessageJSONObject.class);
//        String asymmClazzName = cryptoMessageJSONObject.getAsymmClazzName();
//        String hashClazzName = cryptoMessageJSONObject.getHashClazzName();
//        String messageJson = cryptoMessageJSONObject.getMessageJson();
//        String symmClazzName = cryptoMessageJSONObject.getSymmClazzName();
//        JSONObject jsonObject = JSONObject.parseObject(messageJson);
//        CryptoMessage cryptoMessage = new CryptoMessage();
//
//
//        Class<?> aClass = null;
//        try
//        {
//            aClass = Class.forName(hashClazzName);
//            Object object = jsonObject.getObject("hashOpts", new TypeReference<IHashOpts>(){});
//            System.out.println(object);
//        } catch (ClassNotFoundException e)
//        {
//            e.printStackTrace();
//        }
//
//        return cryptoMessage;
    }
}
