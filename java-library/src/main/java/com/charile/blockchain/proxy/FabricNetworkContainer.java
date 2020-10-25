package com.charile.blockchain.proxy;

import com.charile.base.AbstractInitOnce;
import com.charile.base.IKeyImporter;
import com.charile.blockchain.configuration.*;
import com.charile.blockchain.configuration.ChannelConfiguration;
import com.charile.blockchain.model.NewClientReq;
import com.charile.blockchain.model.UserInfo;
import com.charile.blockchain.wrapper.HFClientWrapper;
import com.charile.exception.ConfigException;
import com.charile.utils.Base64Utils;
import lombok.Data;
import oracle.jvm.hotspot.jfr.JFR;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.protos.peer.Query;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric.sdk.security.CryptoSuiteFactory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import sun.security.krb5.Config;

import java.io.File;
import java.time.Period;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @author Charlie
 * @When
 * @Description 用于获取 client 对象,HFClient等,持久化存储相关信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-23 23:10
 */
@Data
public class FabricNetworkContainer extends AbstractInitOnce
{
    private HFClientWrapper clientWrapper;
    private Logger logger = LoggerFactory.getLogger(this.getClass());


    private FabricNetworkContainer()
    {
        try
        {
            this.initOnce();
        } catch (Exception e)
        {
            throw new RuntimeException(e);
        }
    }

    public static HFClient createNewClient(NewClientReq req)
    {
        HFClient instance = HFClient.createNewInstance();
        try
        {
            instance.setUserContext(new UserInfo(req.getKeyImporter(), req.getMspId(), req.getName(), req.getKeyBytes(), req.getCertBytes()));
        } catch (InvalidArgumentException e)
        {
            throw new RuntimeException("创建新的客户端失败", e);
        }
        return instance;
    }


    @Data
    private static class FabricClientProxyFactoryHolder
    {
        private static final FabricNetworkContainer INSTANCE = new FabricNetworkContainer();

        private FabricClientProxyFactoryHolder() { }
    }

    public static FabricNetworkContainer getInstance()
    {
        return FabricClientProxyFactoryHolder.INSTANCE;
    }


    @Override
    protected void init() throws ConfigException
    {
        try
        {
            // 初始化区块链网络
            logger.info("开始初始化区块链网络");
            BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
            HFClient alphaClient = this.buildAlphaHFClient(blockChainConfiguration);
            this.clientWrapper = new HFClientWrapper();

            blockChainConfiguration.initOnce();
            this.initlizeNetwork(alphaClient, blockChainConfiguration);
            this.installAndInstantiateCC(alphaClient, blockChainConfiguration);
            this.afterInitNetwork(blockChainConfiguration,alphaClient);
        } catch (Exception e)
        {
            throw new ConfigException(e);
        }
    }

    private void afterInitNetwork(BlockChainConfiguration blockChainConfiguration,HFClient alphaClient)throws Exception{
        UserInfo alphaUser = blockChainConfiguration.getAlphaUser();
        // 初始化完毕之后,重新设置为alpha client的证书
        alphaClient.setUserContext(alphaUser);
        this.clientWrapper.setAlphaClient(alphaClient);
    }

    private HFClient createNewHfClient() throws Exception
    {
        HFClient instance = HFClient.createNewInstance();
        CryptoSuite cryptoSuite = CryptoSuiteFactory.getDefault().getCryptoSuite();
        instance.setCryptoSuite(cryptoSuite);
        return instance;
    }

    private HFClient buildAlphaHFClient(BlockChainConfiguration blockChainConfiguration) throws Exception
    {
        HFClient instance = HFClient.createNewInstance();
        CryptoSuite cryptoSuite = CryptoSuiteFactory.getDefault().getCryptoSuite();
        instance.setCryptoSuite(cryptoSuite);
        return instance;
    }


