package com.charlie.blockchain.configuration;

import com.charlie.exception.ConfigException;
import com.charlie.service.IValidator;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 06:49
 */
@Data
public class ChainCodeConfiguration implements IValidator
{
    private List<ChainCodeNode> chainCodes;

    @Override
    public void valid()
    {
        for (ChainCodeNode chainCode : chainCodes)
        {
            chainCode.valid();
        }
    }


    @Data
    public static class ChainCodeNode implements IValidator
    {
        private String chainCodeId;
        private String chainCodePath;
        private boolean needUpdate;
        private boolean needListOnBlockEvent;
        private String version;
        private String policyFile;

        @Override
        public void valid()
        {
            String prefixPath = ConfigurationFactory.getInstance().getBlockChainConfiguration().getPrefixPath();
            if (StringUtils.isNotEmpty(this.policyFile))
            {
                this.policyFile = prefixPath + this.policyFile;
            }
        }
    }

    protected void init() throws ConfigException
    {


    }


}
