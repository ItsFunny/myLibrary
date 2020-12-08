package com.charlie.blockchain.configuration;

import com.charlie.exception.ConfigException;
import com.charlie.service.IValidator;
import com.charlie.utils.CollectionUtils;
import com.charlie.utils.FileUtils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Orderer;

import java.util.*;
import java.util.stream.Collectors;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:06
 */
@Data
public class ChannelConfiguration implements IValidator
{
    private ChannelConfiguration() {}

    private List<ChannelNode> channels;

    public boolean contains(List<String> chs)
    {
        int count = 0;
        for (ChannelNode ch : channels)
        {
            for (String name : chs)
            {
                if (name.equalsIgnoreCase(ch.getChannelId()))
                {
                    count++;
                }
            }
        }
        return count == chs.size();
    }

    @Data
    public static class ChannelNode implements IValidator
    {
        private String channelId;
        private String channelConfigPath;
        private boolean needListOnBlockEvent;

        private byte type;

        // 该channel下的所有orderer组织,默认是只有1个的
        private List<ChannelOrderInfo> orderers;

        // 该channel下的所有peer
        private List<ChannelPeerInfo> peers;

//        // 该channel下的所有peer
//        private List<String>peers;

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
            this.channelConfigPath = ConfigurationFactory.getInstance().getBlockChainConfiguration().getPrefixPath() + this.channelConfigPath;
            if (!CollectionUtils.isEmpty(this.orderers))
            {
                List<String> collect = this.orderers.stream().map(o -> o.getDomain()).collect(Collectors.toList());
                if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().containsOrderers(collect))
                {
                    throw new ConfigException("orderer不匹配");
                }
            }
            if (!CollectionUtils.isEmpty(this.peers))
            {
                List<String> collect = this.peers.stream().map(p -> p.getDomain()).collect(Collectors.toList());
                if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().containsPeers(collect))
                {
                    throw new ConfigException("channel 配置,peer不匹配");
                }
            }
        }
    }


    @Data
    public static class ChannelOrderInfo implements IValidator
    {
        private String domain;
        private byte type;

        @Override
        public void valid()
        {

        }
    }

    @Data
    public static class ChannelPeerInfo implements IValidator
    {
        private String domain;
        private byte type;

        @Override
        public void valid()
        {

        }
    }
//    @Deprecated
//    private String channelId;
//    @Deprecated
//    private String channelConfigPath;
//    @Deprecated
//    private boolean needListOnBlockEvent;
//
//    // 该channel下的所有orderer组织,默认是只有1个的
//    private List<String> orderers;

//    public Orderer getOneOrderer(boolean tls, HFClient client)
//    {
//        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
//        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
//        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
//        {
//            for (OrdererConfiguration configuration : ordererConfigurations)
//            {
//                List<OrdererConfiguration.OrdererNode> ordererNodes = configuration.getOrderers();
//                for (OrdererConfiguration.OrdererNode ordererNode : ordererNodes)
//                {
//                    if (ordererNode.getDomain().equalsIgnoreCase(this.orderers.get(0)))
//                    {
//                        return ordererNode.buildOrderer(tls, client);
//                    }
//                }
//            }
//        }
//        throw new ConfigException("找不到匹配的orderer");
//    }

    public Collection<Orderer> buildOrders(boolean tls, HFClient client)
    {
        Set<Orderer> orders = new HashSet<>();
        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
        Set<Orderer> orderers = new HashSet<>();
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
//            orderers.addAll(ordererConfiguration.(this.getOrderers(), tls, client));
        }
        return orders;
    }

    @Override
    public void valid()
    {

        for (ChannelNode channel : channels)
        {
            channel.valid();
        }

        // FIXME 校验是否存在
//        Set<Orderer> orders = new HashSet<>();
//        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
//        List<OrdererConfiguration> ordererConfigurations = blockChainConfiguration.getOrdererConfigurations();
    }
}
