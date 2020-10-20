package com.charile.blockchain;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:06
 */
@Data
public class ChannelConfiguration
{
    private String channelId;
    private String channelConfigPath;
    private boolean needListOnBlockEvent;

    // 该channel下的所有orderer组织,默认是只有1个的
    private List<OrdererConfiguration> ordererConfiguration;

    // 该channel下的所有组织
    private List<OrganizationConfiguration> organizationConfigurations;

}
