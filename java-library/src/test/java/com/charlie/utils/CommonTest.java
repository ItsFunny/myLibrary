package com.charlie.utils;

import lombok.extern.slf4j.Slf4j;
import org.junit.Test;

import java.util.ArrayList;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-10 08:43
 */
@Slf4j
public class CommonTest
{
    public static void main(String[] args)
    {
        ArrayList<Integer> arrayList = new ArrayList<>();
        arrayList.remove(1);
    }

    @Test
    public void testEmpty()
    {
        log.debug("1");
        log.info("2");
        log.error("3");
        String storePath = "/sqwe/";
        System.out.println(storePath.charAt(0));
        Character c = '/';
        if (c.equals(storePath.charAt(0)))
        {
            storePath = storePath.substring(1);
        }
        if (c.equals(storePath.charAt(storePath.length() - 1)))
        {
            storePath = storePath.substring(0, storePath.length() - 1);
        }
        System.out.println(storePath);
    }

    @Test
    public void test11()
    {
        System.out.println(1<<0);
        System.out.println((2 & 8194) >= 1);
    }


}
