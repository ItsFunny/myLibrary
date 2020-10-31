package com.charlie.template;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-05-18 13:22
 */
public abstract class BaseTemplatePatternV3<T>
{
    // 校验是否是当前执行者该执行的
    public abstract Boolean validIsMine(T type);
}
