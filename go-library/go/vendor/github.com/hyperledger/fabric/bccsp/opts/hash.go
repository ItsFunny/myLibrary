/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-12 19:28
# @File : hash.go
# @Description :    hash 算反返回的都是一样的,当然最好是有一个hashType
# @Attention :
*/
package opts

import (
	"github.com/hyperledger/fabric/base"
)

type DefaultSHA256HashOptsImpl struct {
}

func (DefaultSHA256HashOptsImpl) GetMinExcepted() base.BaseTypeFunc {
	return func() base.BaseType {
		return base.BaseType(base.HASH_SHA256)
	}
}

func (DefaultSHA256HashOptsImpl) Algorithm() string {
	return "SHA256"
}

func (DefaultSHA256HashOptsImpl) GetDetailType() base.DetailTypeFunc {
	return func() base.DetailType {
		return base.DetailType(base.HASH_SHA256)
	}
}

func (DefaultSHA256HashOptsImpl) GetBaseType() base.BaseTypeFunc {
	return func() base.BaseType {
		return base.BaseType(base.HASH_SHA256)
	}
}

type DefaultSM3HashOptsImpl struct {
}

func (DefaultSM3HashOptsImpl) GetMinExcepted() base.BaseTypeFunc {
	return func() base.BaseType {
		return base.BaseType(base.HASH_SM3_WITH_NO)
	}
}

func (DefaultSM3HashOptsImpl) Algorithm() string {
	return "SM3"
}

func (DefaultSM3HashOptsImpl) GetDetailType() base.DetailTypeFunc {
	return func() base.DetailType {
		return base.DetailType(base.HASH_SM3_WITH_NO)
	}
}

// FIXME 应该是gm 而非sm2
func (DefaultSM3HashOptsImpl) GetBaseType() base.BaseTypeFunc {
	return func() base.BaseType {
		return base.BaseType(base.HASH_SM3_WITH_NO)
	}
}
