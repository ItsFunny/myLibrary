package com.charlie.blockchain.configuration;

import com.charlie.exception.ConfigException;
import com.charlie.service.IValidator;
import org.springframework.util.CollectionUtils;
import org.springframework.util.StringUtils;

import java.util.List;
import java.util.Properties;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:03
 */
public class CaConfiguration implements IValidator
{
    private CaConfiguration() {}

    private List<CaNode> caNodes;

    @Override
    public void valid()
    {
        if (CollectionUtils.isEmpty(this.caNodes))
        {
            throw new ConfigException("ca节点信息不可为空");
        }
        for (CaNode caNode : caNodes)
        {


        }

    }

    public List<CaNode> getCaNodes()
    {
        return caNodes;
    }

    public void setCaNodes(List<CaNode> caNodes)
    {
        this.caNodes = caNodes;
    }

    public static class CaNode implements IValidator
    {
        // 可以认为与组织的mspid相同
        private String caName;
        private String url;
        private String adminUserName;
        private String adminPassword;

        public String getAdminUserName()
        {
            return adminUserName;
        }

        public void setAdminUserName(String adminUserName)
        {
            this.adminUserName = adminUserName;
        }

        public String getAdminPassword()
        {
            return adminPassword;
        }

        public void setAdminPassword(String adminPassword)
        {
            this.adminPassword = adminPassword;
        }

        public String getCaName()
        {
            return caName;
        }

        public void setCaName(String caName)
        {
            this.caName = caName;
        }

        public String getUrl()
        {
            return url;
        }

        public void setUrl(String url)
        {
            this.url = url;
        }

        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.caName))
            {
                throw new ConfigException("ca的mspId不可为空");
            }
            if (StringUtils.isEmpty(this.url))
            {
                throw new ConfigException("ca的url不可为空");
            }
        }

        public Properties buildProperties(){
            return null;
        }
    }
}
