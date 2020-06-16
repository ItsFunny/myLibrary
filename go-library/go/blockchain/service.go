/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:01 
# @File : service.go
# @Description : 
# @Attention : 
*/
package config

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"myLibrary/go-library/go/base/service"
)

type BaseLogicServiceImpl struct {
	*service.BaseServiceImpl
}

type BlockWrapper struct {
	*fab.BlockEvent
	ChainCodes []string
}
type IBlockHandler interface {
	Handle(BlockWrapper) error
}



func NewBaseLogicServiceImpl()*BaseLogicServiceImpl{
	l:=new(BaseLogicServiceImpl)
	l.BaseServiceImpl=service.NewBaseServiceImplWithLog4goLogger()
	return l
}

