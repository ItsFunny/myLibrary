package com.charile.exception;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 11:03
 */
public class ConfigException extends  BaseBussException
{

    public ConfigException(Integer code, String msg)
    {
        super(code, msg);
    }

    public ConfigException(String msg)
    {
        super(msg);
    }

    public ConfigException(Exception e)
    {
        super(e);
    }

    public ConfigException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
