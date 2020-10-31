package com.charlie.dto;

import lombok.Data;
import org.apache.commons.lang3.StringUtils;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-21 11:23
 */
@Data
public class BaseWrapper
{
    protected String tableName;
    protected List<String> columns;


    public CommonDTO conv2Common()
    {
        CommonDTO commonDTO = new CommonDTO();
        CommonDTO result = new CommonDTO();
        result.setTableName(this.tableName);
        result.setColumns(" ( " + StringUtils.join(this.columns, ",") + " )");
        return result;
    }


}
