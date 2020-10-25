package com.charile.blockchain.configuration;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 06:49
 */
@Data
public class ChainCodeConfiguration
{
    private List<ChainCodeNode>chainCodes;
    @Data
    public static  class ChainCodeNode{
        private String chainCodeId;
        private String chainCodePath;
        private boolean needUpdate;
        private boolean needListOnBlockEvent;
        private String version;
        private String policyFile;
    }

}
