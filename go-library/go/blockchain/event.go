/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:00 
# @File : event.go
# @Description : 
# @Attention : 
*/
package config

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"go.uber.org/atomic"
	"myLibrary/go-library/go/blockchain/handler"
)

var (
	executor    *TaskExecutor
)

func init() {
	executor = newVlinkTaskExecutor()
}

func newVlinkTaskExecutor() *TaskExecutor {
	v := new(TaskExecutor)
	v.BaseLogicServiceImpl = NewBaseLogicServiceImpl()
	return v
}

type TaskExecutor struct {
	*BaseLogicServiceImpl
	BlockEventListeners []*blockEventExecutor
}

type blockEventExecutor struct {
	*BaseLogicServiceImpl
	ChannelID ChannelID

	ChainCodes         []string
	InterestChainCodes []string
	// 期望的block 编号,用于以防数据丢失
	ExpectedBlockIndex *atomic.Int64
	events             <-chan *fab.BlockEvent

	stop chan struct{}

	BlockHandler IBlockHandler
}

func RegisterBlockEvent(cid ChannelID, interestChainCodes []string, events <-chan *fab.BlockEvent, stop chan struct{}) {
	b := new(blockEventExecutor)
	b.BlockHandler = handler.NewLogBlockHandler()
	b.BaseLogicServiceImpl = NewBaseLogicServiceImpl()
	b.ChannelID = cid
	b.stop = stop
	b.InterestChainCodes = interestChainCodes
	b.ExpectedBlockIndex = atomic.NewInt64(1)
	b.events = events
	if executor.BlockEventListeners == nil {
		executor.BlockEventListeners = make([]*blockEventExecutor, 0)
	}
	executor.BlockEventListeners = append(executor.BlockEventListeners, b)
}

func RunTasks() {
	for _, listner := range executor.BlockEventListeners {
		go listner.handleEventWithDetailBlock()
	}
}

func (this *blockEventExecutor) handleEventWithDetailBlock() {
	for {
		select {
		case <-this.stop:
			break
		case e := <-this.events:
			w := BlockWrapper{
				BlockEvent: e,
				ChainCodes: this.ChainCodes,
			}
			this.BlockHandler.Handle(w)
		}
	}
}
