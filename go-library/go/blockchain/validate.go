/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:18 
# @File : validate.go
# @Description : 
# @Attention : 
*/
package config


type ICrypter interface {
	Encrypt(data ...interface{}) (interface{}, error)
	Decrypt(data ...interface{}) (interface{}, error)
}
