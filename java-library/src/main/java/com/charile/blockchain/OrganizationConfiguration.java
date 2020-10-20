package com.charile.blockchain;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:57
 */
@Data
public class OrganizationConfiguration
{
    private List<OrganizationNode> nodes;

    @Data
    class OrganizationNode
    {
        private String organizationId;
        // 用户信息
        private List<UserNode> users;
        // peer信息
        private PeerConfiguration peerConfiguration;
        // ca信息
        private CaConfiguration caConfiguration;
    }


    @Data
    class UserNode
    {
        private String keyFile;
        private String certFile;
        private String userName;
        private boolean admin;
    }

}
