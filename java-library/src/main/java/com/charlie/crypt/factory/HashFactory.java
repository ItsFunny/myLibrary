package com.charlie.crypt.factory;

import com.charlie.crypt.IHash;
import com.charlie.crypt.impl.MD5Hasher;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-12-08 10:58
 */
public class HashFactory
{
    public static IHash defaultHashChain(){
        return new MD5Hasher();
    }
}
