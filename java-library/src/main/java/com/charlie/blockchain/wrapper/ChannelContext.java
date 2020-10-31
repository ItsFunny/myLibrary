package com.charlie.blockchain.wrapper;

import lombok.Data;
import org.hyperledger.fabric.sdk.Peer;

import java.util.Collection;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 10:06
 */
@Data
public class ChannelContext
{
    private String sequenceId;
    private boolean waitAllPeerResponse = false;
    private Collection<Peer> peers;
    private String channelName;
    private String chainCodeName;
    private String funcName;
    private long beginTime = System.currentTimeMillis();


}
