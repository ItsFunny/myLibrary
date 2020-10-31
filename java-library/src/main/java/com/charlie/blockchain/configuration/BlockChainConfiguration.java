package com.charlie.blockchain.configuration;

import com.charlie.base.AbstractInitOnce;
import com.charlie.base.IKeyImporter;
import com.charlie.blockchain.model.UserInfo;
import com.charlie.exception.ConfigException;
import com.charlie.service.IValidater;
import com.charlie.utils.FileUtils;
import lombok.Data;
import lombok.extern.java.Log;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.*;
import org.springframework.util.CollectionUtils;

import java.io.File;
import java.util.*;
import java.util.stream.Collectors;

/**
 * @author Charlie
 * @When
 * @Description 以channel为主, channel中有多少个组织, 对应多少个chaincode等信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:48
 */

@Data
@Log
public class BlockChainConfiguration extends AbstractInitOnce implements IValidater
{
    // 是否需要deploy 区块链网络 ,true的话,部署整个网络,否则只是初始化channel等信息
    private boolean deploy;


    private String version;
    // 是否启用tls
    private boolean tls;
    private String prefixPath;
    //    private String cryptoConfigPrefixPath;
    // tx 文件路径前缀
//    private String artifactsPrefixPath;
    //
    private String chainCodeRootDir;


    private ClientConfiguration clientConfiguration;
    // 以组织为主
    private OrganizationConfiguration organizationConfiguration;

    // orderer的配置
    private List<OrdererConfiguration> ordererConfigurations;

    // 所有的channel
    private ChannelConfiguration channelConfiguration;

    // 所有的peer
    private PeerConfiguration peerConfiguration;

    // chaincode
    private ChainCodeConfiguration chaincodeConfiguration;


    // ca信息
    private CaConfiguration caConfiguration;


    private BlockChainConfiguration()
    {

    }

    public boolean containsChannels(List<String> channels)
    {
        return this.channelConfiguration.contains(channels);
    }


    @Override
    protected void init() throws ConfigException
    {
        this.valid();
    }

//    public UserInfo getAlphaUser()
//    {
//        List<OrganizationConfiguration.OrganizationNode> organizationNodes = this.organizationConfiguration.getOrganizations();
//        for (OrganizationConfiguration.OrganizationNode organizationNode : organizationNodes)
//        {
//            if (organizationNode.isAlpha())
//            {
//                return organizationNode.getAdminUserInfo(Arrays.asList(IKeyImporter.COMMON_PEM_KEY_IMPORTER, IKeyImporter.STANDARD_SM2_KEY_IMPORTER));
//            }
//        }
//        throw new ConfigException("不可能会调到这里");
//    }

    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(this.version)) this.version = "1";

        if (StringUtils.isEmpty(this.prefixPath))
        {
            throw new ConfigException("prefixPath 不可为空");
        }
        this.prefixPath = FileUtils.appendFilePathIfNone(this.prefixPath);
