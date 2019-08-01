/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-18 09:08 
# @File : base.go
# @Description : 
*/
package job

import "errors"

type BaseJobExecutor struct {
	ConcreteExecutor IConcreteExecutor
	Type             JobType
	NextExecutor     IJobExecutor
}

func NewBaseJobExecutor(Type JobType, conreteExecutor IConcreteExecutor, nextExecutor IJobExecutor) *BaseJobExecutor {
	b := new(BaseJobExecutor)
	b.Type = Type
	b.ConcreteExecutor = conreteExecutor
	b.NextExecutor = nextExecutor
	return b
}

// func (receiver *BaseJobExecutor) ValidJobType(jobType JobType) bool {
// 	return receiver.Type == jobType
// }

func (receiver *BaseJobExecutor) Execute(job IAppEvent) (interface{}, error) {
	if nil == receiver.ConcreteExecutor {
		return nil, errors.New("executor为nil")
	}
	if job.ValidJobType(receiver.Type) {
		data, b, err := receiver.ConcreteExecutor.DoExecute(job)
		if nil != err {
			return nil, err
		}
		if !b || nil == receiver.NextExecutor {
			return data, nil
		} else if b {
			return receiver.NextExecutor.Execute(job.Copy(data))
		}
		return data, nil
	} else if nil != receiver.NextExecutor {
		return receiver.NextExecutor.Execute(job)
	} else {
		return nil, nil
	}
}

type BaseJobExecutorHandler struct {
	*BaseJobExecutor
}

type IJobConcreteHandler interface {
	DoHandle(job IAppEvent) (interface{}, error)
}

type BaseJobHandler struct {
	ConcreteHandler IJobConcreteHandler
	Type            JobType
	NextHandler     IJobHandler
}

func NewBaseJobHandler(Type JobType, concereteHandler IJobConcreteHandler, nextHandler IJobHandler) *BaseJobHandler {
	j := new(BaseJobHandler)
	j.Type = Type
	j.ConcreteHandler = concereteHandler
	j.NextHandler = nextHandler

	return j
}

func (receiver *BaseJobHandler) Handle(job IAppEvent) (interface{}, error) {
	if nil == receiver.ConcreteHandler {
		return nil, errors.New("配置错误,handler为nil")
	}
	if job.ValidJobType(receiver.Type) {
		data, err := receiver.ConcreteHandler.DoHandle(job)
		return data, err
	} else if nil != receiver.NextHandler {
		return receiver.NextHandler.Handle(job)
	} else {
		return nil, nil
	}
}
