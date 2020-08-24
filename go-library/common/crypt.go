/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 16:58 
# @File : crypt.go
# @Description : 
# @Attention : 
*/
package common

import "sync"

type ICrypter interface {
	Encrypt(data ...interface{}) (interface{}, error)
	Decrypt(data ...interface{}) (interface{}, error)
}

func main() {
	m:=sync.Map{}
	m.Load("")
}