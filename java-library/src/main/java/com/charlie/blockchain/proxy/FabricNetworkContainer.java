package com.charlie.blockchain.proxy;

import com.charlie.base.AbstractInitOnce;
import com.charlie.base.IKeyImporter;
import com.charlie.blockchain.ResultInfo;
import com.charlie.blockchain.configuration.ChannelConfiguration;
import com.charlie.blockchain.configuration.*;
import com.charlie.blockchain.model.NewClientReq;
import com.charlie.blockchain.model.Pair;
import com.charlie.blockchain.model.UserInfo;
import com.charlie.blockchain.util.ConvertUtils;
import com.charlie.blockchain.util.ProposalUtils;
import com.charlie.blockchain.wrapper.HFCaClientWrapper;
import com.charlie.blockchain.wrapper.HFClientWrapper;
import com.charlie.exception.ConfigException;
import com.charlie.utils.Base64Utils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.protos.peer.Query;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric.sdk.security.CryptoSuiteFactory;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.io.IOException;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

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
    private HFCaClientWrapper caClientWrapper;
    private static final Logger logger = LoggerFactory.getLogger(FabricNetworkContainer.class);

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
        BlockChainConfiguration blockChainConfiguration = ConfigurationFactory.getInstance().getBlockChainConfiguration();
        if (blockChainConfiguration.isDeploy())
        {
            this.deployNetwork(blockChainConfiguration);
        } else
        {
            this.initSelf(blockChainConfiguration);
        }
        this.initHFCaClient(blockChainConfiguration);
    }

    private void deployNetwork(BlockChainConfiguration blockChainConfiguration)
    {
        try
        {
            // 初始化区块链网络
            logger.info("开始初始化区块链网络");
            HFClient alphaClient = this.createNewHfClient();
            this.clientWrapper = new HFClientWrapper();
            blockChainConfiguration.initOnce();
            try
            {
                this.initlizeNetwork(alphaClient, blockChainConfiguration);
            } catch (InvalidArgumentException e)
            {
                throw new ConfigException(e, ResultInfo.NETWORK_CREATE_PEER_ERROR);
            }

            List<ChannelConfiguration.ChannelNode> channels = blockChainConfiguration.getChannelConfiguration().getChannels();
            for (ChannelConfiguration.ChannelNode c : channels)
            {
                Channel channel = alphaClient.getChannel(c.getChannelId());
                logger.info("begin 初始化channelId=[{}]的channel", c.getChannelId());
                channel.initialize();
                logger.info("end 初始化channelId=[{}]的channel", c.getChannelId());
            }

            this.installAndInstantiateCC(alphaClient, blockChainConfiguration);
            this.afterInitNetwork(blockChainConfiguration, alphaClient);
            // 需要初始化channel
            alphaClient = this.clientWrapper.getAlphaClient();
        } catch (Exception e)
        {
            throw new ConfigException(e);
        }
    }

    private void afterInitNetwork(BlockChainConfiguration blockChainConfiguration, HFClient alphaClient) throws Exception
    {
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

        UserInfo alphaUser = blockChainConfiguration.getAlphaUser();
        instance.setUserContext(alphaUser);
        this.clientWrapper = new HFClientWrapper();
        this.clientWrapper.setAlphaClient(instance);
        this.clientWrapper.put(alphaUser.getMspId(), alphaUser);

        return instance;
    }


    private void initlizeNetwork(HFClient client, BlockChainConfiguration blockChainConfiguration) throws InvalidArgumentException
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
            List<ChannelConfiguration.ChannelNode> channels = blockChainConfiguration.getChannelsByPeers(node.getPeers());
