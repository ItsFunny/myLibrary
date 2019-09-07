/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-21 09:20 
# @File : BaseServiceInitImpl.go
# @Description : 
*/
package baseImpl

import (
	"myLibrary/library/src/main/go/common/log"
)

type BaseServiceInitImpl struct {
	ReqID string
	Log   *log.Log
}

func (receiver *BaseServiceInitImpl) GetReqId() string {
	return receiver.ReqID
}
func (receiver *BaseServiceInitImpl) SetReqId(id string) {
	receiver.ReqID = id
}

func (receiver *BaseServiceInitImpl) SetLogger(l *log.Log) {
	receiver.Log = l
}
func (receiver *BaseServiceInitImpl) GetLogger() *log.Log {
	return receiver.Log
}
func NewBaseServiceInitImpl() *BaseServiceInitImpl {
	iml := new(BaseServiceInitImpl)
	iml.Log = log.NewLog(log.InitLog{
		ReqID: "",
	})
	return iml
}
