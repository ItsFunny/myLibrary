package com.charlie.blockchain;

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
    NETWORK_PARSE_USER_CONTEXT_ERROR(100,"获取用户context失败"),
    NETWORK_CREATE_PEER_ERROR(101,"创建peer失败"),
    NETWORK_CREATE_CHANNEL_ERROR(102,"创建channel失败"),
    NETWORK_ADD_NODE_FAIL(103,"添加节点失败"),

    NETWORK_MAYBE_READ_FILE_ERROR(104,"或许读取文件失败"),
    NETWORK_JOIN_CHANNEL_ERROR(105,"加入channel失败"),

    ILLEGAL_ARGUMENT_ERROR(1000,"参数不合法"),
    UNKNOWN(2000,"未知错误"),



    ///
    INVOKE_FAILED(3000,"交易失败"),
    INVOKE_TIME_OUT(3001,"超时"),
    QUERY_FAILED(3002,"查询失败"),

    //////////////
    CONCURRENT_ERROR(4001,"并发错误"),

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
