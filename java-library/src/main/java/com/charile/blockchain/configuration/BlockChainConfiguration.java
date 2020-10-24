package com.charile.blockchain.configuration;

import com.charile.base.AbstractInitOnce;
import com.charile.base.IKeyImporter;
import com.charile.blockchain.model.UserInfo;
import com.charile.exception.ConfigException;
import com.charile.service.IValidater;
import com.charile.utils.FileUtils;
import com.charile.utils.SystemUtils;
import lombok.Data;
import lombok.extern.java.Log;
import org.apache.commons.lang3.StringUtils;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.HFClient;
import org.hyperledger.fabric.sdk.InstallProposalRequest;

import java.io.File;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description 以channel为主, channel中有多少个组织, 对应多少个chaincode等信息
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:48
 */

@Data
@Log
public class BlockChainConfiguration extends AbstractInitOnce implements IValidater
{
    private String version;
    private boolean alpha;
    // 是否启用tls
    private boolean tls;

    private String cryptoConfigPrefixPath;
    // tx 文件路径前缀
    private String artifactsPrefixPath;
    //
    private String chainCodeRootDir;

    // 以组织为主
    private OrganizationConfiguration organizationConfiguration;

    // orderer的配置
    private List<OrdererConfiguration> ordererConfigurations;


    private BlockChainConfiguration()
    {

    }


    @Override
    protected void init() throws ConfigException
    {
        this.valid();
    }

    public UserInfo getAlphaUser(){
        List<OrganizationConfiguration.OrganizationNode> organizationNodes = this.organizationConfiguration.getOrganizationNodes();
        for (OrganizationConfiguration.OrganizationNode organizationNode : organizationNodes)
        {
            if (organizationNode.isAlpha())
            {
                return organizationNode.getAdminUserInfo(Arrays.asList(IKeyImporter.COMMON_PEM_KEY_IMPORTER, IKeyImporter.STANDARD_SM2_KEY_IMPORTER));
            }
        }
        throw new ConfigException("不可能会调到这里");
    }

    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(this.version)) this.version = "1";
        if (StringUtils.isEmpty(this.artifactsPrefixPath))
        {
            throw new ConfigException("artifactsPrefixPath 不可为空");
        }
        this.artifactsPrefixPath = FileUtils.appendFilePathIfNone(this.artifactsPrefixPath);
        if (StringUtils.isEmpty(this.cryptoConfigPrefixPath))
        {
            throw new ConfigException("cryptoConfigPrefixPath 不可为空");
        }
        this.cryptoConfigPrefixPath = FileUtils.appendFilePathIfNone(this.cryptoConfigPrefixPath);
        this.organizationConfiguration.valid();
    }
}
