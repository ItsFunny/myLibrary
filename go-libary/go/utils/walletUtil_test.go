/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-17 16:30 
# @File : walletUtilTest.go
# @Description : 
*/
package utils

import (
	"fmt"
	"myLibrary/library/src/main/go/crypt"
	"testing"
)

func TestNewWallet(t *testing.T) {
	wallet, e := NewWallet()
	if nil != e {
		panic(e)
	}
	str := "joker"
	strs, e := encrypt.ECCSignWithHex(str, wallet.PrivateKeyBytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(strs)
	hex := encrypt.ECCVerifySignWithHex(str, strs, wallet.PublicKeyBytes)
	fmt.Println(hex)
}

func TestWalletEncrypt(t *testing.T) {
	wallet, e := NewWallet()
	if nil != e {
		panic(e)
	}
	str := "joker"
	s, e := encrypt.ECCEncryptWithHex(str, wallet.PublicKeyBytes)
	if nil != e {
		panic(e)
	}
	hex, e := encrypt.ECCDecryptWithHex(s, wallet.PrivateKeyBytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(hex)

}

func TestNewWallet2(t *testing.T) {
	fmt.Println(88&64 == 64)
}
