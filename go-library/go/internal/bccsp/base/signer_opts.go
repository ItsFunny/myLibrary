/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-29 09:36 
# @File : signer_opts.go
# @Description : 
# @Attention : 
*/
package base

import (
	"crypto"
)

type UnImplementHashFunc struct {

}

func (this UnImplementHashFunc) HashFunc() crypto.Hash {
	panic("implement me")
}

type SM2SignerOpts struct {
	UnImplementHashFunc
	SM2Opts
}


type ECDSASignerOpts struct {
	UnImplementHashFunc
	EcdsaOpts
}

type RSASignerOpts struct {
	UnImplementHashFunc
	RSAOpts
}
