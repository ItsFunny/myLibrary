package com.charlie.blockchain.proxy;

import cn.bidsun.blockchain.model.RegisterEnrollReq;
import com.charlie.blockchain.ResultInfo;
import com.charlie.blockchain.configuration.CaConfiguration;
import com.charlie.blockchain.configuration.ConfigurationFactory;
import com.charlie.blockchain.constants.ResultConstants;
import com.charlie.blockchain.exception.CAException;
import com.charlie.blockchain.exception.EnrollException;
import com.charlie.blockchain.exception.InvokeException;
import com.charlie.blockchain.exception.QueryException;
import com.charlie.blockchain.filter.IPeerFilter;
import com.charlie.blockchain.model.*;
import com.charlie.blockchain.util.KeyUtils;
import com.charlie.blockchain.util.ProposalUtils;
import com.google.common.base.Stopwatch;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.util.Base64Util;
import org.hyperledger.fabric_ca.sdk.EnrollmentRequest;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.hyperledger.fabric_ca.sdk.RegistrationRequest;
import org.hyperledger.fabric_ca.sdk.exception.EnrollmentException;
import org.hyperledger.fabric_ca.sdk.exception.RegistrationException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.security.PrivateKey;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.concurrent.*;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 11:06
 */
public class DefaultFabricClientHandler extends AbstractFabricClient
{
    private static final Logger logger = LoggerFactory.getLogger(DefaultFabricClientHandler.class);

    @Override
    protected void init() throws Exception
    {

    }

    @Override
    public Boolean validIsMine(Byte type)
    {
        return null;
    }

    @Override
    public InvokeResp invokeBlockChain(InvokeReq req, List<IInvokeOption> invokeOptions) throws InvokeException
    {
        req.valid();
        HFClient client = getClient();
        User userContext = req.getUserContext();
        TransactionProposalRequest transactionProposalRequest = client.newTransactionProposalRequest();
        req.to(transactionProposalRequest);

        if (userContext != null && userContext.getEnrollment() != null)
        {
            transactionProposalRequest.setUserContext(userContext);
        }
        transactionProposalRequest.setArgs(req.getArgs());
        logger.info("开始发送交易信息");

        Channel channel = client.getChannel(req.getChannelName());
        Collection<Peer> channelPeers = channel.getPeers();
        if (null != req.getPeerFilter())
        {
            final IPeerFilter peerFilter = req.getPeerFilter();
            List<Peer> peers = new ArrayList<>();
            for (Peer channelPeer : channelPeers)
            {
                if (!peerFilter.filter(channelPeer))
                {
                    peers.add(channelPeer);
                }
            }
            channelPeers = peers;
        }
        Stopwatch stopwatch = Stopwatch.createStarted();

        Collection<ProposalResponse> proposalResponses = null;
        try
        {
            proposalResponses = channel.sendTransactionProposal(transactionProposalRequest, channelPeers);
        } catch (ProposalException e)
        {
            throw new InvokeException("交易失败", e);
        } catch (InvalidArgumentException e)
        {
            throw new IllegalArgumentException("参数错误:" + e.getMessage());
        }

        stopwatch.stop();
        logger.info("交易发送完毕,耗时:" + stopwatch.elapsed(TimeUnit.SECONDS) + "s");
        Pair pair = null;
        try
        {
            pair = ProposalUtils.parseResp(proposalResponses);
            if (pair.getStatusCode() != ResultConstants.SUCCESS)
            {
                logger.error("交易执行失败,msg:" + pair.getDescription());
                throw new InvokeException(ResultInfo.INVOKE_FAILED.getCode(), pair.getDescription());
            }
        } catch (InvalidArgumentException e)
        {
            throw new InvokeException(ResultInfo.ILLEGAL_ARGUMENT_ERROR.getCode(), e.getMessage());
        } catch (IOException e)
        {
            throw new InvokeException(e, ResultInfo.UNKNOWN);
        }
        Collection<ProposalResponse> successResponse = pair.getSuccessResponse();
        Collection<ProposalResponse> failResponse = pair.getFailResponse();
        logger.info("成功的交易数为:{},失败的交易数为:{}", successResponse.size(), failResponse.size());

        try
        {
            BlockEvent.TransactionEvent transactionEvent = sendTransaction(channel, successResponse, userContext);
        } catch (InterruptedException e)
        {
            throw new InvokeException(e, ResultInfo.UNKNOWN);
        } catch (ExecutionException e)
        {
            throw new InvokeException(e, ResultInfo.UNKNOWN);
        } catch (TimeoutException e)
        {
            throw new InvokeException(e, ResultInfo.INVOKE_TIME_OUT);
        }

        InvokeResp resp = new InvokeResp();
        resp.setReturnBytes(pair.getValue());
        return resp;
    }