    private void initlizeNetwork(HFClient client, BlockChainConfiguration blockChainConfiguration) throws Exception
    {
        final HFClient hfClient = client;
        logger.info("查询已经存在的channel");

        OrganizationConfiguration organizationConfiguration = blockChainConfiguration.getOrganizationConfiguration();
        List<OrganizationConfiguration.OrganizationNode> organizationNodes = organizationConfiguration.getOrganizations();

        for (OrganizationConfiguration.OrganizationNode node : organizationNodes)
        {
            logger.info("开始组装当前组织的admin用户信息,organizationId为:{}", node.getMspId());
            User adminUserInfo = null;
            if (clientWrapper.getUserContextMap().containsKey(node.getMspId()))
            {
                adminUserInfo = clientWrapper.getUserContextMap().get(node.getMspId());
            } else
            {
                adminUserInfo = node.getAdminUserInfo(Arrays.asList(IKeyImporter.COMMON_PEM_KEY_IMPORTER, IKeyImporter.STANDARD_SM2_KEY_IMPORTER, IKeyImporter.ANS1_PEM_KEY_IMPORTER));
                this.clientWrapper.put(node.getMspId(), adminUserInfo);
            }
            logger.info("adminUser:" + adminUserInfo);
            hfClient.setUserContext(adminUserInfo);
            logger.info("结束组装admin用户信息");


            String organizationId = node.getMspId();
            logger.info("开始初始化组织id为[{}]的相关区块链信息", organizationId);
            List<String> channels = node.getChannels();
            logger.info("该组织{}加入的channel数量为:{}", organizationId, channels.size());
            for (String cid : channels)
            {
//            for (ChannelConfiguration channelConfiguration : channelConfigurations)
//            {
                String channelId = cid;

                PeerConfiguration.PeerNode anchorPeer = blockChainConfiguration.getOrganizationAnchorPeer(node.getMspId());
                logger.info("anchorPeer:{}", anchorPeer);
                String grpcUrl = blockChainConfiguration.isTls() ? "grpcs://" : "grpc://";
                grpcUrl += anchorPeer.getIpWithPort();
                Peer peer = hfClient.newPeer(anchorPeer.getDomain(), grpcUrl, anchorPeer.buildProperties());

                logger.info("begin 查询该组织[{}]已经加入的channel信息", channelId);
                Set<String> channelStrings = hfClient.queryChannels(peer);
                logger.info("end 查询该组织[{}]已经加入的channel信息,:{}", channelId, channelStrings);

                boolean channelCreated = false;
                logger.info("begin 判断该channel是否已经创建");
                for (String channelString : channelStrings)
                {
                    if (channelString.equalsIgnoreCase(channelId))
                    {
                        channelCreated = true;
                    }
                }
                logger.info("end 判断该channel是否已经创建:{}", channelCreated);
                Channel channel = null;
                logger.info("获取channel的所有peer");
                // 获取该组织的所有配置着的peer,通通加入
                Collection<PeerConfiguration.PeerNode> allPeers = blockChainConfiguration.getOrganizationAllPeers(node.getPeers());
                logger.info("获取channel的所有orderer");
                List<Orderer> allOrderers = blockChainConfiguration.getChannelAllOrderers(cid, hfClient);

                if (channelCreated)
                {
                    logger.info("该channel:{}已经创建", channelId);
                    channel = client.newChannel(channelId);
                    for (PeerConfiguration.PeerNode peerNode : allPeers)
                    {
                        channel.addPeer(peerNode.conv2TPeer(blockChainConfiguration.isTls(), hfClient));
                    }
                    for(Orderer orderer : allOrderers){
                        channel.addOrderer(orderer);
                    }
                } else
                {
                    logger.info("该channel 可能未创建,所以开始准备创建,组织id:{},chanelId:{}", node.getMspId(), cid);
                    logger.info("begin 创建channel");
                    ChannelConfiguration.ChannelNode channelNode = blockChainConfiguration.getChannelNode(cid);
                    org.hyperledger.fabric.sdk.ChannelConfiguration configuration = new org.hyperledger.fabric.sdk.ChannelConfiguration(new File(channelNode.getChannelConfigPath()));
                    byte[] channelConfigurationSignature = client.getChannelConfigurationSignature(configuration, adminUserInfo);
                    logger.info("签名信息为:" + Base64Utils.encode(channelConfigurationSignature));
                    channel = hfClient.newChannel(channelId,allOrderers.get(0), configuration, channelConfigurationSignature);
                    TimeUnit.SECONDS.sleep(8);

                    logger.info("将所有的peer都加入到channel中");
                    joinAllPeers(hfClient, blockChainConfiguration.isTls(), channel, allPeers);
                }
                try
                {
                    logger.info("begin channel 开始初始化");
                    channel.initialize();
                    logger.info("end channel 结束初始化");
                } catch (Exception e)
                {
                    logger.error("channel初始化错误" + e.getMessage());
                    throw new ConfigException(e);
                }
            }

        }
    }

