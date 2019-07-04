/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:52 
# @File : url.go
# @Description : 
*/
package utils

import "strings"

func GetLowerSuffixFromUrl(url string)string{
	index := strings.LastIndex(url, ".")
	if index == -1 {
		return ""
	}
	l := len(url)
	if index == l-1 {
		return ""
	}
	suffix := SubStringBetween(url, index, l)
	suffix = strings.ToLower(suffix)
	return suffix
}