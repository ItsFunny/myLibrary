package com.charlie.blockchain.model;

import com.charlie.base.IKeyImporter;
import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-25 05:49
 */
@Data
public class NewClientReq
{
    private String name;
    private String keyBytes;
    private String certBytes;

    private IKeyImporter keyImporter;
    private String mspId;
}
