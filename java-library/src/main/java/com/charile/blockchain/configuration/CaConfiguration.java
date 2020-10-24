package com.charile.blockchain.configuration;

import com.charile.service.IValidater;
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
    class CaNode
    {
        private String caName;
        private String ip;
        private Integer port;
    }
}
