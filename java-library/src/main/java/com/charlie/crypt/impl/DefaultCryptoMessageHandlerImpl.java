package com.charlie.crypt.impl;

import com.charlie.crypt.CryptoMessage;
import com.charlie.crypt.CryptoMessageBO;
import com.charlie.crypt.EnvelopBO;
import com.charlie.crypt.Envelope;
import com.charlie.crypt.container.CryptoContainer;
import com.charlie.crypt.container.CryptoContainerFactory;
import com.charlie.crypt.factory.EnvelopeCreatorFactory;
import com.charlie.exception.DecryptException;
import com.charlie.utils.Base64Utils;
import com.charlie.utils.DebugUtil;
import com.charlie.utils.JSONUtil;

import java.util.Arrays;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:57
 */
public class DefaultCryptoMessageHandlerImpl extends AbsCryptoMessageHandlerImpl
{
    private DefaultCryptoMessageHandlerImpl(){

    }
    public static DefaultCryptoMessageHandlerImpl newInstance(){
        DefaultCryptoMessageHandlerImpl defaultCryptoMessageCreator = new DefaultCryptoMessageHandlerImpl();

        defaultCryptoMessageCreator.envelopeHandler= EnvelopeCreatorFactory.defaultEnvelopeCreator();
//        defaultCryptoMessageCreator.symmetricCrypto= SymmetricCryptoFactory.defaultSymmetricCryptoChain();

        return defaultCryptoMessageCreator;
    }

    @Override
    protected CryptoMessage doEncrypt(CryptoMessageBO req)
    {
        CryptoContainer cryptoContainer = CryptoContainerFactory.getInstance().getCryptoContainer();
        CryptoMessage result = new CryptoMessage();

        byte[] originData = req.getEnvelope().getOriginData();
        logger.warn("1. 获取到原始的数据,base64为:" + Base64Utils.encode(originData));

        byte[] hashBeforeEncrypt = cryptoContainer.hash(req.getHashOpts(), originData);
        logger.warn("2. hash方法进行加密,hash函数为:{},结果为:{}", req.getHashOpts(), Base64Utils.encode(hashBeforeEncrypt));

        byte[] symmData = cryptoContainer.symmEncrypt(req.getSymmetricOpts(), originData);
        logger.warn("3. 对称加密,对originData进行对称加密,对称方法为:{},返回的数据为:{}", req.getSymmetricOpts(), Base64Utils.encode(symmData));

        req.getEnvelope().setOriginData(symmData);
        Envelope envelope = envelopeHandler.create(req.getEnvelope());
        logger.warn("4. 非对称加密,对对称加密后的数据进行非对称加密,非对称方法为:{},返回的数据为:{}", req.getEnvelope().getAsymmetricOpts(), JSONUtil.toFormattedJson(envelope));

        result.setPlatformId(req.getPlatformId());
        result.setHashBeforeEncrypt(hashBeforeEncrypt);
        result.setSymmetricOpts(req.getSymmetricOpts());
        result.setHashOpts(req.getHashOpts());
        result.setEnvelopInfo(envelope);
        DebugUtil.infoPrint("加密message为:", result);
        return result;
    }

    @Override
    protected CryptoMessageBO doDecrypt(CryptoMessage cryptoMessage)
    {
        CryptoContainer cryptoContainer = CryptoContainerFactory.getInstance().getCryptoContainer();
        Envelope envelopInfo = cryptoMessage.getEnvelopInfo();

        EnvelopBO envelopBO = this.envelopeHandler.decrypt(envelopInfo);

        byte[] originData = cryptoContainer.symmDecrypt(cryptoMessage.getSymmetricOpts(), envelopBO.getOriginData());
        logger.warn(" 对称密钥解内部数据,最初的数据为:{}",new String(originData));

        byte[] hashAfterDecrypt = cryptoContainer.hash(cryptoMessage.getHashOpts(), originData);
        byte[] hashBeforeEncrypt = cryptoMessage.getHashBeforeEncrypt();
        logger.warn(" hash算法校验数据是否一致,校验后的hash码为:{},之前的hash码为:{},是否一致:{}",Base64Utils.encode(hashAfterDecrypt),Base64Utils.encode(hashBeforeEncrypt), Arrays.equals(hashAfterDecrypt, hashBeforeEncrypt));
        if (!Arrays.equals(hashAfterDecrypt,hashBeforeEncrypt))
        {
            throw new DecryptException("解密失败,hash不一致");
        }

        envelopBO.setOriginData(originData);

        CryptoMessageBO result=new CryptoMessageBO();
        result.setEnvelope(envelopBO);
        result.setSymmetricOpts(cryptoMessage.getSymmetricOpts());
        result.setHashOpts(cryptoMessage.getHashOpts());
        result.setPlatformId(cryptoMessage.getPlatformId());

        return result;
    }

    @Override
    protected void init() throws Exception
    {
        this.envelopeHandler.initOnce();
    }
}
