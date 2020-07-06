/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-06 14:15 
# @File : tx.go
# @Description : 
# @Attention : 
*/
package model

type TransactionGetByIdReq struct {
	// 通过交易id获取详情信息
	NeedArgs  bool   `json:"needArgs"`
	TxID      string `json:"txId"`
	ChannelId string `json:"channelId"`
}

type  TransactionGetByIdResp struct {
	// 块高度
	BlockHeight uint64
	// 区块ID
	BlockHash string
	// 交易类型
	Signature string
}