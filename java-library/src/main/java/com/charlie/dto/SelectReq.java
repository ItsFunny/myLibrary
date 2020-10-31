package com.charlie.dto;

import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-21 11:21
 */
@Data
public class SelectReq extends CommonDTO
{
    private String condition;

    @Override
    public String toString()
    {
        return "SELECT " + this.columns + " FROM " + this.tableName + "  " + this.condition;
    }
}
