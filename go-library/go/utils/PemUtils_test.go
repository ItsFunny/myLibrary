/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-06 16:16
# @File : PemUtils.go
# @Description :
# @Attention :
*/
package utils

import (
	"fmt"
	"testing"
)

func Test_replace(t *testing.T) {
	str := `-----BEGIN PRIVATE KEY-----
lYwhNL2TgYmBI9Okg6ouSvDRo8wjcLrNj75oPeQ6/uY=
-----END PRIVATE KEY-----
`
	s := replace(str, "PRIVATE KEY")
	fmt.Println(s)
}
