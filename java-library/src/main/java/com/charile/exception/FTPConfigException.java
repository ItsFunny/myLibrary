package com.charile.exception;

/**
 * @author Charlie
 * @When
 * @Description 配置错误
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-09 09:59
 */
public class FTPConfigException extends BaseBussException
{

    public FTPConfigException(Integer code, String msg)
    {
        super(code, msg);
    }

    public FTPConfigException(String msg)
    {
        super(msg);
    }

    public FTPConfigException(Exception e)
    {
        super(e);
    }

    public FTPConfigException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
