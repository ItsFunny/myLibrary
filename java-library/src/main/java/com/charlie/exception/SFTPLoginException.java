package com.charlie.exception;

import com.jcraft.jsch.JSchException;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-09 09:53
 */
public class SFTPLoginException extends  RuntimeException
{

    public SFTPLoginException(JSchException e) {super(e);}
}
