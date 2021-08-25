/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-16 23:48
# @File : key.go
# @Description :
# @Attention :
*/
package base

import (
	"crypto/ecdsa"
	"github.com/tjfoc/gmsm/sm2"
)

func ConvEcdsaPubK2Sm2PubK(pubK *ecdsa.PublicKey) *sm2.PublicKey {
	sm2Pubk := &sm2.PublicKey{
		Curve: pubK.Curve,
		X:     pubK.X,
		Y:     pubK.Y,
	}
	return sm2Pubk
}
