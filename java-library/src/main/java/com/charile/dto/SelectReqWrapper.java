package com.charile.dto;

import lombok.Data;

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
public class SelectReqWrapper extends BaseWrapper
{
    private List<String> condition;
    private List<Object> conditionValues;

    public SelectReq conv2SelectReq()
    {
        SelectReq req = new SelectReq();
        CommonDTO c = this.conv2Common();
        req.setColumns(c.columns);
        req.setTableName(c.tableName);
        if (this.condition == null && this.conditionValues == null)
        {
            return req;
        }

        if (this.condition.size() != this.conditionValues.size())
        {
            throw new RuntimeException("condition 和 conditionValues 配置错误");
        }
        StringBuilder sb = new StringBuilder();
        sb.append(" WHERE 1=1 ");
        final List<Object> conditionValues = this.conditionValues;
        final List<String> conditions = this.condition;
        int size = conditions.size();
        String colunm = null;
        Object value = null;
        for (int i = 0; i < size - 1; i++)
        {
            colunm = conditions.get(i);
            value = conditionValues.get(i);
            buildCondition(sb, colunm, value);
        }

        colunm = conditions.get(size - 1);
        value = conditionValues.get(size - 1);
        buildCondition(sb, colunm, value);
        req.setCondition(sb.toString());

        return req;
    }

    private void buildCondition(StringBuilder sb, String colunm, Object value)
    {
        if (colunm == null || value == null)
        {
            throw new RuntimeException("条件不可为空,colunm=" + colunm + " value=" + value);
        }
        sb.append("AND ").append(colunm).append(value).append(" ");
    }


}
