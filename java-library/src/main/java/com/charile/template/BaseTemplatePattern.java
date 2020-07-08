package com.charile.template;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2019-12-26 22:55
 */
@Data
public abstract class BaseTemplatePattern
{
    // 校验是否是当前执行者该执行的
    protected abstract Boolean validIsMine(Long type);

}
