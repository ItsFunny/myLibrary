/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:53 
# @File : security.go
# @Description : 
# @Attention : 
*/
package base



type BaseFabricAfterValidModel struct {
	Req           interface{}
	Version       uint64
	BaseTransType TransBaseTypeV2
	// 描述
	BaseTransDescription string
}

type ICrypter interface {
	Encrypt(data ...interface{}) (interface{}, error)
	Decrypt(data ...interface{}) (interface{}, error)
}

