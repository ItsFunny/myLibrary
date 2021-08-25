/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 17:02
# @File : key.go
# @Description :
# @Attention :
*/
package bccsp

import (
	"github.com/hyperledger/fabric/base"
)

type PrivateKeyWrapper struct {
	Key           IKeyConverter
	AlgorithmType base.BaseType
}

type PublicKeyWrapper struct {
}