    private void installAndInstantiateCC(HFClient client, BlockChainConfiguration blockChainConfiguration) throws Exception
    {
        logger.info("begin 安装并且实例化链码");
        final HFClient hfClient = client;
        OrganizationConfiguration organizationConfiguration = blockChainConfiguration.getOrganizationConfiguration();
        List<OrganizationConfiguration.OrganizationNode> organizationNodes = organizationConfiguration.getOrganizations();
        for (OrganizationConfiguration.OrganizationNode node : organizationNodes)
        {
            List<String> channels = node.getChannels();

            for (String cid : channels)
            {
                logger.info("begin 对组织[{}]进行相关操作", node.getMspId());
                User user = this.clientWrapper.getUser(node.getMspId());
                hfClient.setUserContext(user);
                List<String> peers = node.getPeers();
                List<PeerConfiguration.PeerNode> endorserPeers = blockChainConfiguration.getOrgAllEndorserPeers(peers);
                for (PeerConfiguration.PeerNode endorserPeer : endorserPeers)
                {
                    Channel channel = hfClient.getChannel(cid);
                    Collection<Peer> channelPeers = channel.getPeers();
                    Peer peer = null;
                    for (Peer channelPeer : channelPeers)
                    {
                        if (endorserPeer.getDomain().equalsIgnoreCase(channelPeer.getName()))
                        {
                            peer = channelPeer;
                            break;
                        }
                    }
                    if (null == peer)
                    {
                        throw new ConfigException("配置错误,无法在channel中匹配的peer节点,可能是实例化网络的时候,并没有将该peer加入");
                    }
                    List<String> chainCodeIDList = endorserPeer.getChainCodes();
                    List<ChainCodeConfiguration.ChainCodeNode> chainCodes = blockChainConfiguration.getChainCodes(chainCodeIDList);
                    for (ChainCodeConfiguration.ChainCodeNode chainCode : chainCodes)
                    {
                        boolean ccHasInstall = false;
                        boolean needUpgrade = false;
                        String chainCodeId = chainCode.getChainCodeId();
                        logger.info("查询该节点[{}]已经安装的链码:", endorserPeer.getDomain());
                        List<Query.ChaincodeInfo> chaincodeInfos = hfClient.queryInstalledChaincodes(peer);

                        double installVersion = StringUtils.isEmpty(chainCode.getVersion()) ? 0.01 : Double.parseDouble(chainCode.getVersion());
                        for (Query.ChaincodeInfo chaincodeInfo : chaincodeInfos)
                        {
                            double prevVersion = Double.parseDouble(chaincodeInfo.getVersion());
                            String name = chaincodeInfo.getName();
                            if (name.equalsIgnoreCase(chainCodeId))
                            {
                                logger.info("该chaincode:[{}]已经安装,查看是否需要升级", chainCodeId);
                                ccHasInstall = true;
                                if (chainCode.isNeedUpdate())
                                {
                                    needUpgrade = true;
                                    installVersion = installVersion <= prevVersion ? installVersion + 0.01 : installVersion;
                                    logger.info("该chaincode已经安装,同时需要升级,newVersion=[{}],prevVersion=[{}]", installVersion, prevVersion);
                                }
                            }
                        }

                        ChaincodeID.Builder chaincodeIDBuilder = ChaincodeID.newBuilder().setName(chainCodeId).setVersion(installVersion + "")
                                .setPath(chainCode.getChainCodePath());
                        ChaincodeID chaincodeID = chaincodeIDBuilder.build();

                        if (ccHasInstall)
                        {
                            if (!needUpgrade)
                            {
                                logger.info("该chaincode=[{}],不需要升级", chainCodeId);
                            } else
                            {
                                logger.info("该chaincode需要升级,因此需要安装新版本,再升级");

                                InstallProposalRequest request = hfClient.newInstallProposalRequest();
                                request.setChaincodeID(chaincodeID);
                                request.setUserContext(hfClient.getUserContext());
                                request.setChaincodeSourceLocation(new File(blockChainConfiguration.getChainCodeRootDir()));
                                request.setChaincodeVersion(installVersion + "");
                                Collection<ProposalResponse> responses = hfClient.sendInstallProposal(request, Arrays.asList(peer));
                                handlerResponse(responses, "链码升级前安装");


                                UpgradeProposalRequest upgradeProposalRequest = hfClient.newUpgradeProposalRequest();
                                upgradeProposalRequest.setChaincodeID(chaincodeID);
                                upgradeProposalRequest.setArgs("a", "b", "c");
//                            upgradeProposalRequest.setProposalWaitTime();
                                upgradeProposalRequest.setFcn("init");
                                upgradeProposalRequest.setUserContext(user);
                                Collection<ProposalResponse> proposalResponses = channel.sendUpgradeProposal(upgradeProposalRequest, Arrays.asList(peer));
                                handlerResponse(proposalResponses, "链码升级");
                            }
                        } else
                        {
                            logger.info("该chaincodeid=[{}],并未安装,开始安装", chainCodeId);
                            InstallProposalRequest request = hfClient.newInstallProposalRequest();

                            request.setChaincodeID(chaincodeID);
                            request.setUserContext(hfClient.getUserContext());
                            request.setChaincodeSourceLocation(new File(blockChainConfiguration.getChainCodeRootDir()));
                            request.setChaincodeVersion(installVersion + "");
                            Collection<ProposalResponse> responses = hfClient.sendInstallProposal(request, Arrays.asList(peer));
                            for (ProposalResponse resp : responses)
                            {
                                if (resp.getStatus().getStatus() == 200)
                                {
                                    logger.info("该peer节点:{},成功安装该chaincode:{},version:{}", endorserPeer.getDomain(), chaincodeID, installVersion);
                                } else
                                {
                                    logger.error("该peer节点:{},安装该chaincode:{},version:{} 失败,原因:{}", endorserPeer.getDomain(), chaincodeID, installVersion, resp.getProposalResponse().getPayload());
                                    throw new ConfigException("安装链码失败:" + resp.getProposalResponse().getPayload());
                                }
                            }
                        }

                        logger.info("开始实例化该链码");
                        InstantiateProposalRequest instantiateProposalRequest = hfClient.newInstantiationProposalRequest();
                        instantiateProposalRequest.setProposalWaitTime(180000);
                        instantiateProposalRequest.setChaincodeID(chaincodeID);
                        instantiateProposalRequest.setChaincodeLanguage(TransactionRequest.Type.GO_LANG);
                        instantiateProposalRequest.setFcn("init");
                        instantiateProposalRequest.setArgs("a", "b", "c");

//                            Map<String, byte[]> tm = new HashMap<>();
//                            tm.put("HyperLedgerFabric", "InstantiateProposalRequest:JavaSDK".getBytes(UTF_8));
//                            tm.put("method", "InstantiateProposalRequest".getBytes(UTF_8));
//                            instantiateProposalRequest.setTransientMap(tm);
                        if (!StringUtils.isEmpty(chainCode.getPolicyFile()))
                        {
                            ChaincodeEndorsementPolicy chaincodeEndorsementPolicy = new ChaincodeEndorsementPolicy();
                            chaincodeEndorsementPolicy.fromYamlFile(new File(chainCode.getPolicyFile()));
                            instantiateProposalRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
                        }
//                    List<ChannelConfiguration> channelConfigurations = node.getChannelConfigurations();
//                    for (ChannelConfiguration channelConfiguration : channelConfigurations)
//                    {
                        String channelId = cid;
                        boolean ccHasInstaned = false;
                        logger.info("begin 判断该chaincode ,在该channelId=[{}]上是否已经被实例化", channelId);
                        List<Query.ChaincodeInfo> instantiatedChaincodes = channel.queryInstantiatedChaincodes(peer, user);
                        logger.info("在 该channel上查询到的实例化的chaincode数量为:" + instantiatedChaincodes.size());
                        for (Query.ChaincodeInfo instantiatedChaincode : instantiatedChaincodes)
                        {
                            if (instantiatedChaincode.getName().equalsIgnoreCase(chainCodeId))
                            {
                                ccHasInstaned = true;
                                continue;
                            }
                        }
                        if (ccHasInstaned)
                        {
                            logger.info("该chaincode已经被实例化");
                        } else
                        {
                            logger.info("该chaincode未被实例化,开始实例化");
                            logger.info("begin 往channel 请求实例化链码,channelId=[{}]", cid);
                            Collection<ProposalResponse> instantiationProposal = channel.sendInstantiationProposal(instantiateProposalRequest);
                            boolean success = true;
                            String msg = null;
                            for (ProposalResponse proposalResponse : instantiationProposal)
                            {
                                if (proposalResponse.getStatus().getStatus() != 200)
                                {
                                    success = false;
                                    msg += proposalResponse.getMessage();
                                }
                            }
                            if (success)
                            {
                                logger.info("成功实例化链码,往orderer发起交易请求");
                                CompletableFuture<BlockEvent.TransactionEvent> cf = channel.sendTransaction(instantiationProposal);
                                List<String> msgList = new ArrayList<>();
                                cf.thenApply((event) ->
                                {
                                    logger.info("链码实例化成功,并且往orderer发送了信息成功,接收到event,块编号为:" + event.getBlockEvent().getBlockNumber());
                                    return null;
                                }).exceptionally(e ->
                                {
                                    logger.error("链码实例化失败,msg=" + e.getMessage());
                                    msgList.add(e.getMessage());
                                    return null;
                                }).get();
                                if (msgList.isEmpty())
                                {
                                    logger.info("链码实例化成功,并且往orderer发送了信息成功");
                                } else
                                {
                                    logger.error("链码实例化失败:" + msgList.get(0));
                                    throw new ConfigException("链码实例化失败:" + msgList.get(0));
                                }
                            } else
                            {
                                throw new ConfigException("实例化链码失败:" + msg);
                            }
                        }
//                    }
                    }
                }
            }
        }
    }

    private void handlerResponse(Collection<ProposalResponse> responses, String notify)
    {
        for (ProposalResponse resp : responses)
        {
            if (resp.getStatus().getStatus() != 200)
            {
                logger.error("执行该proposal失败,失败原因为:" + resp);
                throw new RuntimeException(notify + ",执行错误,错误:" + resp.getMessage());
            }
        }
    }

    private void joinAllPeers(HFClient client, boolean tls, Channel channel, Collection<PeerConfiguration.PeerNode> peerNodes)
    {
        final HFClient hfClient = client;
        Peer peer = null;
        try
        {
            for (PeerConfiguration.PeerNode p : peerNodes)
            {
                String grpcUrl = tls ? "grpcs://" : "grpc://";
                grpcUrl += p.getIpWithPort();
                peer = hfClient.newPeer(p.getDomain(), grpcUrl, p.buildProperties());
                channel.joinPeer(peer);
            }
        } catch (InvalidArgumentException e)
        {
            throw new ConfigException("证书解析失败");
        } catch (ProposalException e)
        {
            throw new ConfigException("加入channel失败", e);
        }

    }

}
