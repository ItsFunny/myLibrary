package com.charile.blockchain;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-20 23:07
 */
@Data
public class OrdererConfiguration
{
    private List<OrdererNode> nodes;

    @Data
    class OrdererNode
    {
        private String ip;
        private Integer port;
        private String ordererId;
    }
}
