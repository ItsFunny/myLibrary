package com.charile.blockchain;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 22:58
 */
@Data
public class PeerConfiguration
{
    private List<PeerNode> anchorPeers;
    private List<EndorserPeer> endorserPeers;

    @Data
    class PeerNode
    {
        private String ip;
        private Integer port;
    }

    @Data
    class EndorserPeer extends PeerNode
    {
        private List<ChaincodeNode> chainCodes;
    }

    @Data
    class ChaincodeNode
    {
        private String chainCodeId;
        private String chainCodePath;
        private boolean needUpdate;
        private boolean needListOnBlockEvent;
        private String version;
        private Byte policyType;
        private String policy;
    }


}
