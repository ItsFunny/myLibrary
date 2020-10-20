package com.charile.utils;

import org.apache.commons.io.FileUtils;
import org.apache.commons.lang3.StringUtils;

import java.io.File;
import java.io.IOException;
import java.util.Properties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:15
 */
public class SystemUtils
{
    public static Properties getOSEnveiroments()
    {
        Properties properties = System.getProperties();
        return properties;
    }

    public static Object getEnviroment(String key)
    {
        return getOSEnveiroments().get(key);
    }

    public static <T> T parseConfigFile(String env, Class<T> tClass)
    {
        String path = (String) getEnviroment(env);
        if (StringUtils.isEmpty(path))
        {
            throw new RuntimeException("环境变量:" + env + ",为空");
        }
        try
        {
            return JSONUtil.json2Object(FileUtils.readFileToString(new File(path), "UTF-8"), tClass);
        } catch (IOException e)
        {
            throw new RuntimeException("读取文件错误", e);
        }
    }

}
