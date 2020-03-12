package com.charile.exception;

/**
 * @author Charlie
 * @When
 * @Description sftp 数据传输发生失败
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-09 10:01
 */
public class SFTPDataTransferException extends BaseBussException
{

    public SFTPDataTransferException(Integer code, String msg)
    {
        super(code, msg);
    }

    public SFTPDataTransferException(String msg)
    {
        super(msg);
    }

    public SFTPDataTransferException(Exception e)
    {
        super(e);
    }

    public SFTPDataTransferException(String msg, Throwable t)
    {
        super(msg, t);
    }
}
