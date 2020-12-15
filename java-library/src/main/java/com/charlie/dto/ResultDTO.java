/**
 * @author joker
 * @date 创建时间：2018年5月25日 下午1:32:56
 */
package com.charlie.dto;


import com.charlie.error.ResultInfo;

import java.io.Serializable;


/**
 * @author joker
 * @date 创建时间：2018年5月25日 下午1:32:56
 */
public class ResultDTO<T> implements Serializable
{
    /**
     * @author joker
     * @date 创建时间：2018年8月15日 下午9:05:02
     */
    private static final long serialVersionUID = -2218913589087908206L;
    private T data;
    private String msg;
    private Integer code;


    public String getMsg()
    {
        return msg;
    }

    public void setMsg(String msg)
    {
        this.msg = msg;
    }

    public T getData()
    {
        return data;
    }

    public void setData(T data)
    {
        this.data = data;
    }

    public Integer getCode()
    {
        return code;
    }

    public void setCode(Integer code)
    {
        this.code = code;
    }

    public static <T> ResultDTO<T> success(T data, String msg)
    {
        ResultDTO<T> ResultDTO = new ResultDTO<>();
        ResultDTO.setCode(1);
        ResultDTO.setData(data);
        ResultDTO.setMsg(msg);
        return ResultDTO;
    }

    public static <T> ResultDTO<T> success(T data)
    {
        return success(data, "success");
    }


    public static <T> ResultDTO<T> result(T data, ResultInfo resultInfo)
    {
        ResultDTO<T> ResultDTO = new ResultDTO();
        ResultDTO.setData(data);
        ResultDTO.setCode(resultInfo.getCode());
        ResultDTO.setMsg(resultInfo.getMsg());
        return ResultDTO;
    }

    public static <T> ResultDTO<T> result(T data, ResultInfo resultInfo, String msg)
    {
        resultInfo.setMsg(resultInfo.getMsg() + "," + msg);
        return result(data, resultInfo);
    }


    public static <T> ResultDTO<T> fail(T data, String msg)
    {
        ResultDTO<T> ResultDTO = new ResultDTO();
        ResultDTO.setData(data);
        ResultDTO.setCode(2);
        ResultDTO.setMsg(msg);
        return ResultDTO;
    }

    public static <T> ResultDTO<T> unauth(String msg)
    {
        ResultDTO<T> ResultDTO = new ResultDTO();
        ResultDTO.setCode(ResultInfo.ILLEGAL_ARGUMENT_ERROR.getCode());
        ResultDTO.setMsg(msg);
        return ResultDTO;
    }
}
