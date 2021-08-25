/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-20 17:31
# @File : opts.go
# @Description :
# @Attention :
*/
package base

type SerializeOpts interface {
	GetBaseTypeFunc() BaseTypeFunc
}

type DefaultSerializeOpts struct {
	Func BaseTypeFunc
}

func (d DefaultSerializeOpts) GetBaseTypeFunc() BaseTypeFunc {
	return d.Func
}
