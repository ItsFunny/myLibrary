package com.charlie.crypt.impl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 11:03
 */
public class DefaultEnvelopHandlerImpl extends  AbsEnvelopeHandlerImpl
{
    private DefaultEnvelopHandlerImpl(){

    }

    public static  DefaultEnvelopHandlerImpl newInstance(){
        DefaultEnvelopHandlerImpl impl=new DefaultEnvelopHandlerImpl();
        impl.asymmetricCrypto=DefaultSM2AsymmetricCryptoImpl.newInstance();
        return impl;
    }

    @Override
    protected void init()throws Exception
    {
        this.asymmetricCrypto.initOnce();
    }
}
