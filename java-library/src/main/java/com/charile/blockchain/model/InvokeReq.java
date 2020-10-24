package com.charile.blockchain.model;

import lombok.Data;

import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 09:34
 */
@Data
public class InvokeReq
{
    private String channelName;
    private String chainCodeName;
    private String funcName;
    private List<String> args;
    // 哪个handler处理
    private Byte handlerType;
}
