package com.charlie.factory;

/**
 * @author Charlie
 * @When
 * @Description 只用于工厂类,将实体对象转化为XXX
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-10 13:05
 */
public interface IFactortyTO<F,T>
{
    T to(F f);
}
