/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-19 13:36
# @File : logic_util.go
# @Description :
# @Attention :
*/
package utils

import (
	"fmt"
	"testing"
)

func TestGenerateOrgCode(t *testing.T) {
	fmt.Println(GenerateOrgCode())
}

func TestGenearateSocialCreditCode(t *testing.T) {
	code := GenerateSocialCreditCode()
	fmt.Println(code)
}
