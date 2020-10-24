package com.charile.blockchain.configuration;

import com.charile.blockchain.util.TLSUtils;
import com.charile.exception.ConfigException;
import com.charile.service.IValidater;
import com.charile.utils.FileUtils;
import lombok.Data;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.springframework.util.CollectionUtils;
import org.springframework.util.StringUtils;
import sun.security.krb5.Config;

import java.io.IOException;
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
public class PeerConfiguration implements IValidater
{
    private PeerConfiguration() {}

    private PeerNode anchorPeer;
    private List<EndorserPeer> endorserPeers;

    @Override
    public void valid()
    {
        if (CollectionUtils.isEmpty(endorserPeers))
        {
            throw new ConfigException("endorser 节点不可为空");
        }
        // FIXME 校验 格式
        if (anchorPeer == null)
        {
            PeerNode peerNode = new PeerNode();
            peerNode.setDomain(endorserPeers.get(0).getDomain());
            peerNode.setIpWithPort(endorserPeers.get(0).getIpWithPort());
            peerNode.setTlsCertFile(endorserPeers.get(0).getTlsCertFile());
            anchorPeer = peerNode;
        }
        anchorPeer.valid();
        for (EndorserPeer endorserPeer : endorserPeers)
        {
            endorserPeer.valid();
        }
    }

    @Data
    public static class PeerNode implements IValidater
    {
        // 域名
        protected String domain;
        protected String ipWithPort;
        //        private Integer port;
        protected String tlsCertFile;

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

        public Peer conv2TPeer(boolean tls, HFClient client)
        {
            String grpcUrl = tls ? "grpcs://" : "grpc://";
            grpcUrl += this.getIpWithPort();
            Peer peer = null;
            try
            {
                peer = client.newPeer(this.getDomain(), grpcUrl, this.buildProperties());
            } catch (InvalidArgumentException e)
            {
                throw new ConfigException("解析peer失败:" + e.getMessage());
            }
            return peer;
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
            if (StringUtils.isEmpty(this.ipWithPort))
            {
                throw new ConfigException("peer 的ipwithport 不可为空");
            }
            this.tlsCertFile = ConfigurationFactory.getInstance().getBlockChainConfiguration().getCryptoConfigPrefixPath() + this.tlsCertFile;
        }
    }

    @Data
    public static class EndorserPeer extends PeerNode
    {
        private List<ChaincodeNode> chainCodes;

        @Override
        public void valid()
        {
            super.valid();
            if (!CollectionUtils.isEmpty(this.chainCodes))
            {
                for (ChaincodeNode chainCode : chainCodes)
                {
                    chainCode.valid();
                }
            }

        }
    }

    @Data
    public static class ChaincodeNode implements IValidater
    {
        private String chainCodeId;
        private String chainCodePath;
        private boolean needUpdate;
        private boolean needListOnBlockEvent;
        private String version;
        private String policyFile;

        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.chainCodeId))
            {
                throw new ConfigException("chainCodeId 不可为空");
            }
            if (StringUtils.isEmpty(this.chainCodePath))
            {
                throw new ConfigException("chainCodePath 不可为空");
            }
        }
    }


}