    private BlockEvent.TransactionEvent sendTransaction(Channel channel, Collection<ProposalResponse> successful, User userContext) throws InterruptedException, ExecutionException, TimeoutException
    {
        // 等待时间
        logger.info("开是发送交易信息给orderer节点");
        CompletableFuture<BlockEvent.TransactionEvent> future = null;
        if (userContext == null)
        {
            future = channel.sendTransaction(successful);
        } else
        {
            future = channel.sendTransaction(successful, userContext);
        }
        BlockEvent.TransactionEvent transactionEvent = future.thenApply(event ->
        {
            logger.info("收到交易事件,区块编号为:{}", event.getBlockEvent().getBlockNumber());
            return event;
        }).exceptionally(e ->
        {
            logger.error("交易失败:{}", e.getMessage());
            return null;
        }).get(5, TimeUnit.SECONDS);
        return transactionEvent;
    }

    @Override
    public Future<InvokeResp> invokeBlockChainAsync(InvokeReq req, List<IInvokeOption> options) throws InvokeException
    {


        return null;
    }

    @Override
    public InstallChaincodeResp installChainCode(InstallChaincodeReq req)
    {
        return null;
    }

    @Override
    public QueryResp queryBlockChain(QueryReq req, List<IQueryOption> queryOptions) throws QueryException
    {
        req.valid();
        String channelName = req.getChannelName();
        String chainCodeName = req.getChainCodeName();
        String funcName = req.getFuncName();
        ArrayList<String> args = req.getArgs();

        HFClient client = getClient();
        Channel channel = client.getChannel(channelName);
        if (channel == null)
        {
            throw new QueryException(ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        }
        ChaincodeID chaincodeID = ChaincodeID.newBuilder().setName(chainCodeName).build();
        // 创建交易提案请求
        QueryByChaincodeRequest request = client.newQueryProposalRequest();
        request.setChaincodeID(chaincodeID);
        request.setFcn(funcName);
        request.setArgs(args);
        Collection<Peer> queryPeers = channel.getPeers();
        if (req.getPeerFilter() != null)
        {
            List<Peer> peers = new ArrayList<>(queryPeers.size());
            final IPeerFilter peerFilter = req.getPeerFilter();
            for (Peer channelPeer : queryPeers)
            {
                if (!peerFilter.filter(channelPeer))
                {
                    peers.add(channelPeer);
                }
                continue;
            }
            queryPeers = peers;
        }
        Collection<ProposalResponse> proposalResponses = null;
        try
        {
            proposalResponses = channel.queryByChaincode(request, queryPeers);
        } catch (InvalidArgumentException e)
        {
            throw new QueryException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        } catch (ProposalException e)
        {
            throw new QueryException(e, ResultInfo.QUERY_FAILED);
        }
        Pair pair = null;
        try
        {
            pair = ProposalUtils.parseResp(proposalResponses);
        } catch (InvalidArgumentException e)
        {
            throw new QueryException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        } catch (IOException e)
        {
            throw new QueryException(e, ResultInfo.UNKNOWN);
        }
        if (pair.getStatusCode()!=ResultConstants.SUCCESS)
        {
            throw new QueryException(ResultInfo.UNKNOWN.getCode(),"未抛出异常,但是查询失败");
        }
        Collection<ProposalResponse> successResponse = pair.getSuccessResponse();
        Collection<ProposalResponse> failResponse = pair.getFailResponse();
        logger.info("成功的次数:{},失败的次数:{}", successResponse.size(), failResponse.size());

        QueryResp result = new QueryResp();
        result.setReturnBytes(pair.getValue());
        return result;
    }

    @Override
    public RegisterEnrollResp registerAndEnroll(RegisterEnrollReq req, List<IRegisterOption> registerOptions, List<IEnrollOption> enrollOptions) throws CAException
    {
        RegisterEnrollResp result = new RegisterEnrollResp();
        final HFCAClient caClient = getCaClient();
        final HFClient client = getClient();

        CaConfiguration.CaNode caNodeByName = ConfigurationFactory.getInstance().getBlockChainConfiguration().getCaNodeByMspId(req.getMspId());
        // 默认不为空
        Enrollment adminEnroll = null;
        try
        {
            adminEnroll = caClient.enroll(caNodeByName.getAdminUserName(), caNodeByName.getAdminPassword());
        } catch (EnrollmentException e)
        {
            throw new EnrollException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        } catch (org.hyperledger.fabric_ca.sdk.exception.InvalidArgumentException e)
        {
            throw new EnrollException(e, ResultInfo.CA_ENROLL_ERROR);
        }


        cn.bidsun.blockchain.model.UserRegisterReq userRegisterReq = new cn.bidsun.blockchain.model.UserRegisterReq();
        userRegisterReq.setPassword(req.getPassword());
        userRegisterReq.setUserName(req.getUserName());
        UserRegisterResp register = register(caClient, client, userRegisterReq, registerOptions, adminEnroll);
        result.setPassword(register.getPassword());
        result.setUserName(req.getUserName());

        EnrollReq enrollReq = new EnrollReq();
        enrollReq.setPassword(register.getPassword());
        enrollReq.setUserName(req.getUserName());
        try
        {
            cn.bidsun.blockchain.model.EnrollResp enroll = enroll(caClient, client, enrollReq, enrollOptions);
            result.setPrivateKeyPem(enroll.getPrivateKeyPem());
            result.setSignCertPem(enroll.getSignCertPem());
        } catch (org.hyperledger.fabric_ca.sdk.exception.InvalidArgumentException e)
        {
            throw new EnrollException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        }

        return result;
    }

    public UserRegisterResp register(HFCAClient caClient, HFClient client, cn.bidsun.blockchain.model.UserRegisterReq req, List<IRegisterOption> registerOptions, Enrollment adminEnrollMent)
    {
        UserRegisterResp result = new UserRegisterResp();
        RegistrationRequest request = null;
        try
        {
            request = new RegistrationRequest(req.getUserName());
            // 默认 admin
            request.setType("admin");
            if (!StringUtils.isEmpty(req.getPassword()))
            {
                request.setSecret(req.getPassword());
            }
        } catch (Exception e)
        {
            throw new CAException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        }
        UserInfo newUserInfo = new UserInfo(req.getUserName(), req.getAffiliation(), req.getMspId(), adminEnrollMent);
        try
        {
            result.setPassword(caClient.register(request, newUserInfo));
        } catch (RegistrationException e)
        {
            throw new CAException(e, ResultInfo.CA_REGISTER_ERROR);
        } catch (org.hyperledger.fabric_ca.sdk.exception.InvalidArgumentException e)
        {
            throw new CAException(e, ResultInfo.ILLEGAL_ARGUMENT_ERROR);
        }

        return result;
    }

    public cn.bidsun.blockchain.model.EnrollResp enroll(HFCAClient caClient, HFClient client, EnrollReq req, List<IEnrollOption> enrollOptions) throws org.hyperledger.fabric_ca.sdk.exception.InvalidArgumentException
    {
        req.valid();
        final EnrollmentRequest enrollmentRequest = new EnrollmentRequest();
//        KeyPair keyPair = null;
//        try
//        {
//            keyPair = req.getKeyPairGenerator().generateKeyPair();
//        } catch (Exception e)
//        {
//            throw new CAException(e, ResultInfo.KEYPAIR_GENERATE_ERROR);
//        }
//        enrollmentRequest.setKeyPair(keyPair);
//        enrollmentRequest.setProfile(req.getProfile());

        Enrollment enroll = null;
        try
        {
            enroll = caClient.enroll(req.getUserName(), req.getPassword(), enrollmentRequest);
        } catch (EnrollmentException e)
        {
            throw new CAException(e, ResultInfo.CA_ENROLL_ERROR);
        }
        String signCert = enroll.getCert();
        PrivateKey key = enroll.getKey();
        byte[] bytes = KeyUtils.formatPrvKey(key.getEncoded());
        String encode = Base64Util.encode(bytes);
        System.out.println("返回的私钥为:" + encode);
        System.out.println("签发证书为:" + signCert);
        cn.bidsun.blockchain.model.EnrollResp enrollResp = new cn.bidsun.blockchain.model.EnrollResp();
        enrollResp.setSignCertPem(signCert);
        enrollResp.setPrivateKeyPem(encode);
//        try
//        {
//            enrollResp.setPrivateKeyPem(KeyUtils.convKey2TPem(key));
//        } catch (IOException e)
//        {
//            throw new EnrollException(e, ResultInfo.CA_ENROLL_ERROR);
//        }
        // FIXME key 转换成标准的32字节的


        return enrollResp;
    }
    private HFCAClient getCaClient()
    {
        return FabricNetworkContainer.getInstance().getCaClientWrapper().getCaClient();
    }


    private HFClient getClient()
    {
        return FabricNetworkContainer.getInstance().getClientWrapper().getAlphaClient();
    }
}
