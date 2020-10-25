package com.charile.blockchain.model;

import lombok.Data;
import org.hyperledger.fabric.sdk.ProposalResponse;

import java.util.Collection;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 10:05
 */
@Data
public class Pair
{
    int statusCode;
    byte[] value;
    String description;

    // 成功的交易
    private Collection<ProposalResponse>successResponse;
    // 失败的交易
    private Collection<ProposalResponse>failResponse;

    public Pair(int statusCode, byte[] value, String description) {
        this.statusCode = statusCode;
        this.value = value;
        this.description = description;
    }

}
