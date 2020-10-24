package com.charile.base;

import com.charile.constants.PemConstant;
import com.charile.utils.Base64Utils;
import com.charile.utils.PemUtils;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-05 17:33
 */
public interface IDataReadDecorator<T> extends IDataDecorator<byte[]>
{

    IDataReadDecorator<byte[]> PUBLICKEY_READER = (data) -> PemUtils.replace(PemConstant.PUBLICKEY, data);

    IDataReadDecorator<byte[]> BASE64_READER = (data) -> Base64Utils.decode(new String(data));
}
