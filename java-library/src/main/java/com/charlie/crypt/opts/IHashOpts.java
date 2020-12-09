package com.charlie.crypt.opts;

import com.charlie.crypt.EnumHashMethod;
import com.charlie.crypt.opts.IBaseOpts;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 09:57
 */
public interface IHashOpts extends IBaseOpts
{
    EnumHashMethod getDetailHashOpts();
}
