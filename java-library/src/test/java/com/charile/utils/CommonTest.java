package com.charile.utils;

import org.junit.Test;

import java.io.File;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-10 08:43
 */
public class CommonTest
{
    @Test
    public void testEmpty()
    {
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

}
