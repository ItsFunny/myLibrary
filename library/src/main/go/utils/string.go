/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-04-24 10:56 
# @File : string.go
# @Description : 
*/
package utils

import (
	"regexp"
	"strings"
)

// 消除多余的空格
func ClearEmpltyBlank(str string) string {
	return strings.Replace(str, " ", "", -1)
}

// 从0开始
func SubString(str string, index int) string {
	str = string([]rune(str)[:index])
	return str
}
func SubStringBetween(str string,begin ,end int )string{
	// 注意下标从0开始 前闭后开 [begin,end)
	return  string([]rune(str)[begin:end])
}

// 判断是否有空格
func ContainEmptyBlank(str ...string) bool {
	for _, s := range str {
		if strings.Index(s, " ") > 0 {
			return true
		}
	}

	return false
}

// 合并多个空格为1个
func MergeExtraEmptyBlank(s string) string {
	// 删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 { // 找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}
