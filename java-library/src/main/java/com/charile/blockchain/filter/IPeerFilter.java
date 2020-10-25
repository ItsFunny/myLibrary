package com.charile.blockchain.filter;

import org.hyperledger.fabric.sdk.Peer;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 09:55
 */
public interface  IPeerFilter
{
    boolean filter(Peer peer);

}
