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
public class BatchInsertReq extends  CommonDTO
{
    protected String values;

    @Override
    public String toString()
    {
        return "INSERT INTO " + this.tableName + " " + this.columns + " VALUES " + this.values;
    }
}
