/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-10 22:47 
# @File : base_service_test.go
# @Description : 
# @Attention : 
*/
package service

import "testing"

func TestNewBaseServiceImplWithLog4goLogger(t *testing.T) {
	logger := NewBaseServiceImplWithLog4goLogger()
	logger.BeforeStart("aaa")
	logger.BeforeStart("bbbbb")
	logger.BeforeStart("ccccc")
	logger.AfterEnd()
}
