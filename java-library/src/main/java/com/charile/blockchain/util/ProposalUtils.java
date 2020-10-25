package com.charile.blockchain.util;

import com.charile.blockchain.constants.ResultConstants;
import com.charile.blockchain.model.Pair;
import com.charile.blockchain.wrapper.ChannelContext;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 10:05
 */
public class ProposalUtils
{
//    public static Pair parseResp(Collection<ProposalResponse> propResp, List<ProposalResponse> successful,
//                                 List<ProposalResponse> failed, List<byte[]> resList, Channel channel, String channelName,
//                                 String chaincodeName, String funcName, ChannelContext channelContext)
//            throws InvalidArgumentException, IOException
//    {
//        int statusCode = 200;
//        byte[] res = null;
//        String desc = "";
//        for (ProposalResponse proposalResponse : propResp)
//        {
//            if (proposalResponse.isVerified() && proposalResponse.getStatus().equals(ProposalResponse.Status.SUCCESS))
//            {
//                res = proposalResponse.getChaincodeActionResponsePayload();
//            } else
//            {
//                desc += proposalResponse.getMessage();
//            }
//        }
////        return parseResp(propResp, successful, failed, resList, channel, channelName, chaincodeName, funcName, channelContext.getSequenceId(),
////                channelContext);
//        return new Pair(statusCode, res, desc);
//    }

    public static Pair parseResp(Collection<ProposalResponse> propResp)
            throws InvalidArgumentException, IOException
    {
        int statusCode = ResultConstants.SUCCESS;
        byte[] res = null;
        String desc = "";
        List<ProposalResponse>successful=new ArrayList<>();
        List<ProposalResponse>fail=new ArrayList<>();
        for (ProposalResponse proposalResponse : propResp)
        {
            if (proposalResponse.isVerified() && proposalResponse.getStatus().equals(ProposalResponse.Status.SUCCESS))
            {
                res = proposalResponse.getChaincodeActionResponsePayload();
                successful.add(proposalResponse);
            } else
            {
                statusCode=ResultConstants.FAIL;
                desc += proposalResponse.getMessage();
                fail.add(proposalResponse);
            }
        }
        Pair pair = new Pair(statusCode, res, desc);
//        return parseResp(propResp, successful, failed, resList, channel, channelName, chaincodeName, funcName, channelContext.getSequenceId(),
//                channelContext);
        pair.setSuccessResponse(successful);
        pair.setFailResponse(fail);
        return pair;
    }

}
