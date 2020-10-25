package com.charile.blockchain;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 10:14
 */
public enum  ResultInfo
{
    ILLEGAL_ARGUMENT_ERROR(1000,"参数不合法"),
    UNKNOWN(2000,"未知错误"),



    ///
    INVOKE_FAILED(3000,"交易失败"),
    INVOKE_TIME_OUT(3001,"超时"),
    ;
    private Integer code;
    private String msg;

    public Integer getCode()
    {
        return code;
    }

    public void setCode(Integer code)
    {
        this.code = code;
    }

    public String getMsg()
    {
        return msg;
    }

    public void setMsg(String msg)
    {
        this.msg = msg;
    }

    ResultInfo(Integer code, String msg)
    {
        this.code = code;
        this.msg = msg;
    }}
