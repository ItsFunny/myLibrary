/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-14 16:37 
# @File : base.go
# @Description : 
# @Attention : 
*/
package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"myLibrary/go-library/go/authentication"
	"myLibrary/go-library/go/constants"
	"myLibrary/go-library/go/converters"
)

type TransBaseType int
type ChannelID string
type OrganizationID string
type ChainCodeID string
type MethodName string
// 代表区块链上的key
type Key string
// fromWalletAddress 从哪个钱包过来的
type From string
// toWalletAddress 去往哪个钱包的
type To string
// token 交易coin
type Token float64
type Version uint64

type TransBaseTypeV2 authentication.Authority
type TransBaseTypeV2Value authentication.AuthValue

func (this TransBaseTypeV2) String() string {
	return hex.EncodeToString(this.BigEndianConvtBytes())
}

func (this TransBaseTypeV2) BigEndianConvtBytes() []byte {
	return authentication.Authority(this).BigEndianConvt2Bytes()
}
// 用于存储返回值公有的一些属性
type BaseRespCommonAttribute struct {
	// 存储的是本次交易的主目标 :如 作品上链
	BaseType      TransBaseTypeV2 `json:"baseType"`
	TxDescription string          `json:"txDescription"`
	// 在哪条链上
	ChannelID ChannelID `json:"channelId"`
	// 2020-02-27
	// 10:23 add
	// 交易ID
	TransactionID string `json:"transactionId"`
}

// 区块链调用的返回值
type BaseFabricResp struct {
	// 具体的业务返回值
	DataBytes []byte `json:"dataBytes"`
	// 代表code
	CodeBytes []byte `json:"codeBytes"`
	// 代表消息 ,直接String强转即可
	MsgBytes []byte `json:"msgBytes"`
	// 遗留字节,可能有其他用处
	OtherBytes []byte `json:"otherBytes"`
	// 日志字节,包含了是否需要记录等信息
	LogBytes []byte `json:"logBytes"`
	//
}
type TxRecordInfoNode struct {
	From          From            `json:"from"`
	To            To              `json:"to"`
	Token         Token           `json:"token"`
	BaseType      TransBaseTypeV2 `json:"baseType"`
	TxDescription string          `json:"txDescription"`
	// 在哪条链上
	ChannelID ChannelID `json:"channelId"`
	// 2020-02-27
	// 10:23 add
	// 交易ID
	TransactionID string `json:"transactionId"`
}

type ResultInfo struct {
	// code 要么 与success 做了或运算代表成功,
	// 要么与 fail 做了或运算代表逻辑失败,至于是否显示抛出取决于是否与out_put做了或运算
	LogicCode  int    `json:"code"`
	LogicMsg   string `json:"msg"`
	From       From   `json:"from"`
	To         To     `json:"to"`
	Token      Token  `json:"token"`
	DataBytes  []byte `json:"data"`
	OtherBytes []byte `json:"otherBytes"`
	// 2020-03-16 13:16 add
	// 为了后续业务的变更,这里统一返回数组的形式返回,而非返回单个数据的形式,便于后续维护
	TxRecords []*TxRecordInfoNode `json:"txRecords"`
}

func (r ResultInfo) String() string {
	str := "code=[%d],msg=[%s],from=[%v],to=[%v],token=[%v],data=[%s],other=[%s]"
	return fmt.Sprintf(str, r.LogicCode, r.LogicMsg, r.From, r.To, r.Token, string(r.DataBytes), string(r.OtherBytes))
}

func (r ResultInfo) Success() bool {
	return (r.LogicCode & constants.SUCCESS) > 0
}

// "from":"13Yrapjusm3ic9ncVGrCHXiNJPXJpwMQCD",
// "to":"",
// "token":0,
// "Data":{"dna":"840aa0eb-96f5-4960-914c-1b313570ccca","prvKey":"","coinAddress":"1MxWsr2qXr63JVfLghR3u5E3asJP7tutr9"},
// "Code":1,"Msg":"SUCCESS
func (r BaseFabricResp) Convt2ResultInfo() ResultInfo {
	res := ResultInfo{}
	res.LogicCode = int(converter.BigEndianBytes2Int64(r.CodeBytes))
	res.LogicMsg = string(r.MsgBytes)
	res.OtherBytes = r.OtherBytes
	res.DataBytes = r.DataBytes
	json.Unmarshal(r.LogBytes, &res.TxRecords)

	return res
}
func (this TxRecordInfoNode) String() string {
	rStr := ""
	rStr += " [ "
	rStr += " baseType: %s  ,"
	rStr += " ChannelID=%s , "
	rStr += "FROM =%s ,"
	rStr += " TO=%s ,"
	rStr += "Token=%f ,"
	rStr += "TransactionID=%s,"
	rStr += " TxDescription=%s"
	rStr += " ] \n"
	rStr = fmt.Sprintf(rStr, this.BaseType.String(), this.ChannelID, this.From, this.To, this.Token, this.TransactionID, this.TxDescription)
	return rStr
}

// 用于底层的logicService的返回值
type ServiceLogicBaseResp struct {
	LogicCode int    `json:"code"`
	LogicMsg  string `json:"msg"`
	LogBytes  []byte `json:"logBytes"`

	CommAttribute BaseRespCommonAttribute
}

type ChainBaseReq struct {
	MethodName  MethodName
	ChannelID   ChannelID
	ChainCodeID ChainCodeID
}

// 通过交易id获取详情信息
type BSBlockChainGetTransactionDetailReqBO struct {
	NeedArgs  bool   `json:"needArgs"`
	TxID      string `json:"txId"`
	ChannelId string `json:"channelId"`
}
