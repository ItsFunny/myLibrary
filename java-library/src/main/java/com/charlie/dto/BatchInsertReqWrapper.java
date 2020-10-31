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
 * @Date 创建时间：2020-04-06 14:28
 */
@Data
public class BatchInsertReqWrapper extends BaseWrapper
{
    protected List<Object> values;


    public BatchInsertReq conv2TInsertDBReq()
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

    protected int getIndex(int onceValuesCount, StringBuilder sb, int index)
    {
        for (int j = 0; j < onceValuesCount - 1; j++)
        {
            index = judgeIsNullOrEmpty(sb, index);
        }
        Object o = this.values.get(index);
        if (null == o || (o instanceof String && ((String) o).isEmpty()))
        {
            sb.append("''");
        } else
        {
            sb.append("'" + o + "'");
        }
        index++;
        return index;
    }

    protected int judgeIsNullOrEmpty(StringBuilder sb, int index)
    {
        Object o = this.values.get(index);
        if (null == o || (o instanceof String && ((String) o).isEmpty()))
        {
            sb.append("''");
            sb.append(",");
        } else
        {
            sb.append("'" + o + "'");
            sb.append(",");
        }
        index++;
        return index;
    }
}
