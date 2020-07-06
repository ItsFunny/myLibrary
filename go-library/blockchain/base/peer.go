/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-25 17:31 
# @File : peer.go
# @Description : 
# @Attention : 
*/
package base

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"myLibrary/go-library/go/constants"
	error2 "myLibrary/go-library/blockchain/error"
	error3 "myLibrary/go-library/go/error"
	"myLibrary/go-library/go/utils"
)
type BasePeerResponse struct {
	peer.Response
}


func SuccessPeerResponse(bytes []byte) BasePeerResponse {
	return SuccessWithDetail(bytes, nil, constants.SUCCESS, "SUCCESS")
}
func Fail(e error3.IBaseError) BasePeerResponse {
	resp := NewBaseFabricResp(int(e.GetCode()), e.GetMsg())
	bbs, _ := json.Marshal(resp)
	r := BasePeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}

func SuccessWithDetail(bytes []byte, logBytes []byte, code int, msg string) BasePeerResponse {
	resp := *NewBaseFabricResp(code, msg)
	resp.DataBytes = bytes
	resp.LogBytes = logBytes
	bbs, _ := json.Marshal(resp)
	r := BasePeerResponse{
		Response: shim.Success(bbs),
	}
	return r
}

func SuccessWithDetailTransfer(transfer TempTransfer) BasePeerResponse {
	resp := *NewBaseFabricResp(transfer.Code, transfer.Msg)
	if transfer.ReturnData != nil {
		dataBytes, e := json.Marshal(transfer.ReturnData)
		if nil != e {
			return Fail(error2.NewJSONSerializeError(e, "ReturnData"))
		}
		resp.DataBytes = dataBytes
	}

	if transfer.TxRecords != nil && len(transfer.TxRecords) > 0 {
		logBytes, e := json.Marshal(transfer.TxRecords)
		if nil != e {
			return Fail(error2.NewJSONSerializeError(e, "TxRecords"))
		}
		resp.LogBytes = logBytes
	}

	otherBytes, e := json.Marshal(transfer.BaseRespCommonAttribute)
	if nil != e {
		return Fail(error2.NewJSONSerializeError(e, "BaseRespCommonAttribute"))
	}
	resp.OtherBytes = otherBytes

	bytes, e := json.Marshal(resp)
	if nil != e {
		return Fail(error2.NewJSONSerializeError(e, "BaseFabricResp"))
	}

	utils.DebugPrintDetail("*", "返回值", string(bytes))
	r := BasePeerResponse{
		Response: shim.Success(bytes),
	}

	return r

}
