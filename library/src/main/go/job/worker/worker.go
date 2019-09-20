/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-31 17:26 
# @File : worker.go
# @Description : 
*/
package worker

import (
	"io"
	"log"
)

type Worker interface {
	io.Closer
	Start()

	SetWorkerPool(workerPool chan chan interface{})
	GetWorkerPool() chan chan interface{}
}

// 2019-01-08 16:07 面向接口编程,干脆把这个给单独抽出
type Consumer interface {
	Consume(interface{}) (interface{}, error)
}

type ConcreteConsumer interface {
	DoConsume(interface{}) (interface{}, error)
}

type BaseWoker struct {
	// 共享全局的工作池
	WorkPool chan chan interface{}
	// 这个worker的工作队列,当有新的任务
	jobQueue         chan interface{}
	closed           bool
	Config           func() // 启动之前的先初始化,或者说是配置
	ConcreteConsumer ConcreteConsumer
	// 抽象方法 ,2019-01-08 15:54 将这个抽离为单独的方法,因为不拆分的话会变成一坨代码
	// Consume  func(interface{}) error
}

func (w *BaseWoker) InitJobQueue(maxJobCount int) {
	w.jobQueue = make(chan interface{}, maxJobCount)
}

func (w *BaseWoker) SetWorkerPool(workerPool chan chan interface{}) {
	w.WorkPool = workerPool
}

func (w *BaseWoker) GetWorkerPool() chan chan interface{} {
	return w.WorkPool
}

func NewBaseWorker(workPool chan chan interface{}, jobQueue chan interface{}) *BaseWoker {
	b := new(BaseWoker)
	b.jobQueue = jobQueue
	b.WorkPool = workPool
	return b
}
func NewBaseWorkerWithOutConfig(consumer ConcreteConsumer, confFunc func()) *BaseWoker {
	b := new(BaseWoker)
	b.ConcreteConsumer = consumer
	return b
}

// 2019-01-08 16:11 交给子类去实现
func (w *BaseWoker) Consume(d interface{}) (interface{}, error) {
	return w.ConcreteConsumer.DoConsume(d)
}

func (w *BaseWoker) Close() error {
	if !w.closed {
		close(w.jobQueue)
		w.closed = true
	}
	return nil
}

func (w *BaseWoker) Start() {
	go func() {
		for {
			w.WorkPool <- w.jobQueue
			select {
			case v := <-w.jobQueue:
				log.Println("[Worker]收到了新的消息:", v)
				_, err := w.Consume(v)
				if nil != err {
					log.Println("[Worker]consume occur error; ", err)
				}
			}
		}
	}()
}
