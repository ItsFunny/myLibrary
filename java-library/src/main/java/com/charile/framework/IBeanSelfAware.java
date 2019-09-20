package com.charile.framework;

/**
 * @author Charlie
 * @When
 * @Description 用于aop 不生效的情况,注入自身
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-01-14 16:46
 */
public interface IBeanSelfAware
{
    void setSelf(Object proxy);
}
