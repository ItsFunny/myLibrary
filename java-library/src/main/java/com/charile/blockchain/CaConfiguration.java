package com.charile.blockchain;

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
public class CaConfiguration
{
    private List<CaNode>caNodes;

    @Data
    class CaNode
    {
        private String caName;
        private String ip;
        private Integer port;
    }
}
