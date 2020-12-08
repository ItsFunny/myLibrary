package com.charlie.crypt.impl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 11:03
 */
public class DefaultEnvelopCreatorImpl extends  AbsEnvelopeCreatorImpl
{
    private DefaultEnvelopCreatorImpl(){

    }

    public static  DefaultEnvelopCreatorImpl newInstance(){
        DefaultEnvelopCreatorImpl impl=new DefaultEnvelopCreatorImpl();
        impl.asymmetricCrypto=DefaultSM2AsymmetricCryptoImpl.newInstance();
        return impl;
    }

    @Override
    protected void init()throws Exception
    {
        this.asymmetricCrypto.initOnce();
    }
}
