package com.charlie.dto;

import org.junit.Test;

import java.util.Arrays;
import java.util.List;

public class SelectReqWrapperTest
{

    @Test
    public void conv2SelectReq()
    {
        SelectReqWrapper wrapper = new SelectReqWrapper();
        wrapper.setColumns(Arrays.asList("ID", "COPYRIGHT_KEY", "PRE_VERSION", "ITEM_CATEGORY", "VERSION",
                "REASON", "TYPE", "CREATE_DATE", "CREATE_USER", "CREATE_USER_ID"));
        wrapper.setTableName("VLINK_COPYRIGHT_HISTORY");
        List<String> conditon = Arrays.asList("TYPE", "COPYRIGHT_KEY");
        List<Object> cvalues = Arrays.asList("=1", "=3");
        wrapper.setCondition(conditon);
        wrapper.setConditionValues(cvalues);
        SelectReq selectReq = wrapper.conv2SelectReq();

        System.out.println(selectReq);
    }
}