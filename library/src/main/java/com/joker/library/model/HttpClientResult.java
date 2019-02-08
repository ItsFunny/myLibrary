package com.joker.library.model;


import lombok.Data;

import java.io.Serializable;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-05 09:18
 */
@Data
public class HttpClientResult implements Serializable
{
    private static final long serialVersionUID = 2168152194164783950L;

    public static final int SUCCESS = 0;
    public static final int FAIL = -1;

    private int status;
    /**
     * 响应状态码
     */
    private int code;

    /**
     * 响应数据
     */
    private String content;

    public HttpClientResult()
    {
    }

    public static HttpClientResult buildSuccess(int code, String content)
    {
        HttpClientResult result = new HttpClientResult();
        result.code = code;
        result.content = content;
        result.status = SUCCESS;
        return result;
    }

    public static HttpClientResult buildSuccess(int code)
    {
        HttpClientResult result = new HttpClientResult();
        result.code = code;
        result.content = "ok";
        result.status = SUCCESS;
        return result;
    }

    public static HttpClientResult buildSuccess()
    {
        HttpClientResult result = new HttpClientResult();
        result.code = 200;
        result.content = "ok";
        result.status = SUCCESS;
        return result;
    }

    public HttpClientResult(int code)
    {
        this.code = code;
    }

    public HttpClientResult(String content)
    {
        this.content = content;
    }

    public HttpClientResult(int code, String content)
    {
        this.code = code;
        this.content = content;
    }

    public boolean isSuccess()
    {
        return this.status == SUCCESS;
    }

}
