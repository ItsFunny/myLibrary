package com.charlie.framework;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.BeanPostProcessor;

/**
 * @author Charlie
 * @When
 * @Description 用于aop 注入自身
 * @Detail
 * @Attention: 需要以bean的形式写入的ioc中
 * @Date 创建时间：2020-01-14 16:47
 */
public class BeanSelefAwarePostProcessor implements BeanPostProcessor
{
    @Override
    public Object postProcessBeforeInitialization(Object bean, String beanName) throws BeansException
    {
        return BeanPostProcessor.super.postProcessBeforeInitialization(bean, beanName);
    }


    @Override
    public Object postProcessAfterInitialization(Object bean, String beanName) throws BeansException
    {
        if (bean instanceof IBeanSelfAware)
        {
            IBeanSelfAware beanSelfAware = (IBeanSelfAware) bean;
            beanSelfAware.setSelf(beanSelfAware);
            return beanSelfAware;
        }
        return BeanPostProcessor.super.postProcessAfterInitialization(bean, beanName);
    }
}
