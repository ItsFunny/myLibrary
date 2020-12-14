/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-14 11:47 
# @File : zap_log.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Ops struct {

}

func main() {

	logger := zap.NewExample(zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Println("这是hook:"+entry.Message)
		return nil
	}))
	logger.Info("123123")

}
