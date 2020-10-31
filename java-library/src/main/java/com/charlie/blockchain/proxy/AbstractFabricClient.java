package com.charlie.blockchain.proxy;

import com.charlie.base.IInitOnce;
import lombok.Data;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-23 23:02
 */
@Data
public abstract class AbstractFabricClient implements IInitOnce, IFabricClientService
{
    private static boolean init;

    private AbstractFabricClient nextHandler;

    @Override
    public void initOnce() throws Exception
    {
        if (this.init)
        {
            return;
        }
        this.init();
        this.init = true;
    }

    protected abstract void init() throws Exception;

    public abstract Boolean validIsMine(Byte type);
}
