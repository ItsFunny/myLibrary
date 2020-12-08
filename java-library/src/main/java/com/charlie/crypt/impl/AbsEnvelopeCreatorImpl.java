package com.charlie.crypt.impl;

import com.charlie.crypt.EnvelopBO;
import com.charlie.crypt.Envelope;
import com.charlie.crypt.IAsymmetricCrypto;
import com.charlie.crypt.IEnvelopeHandler;
import com.charlie.utils.Base64Utils;
import com.charlie.utils.JSONUtil;
import org.bouncycastle.cert.ocsp.Req;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:04
 */
public abstract  class AbsEnvelopeCreatorImpl implements IEnvelopeHandler
{
    private Logger logger= LoggerFactory.getLogger(AbsEnvelopeCreatorImpl.class);
    protected  boolean inited;
    protected IAsymmetricCrypto asymmetricCrypto;


    @Override
    public Envelope create(EnvelopBO req)
    {
        Envelope envelope=new Envelope();
        envelope.setExtension(req.getExtension());
        envelope.setEnvelopeIdentifier(req.getEnvelopeIdentifier());
        envelope.setEncryptPublicKey(asymmetricCrypto.getPublicKey(req.getCertAlgorithm()));
        envelope.setEncryptMethod(req.getCertAlgorithm());
        envelope.setDescription(req.getDescription());
        envelope.setEnvelopeData(asymmetricCrypto.encrypt(req.getCertAlgorithm(),req.getOriginData()));
        return envelope;
    }

    @Override
    public EnvelopBO decrypt(Envelope envelope)
    {
        logger.warn("1. 非对称密钥解密,非对称方法为:{},信封数据为:{}",envelope.getEncryptMethod(), JSONUtil.toFormattedJson(envelope));

        byte[] originData = this.asymmetricCrypto.decrypt(envelope.getEncryptMethod(), envelope.getEnvelopeData());
        logger.warn("2. 非对称解密获取得到的对称加密后的数据为:{}", Base64Utils.encode(originData));

        EnvelopBO result=new EnvelopBO();
        result.setOriginData(originData);
        result.setExtension(envelope.getExtension());
        result.setDescription(envelope.getDescription());
        result.setCertAlgorithm(envelope.getEncryptMethod());
        result.setEnvelopeIdentifier(envelope.getEnvelopeIdentifier());
        logger.warn("信封数据为:{}",JSONUtil.toFormattedJson(result));

        return result;
    }

    @Override
    public void initOnce() throws Exception
    {
        if (inited)return;
        this.init();
        this.inited=true;
    }

    protected abstract void init() throws Exception;
}
