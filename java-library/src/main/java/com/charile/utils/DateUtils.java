/**
 * @author joker
 * @date 创建时间：2018年8月8日 上午9:15:39
 */
package com.charile.utils;

import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;

/**
 * @author joker
 * @date 创建时间：2018年8月8日 上午9:15:39
 */
public class DateUtils
{
    public static final int DAY_OF_MILLSECONDS = 1000 * 3600 * 24;

    public static void main(String[] args)
    {
        Long a = 1594452019403l;
        SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        String format = dateFormat.format(new Date(a));
        System.out.println(format);
        System.out.println(Long.parseLong(format));
        //        //获取当前的年月日
//        SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMdd");
//        Calendar c = Calendar.getInstance();
//        c.add(Calendar.DAY_OF_MONTH, 0);
//        c.add(Calendar.DAY_OF_WEEK, 1);
//        System.out.println(c.getTime());
//        System.out.println(c.getTimeInMillis());
//        long rightNow = Long.parseLong(sdf.format(c.getTime()));
//        System.out.println(rightNow);


        // 获取当前年份的最后一天
//        Date lastDayOrTheYear = getLastDayOrTheYear(2020);
//        System.out.println(lastDayOrTheYear);
        System.out.println(new Date().getTime());
        System.out.println(getCurrentDay());
    }

    public static Long getCurrentDay()
    {
        SimpleDateFormat sdf = new SimpleDateFormat("yyyyMMdd");
        Calendar c = Calendar.getInstance();
        c.add(Calendar.DAY_OF_MONTH, 0);
        long rightNow = Long.parseLong(sdf.format(c.getTime()));
        return rightNow;
    }


    public static int calcDateDays(Date early, Date later)
    {
        long between = later.getTime() - early.getTime();
        return (int) (between / DAY_OF_MILLSECONDS);
    }


    /**
     * 获取某年最后一天日期
     *
     * @param year 年份
     * @return Date
     */
    public static Date getLastDayOrTheYear(int year)
    {
        Calendar calendar = Calendar.getInstance();
        calendar.clear();
        calendar.set(Calendar.YEAR, year);
        calendar.roll(Calendar.DAY_OF_YEAR, -1);
        Date currYearLast = calendar.getTime();

        return currYearLast;
    }


    public static String formatStandardTemplate(Date date)
    {
        if (null == date)
        {
            return null;
        }
        SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        return dateFormat.format(date);
    }
}


