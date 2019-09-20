package com.charile.page;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-21 15:48
 */
@Data
public class SingleDBPageRequest
{
    private Integer pageSize;

    private Integer pageNum;

    // 通过单个排序
    private String singleOrderBy;

    // 是否降序
    private boolean desc;

}
