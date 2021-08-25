/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/2/25 8:51 上午
# @File : wrapper.go
# @Description :
# @Attention :
*/
package bccsp

import (
	config2 "github.com/cloudflare/cfssl/config"
	"time"
)

type CrtCsrWrapper struct {
	Policy    *config2.Signing
	Profile   string
	NotBefore time.Time
	// If provided, NotAfter will be used without modification (except
	// for canonicalization) as the value of the notAfter field of the
	// certificate.
	NotAfter time.Time
}
