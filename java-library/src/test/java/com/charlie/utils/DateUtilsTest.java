package com.charlie.utils;

import org.junit.Test;

import java.util.Date;

public class DateUtilsTest
{

    @Test
    public void formatStandardTemplate()
    {
        System.out.println(new Date().getTime() / 1000);
//        1589525808
        Long date = 1589525808L * 1000;
        Date d = new Date(date);
        System.out.println(d.toString());
        String s = DateUtils.formatStandardTemplate(new Date(date));
        System.out.println(s);
    }
}