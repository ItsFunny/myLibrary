package com.charile.utils;

import lombok.extern.java.Log;
import lombok.extern.log4j.Log4j;
import lombok.extern.log4j.Log4j2;
import lombok.extern.slf4j.Slf4j;
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
@Slf4j
public class CommonTest
{
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

}
