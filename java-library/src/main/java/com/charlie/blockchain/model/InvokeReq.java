package com.charlie.blockchain.model;

import com.charlie.blockchain.filter.IPeerFilter;
import com.charlie.service.IValidator;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.ChaincodeID;
import org.hyperledger.fabric.sdk.TransactionProposalRequest;
import org.hyperledger.fabric.sdk.User;

import java.util.ArrayList;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 09:34
 */
@Data
public class InvokeReq implements IValidator
{
    private String channelName;
    private String chainCodeName;
    private String funcName;
    private ArrayList<String> args;

    private IPeerFilter peerFilter;
    private User userContext;
    // 哪个handler处理
    private Byte handlerType;

    public void to(TransactionProposalRequest request){
        request.setArgs(this.getArgs());
        ChaincodeID chaincodeID = ChaincodeID.newBuilder().setName(this.getChainCodeName()).build();
        request.setChaincodeID(chaincodeID);
        request.setFcn(this.funcName);
    }
    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(channelName))
        {
            throw new IllegalArgumentException("参数错误,channelName不可为空");
        }
        if (StringUtils.isEmpty(chainCodeName))
        {
            throw new IllegalArgumentException("参数错误,channelName不可为空");
        }
        if (StringUtils.isEmpty(funcName))
        {
            throw new IllegalArgumentException("参数错误,channelName不可为空");
        }


    }
}
