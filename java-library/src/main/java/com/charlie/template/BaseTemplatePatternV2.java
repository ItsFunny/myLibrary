package com.charlie.template;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-10 10:35
 */
@Data
public abstract  class BaseTemplatePatternV2
{
    // 校验是否是当前执行者该执行的
    protected abstract Boolean validIsMine(Number type);
}
