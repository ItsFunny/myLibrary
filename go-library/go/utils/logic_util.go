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
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

func GenerateOrgCode() string {
	in := []int{3, 7, 9, 10, 5, 8, 4, 2}
	data := ""
	yz := ""
	a := 0
	for i := 0; i < len(in); i++ {
		word := getCharAndNumr(1, 0)
		word = strings.ToUpper(word)
		if ok, e := regexp.Match("[A-Z]", []byte(word)); nil != e || !ok {
			a += in[i] * getAsc(word)
		} else {
			atoi, _ := strconv.Atoi(word)
			a += in[i] * atoi;
		}
		data += word
	}
	// 确定序列
	c9 := 11 - a%11
	// 判断c9大小，安装 X 0 或者C9
	if c9 == 10 {
		yz = "X"
	} else if c9 == 11 {
		yz = "0"
	} else {
		yz = strconv.Itoa(c9)
	}
	data += "-" + yz

	return strings.ToUpper(data)
}

func getAsc(s string) int {
	gc := []byte(s)
	ascNum := int(gc[0] - 55)
	return ascNum
}

func getCharAndNumr(length int, status int) string {
	charStr := "0123456789abcdefghijklmnopqrstuvwxy";
	sb := strings.Builder{}
	if status == 1 {
		charStr = "0123456789"
	} else if status == 2 {
		charStr = "0123456789"
	} else if status == 3 {
		charStr = "0123456789ABCDEFGHJKLMNPQRTUWXY"
	}
	charLength := len(charStr)
	for i := 0; i < length; i++ {
		index, _ := IntRange(0, charLength)
		if status == 1 && index == 0 {
			index = 3
		}
		sb.WriteByte(charStr[index])
	}
	return sb.String()
}

//
// // 生成统一社会信用代码
func GenerateSocialCreditCode() string {
	data := ""
	first := "Y2" + getCharAndNumr(6, 3) + getCharAndNumr(9, 3)
	firstBytes:=[]byte(first)
	firstBytes=append(firstBytes,getUSCCCheckBit(first))
	data = string(firstBytes)
	data = strings.ToUpper(data)
	if !test4(data) {
		return GenerateSocialCreditCode()
	}
	return data
}

func getUSCCCheckBit(businessCode string) byte {
	if len(businessCode) == 0 || len(businessCode) != 17 {
		return 0
	}
	baseCode := "0123456789ABCDEFGHJKLMNPQRTUWXY"
	baseCodeArray := []byte(baseCode)
	codes := make(map[byte]int, 0)
	for i := 0; i < len(baseCode); i++ {
		codes[baseCodeArray[i]] = i
	}
	businessCodeArray := []byte(businessCode)
	wi := []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
	sum := 0
	for i := 0; i < 17; i++ {
		key := businessCodeArray[i]
		if !bytes.Contains(baseCodeArray,[]byte{key}){
			return 0
		}
		// exist := false
		// for _, b := range baseCode {
		// 	if b == int32(key) {
		// 		exist = true
		// 	}
		// }
		// if !exist {
		// 	return 0
		// }
		sum += codes[key] * wi[i]
	}
	value := 31 - sum%31
	if value == 31 {
		value = 0
	}
	return baseCodeArray[value]
}

func test4(data string) bool {
	if len(data) == 0 {
		return false
	}
	// Y2QY94Q2F8BHNTRXX3
	// Y24B6JNYX9M9XP1W880
	if len(data) != 18 {
		return false
	}
	if ok, err := regexp.Match("[a-zA-Z0-9]+", []byte(data)); nil != err || !ok {
		return false
	}
	regex := "^([159Y]{1})([1239]{1})([0-9ABCDEFGHJKLMNPQRTUWXY]{6})([0-9ABCDEFGHJKLMNPQRTUWXY]{9})([0-90-9ABCDEFGHJKLMNPQRTUWXY])$"
	if ok, err := regexp.Match(regex, []byte(data)); nil != err || !ok {
		return false
	}
	return true
}