//            List<String> channels = node.getChannels();
            logger.info("该组织{}加入的channel数量为:{}", organizationId, channels.size());
            for (ChannelConfiguration.ChannelNode channelNode : channels)
            {
                String channelId = channelNode.getChannelId();
                logger.info("获取channel的所有peer");
                // 获取该组织的所有配置着的peer,通通加入
                Collection<PeerConfiguration.PeerChannelBO> allPeers = blockChainConfiguration.getOrganizationAllPeersByChannel(channelId, node.getMspId());
                logger.info("获取channel的所有orderer");
                List<Orderer> allOrderers = blockChainConfiguration.getChannelAllOrderers(channelId, hfClient);
                Channel channel = client.getChannel(channelId);
                // 需要先get 再new ,不然会报重复的错误
                if (channel == null)
                {
                    channel = client.newChannel(channelId);
                    channel.decorate(Arrays.asList((IDataDecorator<Channel>) channel1 ->
                    {
                        channel1.setType(channelNode.getType());
                        return channel1;
                    }));
                }
                for (PeerConfiguration.PeerChannelBO peerNode : allPeers)
                {
                    channel.addPeer(peerNode.conv2TPeer(blockChainConfiguration.isTls(), hfClient));
                }
                for (Orderer orderer : allOrderers)
                {
                    channel.addOrderer(orderer);
                }


                PeerConfiguration.PeerNode anchorPeer = blockChainConfiguration.getOrganizationAnchorPeer(node.getMspId());
                logger.info("anchorPeer:{}", anchorPeer);
                Peer peer = SDKUtils.getSpecificPeer(channel, anchorPeer.getDomain());

                logger.info("begin 查询该组织[{}]已经加入的channel信息", channelId);
                Set<String> channelStrings = null;
                try
                {
                    channelStrings = hfClient.queryChannels(peer);
                } catch (ProposalException e)
                {
                    throw new ConfigException(e, ResultInfo.INVOKE_FAILED);
                }
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
//                Channel channel = null;


                if (channelCreated)
                {
                    logger.info("该channel:{}已经创建", channelId);
//                    // 需要先get 再new ,不然会报重复的错误
//                    channel = client.getChannel(channelId);
//                    if (channel == null)
//                    {
//                        channel = client.newChannel(channelId);
//                        channel.decorate(Arrays.asList((IDataDecorator<Channel>) channel1 ->
//                        {
//                            channel1.setType(channelNode.getType());
//                            return channel1;
//                        }));
//                    }
//                    for (PeerConfiguration.PeerNode peerNode : allPeers)
//                    {
//                        channel.addPeer(peerNode.conv2TPeer(blockChainConfiguration.isTls(), hfClient));
//                    }
//                    for (Orderer orderer : allOrderers)
//                    {
//                        channel.addOrderer(orderer);
//                    }
                } else
                {
                    logger.info("该channel 可能未创建,所以开始准备创建,组织id:{},chanelId:{}", node.getMspId(), channelId);
                    logger.info("begin 创建channel");
//                    ChannelConfiguration.ChannelNode channelNode = blockChainConfiguration.getChannelNode(cid);
                    org.hyperledger.fabric.sdk.ChannelConfiguration configuration = null;
                    try
                    {
                        configuration = new org.hyperledger.fabric.sdk.ChannelConfiguration(new File(channelNode.getChannelConfigPath()));
                    } catch (IOException e)
                    {
                        throw new ConfigException(new IllegalArgumentException(e), ResultInfo.NETWORK_MAYBE_READ_FILE_ERROR);
                    } catch (InvalidArgumentException e)
                    {
                        throw new ConfigException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
                    }
                    byte[] channelConfigurationSignature = new byte[0];
                    channelConfigurationSignature = client.getChannelConfigurationSignature(configuration, adminUserInfo);

                    logger.info("签名信息为:" + Base64Utils.encode(channelConfigurationSignature));
                    boolean joined = false;
                    try
                    {
                        channel = hfClient.newChannel(channelId, allOrderers.get(0), configuration, channelConfigurationSignature);
                        channelCreated = true;
                    } catch (Exception e)
                    {
                        logger.info("可能是其他组织已经创建了该channel,所以尝试直接加入");
                        channel = client.getChannel(channelId);
                        try
                        {
                            tryJoinAllPeersAfterAlreadCreated(hfClient, blockChainConfiguration.isTls(), channel, allPeers);
                        } catch (ProposalException e1)
                        {
                            logger.error("尝试加入失败,直接退出,抛出错误:" + e.getMessage());
                            throw new ConfigException(e, ResultInfo.UNKNOWN);
                        }
                        joined = true;
                    }
                    if (!joined)
                    {
                        try
                        {
                            joinAllPeers(client, blockChainConfiguration.isTls(), channel, allPeers);
                        } catch (ProposalException e)
                        {
                            throw new ConfigException(e, ResultInfo.NETWORK_JOIN_CHANNEL_ERROR);
                        }
                    } else
                    {
                        logger.info("channel已由其他组织创建,该组织{}的peer加入成功");
                    }
                    try
                    {
                        if (channelCreated)
                        {
                            TimeUnit.SECONDS.sleep(8);
                        }
                    } catch (InterruptedException e)
                    {
                        throw new ConfigException(e, ResultInfo.CONCURRENT_ERROR);
                    }

                    logger.info("将所有的peer都加入到channel中");

                }
            }

        }
    }

    @Data
    class ChaincodeInstallFlager
    {
        private List<String> peers;
        private Double latestVersion;

        public ChaincodeInstallFlager(List<String> peers)
        {
            this.peers = peers;
        }
    }

    // FIXME 需要单独的抽成接口
    // 支持当chaincode升级的时候只需要升级一次,并且其他的peer会同步安装最新的版本,并且,升级的话,默认第一个peer会升级,后面的就不需要
    // 再次升级了
    private void installAndInstantiateCC(HFClient client, BlockChainConfiguration blockChainConfiguration) throws Exception
    {
        Map<String, List<String>> nodeChainCodeMapList = new HashMap<>();
        logger.info("begin 安装并且实例化链码");
        final HFClient hfClient = client;
        OrganizationConfiguration organizationConfiguration = blockChainConfiguration.getOrganizationConfiguration();
        List<OrganizationConfiguration.OrganizationNode> organizationNodes = organizationConfiguration.getOrganizations();
        for (OrganizationConfiguration.OrganizationNode node : organizationNodes)
        {
            List<ChannelConfiguration.ChannelNode> channelsByPeers = blockChainConfiguration.getChannelsByPeers(node.getPeers());

            for (ChannelConfiguration.ChannelNode channelNode : channelsByPeers)
            {
                String cid = channelNode.getChannelId();
                logger.info("begin 对组织[{}]进行相关操作", node.getMspId());
                User user = this.clientWrapper.getUser(node.getMspId());
                hfClient.setUserContext(user);
                List<String> peers = node.getPeers();
                List<PeerConfiguration.PeerNode> endorserPeers = blockChainConfiguration.getOrgAllEndorserPeers(peers);
                // 同一个channel下的
                // 如果说chaincode 已经升级过了或者实例化过了,则不需要进行再次升级
                for (PeerConfiguration.PeerNode endorserPeer : endorserPeers)
                {
                    Channel channel = hfClient.getChannel(cid);
                    Collection<Peer> channelPeers = channel.getPeers();
                    // FIXME ,获取该组织下的要安装该chaincode的所有peer,既 chaincode在外层遍历,而peer在内层遍历
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
                        // 可能存在相同的chaincode安装在不同的channel中的peer中,
                        // 所以需要找到该chaincode+channel的所有peer
                        // 其中一个peer 实例化|升级过了,就不需要升级,只需要安装最新版本即可
                        // 1. 判断 该chaincode所属的peer中有没有
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
                                    if (!StringUtils.isEmpty(chainCode.getVersion()))
                                    {
                                        double configVersion = Double.parseDouble(chainCode.getVersion());
                                        installVersion = configVersion <= prevVersion ? prevVersion + 0.01 : configVersion;
                                    } else
                                    {
                                        installVersion = installVersion <= prevVersion ? prevVersion + 0.01 : installVersion;
                                    }
                                    logger.info("该chaincode已经安装,同时需要升级,newVersion=[{}],prevVersion=[{}]", installVersion, prevVersion);
                                }
                            }
                        }

                        ChaincodeID.Builder chaincodeIDBuilder = ChaincodeID.newBuilder().setName(chainCodeId).setVersion(installVersion + "")
                                .setPath(chainCode.getChainCodePath());
                        ChaincodeID chaincodeID = chaincodeIDBuilder.build();

                        // FIXME 升级有问题,问题在于 就是需要判断实例化的最大版本是否大于当前版本
                        // 以及需要获取到该chaincode 在该channel中最新的版本,避免peer1装了1.1之后,就不执行了
                        // 导致下次安装的时候,peer1位1.2 而peer2位1.1
                        if (ccHasInstall)
                        {
                            if (!needUpgrade)
                            {
                                logger.info("该chaincode=[{}],不需要升级", chainCodeId);
                            } else
                            {
//                                logger.info("该chaincode需要升级,因此需要安装新版本,再升级");
//                                InstallProposalRequest request = hfClient.newInstallProposalRequest();
//                                request.setChaincodeID(chaincodeID);
//                                request.setUserContext(hfClient.getUserContext());
//                                request.setChaincodeSourceLocation(new File(blockChainConfiguration.getChainCodeRootDir()));
//                                request.setChaincodeVersion(installVersion + "");
//                                Collection<ProposalResponse> responses = hfClient.sendInstallProposal(request, Arrays.asList(peer));
//                                handlerResponse(responses, "链码升级前安装");
//                                UpgradeProposalRequest upgradeProposalRequest = hfClient.newUpgradeProposalRequest();
//                                upgradeProposalRequest.setChaincodeID(chaincodeID);
//                                upgradeProposalRequest.setArgs("a", "b", "c");
////                            upgradeProposalRequest.setProposalWaitTime();
//                                upgradeProposalRequest.setFcn("init");
//                                upgradeProposalRequest.setUserContext(user);
//                                Collection<ProposalResponse> proposalResponses = channel.sendUpgradeProposal(upgradeProposalRequest, Arrays.asList(peer));
//                                handlerResponse(proposalResponses, "链码升级");
//
//                                Pair pair = ProposalUtils.parseResp(proposalResponses);
//                                Collection<ProposalResponse> successResponse = pair.getSuccessResponse();
//                                Collection<ProposalResponse> failResponse = pair.getFailResponse();
//                                logger.info("成功:[{}],失败:[{}]", successResponse.size(), failResponse.size());
//                                if (successResponse.size() < failResponse.size())
//                                {
//                                    throw new ConfigException("成功的小于失败的");
//                                }
//                                send2Orderer("升级链码", channel, successResponse);
                            }
                        } else
                        {
                            logger.info("该chaincodeid=[{}],并未安装,开始安装", chainCodeId);
                            InstallProposalRequest request = hfClient.newInstallProposalRequest();
                            // FIXME 可能需要查询安装该chaincode的所有peer,然后统一安装,实例化的时候只需要实例化一个就可以
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
                        if (!StringUtils.isEmpty(chainCode.getPolicyFile()))
                        {
                            ChaincodeEndorsementPolicy chaincodeEndorsementPolicy = new ChaincodeEndorsementPolicy();
                            chaincodeEndorsementPolicy.fromYamlFile(new File(chainCode.getPolicyFile()));
                            instantiateProposalRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
                        }
                        String channelId = cid;
                        boolean ccHasInstaned = false;
                        // peer是否需要更新安装到最新的版本
                        boolean peerNeedUpdateInstallCC = false;
                        double updateInstallVersion = 0.01;
                        logger.info("begin 判断该chaincode ,在该channelId=[{}]上是否已经被实例化", channelId);
                        List<Query.ChaincodeInfo> instantiatedChaincodes = channel.queryInstantiatedChaincodes(peer, user);
                        logger.info("在 该channel上查询到的实例化的chaincode数量为:" + instantiatedChaincodes.size());
                        for (Query.ChaincodeInfo instantiatedChaincode : instantiatedChaincodes)
                        {
                            // 实例化的版本号,如果需要升级,则会判断当前已经实例化的链码是否版本比要升级的链码低,低的话升级
                            double instantiatedVersion = Double.parseDouble(instantiatedChaincode.getVersion());

                            if (needUpgrade && instantiatedVersion > installVersion)
                            {
                                logger.info("当前已经实例化的链码为{},比安装的链码版本要高,所以不需要升级链码,但是需要安装链码的当前最新版本", instantiatedVersion, installVersion);
                                needUpgrade = false;
                                peerNeedUpdateInstallCC = true;
                                updateInstallVersion = instantiatedVersion;
                            }

                            if (instantiatedChaincode.getName().equalsIgnoreCase(chainCodeId))
                            {
                                ccHasInstaned = true;
                            }
                        }

                        if (ccHasInstaned)
                        {
                            logger.info("该chaincode已经被实例化");
                        } else
                        {
                            logger.info("该chaincode未被实例化,开始实例化");
                            logger.info("begin 往channel 请求实例化链码,channelId=[{}]", cid);
                            List<Peer> instanPeers = Arrays.asList(peer);
                            Collection<ProposalResponse> instantiationProposal = channel.sendInstantiationProposal(instantiateProposalRequest, instanPeers);
                            int successCount = 0;
                            String msg = null;
                            List<ProposalResponse> successful = new ArrayList<>();
                            for (ProposalResponse proposalResponse : instantiationProposal)
                            {
                                if (proposalResponse.getStatus().getStatus() != 200)
                                {
                                    msg += proposalResponse.getMessage();
                                } else
                                {
                                    successCount++;
                                    successful.add(proposalResponse);
                                }
                            }
                            if (successCount == instanPeers.size())
                            {
                                logger.info("成功实例化链码,往orderer发起交易请求");
                                // FIXME 抽出来
                                CompletableFuture<BlockEvent.TransactionEvent> cf = channel.sendTransaction(successful,null);
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

                        logger.info("查询需要装该chaincode的所有peer,如果已经装了一个,则剩下的不需要再安装");


                        if (!needUpgrade)
                        {
                            logger.info("该chaincode=[{}],不需要升级", chainCodeId);
                            if (peerNeedUpdateInstallCC)
                            {
                                logger.info("该peer的chaincode不需要升级,但是需要同步更新安装的版本,当前最新版本为:" + updateInstallVersion);
                                InstallProposalRequest request = hfClient.newInstallProposalRequest();
                                request.setChaincodeID(chaincodeID);
                                request.setUserContext(hfClient.getUserContext());
                                request.setChaincodeSourceLocation(new File(blockChainConfiguration.getChainCodeRootDir()));
                                request.setChaincodeVersion(updateInstallVersion + "");
                                Collection<ProposalResponse> responses = hfClient.sendInstallProposal(request, Arrays.asList(peer));
                                handlerResponse(responses, "peer同步链码安装版本");
                            }
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

                            Pair pair = ProposalUtils.parseResp(proposalResponses);
                            Collection<ProposalResponse> successResponse = pair.getSuccessResponse();
                            Collection<ProposalResponse> failResponse = pair.getFailResponse();
                            logger.info("成功:[{}],失败:[{}]", successResponse.size(), failResponse.size());
                            if (successResponse.size() < failResponse.size())
                            {
                                throw new ConfigException("成功的小于失败的");
                            }
                            send2Orderer("升级链码", channel, successResponse);
                        }
                    }
                }
            }
        }
    }

    private void send2Orderer(String msg, Channel channel, Collection<ProposalResponse> successful) throws ExecutionException, InterruptedException
    {
        CompletableFuture<BlockEvent.TransactionEvent> cf = channel.sendTransaction(successful,null);
        List<String> msgList = new ArrayList<>();
        cf.thenApply((event) ->
        {
            logger.info("并且往orderer发送了信息成功,接收到event,块编号为:" + event.getBlockEvent().getBlockNumber());
            return null;
        }).exceptionally(e ->
        {
            logger.error("发送orderer失败,msg=" + e.getMessage());
            msgList.add(e.getMessage());
            return null;
        }).get();
        if (msgList.isEmpty())
        {
            logger.info("发送orderer成功,并且往orderer发送了信息成功");
        } else
        {
            logger.error(msg + ":" + msgList.get(0));
            throw new ConfigException("链码实例化失败:" + msgList.get(0));
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

    // 在已经创建过了的话,直接尝试全部加入
    private void tryJoinAllPeersAfterAlreadCreated(HFClient client, boolean tls, Channel channel, Collection<PeerConfiguration.PeerChannelBO> peerNodes) throws ProposalException, InvalidArgumentException
    {
        joinAllPeers(client, tls, channel, peerNodes);
    }

    private void joinAllPeers(HFClient client, boolean tls, Channel channel, Collection<PeerConfiguration.PeerChannelBO> peerNodes) throws InvalidArgumentException, ProposalException
    {
        final HFClient hfClient = client;
        Peer peer = null;
        for (PeerConfiguration.PeerChannelBO p : peerNodes)
        {
            String grpcUrl = p.getUrl();
            peer = hfClient.newPeer(p.getDomain(), grpcUrl, p.buildProperties());
            channel.joinPeer(peer);
        }

    }

    private void initHFCaClient(BlockChainConfiguration blockChainConfiguration) throws ConfigException
    {
        if (blockChainConfiguration.getCaConfiguration() == null)
        {
            return;
        }
        this.caClientWrapper = new HFCaClientWrapper();
        HFClient alphaClient = this.clientWrapper.getAlphaClient();
        String mspId = alphaClient.getUserContext().getMspId();
        CryptoSuite cryptoSuite = alphaClient.getCryptoSuite();
        try
        {
            CaConfiguration.CaNode caNodeByMspId = blockChainConfiguration.getCaNodeByMspId(mspId);
            HFCAClient hfcaClient = HFCAClient.createNewInstance(caNodeByMspId.getCaName(), caNodeByMspId.getUrl(), caNodeByMspId.buildProperties());
            hfcaClient.setCryptoSuite(cryptoSuite);
            this.caClientWrapper.setCaClient(hfcaClient);
        } catch (Exception e)
        {
            throw new ConfigException("配置hfCaClient失败:" + e.getMessage(), e);
        }
    }

    private void initSelf(BlockChainConfiguration blockChainConfiguration)
    {
        logger.info("只是实例化channel等信息");
        try
        {
            HFClient alphaClient = this.buildAlphaHFClient(blockChainConfiguration);


            ChannelConfiguration channelConfiguration = blockChainConfiguration.getChannelConfiguration();
            List<ChannelConfiguration.ChannelNode> channels = channelConfiguration.getChannels();

            for (ChannelConfiguration.ChannelNode channelNode : channels)
            {
                Channel channel = alphaClient.newChannel(channelNode.getChannelId());
                channel.decorate(Arrays.asList((IDataDecorator<Channel>) channel1 ->
                {
                    channel1.setType(channelNode.getType());
                    return channel1;
                }));
                List<ChannelConfiguration.ChannelOrderInfo> orderInfos = channelNode.getOrderers();
                List<String> orderers = orderInfos.stream().map(o -> o.getDomain()).collect(Collectors.toList());
                List<Orderer> channelAllOrderers = blockChainConfiguration.getChannelAllOrderers(channelNode.getChannelId(), alphaClient);
                logger.info("查询该channelId=[{}]的所有peer,结果为{}条记录", channelNode.getChannelId(), channelAllOrderers.size());
                for (Orderer o : channelAllOrderers)
                {
                    channel.addOrderer(o);
                }
                logger.info("查询该channelId=[{}]的所有peer", channelNode.getChannelId());
                List<PeerConfiguration.PeerChannelBO> channelAllPeers = blockChainConfiguration.getChannelAllPeers(channelNode.getChannelId());
                List<Peer> peers = ConvertUtils.convPeerNodes2TPeers(channelAllPeers, blockChainConfiguration.isTls(), alphaClient);
                for (Peer peer : peers)
                {
                    channel.addPeer(peer);
                }
                channel.initialize();
            }
        } catch (Exception e)
        {
            throw new ConfigException(e);
        }
    }


}
