package com.charile.blockchain.configuration;

import com.charile.blockchain.util.TLSUtils;
import com.charile.exception.ConfigException;
import com.charile.service.IValidater;
import com.charile.utils.FileUtils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;

import java.io.IOException;
import java.util.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:07
 */
@Data
public class OrdererConfiguration implements IValidater
{
    private OrdererConfiguration() {}

    private String orderMspId;
    private List<OrdererNode> orderers;


//    public Collection<Orderer> searchOrders(List<String> orderDomain, boolean tls, HFClient client)
//    {
//        Set<Orderer> orderers = new HashSet<>();
//        for (OrdererNode node : orderers)
//        {
//            for (String s : orderDomain)
//            {
//                if (node.getDomain().equalsIgnoreCase(s))
//                {
//                    orderers.add(node.buildOrderer(tls, client));
//                }
//            }
//        }
//        return orderers;
//    }

//    public Collection<Orderer> buildOrders(boolean tls, HFClient client)
//    {
//        Set<Orderer> result = new HashSet<>();
//        for (OrdererNode ordererNode : orderers)
//        {
//            orderers.add(ordererNode.buildOrderer(tls, client));
//        }
//        return result;
//    }

    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(this.orderMspId))
        {
            throw new ConfigException("orderMsp不可为空");
        }
        for (OrdererNode ordererNode : orderers)
        {
            ordererNode.valid();
        }

    }

    @Data
    class OrdererNode implements IValidater
    {
        private String domain;
        private String ipWithPort;
        private String tlsCertFile;

        public Orderer buildOrderer(boolean tls, HFClient client)
        {
            String grpcUrl = tls ? "grpcs://" : "grpc://";
            grpcUrl += this.getIpWithPort();
            Properties properties = null;
            try
            {
                properties = TLSUtils.loadTLSFile(this.tlsCertFile, this.domain);
            } catch (IOException e)
            {
                throw new ConfigException("无法解析tls证书");
            }
            try
            {
                Orderer orderer = client.newOrderer(this.domain, grpcUrl, properties);
                return orderer;
            } catch (InvalidArgumentException e)
            {
                throw new ConfigException("配置获取ordrer信息失败");
            }
        }

        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.domain))
            {
                throw new ConfigException("domain 不可为空");
            }
            if (StringUtils.isEmpty(this.ipWithPort))
            {
                throw new ConfigException("ipWithPort不可为空");
            }
            if (StringUtils.isEmpty(this.tlsCertFile))
            {
                throw new ConfigException("证书不可为空");
            }
            this.tlsCertFile = ConfigurationFactory.getInstance().getBlockChainConfiguration().getCryptoConfigPrefixPath() + FileUtils.appendFilePathIfNone(this.tlsCertFile);
        }
    }
}
