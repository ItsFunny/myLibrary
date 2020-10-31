package com.charlie.blockchain.model;

import com.charlie.blockchain.filter.IPeerFilter;
import com.charlie.service.IValidater;
import lombok.Data;

import java.util.ArrayList;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-26 08:50
 */
@Data
public class QueryReq implements IValidater
{
    private String channelName;
    private String chainCodeName;
    private String funcName;
    private ArrayList<String>args;
    private IPeerFilter peerFilter;

    @Override
    public void valid()
    {

    }
}
