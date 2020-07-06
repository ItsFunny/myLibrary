/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-25 17:34 
# @File : resdp.go
# @Description : 
# @Attention : 
*/
package base

import (
	"encoding/json"
	"myLibrary/go-library/go/converters"
)

type ITxBaseResper interface {
	// 用于后续的钩子记录
	GetTXRecordInfoList() []*TxRecordInfoNode
	// 用于获取内部数据
	GetReturnData() interface{}

	// 用于获取返回code,判断是否成功
	GetCode() int
	GetMsg() string

	// 2020-02-24
	// 15:37 update
	// 设置相关的 tx信息
	SetTxBaseType(baseType TransBaseTypeV2)
	SetTxDescription(d string)

	// 2020-03-16
	// 14:57 add
	SetChannelID(c ChannelID)
	SetTransactionID(tid string)

	// 2020-04-17
	// 16:32 add 用于对为空的交易信息补全
	InfoFix()
	// 21:29 用于获取一些共有属性,原因在于 返回web的所有tx 记录为空,该返回和Infofix 就是为这些而存在的
	GetCommAttribute() BaseRespCommonAttribute
}


func NewBaseFabricResp(code int, msg string) *BaseFabricResp {
	r := new(BaseFabricResp)
	r.CodeBytes = converter.BigEndianInt642Bytes(int64(code))
	r.MsgBytes = []byte(msg)
	return r
}


// 用于临时传输数据
type TempTransfer struct {
	Code       int
	Msg        string
	TxRecords  []*TxRecordInfoNode
	ReturnData interface{}

	BaseRespCommonAttribute BaseRespCommonAttribute
}



func Success() *BaseFabricResp {
	resp := NewSuccessBaseFabricResp()
	return resp
}

func SuccessWithEmptyData() []byte {
	bytes, _ := json.Marshal(*Success())
	return bytes
}