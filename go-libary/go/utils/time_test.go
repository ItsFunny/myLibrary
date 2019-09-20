/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 11:24 
# @File : time_test.go
# @Description : 
*/
package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestInt64ConvT2TimeStrTilSec(t *testing.T) {
	// unix := time.Now().Unix()
	// timeunix:=156233873
	sec := Int64ConvT2TimeStrTilSec(1567442974)
	fmt.Println(sec)
}
func TestFormatTime2StringByTemplate(t *testing.T) {
	fmt.Println(FormatTime2StringByTemplate("20060102150405", time.Now().Unix()))
}
func TestTimeConvString2Int64(t *testing.T) {
	str := "2019-09-20 12:54:21"
	string2Int64 := TimeConvStringWithOtherWay(str).Unix()
	sec := Int64ConvT2TimeStrTilSec(string2Int64)
	fmt.Println(sec)
	fmt.Println(string2Int64)
}
