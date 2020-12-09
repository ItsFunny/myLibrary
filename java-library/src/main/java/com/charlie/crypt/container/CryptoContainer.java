package com.charlie.crypt.container;

import com.charlie.crypt.IAsymmetricCrypto;
import com.charlie.crypt.IHash;
import com.charlie.crypt.ISymmetricCrypto;

/**
 * @author Charlie
 * @When
 * @Description 包含了所有的hash,symm和asymm方法
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-09 10:23
 */
public class CryptoContainer
{
    private IHash hasher;
    private ISymmetricCrypto symmetricCrypto;
    private IAsymmetricCrypto asymmetricCrypto;


}
