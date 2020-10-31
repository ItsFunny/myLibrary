package com.charlie.blockchain.configuration;

import com.charlie.exception.ConfigException;
import com.charlie.service.IValidater;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 06:43
 */
@Data
public class ClientConfiguration implements IValidater
{
    // 这个客户端以哪个为主
    private String organizationMspId;

    @Override
    public void valid()
    {
        if (StringUtils.isEmpty(organizationMspId))
        {
            throw new ConfigException("组织id不可为空");
        }
    }

    @Data
    public static class ConnectionConfiguration
    {

    }

}
