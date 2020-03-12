package com.charile.utils;

import org.junit.Test;

import static org.junit.Assert.*;

public class FileUtilTest
{

    @Test
    public void getSuffix()
    {
        String path = "/a/b/c.jpg";
        String suffix = FileUtil.getSuffix(path);
        System.out.println(suffix);
    }
}