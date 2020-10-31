package com.charlie.blockchain.proxy;

import com.charlie.blockchain.exception.InvokeException;
import com.charlie.blockchain.exception.QueryException;
import com.charlie.blockchain.model.*;

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

    QueryResp queryBlockChain(QueryReq req,List<IQueryOption>queryOptions)throws QueryException;

}
