package com.charile.blockchain.exception;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-23 23:28
 */
public class InvokeException extends  BaseFabricException
{

    public InvokeException(Integer errorCode, String msg)
    {
        super(errorCode, msg);
    }
    public InvokeException(String message, Integer errorCode, String msg)
    {
        super(message, errorCode, msg);
    }

    public InvokeException(String message, Throwable cause, Integer errorCode, String msg)
    {
        super(message, cause, errorCode, msg);
    }

    public InvokeException(Throwable cause, Integer errorCode, String msg)
    {
        super(cause, errorCode, msg);
    }

    public InvokeException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace, Integer errorCode, String msg)
    {
        super(message, cause, enableSuppression, writableStackTrace, errorCode, msg);
    }
}
