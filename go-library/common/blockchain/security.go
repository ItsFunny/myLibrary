/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:53 
# @File : security.go
# @Description : 
# @Attention : 
*/
package blockchain

import "myLibrary/go-library/common/blockchain/base"

type BaseFabricAfterValidModel struct {
	Req           interface{}
	Version       uint64
	BaseTransType base.TransBaseTypeV2
	// 描述
	BaseTransDescription string
}


