/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-24 12:43
# @File : blockchian.go
# @Description :
# @Attention :
*/
package config

import (
	"fmt"
	"testing"
)

func Test_generateNewVersion(t *testing.T) {
	fmt.Println(generateNewVersion("1.34"))
}
