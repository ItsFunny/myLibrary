package com.charile.blockchain.proxy;

import com.charile.blockchain.exception.InvokeException;
import com.charile.blockchain.model.InstallChaincodeReq;
import com.charile.blockchain.model.InstallChaincodeResp;
import com.charile.blockchain.model.InvokeReq;
import com.charile.blockchain.model.InvokeResp;

import java.util.List;
import java.util.concurrent.Future;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-10-24 11:06
 */
public class DefaultFabricClientHandler extends AbstractFabricClient
{

    @Override
    protected void init() throws Exception
    {

    }

    @Override
    public Boolean validIsMine(Byte type)
    {
        return null;
    }

    @Override
    public InvokeResp invokeBlockChain(InvokeReq req, List<IInvokeOption> invokeOptions) throws InvokeException
    {
        return null;
    }

    @Override
    public Future<InvokeResp> invokeBlockChainAsync(InvokeReq req, List<IInvokeOption> options) throws InvokeException
    {
        return null;
    }

    @Override
    public InstallChaincodeResp installChainCode(InstallChaincodeReq req)
    {
        return null;
    }
}
