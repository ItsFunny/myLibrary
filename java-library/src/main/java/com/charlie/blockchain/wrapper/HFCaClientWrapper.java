package com.charlie.blockchain.wrapper;

import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.User;
import org.hyperledger.fabric_ca.sdk.HFCAClient;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 00:35
 */
public class HFCaClientWrapper
{
    // caClient: 该组织的caClient
    private HFCAClient caClient;

    public HFCAClient getCaClient()
    {
        return caClient;
    }

    public void setCaClient(HFCAClient caClient)
    {
        this.caClient = caClient;
    }
}
