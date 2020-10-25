package com.charile.blockchain.util;

import com.charile.blockchain.configuration.PeerConfiguration;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Peer;

import java.util.List;
import java.util.stream.Collectors;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 21:56
 */
public class ConvertUtils
{
    public static List<Peer> convPeerNodes2TPeers(List<PeerConfiguration.PeerNode> nodes, boolean tls, HFClient client)
    {
        return  nodes.stream().map(c->c.conv2TPeer(tls,client)).collect(Collectors.toList());
    }

}
