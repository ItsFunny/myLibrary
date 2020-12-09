package com.charlie.crypt;

import com.charlie.base.IInitOnce;
import com.charlie.crypt.opts.IHashOpts;
import com.charlie.exception.HashException;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:20
 */
public interface IHash extends IInitOnce
{
    byte[] hash(IHashOpts hashOpts, byte[] originData)throws HashException;
}
