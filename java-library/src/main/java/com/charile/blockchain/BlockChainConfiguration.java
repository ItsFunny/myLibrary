package com.charile.blockchain;

import com.charile.base.AbstractInitOnce;
import lombok.Data;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.InstallProposalRequest;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description 以channel为主,channel中有多少个组织,对应多少个chaincode等信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:48
 */

@Data
public class BlockChainConfiguration extends AbstractInitOnce
{
    private String version;
    private List<ChannelConfiguration> channelConfiguration;


    public static void main(String[] args) throws Exception
    {
        HFClient hfClient = HFClient.createNewInstance();
        Channel channel = hfClient.newChannel("");
        InstallProposalRequest installProposalRequest = hfClient.newInstallProposalRequest();

    }

    @Override
    protected void init() throws Exception
    {

    }
}
