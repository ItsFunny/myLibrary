package com.charlie.blockchain.configuration;

import com.charlie.blockchain.util.TLSUtils;
import com.charlie.exception.ConfigException;
import com.charlie.service.IValidator;
import lombok.Data;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.IDataDecorator;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.springframework.util.CollectionUtils;
import org.springframework.util.StringUtils;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;
import java.util.Objects;
import java.util.Properties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:58
 */
@Data
public class PeerConfiguration implements IValidator
{
    private PeerConfiguration() {}

    private List<PeerNode> peers;

//    private PeerNode anchorPeer;
//    private List<EndorserPeer> endorserPeers;

    @Override
    public void valid()
    {
        if (CollectionUtils.isEmpty(peers))
        {
            throw new ConfigException("peers 节点不可为空");
        }
        boolean containsAnchor = false;
        for (PeerNode peer : peers)
        {
            if (peer.isAnchorPeer())
            {
                containsAnchor = true;
            }
            peer.valid();
        }
        if (!containsAnchor)
        {
            peers.get(0).setAnchorPeer(true);
        }
        // FIXME 校验 格式
//        if (anchorPeer == null)
//        {
//            PeerNode peerNode = new PeerNode();
//            peerNode.setDomain(endorserPeers.get(0).getDomain());
//            peerNode.setIpWithPort(endorserPeers.get(0).getIpWithPort());
//            peerNode.setTlsCertFile(endorserPeers.get(0).getTlsCertFile());
//            anchorPeer = peerNode;
//        }
//        anchorPeer.valid();
//        for (EndorserPeer endorserPeer : endorserPeers)
//        {
//            endorserPeer.valid();
//        }
    }

    @Data
    public static class PeerNode implements IValidator
    {
        // 域名
        protected String domain;
        protected String url;
        //        private Integer port;
        protected String tlsCertFile;

        protected List<String> chainCodes;

        protected boolean anchorPeer;

        protected boolean endorsingPeer = true;
        protected boolean chaincodeQuery = true;
        protected boolean ledgerQuery = true;
        protected boolean eventSource = true;

//        // 2020-11-13 22:38 add ,判断算法类型,是国密还是ecdsa,默认为0,代表的是ecdsa
//        private byte type;


        @Override
        public boolean equals(Object o)
        {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            PeerNode peerNode = (PeerNode) o;
            return domain.equals(peerNode.domain);
        }

        @Override
        public int hashCode()
        {
            return Objects.hash(domain);
        }

        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.tlsCertFile))
            {
                throw new ConfigException("peer的tlsCertFile不可为空");
            }
            if (StringUtils.isEmpty(this.domain))
            {
                throw new ConfigException("domain 不可为空");
            }
            if (StringUtils.isEmpty(this.url))
            {
                throw new ConfigException("peer 的ipwithport 不可为空");
            }
            this.tlsCertFile = ConfigurationFactory.getInstance().getBlockChainConfiguration().getPrefixPath() + this.tlsCertFile;
            if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().containsChainCodes(this.chainCodes))
            {
                throw new ConfigException("chaincode 不匹配");
            }
//            if (!CollectionUtils.isEmpty(this.channels))
//            {
//                List<String> names = this.channels.stream().map(t -> t.getChannelName()).collect(Collectors.toList());
//                if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().containsChannels(names))
//                {
//                    throw new ConfigException("channels 不匹配,该peer所属的channel在配置中并不存在");
//                }
//            }
        }
        public Properties buildProperties()
        {
            Properties properties = null;
            try
            {
                properties = TLSUtils.loadTLSFile(this.tlsCertFile, this.domain);
            } catch (IOException e)
            {
                throw new ConfigException("无法解析tls证书");
            }
            return properties;
        }
    }

    @Data
    public static class PeerChannelBO {
        // 域名
        protected String domain;
        protected String url;
        //        private Integer port;
        protected String tlsCertFile;

        protected List<String> chainCodes;

        protected boolean anchorPeer;

        protected boolean endorsingPeer = true;
        protected boolean chaincodeQuery = true;
        protected boolean ledgerQuery = true;
        protected boolean eventSource = true;

        // 2020-11-13 22:38 add ,判断算法类型,是国密还是ecdsa,默认为0,代表的是ecdsa
        private byte type;
        public Peer conv2TPeer(boolean tls, HFClient client)
        {
            String grpcUrl = this.getUrl();
            Peer peer = null;
            try
            {
                peer = client.newPeer(this.getDomain(), grpcUrl, this.buildProperties());
                peer.decorate(Arrays.asList((IDataDecorator<Peer>) peer1 ->
                {
                    peer1.setType(type);
                    return peer1;
                }));
            } catch (InvalidArgumentException e)
            {
                throw new ConfigException("解析peer失败:" + e.getMessage());
            }
            return peer;
        }

        public Properties buildProperties()
        {
            Properties properties = null;
            try
            {
                properties = TLSUtils.loadTLSFile(this.tlsCertFile, this.domain);
            } catch (IOException e)
            {
                throw new ConfigException("无法解析tls证书");
            }
            return properties;
        }


    }

    @Data
    public static class PeerChannelTypeInfo implements IValidator
    {
        private String channelName;
        private byte type;

        @Override
        public void valid()
        {

        }
    }

//    @Data
//    public static class EndorserPeer extends PeerNode
//    {
//        private List<ChaincodeNode> chainCodes;
//
//        @Override
//        public void valid()
//        {
//            super.valid();
//            if (!CollectionUtils.isEmpty(this.chainCodes))
//            {
//                for (ChaincodeNode chainCode : chainCodes)
//                {
//                    chainCode.valid();
//                }
//            }
//
//        }
//    }


}
