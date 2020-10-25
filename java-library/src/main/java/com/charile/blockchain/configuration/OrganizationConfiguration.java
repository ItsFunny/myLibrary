package com.charile.blockchain.configuration;

import com.charile.base.IKeyImporter;
import com.charile.blockchain.model.UserInfo;
import com.charile.exception.ConfigException;
import com.charile.service.IValidater;
import com.charile.utils.FileUtils;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;

import java.io.File;
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
public class OrganizationConfiguration implements IValidater
{
    private OrganizationConfiguration() {}

    // 组织中的节点信息
    private List<OrganizationNode> organizations;

    @Override
    public void valid()
    {
//        boolean alphaExist = false;
//        for (OrganizationNode organizationNode : organizations)
//        {
//            if (organizationNode.isAlpha())
//            {
//                alphaExist = true;
//            }
//            organizationNode.valid();
//        }
//        if (!alphaExist)
//        {
//            throw new ConfigException("alpha组织必须存在");
//        }
        for (OrganizationNode organization : organizations)
        {
            organization.valid();
        }
    }


    @Data
    public static class OrganizationNode implements IValidater
    {
        private String mspId;
        //        private boolean alpha;
        // 用户信息
        private List<UserNode> users;
//        private List<String> channels;
        private List<String> peers;
        private String ca;

        // peer信息
        //        private PeerConfiguration peerConfiguration;
        // ca信息
        // 加入了哪些channel
//        private List<ChannelConfiguration> channelConfigurations;

//        public Collection<PeerConfiguration.PeerNode> getAllPeers()
//        {
//            Map<String, PeerConfiguration.PeerNode> peers = new HashMap();
//            PeerConfiguration.PeerNode anchorPeer = this.peerConfiguration.getAnchorPeer();
//            peers.put(anchorPeer.getDomain(), anchorPeer);
//            List<PeerConfiguration.EndorserPeer> endorserPeers = this.peerConfiguration.getEndorserPeers();
//            for (PeerConfiguration.EndorserPeer endorserPeer : endorserPeers)
//            {
//                if (!peers.containsKey(endorserPeer.getDomain()))
//                {
//                    peers.put(endorserPeer.getDomain(), endorserPeer);
//                }
//            }
//            return peers.values();
//        }
//
        public UserInfo getAdminUserInfo(List<IKeyImporter> importers)
        {
            UserNode adminUser = getAdminUser();
            byte[] keyBytes = null;
            byte[] certBytes = null;
            try
            {
                keyBytes = StringUtils.isNotEmpty(adminUser.getKeyString()) ? adminUser.getKeyString().getBytes() : FileUtils.readFileToByteArray(new File(adminUser.getKeyFile()));
                certBytes = StringUtils.isNotEmpty(adminUser.getCertString()) ? adminUser.getCertString().getBytes() : FileUtils.readFileToByteArray(new File(adminUser.getCertFile()));
            } catch (Exception e)
            {
                throw new RuntimeException(e);
            }

            for (IKeyImporter importer : importers)
            {
                UserInfo userInfo = null;
                try
                {
                    userInfo = new UserInfo(importer, this.mspId, adminUser.getName(), keyBytes, certBytes);
                } catch (Exception e)
                {
                    continue;
                }
                return userInfo;
            }
            throw new RuntimeException("找不到admin用户 | 无法导入证书和私钥");
        }

        public UserNode getAdminUser()
        {
            for (UserNode user : users)
            {
                if (user.isAdmin()) return user;
            }
            // 不可能为空
            throw new ConfigException("配置错误,找不到匹配的admin用户");
        }

//        public PeerConfiguration.PeerNode getAnchorPeer()
//        {
//            return this.peerConfiguration.getAnchorPeer();
//        }

        @Override
        public void valid()
        {
            for (UserNode user : users)
            {
                user.valid();
            }
            // 判断是否存在
//            if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().getChannelConfiguration().contains(this.channels))
//            {
//                throw new ConfigException("配置中的channel不匹配");
//            }
            if (!ConfigurationFactory.getInstance().getBlockChainConfiguration().containsPeers(this.getPeers()))
            {
                throw new ConfigException("配置中的peer不匹配");
            }

//            private List<String>peers;
//            private List<String>channels;
//            private String ca;

//            this.peerConfiguration.valid();
//            this.caConfiguration.valid();


//            for (ChannelConfiguration channelConfiguration : this.channelConfigurations)
//            {
//                channelConfiguration.valid();
//            }
        }
    }


    @Data
    public static class UserNode implements IValidater
    {
        private String keyFile;
        private String certFile;

        private String keyString;
        private String certString;
        // 2选1


        private String name;
        // 默认为true
        private boolean admin = true;


        @Override
        public void valid()
        {
            if (StringUtils.isEmpty(this.getCertString()) && StringUtils.isEmpty(this.getCertFile()))
            {
                throw new ConfigException("certFile 和 certString 不可同时为空");
            }
            if (StringUtils.isEmpty(this.getKeyString()) && StringUtils.isEmpty(this.getKeyFile()))
            {
                throw new ConfigException("keyFile 和 keyString 不可同时为空");
            }
            if (StringUtils.isEmpty(this.getName()))
            {
                throw new ConfigException("用户名为空");
            }
        }
    }
}
