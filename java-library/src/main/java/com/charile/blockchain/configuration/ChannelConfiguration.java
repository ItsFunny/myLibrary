package com.charile.blockchain.configuration;

import com.charile.exception.ConfigException;
import com.charile.service.IValidater;
import com.charile.utils.FileUtils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Orderer;
import org.springframework.core.annotation.Order;

import java.util.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:06
 */
@Data
public class ChannelConfiguration implements IValidater
{
    private ChannelConfiguration() {}

    private String channelId;
    private String channelConfigPath;
    private boolean needListOnBlockEvent;

    // 该channel下的所有orderer组织,默认是只有1个的
    private List<String> orderers;

    public Orderer getOneOrderer(boolean tls, HFClient client)
    {
        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
            for (OrdererConfiguration configuration : ordererConfigurations)
            {
                List<OrdererConfiguration.OrdererNode> ordererNodes = configuration.getOrdererNodes();
                for (OrdererConfiguration.OrdererNode ordererNode : ordererNodes)
                {
                    if (ordererNode.getDomain().equalsIgnoreCase(this.orderers.get(0)))
                    {
                        return ordererNode.buildOrderer(tls, client);
                    }
                }
            }
        }
        throw new ConfigException("找不到匹配的orderer");
    }

    public Collection<Orderer> buildOrders(boolean tls, HFClient client)
    {
        Set<Orderer> orders = new HashSet<>();
        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
        Set<Orderer> orderers = new HashSet<>();
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
            orderers.addAll(ordererConfiguration.searchOrders(this.getOrderers(), tls, client));
        }
        return orders;
    }

    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(this.channelId))
        {
            throw new ConfigException("channelid不可为空");
        }
        if (StringUtils.isEmpty(this.channelConfigPath))
        {
            throw new ConfigException("channelConfigPath 不可为空");
        }
        this.channelConfigPath = FileUtils.cutPathIfStartWith(this.channelConfigPath);
        this.channelConfigPath = ConfigurationFactory.getInstance().getBlockChainConfiguration().getArtifactsPrefixPath() + this.channelConfigPath;


        // FIXME 校验是否存在
//        Set<Orderer> orders = new HashSet<>();
//        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
//        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
    }
}
