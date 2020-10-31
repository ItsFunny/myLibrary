package com.charlie.blockchain.configuration;

import com.charlie.service.IValidater;
import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:03
 */
@Data
public class CaConfiguration implements IValidater
{
    private CaConfiguration(){}
    private List<CaNode>caNodes;

    @Override
    public void valid()
    {

    }

    @Data
    public static class CaNode
    {
        private String domain;
        private String ipWithPort;
    }
}
