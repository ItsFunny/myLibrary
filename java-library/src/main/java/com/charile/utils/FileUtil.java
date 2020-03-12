package com.charile.utils;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-09 15:26
 */
public class FileUtil
{

    public static String getSuffix(String path)
    {
        return path.substring(path.lastIndexOf(".") + 1);
    }
}
