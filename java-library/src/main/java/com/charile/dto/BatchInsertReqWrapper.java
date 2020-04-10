package com.charile.dto;

import lombok.Data;
import org.apache.commons.lang3.StringUtils;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-06 14:28
 */
@Data
public class BatchInsertReqWrapper
{
    private String tableName;
    private List<String> columns;
    private List<Object> values;


    public BatchInsertReq conv2TDBReq()
    {
        BatchInsertReq result = new BatchInsertReq();
        result.setTableName(this.tableName);
        result.setColumns(" ( " + StringUtils.join(this.columns, ",") + " )");
        int onceValuesCount = this.columns.size();
        if (this.values.size() % onceValuesCount != 0) throw new RuntimeException("逻辑错误");
        int times = this.values.size() / onceValuesCount;
        StringBuilder sb = new StringBuilder();
        int index = 0;

        for (int i = 0; i < times - 1; i++)
        {
            sb.append(" ( ");
            index = getIndex(onceValuesCount, sb, index);
            sb.append(" ) ,");
        }
        sb.append("( ");
        getIndex(onceValuesCount, sb, index);
        sb.append(" ) ");

        result.setValues(sb.toString());

        return result;
    }

    private int getIndex(int onceValuesCount, StringBuilder sb, int index)
    {
        for (int j = 0; j < onceValuesCount - 1; j++)
        {
            sb.append(this.values.get(index));
            sb.append(",");
            index++;
        }
        sb.append(this.values.get(index));
        index++;
        return index;
    }

}
