package com.charile.exception;

import lombok.Data;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @Description: 基础业务异常包装类
 * @Attention:
 * @Date 创建时间：2019-12-25 13:51
 */
@Data
public class BaseBussException extends RuntimeException
{
    private Integer code;


    public BaseBussException(Integer code, String msg)
    {
        super(msg);
        this.code = code;
    }


    public BaseBussException(String msg)
    {
        super(msg);
    }

    public BaseBussException(Exception e)
    {
        super(e);
    }

    public BaseBussException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
