package com.charlie.blockchain.configuration;

import com.charlie.blockchain.util.TLSUtils;
import com.charlie.exception.ConfigException;
import com.charlie.service.IValidator;
import com.charlie.utils.FileUtils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.IDataDecorator;
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
public class OrdererConfiguration implements IValidator
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
    class OrdererNode implements IValidator
    {
        private String domain;
        private String url;
        private String tlsCertFile;


        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.domain))
            {
                throw new ConfigException("domain 不可为空");
            }
            if (StringUtils.isEmpty(this.url))
            {
                throw new ConfigException("ipWithPort不可为空");
            }
            if (StringUtils.isEmpty(this.tlsCertFile))
            {
                throw new ConfigException("证书不可为空");
            }
            this.tlsCertFile = ConfigurationFactory.getInstance().getBlockChainConfiguration().getPrefixPath() + FileUtils.appendFilePathIfNone(this.tlsCertFile);
        }
    }

    @Data
    public static class OrdererChannelBO {
        private String domain;
        private String url;
        private String tlsCertFile;
        private byte type;

        public Orderer buildOrderer(boolean tls, HFClient client)
        {
            String grpcUrl=this.getUrl();
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
                orderer.decorate(Arrays.asList((IDataDecorator<Orderer>) orderer1 ->
                {
                    orderer1.setType(type);
                    return orderer1;
                }));
                return orderer;
            } catch (InvalidArgumentException e)
            {
                throw new ConfigException("配置获取ordrer信息失败");
            }
        }
    }
}
