package com.charile.blockchain.configuration;

import com.charile.base.AbstractInitOnce;
import com.charile.blockchain.constants.ConfigConstants;
import com.charile.exception.ConfigException;
import com.charile.utils.SystemUtils;
import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 06:30
 */
@Data
public class ConfigurationFactory extends AbstractInitOnce
{
    private BlockChainConfiguration blockChainConfiguration;

    private ConfigurationFactory()
    {
        this.initOnce();
    }

    @Override
    protected void init() throws ConfigException
    {
        BlockChainConfiguration blockChainConfiguration = SystemUtils.parseConfigFile(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, BlockChainConfiguration.class);
        this.blockChainConfiguration = blockChainConfiguration;

//         加载配置文件
//        ConfigurationFactory blockchain = SystemUtils.parseConfigFile(ConfigConstants.CONFIG_BLOCKCHAIN_CONFIGURAITION, this.getClass());
//        this.copy(blockchain);
    }

    public void slowInit() throws ConfigException
    {
        this.blockChainConfiguration.initOnce();
    }

    private void copy(ConfigurationFactory configurationFactory) throws ConfigException
    {
        this.blockChainConfiguration = configurationFactory.getBlockChainConfiguration();
        this.blockChainConfiguration.initOnce();
    }

    private static class ConfigurationHolder
    {
        private static ConfigurationFactory INSTANCE = new ConfigurationFactory();

        private ConfigurationHolder()
        {

        }
    }

    public static ConfigurationFactory getInstance()
    {
        return ConfigurationFactory.ConfigurationHolder.INSTANCE;
    }
}
