package com.charlie.blockchain.proxy;

import com.charlie.blockchain.configuration.ConfigurationFactory;
import com.charlie.blockchain.constants.ConfigConstants;
import com.charlie.blockchain.model.InvokeReq;
import com.charlie.blockchain.model.InvokeResp;
import com.charlie.blockchain.model.QueryReq;
import com.charlie.blockchain.model.QueryResp;
import com.charlie.utils.SystemUtils;
import org.apache.commons.lang3.RandomUtils;
import org.junit.Test;

import java.util.ArrayList;
import java.util.concurrent.TimeUnit;


public class FabricNetworkContainerTest
{
    private String channelName = "demochannel";
    private String chaincodeName = "democc";
    private String setFunc = "setvalue";
    private String getFunc = "getvalue";

    @Test
    public void testCreateChannel() throws Exception
    {

        SystemUtils.setEnviroment(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, "/Users/joker/go/src/myLibrary/java-library/src/main/resources/config_47_103_samecc.json");
//        SystemUtils.setEnviroment(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, "/Users/joker/go/src/myLibrary/java-library/src/main/resources/config_47_103_samecc.json");
        ConfigurationFactory.getInstance().slowInit();
        IFabricClientService fabricClientService = new DefaultFabricClientHandler();
        try
        {
            InvokeResp resp = fabricClientService.invokeBlockChain(buildInvokeReq(), null);
            byte[] returnBytes = resp.getReturnBytes();
            System.out.println(new String(returnBytes));
        } catch (Exception e)
        {

        }

        TimeUnit.HOURS.sleep(3);
    }

    @Test
    public void testQuery() throws Exception
    {
        SystemUtils.setEnviroment(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, "/Users/joker/company/go/src/github.com/hyperledger/fabric-demo/bidsun-fabric-benchmark/src/main/resources/config_47_103_samecc.json");
        ConfigurationFactory.getInstance().slowInit();
        IFabricClientService fabricClientService = new DefaultFabricClientHandler();

        QueryResp queryResp = fabricClientService.queryBlockChain(buildQueryReq(), null);

        String s = new String(queryResp.getReturnBytes());
        System.out.println(s);
    }

    private QueryReq buildQueryReq()
    {
        QueryReq req = new QueryReq();
        ArrayList<String> arrayList = new ArrayList<>();
        req.setArgs(arrayList);
        req.setChainCodeName(chaincodeName);
        req.setChannelName(channelName);
        req.setFuncName(getFunc);
        req.setPeerFilter((p) ->
        {
            if (p.getName().contains("org1"))
            {
                return false;
            }
            return true;
        });
        return req;
    }

    private InvokeReq buildInvokeReq()
    {
        InvokeReq req = new InvokeReq();
        ArrayList<String> arrayList = new ArrayList<>();
        arrayList.add(RandomUtils.nextLong() + "");
        arrayList.add(+RandomUtils.nextLong() + "");
        req.setArgs(arrayList);
        req.setChainCodeName(chaincodeName);
        req.setFuncName(setFunc);
        req.setChannelName(channelName);
        return req;
    }

}