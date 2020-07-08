/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-04-16 14:55 
# @File : date_utils_test.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"testing"
)

func TestGetCurrentTimeStrByTemplate(t *testing.T) {
	// fmt.Println(GetCurrentTimeDefault())
	// a := "/asd/asda/sdas/das/daa//sdd/d.tsa"
	// fmt.Println(len(a))
	// fmt.Println(strings.LastIndex(a, "/"))
	// str := "/Users/joker/Desktop/未闻花名/公司/test//615194096528523264/a.tsa"
	// index := strings.LastIndex(str, "/")
	// fmt.Println(index)
	// s := SubString(str, index)
	// fmt.Println(s)

	type Name struct {
		Name string
	}
	type tempStr struct {
		Names *[]Name
	}
}
