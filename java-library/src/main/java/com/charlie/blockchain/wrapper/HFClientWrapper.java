package com.charlie.blockchain.wrapper;

import lombok.Data;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.User;

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
@Data
public class HFClientWrapper
{
    // alphaClient,这个组织的client 加入了所有的channel,用于监管所有channel
    private HFClient alphaClient;

    public HFClientWrapper(String orgMsp, User user)
    {
        this.userContextMap = new HashMap<>();
//        this.clientCache.put(orgMsp,us);
        this.userContextMap.put(orgMsp, user);
    }


    public ConcurrentHashMap<String, HFClient> clientCache = new ConcurrentHashMap<>();

    // 各个组织的context
    private Map<String, User> userContextMap = new HashMap<>();

    public HFClientWrapper() { this.userContextMap = new HashMap<>(); }

    public  void put(String mspId, User userInfo){
        this.userContextMap.put(mspId,userInfo);
    }
    public User getUser(String mspId){
        return this.userContextMap.get(mspId);
    }
}
