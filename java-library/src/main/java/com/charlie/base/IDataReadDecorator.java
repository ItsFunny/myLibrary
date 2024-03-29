package com.charlie.base;

import com.charlie.constants.PemConstant;
import com.charlie.utils.Base64Utils;
import com.charlie.utils.PemUtils;

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
