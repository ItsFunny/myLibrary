/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-14 22:46
# @File : key.go
# @Description :
# @Attention :
*/
package bccsp

import (
	"crypto"
)

type IKeyConverter interface {
	crypto.Signer
	ToFabricKey() KeyAdapter
}
