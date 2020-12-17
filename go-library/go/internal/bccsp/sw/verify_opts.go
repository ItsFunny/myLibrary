/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-11 19:40 
# @File : verify_opts.go
# @Description : 
# @Attention : 
*/
package sw

import "github.com/hyperledger/fabric/base"

type ECDSAVerifyOpts struct {
	base.EcdsaOpts
}

func (ECDSAVerifyOpts) A() interface{} {
	panic("implement me")
}

type SM2VerifyOpts struct {
	base.SM2Opts
}


func (SM2VerifyOpts) A() interface{} {
	panic("implement me")
}

type RSAVerifyOpts struct {
	base.RSAOpts
}

func (RSAVerifyOpts) A() interface{} {
	panic("implement me")
}
