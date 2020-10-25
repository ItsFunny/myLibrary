package com.charile.blockchain.proxy;

import com.charile.blockchain.configuration.ConfigurationFactory;
import com.charile.blockchain.constants.ConfigConstants;
import com.charile.blockchain.model.InvokeReq;
import com.charile.blockchain.model.InvokeResp;
import com.charile.utils.SystemUtils;
import org.apache.commons.lang3.RandomUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.junit.Test;
import org.springframework.context.annotation.ConfigurationCondition;

import java.util.ArrayList;
import java.util.concurrent.TimeUnit;


public class FabricNetworkContainerTest
{
    private String channelName="demochannel";
    private String chaincodeName="democc";
    private String setFunc="setvalue";
    private String getFunc="getvalue";
    @Test
    public void testCreateChannel() throws Exception
    {
        SystemUtils.setEnviroment(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, "/Users/joker/go/src/myLibrary/java-library/src/main/resources/config_47_103_samecc_just_use.json");
        ConfigurationFactory.getInstance().slowInit();
        IFabricClientService fabricClientService=new DefaultFabricClientHandler();
        InvokeResp resp = fabricClientService.invokeBlockChain(buildInvokeReq(), null);
        byte[] returnBytes = resp.getReturnBytes();
        System.out.println(new String(returnBytes));

        TimeUnit.HOURS.sleep(3);
    }

    private InvokeReq buildInvokeReq(){
        InvokeReq req=new InvokeReq();
        ArrayList<String>arrayList=new ArrayList<>();
        arrayList.add(RandomUtils.nextLong()+"");
        req.setArgs(arrayList);
        req.setChainCodeName(chaincodeName);
        req.setFuncName(setFunc);
        req.setChannelName(channelName);
        return req;
    }

}