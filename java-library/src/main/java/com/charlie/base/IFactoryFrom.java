package com.charlie.base;


import com.charlie.exception.DeserializeException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-10 13:14
 */
public interface IFactoryFrom<F,M>
{
    M  from(F f) throws DeserializeException;
}
