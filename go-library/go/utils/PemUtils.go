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
	"regexp"
	"strings"
)

func replace(str,pemType  string) string {
	str = strings.ReplaceAll(str, "-----BEGIN "+pemType+"-----", "")
	str = strings.ReplaceAll(str, "-----END "+pemType+"-----", "")
	// 匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

