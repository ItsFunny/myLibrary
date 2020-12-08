package com.charlie.crypt.factory;

import com.charlie.crypt.IEnvelopeHandler;
import com.charlie.crypt.impl.DefaultEnvelopCreatorImpl;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 11:03
 */
public class EnvelopeCreatorFactory
{
    public static IEnvelopeHandler defaultEnvelopeCreator(){
        DefaultEnvelopCreatorImpl defaultEnvelopCreator = DefaultEnvelopCreatorImpl.newInstance();
        return defaultEnvelopCreator;
    }

}
