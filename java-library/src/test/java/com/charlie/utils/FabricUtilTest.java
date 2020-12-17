package com.charlie.utils;

import org.hyperledger.fabric.sdk.BlockInfo;
import org.junit.Test;

public class FabricUtilTest
{

    @Test
    public void getRWSetFromBlock()throws Exception
    {
        BlockInfo blockInfo = JSONUtil.jsonFileToObj("/Users/joker/fsdownload/a.block", BlockInfo.class);
        System.out.println(blockInfo);

    }
}