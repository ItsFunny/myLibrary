/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-12 19:28 
# @File : hash.go
# @Description : 
# @Attention : 
*/
package opts

import (
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/constants"
)

type DefaultSHA256HashOptsImpl struct {

}

func (DefaultSHA256HashOptsImpl) Algorithm() string {
	return "SHA256"
}

func (DefaultSHA256HashOptsImpl) GetTypeFunc() base.SerializableFunc {
	return func() base.SerializableInterface {
		return  constants.HASH_SHA256
	}
}

func (DefaultSHA256HashOptsImpl) GetBaseType() base.SerializableFunc {
	return func() base.SerializableInterface {
		return constants.ECDSA
	}
}


type DefaultSM3HashOptsImpl struct {

}

func (DefaultSM3HashOptsImpl) Algorithm() string {
	return "SM3"
}

func (DefaultSM3HashOptsImpl) GetTypeFunc() base.SerializableFunc {
	return func() base.SerializableInterface {
		return constants.HASH_SM3_WITH_NO
	}
}
// FIXME 应该是gm 而非sm2
func (DefaultSM3HashOptsImpl) GetBaseType() base.SerializableFunc {
	return func() base.SerializableInterface {
		return constants.SM2
	}
}
