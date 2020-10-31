package com.charlie.dto;

import org.junit.Test;

import java.util.*;

public class BatchInsertReqWrapperTest
{

    public static void main(String[] args)
    {
        System.out.println("''");
    }

    @Test
    public void conv2TDBReq()
    {
        BatchInsertReqWrapper wrapper = new BatchInsertReqWrapper();
        wrapper.setColumns(Arrays.asList("ID", "COPYRIGHT_KEY", "PRE_VERSION", "ITEM_CATEGORY", "VERSION",
                "REASON", "TYPE", "CREATE_DATE", "CREATE_USER", "CREATE_USER_ID"));
        wrapper.setTableName("VLINK_COPYRIGHT_HISTORY");
        List<Object> values = new ArrayList<>();
        for (int i = 0; i < 4; i++)
        {
            values.add(i);
            values.add(UUID.randomUUID().toString().substring(8));
            if (i % 2 == 0)
            {
                values.add("");
            } else
            {
                values.add(i);
            }
            values.add(i);
            values.add(i + 1);
            values.add("123" + i);
            values.add(i);
            values.add(new Date());
            values.add("user" + i);
            values.add(i);
        }
        wrapper.setValues(values);
//        BatchInsertReq batchInsertReq = wrapper.conv2TDBReq();
//        System.out.println(batchInsertReq);
    }
}