//        if (StringUtils.isEmpty(this.artifactsPrefixPath))
//        {
//            this.artifactsPrefixPath = this.prefixPath + "artifacts" + File.separator;
//        } else
//        {
//            this.artifactsPrefixPath = FileUtils.appendFilePathIfNone(this.artifactsPrefixPath);
//        }
//        if (StringUtils.isEmpty(this.cryptoConfigPrefixPath))
//        {
//            this.cryptoConfigPrefixPath = this.prefixPath + "crypto-config" + File.separator;
//        } else
//        {
//            this.cryptoConfigPrefixPath = FileUtils.appendFilePathIfNone(this.cryptoConfigPrefixPath);
//        }
        this.clientConfiguration.valid();
        this.organizationConfiguration.valid();
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
            ordererConfiguration.valid();
        }
        this.channelConfiguration.valid();
        this.peerConfiguration.valid();
        this.chaincodeConfiguration.valid();
    }

    public boolean containsOrderers(List<String> orderers)
    {
        if (CollectionUtils.isEmpty(orderers))
        {
            throw new ConfigException("orderers不可为空");
        }
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
            Map<String, OrdererConfiguration.OrdererNode> collect = ordererConfiguration.getOrderers().stream().collect(Collectors.toMap(o -> o.getDomain(), ordererNode -> ordererNode));
            for (String orderer : orderers)
            {
                if (!collect.containsKey(orderer))
                {
                    return false;
                }
            }
        }
        return true;
    }

    public boolean containsPeers(List<String> s)
    {
        if (CollectionUtils.isEmpty(s))
        {
            throw new ConfigException("peers不可为空");
        }
        List<PeerConfiguration.PeerNode> peerNodes = this.peerConfiguration.getPeers();
        Map<String, PeerConfiguration.PeerNode> collect = peerNodes.stream().collect(Collectors.toMap(p -> p.getDomain(), p -> p));
        for (String s1 : s)
        {
            if (!collect.containsKey(s1))
            {
                return false;
            }
        }
        return true;
    }

    public boolean containsCa(List<String> ca)
    {
        CaConfiguration caConfiguration = this.getCaConfiguration();
        List<CaConfiguration.CaNode> caNodes = caConfiguration.getCaNodes();
//        caNodes.stream().collect(Collectors.toMap(c->c.getCaName()))

        return true;
    }

    public boolean containsChainCodes(List<String> cs)
    {
        if (CollectionUtils.isEmpty(cs))
        {
            throw new ConfigException("chaincodes 不可为空");
        }

        ChainCodeConfiguration chaincodeConfiguration = this.getChaincodeConfiguration();
        List<ChainCodeConfiguration.ChainCodeNode> chaincodes = chaincodeConfiguration.getChainCodes();
        Map<String, ChainCodeConfiguration> collect = chaincodes.stream().collect(Collectors.toMap(c -> c.getChainCodeId(), chainCodeNode -> chaincodeConfiguration));
        for (String c : cs)
        {
            if (!collect.containsKey(c))
            {
                return false;
            }
        }
        return true;
    }

    public UserInfo getAlphaUser()
    {
        List<OrganizationConfiguration.OrganizationNode> organizations = this.organizationConfiguration.getOrganizations();
        String organizationMspId = this.clientConfiguration.getOrganizationMspId();
        for (OrganizationConfiguration.OrganizationNode organization : organizations)
        {
            if (!organization.getMspId().equalsIgnoreCase(organizationMspId))
            {
                continue;
            }
            OrganizationConfiguration.UserNode adminUser = organization.getAdminUser();
            UserInfo info = new UserInfo(IKeyImporter.COMMON_PEM_KEY_IMPORTER, organizationMspId, adminUser.getName(), adminUser.getKeyString().getBytes(), adminUser.getCertString().getBytes());
            return info;
        }
        throw new ConfigException("不可能到达这里");
    }

    public PeerConfiguration.PeerNode getOrganizationAnchorPeer(String mspId)
    {
        List<OrganizationConfiguration.OrganizationNode> organizations = this.organizationConfiguration.getOrganizations();
        OrganizationConfiguration.OrganizationNode node = null;
        for (OrganizationConfiguration.OrganizationNode organization : organizations)
        {
            if (organization.getMspId().equalsIgnoreCase(mspId))
            {
                node = organization;
            }
        }
        List<String> peers = node.getPeers();
        return getAnchorPeer(peers);
    }


    public PeerConfiguration.PeerNode getAnchorPeer(List<String> peerDomainList)
    {
        List<PeerConfiguration.PeerNode> peers = this.peerConfiguration.getPeers();
        Map<String, PeerConfiguration.PeerNode> collect = peers.stream().collect(Collectors.toMap(p -> p.getDomain(), p -> p));
        PeerConfiguration.PeerNode anchorPeer = null;
        for (String p : peerDomainList)
        {
            PeerConfiguration.PeerNode peerNode = collect.get(p);
            if (peerNode.isAnchorPeer())
            {
                anchorPeer = peerNode;
                break;
            }
        }
        return anchorPeer;
    }

    public List<PeerConfiguration.PeerNode> getOrganizationAllPeers(List<String> ps)
    {
        List<PeerConfiguration.PeerNode> peers = this.peerConfiguration.getPeers();
        Map<String, PeerConfiguration.PeerNode> collect = peers.stream().collect(Collectors.toMap(p -> p.getDomain(), p -> p));
        List<PeerConfiguration.PeerNode> result = new ArrayList<>();
        for (String p : ps)
        {
            PeerConfiguration.PeerNode peerNode = collect.get(p);
            result.add(peerNode);
        }
        return result;
    }

    public List<Orderer> getChannelAllOrderers(String cid, HFClient client)
    {
        List<ChannelConfiguration.ChannelNode> channels = this.channelConfiguration.getChannels();
        ChannelConfiguration.ChannelNode node = null;
        for (ChannelConfiguration.ChannelNode channel : channels)
        {
            if (channel.getChannelId().equalsIgnoreCase(cid))
            {

                node = channel;
                break;
            }
        }
        return this.getChannelAllOrderers(node.getOrderers(), client);
    }

    public List<Orderer> getChannelAllOrderers(List<String> os, HFClient client)
    {
        List<OrdererConfiguration> ordererConfigurations = this.getOrdererConfigurations();
        Map<String, OrdererConfiguration.OrdererNode> nodeMap = new HashMap<>();
        for (OrdererConfiguration ordererConfiguration : ordererConfigurations)
        {
            List<OrdererConfiguration.OrdererNode> orderers = ordererConfiguration.getOrderers();
            for (OrdererConfiguration.OrdererNode orderer : orderers)
            {
                nodeMap.put(orderer.getDomain(), orderer);
            }
        }
        List<Orderer> result = new ArrayList<>();
        for (String o : os)
        {
            OrdererConfiguration.OrdererNode ordererNode = nodeMap.get(o);
            Orderer orderer = ordererNode.buildOrderer(tls, client);
            result.add(orderer);
        }
        return result;
    }

    public ChannelConfiguration.ChannelNode getChannelNode(String cid)
    {
        List<ChannelConfiguration.ChannelNode> channels = this.channelConfiguration.getChannels();
        for (ChannelConfiguration.ChannelNode channel : channels)
        {
            if (channel.getChannelId().equalsIgnoreCase(cid))
            {
                return channel;
            }
        }
        throw new ConfigException("这不可能发生");
    }

    public List<PeerConfiguration.PeerNode> getOrgAllEndorserPeers(List<String> ps)
    {

        List<PeerConfiguration.PeerNode> peers = this.peerConfiguration.getPeers();
        Map<String, PeerConfiguration.PeerNode> collect = peers.stream().collect(Collectors.toMap(p -> p.getDomain(), p -> p));
        List<PeerConfiguration.PeerNode> res = new ArrayList<>();
        for (String p : ps)
        {
            if (!collect.containsKey(p))
            {
                continue;
            }
            PeerConfiguration.PeerNode peerNode = collect.get(p);
            if (peerNode.isEndorsingPeer())
            {
                res.add(peerNode);
            }
        }

        return res;
    }

    public List<ChainCodeConfiguration.ChainCodeNode> getChainCodes(List<String> cids)
    {
        List<ChainCodeConfiguration.ChainCodeNode> res = new ArrayList<>();
        List<ChainCodeConfiguration.ChainCodeNode> chainCodes = this.getChaincodeConfiguration().getChainCodes();
        Map<String, ChainCodeConfiguration.ChainCodeNode> collect = chainCodes.stream().collect(Collectors.toMap(c -> c.getChainCodeId(), c -> c));
        for (String cid : cids)
        {
            ChainCodeConfiguration.ChainCodeNode chainCodeNode = collect.get(cid);
            if (null != chainCodeNode)
            {
                res.add(chainCodeNode);
            }
        }
        return res;
    }

    public List<ChannelConfiguration.ChannelNode> getChannelsByPeers(List<String> ps)
    {
        List<PeerConfiguration.PeerNode> peers = this.peerConfiguration.getPeers();
        Map<String, PeerConfiguration.PeerNode> collect = peers.stream().collect(Collectors.toMap(p -> p.getDomain(), p -> p));
        List<ChannelConfiguration.ChannelNode> result = new ArrayList<>();
        for (String p : ps)
        {
            PeerConfiguration.PeerNode peerNode = collect.get(p);
            if (null == peerNode)
            {
                continue;
            }
            List<String> channels = peerNode.getChannels();
            for (String cid : channels)
            {
                result.add(this.getChannelNode(cid));
            }
        }
        return result;
    }

    public List<PeerConfiguration.PeerNode> getChannelAllPeers(String channelId)
    {
        final PeerConfiguration peerConfiguration = this.getPeerConfiguration();
        List<PeerConfiguration.PeerNode> peers = peerConfiguration.getPeers();
        List<PeerConfiguration.PeerNode> peerNodes = new ArrayList<>();
        for (PeerConfiguration.PeerNode peer : peers)
        {
            if (peer.getChannels().contains(channelId))
            {
                peerNodes.add(peer);
            }
        }
        return peerNodes;
    }

    // 获取该channel下的所有要安装该chaincode的peer
    public List<PeerConfiguration.PeerNode> getAllEndorserPeersByChaincodeIDAndChannelId(String chaincodeId, String channelId)
    {
        List<OrganizationConfiguration.OrganizationNode> organizations = this.getOrganizationConfiguration().getOrganizations();
        List<PeerConfiguration.PeerNode> channelAllPeers = this.getChannelAllPeers(channelId);
        List<PeerConfiguration.PeerNode> result = new ArrayList<>();
        for (PeerConfiguration.PeerNode peer : channelAllPeers)
        {
            List<String> chainCodes = peer.getChainCodes();
            if (CollectionUtils.isEmpty(chainCodes)) continue;
            if (chainCodes.contains(chaincodeId))
            {
                result.add(peer);
            }
        }
        return result;
    }
}
