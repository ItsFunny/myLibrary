/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-25 09:52 
# @File : aes_test.go
# @Description : 
*/
package encrypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesCBCEncrypt(t *testing.T) {
	key := []byte("321423u9y8d2fwfl")
	data := "joker"
	bytes, e := AesCBCEncrypt(key, []byte(data))
	if nil != e {
		panic(e)
	} else {
		fmt.Println(hex.EncodeToString(bytes))
		fmt.Println(bytes)
	}

}

func TestAesCBCDecrypt(t *testing.T) {
	key := []byte("321423u9y8d2fwfl")
	str2 := "7a77b3a999c5724caf1bb32bafce2989e6fc5ddfcff411886e57e7d5d8f66fef"
	// str := "9adabc228faf25211a39904ef7bcb8e104df3030fe"
	bytes, _ := hex.DecodeString(str2)
	aesDecrypt, e := AesDecrypt(key, bytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(string(aesDecrypt))
	// originData, err := AesCBCDecrypt(key, bytes)
	// if nil != err {
	// 	panic(err)
	// } else {
	// 	l := "156156165152165156156"
	// 	decodeString, _ := hex.DecodeString(l)
	// 	fmt.Println(decodeString)
	// 	fmt.Println(originData)
	// }
}
