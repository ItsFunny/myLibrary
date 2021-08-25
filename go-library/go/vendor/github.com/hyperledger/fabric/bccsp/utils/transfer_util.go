/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-22 17:34
# @File : transfer_util.go
# @Description :
# @Attention :
*/
package utils

import (
	"crypto/ecdsa"
	"github.com/tjfoc/gmsm/sm2"
)

func ParseEcdsaPrvKey2SM2PrvKey(ecdsaPrv *ecdsa.PrivateKey) *sm2.PrivateKey {
	return &sm2.PrivateKey{
		PublicKey: sm2.PublicKey{
			Curve: ecdsaPrv.PublicKey.Curve,
			X:     ecdsaPrv.PublicKey.X,
			Y:     ecdsaPrv.PublicKey.Y,
		},
		D: ecdsaPrv.D,
	}
}
func ParseEcdsaPubKey2SM2PubK(ecdsPub *ecdsa.PublicKey) *sm2.PublicKey {
	return &sm2.PublicKey{
		Curve: ecdsPub.Curve,
		X:     ecdsPub.X,
		Y:     ecdsPub.Y,
	}
}
