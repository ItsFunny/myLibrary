/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-06 22:14
# @File : chain.go
# @Description :
# @Attention :
*/
package bccsp

import "github.com/hyperledger/fabric/base"

type IHashChain interface {
	ValidIsMine(baseType base.BaseType) bool
}
