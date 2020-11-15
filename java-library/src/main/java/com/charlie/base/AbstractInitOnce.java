package com.charlie.base;

import com.charlie.blockchain.configuration.BlockChainConfiguration;
import com.charlie.blockchain.configuration.ChannelConfiguration;
import com.charlie.blockchain.configuration.PeerConfiguration;
import com.charlie.blockchain.util.ConvertUtils;
import com.charlie.exception.ConfigException;
import lombok.Data;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:13
 */
@Data
public abstract class AbstractInitOnce implements IInitOnce
{
    private  boolean init;

    @Override
    public void initOnce() throws ConfigException
    {
        if (init)
        {
            return;
        }
        this.init();
        init = true;
    }

    protected abstract void init() throws ConfigException;

}
