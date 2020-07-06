/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 14:01 
# @File : config.go
# @Description : 
# @Attention : 
*/
package base

import "myLibrary/go-library/go/base/service"

type BaseConfigServiceImpl struct {
	*service.BaseServiceImpl
}

func NewBaseConfigServiceImpl() *BaseConfigServiceImpl {
	l := new(BaseConfigServiceImpl)
	l.BaseServiceImpl=service.NewBaseServiceImplWithLog4goLogger()
	return l
}
