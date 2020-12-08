package com.charlie.crypt;

import com.charlie.base.IInitOnce;

import java.io.Serializable;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 08:54
 */
public interface IEnvelopeHandler extends IInitOnce
{
    // FIXME 或许可以是2个参数,第一个参数是type
    Envelope create(EnvelopBO req);
    EnvelopBO decrypt(Envelope envelope);
}
