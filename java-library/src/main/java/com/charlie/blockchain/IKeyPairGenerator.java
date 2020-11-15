package com.charlie.blockchain;

import java.security.KeyPair;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-28 23:35
 */
public interface IKeyPairGenerator
{
    KeyPair generateKeyPair()throws Exception;
}
