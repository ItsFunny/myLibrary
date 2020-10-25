package com.charile.blockchain.proxy;

import com.charile.blockchain.ResultInfo;
import com.charile.blockchain.constants.ResultConstants;
import com.charile.blockchain.exception.InvokeException;
import com.charile.blockchain.filter.IPeerFilter;
import com.charile.blockchain.model.*;
import com.charile.blockchain.util.ProposalUtils;
import com.google.common.base.Stopwatch;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.xml.transform.Result;
import java.io.IOException;
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
    private Logger logger = LoggerFactory.getLogger(DefaultFabricClientHandler.class);

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


    private HFClient getClient()
    {
        return FabricNetworkContainer.getInstance().getClientWrapper().getAlphaClient();
    }
}
