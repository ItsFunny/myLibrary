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
 * @Date 创建时间：2020-10-23 22:58
 */
public interface IFabricClientService
{
    InvokeResp invokeBlockChain(InvokeReq req, List<IInvokeOption> invokeOptions) throws InvokeException;

    Future<InvokeResp> invokeBlockChainAsync(InvokeReq req, List<IInvokeOption> options) throws InvokeException;

    InstallChaincodeResp installChainCode(InstallChaincodeReq req);

}
