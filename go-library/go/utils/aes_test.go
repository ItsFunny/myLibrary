/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-23 14:24
# @File : aes.go
# @Description :
# @Attention :
*/
package utils

import (
	"fmt"
	"testing"
)

func TestAesEncryptCBC(t *testing.T) {
	guangzhouAesKey:="WbUBp5dzlv7gHeST94SqiLhj/IQU0Kw9EN1hmypmyig="
	originData:=[]byte("123")
	bytes, _ := Base64Decode(guangzhouAesKey)
	encrypted := AesEncryptCFB(originData,bytes)
	fmt.Println(Base64Encode(encrypted))

	decrypted := AesDecryptCFB(encrypted, bytes)
	fmt.Println(string(decrypted))

}
