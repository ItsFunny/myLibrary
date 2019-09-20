package com.charile.cache;


/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-01 07:07
 */
public interface ObjectCreateStrategy<T>
{
    T create(Object key);
}
