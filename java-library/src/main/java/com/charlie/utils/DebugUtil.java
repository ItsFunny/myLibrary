package com.charlie.utils;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;

import javax.servlet.ServletRequest;
import javax.servlet.http.HttpServletRequest;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-04-08 10:38
 */
@Slf4j
public class DebugUtil
{

    public static String showHttpParameters(HttpServletRequest request)
    {
        Map<String, String[]> parameterMap = request.getParameterMap();
        StringBuilder sb = new StringBuilder();
        parameterMap.forEach((k, v) ->
                sb.append("    name=[" + k + ":" + StringUtils.join(v, ",") + "   "));
        return sb.toString();
    }

    public static String showHttpBodyString(ServletRequest request)
    {
        StringBuilder sb = new StringBuilder();

        InputStream inputStream = null;
        BufferedReader reader = null;
        try
        {
            inputStream = request.getInputStream();
            reader = new BufferedReader(new InputStreamReader(inputStream, Charset.forName("UTF-8")));
            String line = "";
            while ((line = reader.readLine()) != null)
            {
                sb.append(line);
            }


        } catch (IOException e)
        {
            log.error("读取流错误:" + e.getMessage());
        } finally
        {
            if (inputStream != null)
            {
                try
                {
                    inputStream.close();
                } catch (IOException e)
                {
                    log.error("关闭输入流错误:" + e.getMessage());
                }
            }
            if (reader != null)
            {
                try
                {
                    reader.close();
                } catch (IOException e)
                {
                    log.error("关闭reader失败:", e.getMessage());
                }
            }
        }
        return sb.toString();
    }
}
