/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:14 
# @File : block.go
# @Description : 
# @Attention : 
*/
package handler

import (
	"myLibrary/go-library/go/blockchain"
)

type LogBlockHandler struct {
	*config.BaseLogicServiceImpl
}
func NewLogBlockHandler()*LogBlockHandler{
	l:=new(LogBlockHandler)
	return l
}

func (this *LogBlockHandler) Handle(wrapper config.BlockWrapper) error {
	this.Debug("接收到块:%v", wrapper)
	return nil
	// e:=wrapper.BlockEvent
	// block := e.Block
	// sourceUrl := e.SourceURL
	// blockNumber := block.Header.Number
	// currentHash := block.Header.DataHash
	// prevHash := block.Header.PreviousHash
	//
	// l := len(block.Data.Data)
	// this.Debug("检测到生成了新的block的交易总数为:%d", l)
	// txFilter := util.TxValidationFlags(block.Metadata.Metadata[sdkCommon.BlockMetadataIndex_TRANSACTIONS_FILTER])
	// if len(txFilter) == 0 {
	// 	txFilter = util.NewTxValidationFlags(l)
	// 	block.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER] = txFilter
	// }
	// totalTransaction := 0
	// detailInfo, err := utils.GetBlockDetailV2(wrapper.ChainCodes, block.Data.Data[0])
	// createdTime := detailInfo.CreateTime
	// amount := detailInfo.Amount
	// signature := ""
	// if detailInfo.Signature != nil && len(detailInfo.Signature) > 0 {
	// 	signature = hex.EncodeToString(detailInfo.Signature)
	// }
	// if nil != err {
	// 	this.Error("获取区块详情失败:%s", err.Error())
	// } else {
	// 	for i := 1; i < l; i++ {
	// 		if !txFilter.IsValid(i) {
	// 			logs.Error("检测到无效交易")
	// 			continue
	// 		}
	// 		isLogic, _, am, err := utils.GetBlockDetail(wrapper.ChainCodes, block.Data.Data[i])
	// 		if nil != err {
	// 			this.Error("获取区块详情失败:%s", err.Error())
	// 		} else if isLogic {
	// 			totalTransaction++
	// 			amount += am
	// 		}
	// 	}
	// 	id, _ := idGenerator.NextId()
	// 	b := models.DBVlinkBlock{
	// 		ID:                          int(id),
	// 		CreatedDate:                 createdTime,
	// 		BlockNumber:                 blockNumber,
	// 		BlockCurrentHash:            hex.EncodeToString(currentHash),
	// 		BlockPrevHash:               hex.EncodeToString(prevHash),
	// 		BlockTotalTransaction:       totalTransaction,
	// 		BlockTotalTransactionAmount: amount,
	// 		BlockSourceUrl:              sourceUrl,
	// 		ChannelID:                   string(this.ChannelID),
	// 		Signature:                   signature,
	// 	}
	// }
}
