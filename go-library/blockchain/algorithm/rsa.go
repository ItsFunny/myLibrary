/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-13 22:02 
# @File : rsa.go
# @Description : 
# @Attention : 
*/
package algorithm

import (
	"crypto/rand"
	"crypto/rsa"
)

func RSAPrivateKeyGenerator() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// 生成RSA证书
