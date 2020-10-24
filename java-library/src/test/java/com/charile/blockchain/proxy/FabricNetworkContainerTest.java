package com.charile.blockchain.proxy;

import com.charile.blockchain.configuration.ConfigurationFactory;
import com.charile.blockchain.constants.ConfigConstants;
import com.charile.utils.SystemUtils;
import org.hyperledger.fabric.sdk.HFClient;
import org.junit.Test;
import org.springframework.context.annotation.ConfigurationCondition;


public class FabricNetworkContainerTest
{
    @Test
    public void testCreateChannel() throws Exception
    {
        SystemUtils.setEnviroment(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, "/Users/joker/go/src/myLibrary/java-library/src/main/resources/config.json");
        ConfigurationFactory.getInstance().slowInit();

        HFClient client = FabricNetworkContainer.getInstance().getClient();

    }

}