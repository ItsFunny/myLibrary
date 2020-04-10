package com.charile.dto;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-06 14:21
 */
@Data
public class BatchInsertReq
{
    private String tableName;

    private String columns;

    private String values;
}
