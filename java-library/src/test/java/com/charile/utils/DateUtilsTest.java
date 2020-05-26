package com.charile.utils;

import org.junit.Test;

import java.util.Calendar;
import java.util.Date;

import static org.junit.Assert.*;

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