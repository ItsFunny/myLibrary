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
	str := "1997-04-02 09:25:41"
	tt := TimeConvStringWithOtherWay(str)
	now := time.Now()
	i := now.Unix() - tt.Unix()
	fmt.Println("秒:", i)
	fmt.Println("分:", i/60)
	fmt.Println("时:", i/60/60)
	fmt.Println("天:", i/60/60/24)
	fmt.Println("年:", i/60/60/24/365)
	fmt.Println("==============")
	// 14350 *
	fen := 14350 * 3.5
	miao := fen * 60
	shi := fen / 60
	tian := shi / 24
	fmt.Println("分:", fen)
	fmt.Println("秒:", int(miao))
	fmt.Println("时:", shi)
	fmt.Println("天:", tian)
	fmt.Println(float64(837 / 199540))

	fmt.Println("========")
	f := float64(1006 * 3.37)
	m := f * 60
	s := f / 60
	ttt := s / 24
	fmt.Println("秒:", m)
	fmt.Println("分:", f)
	fmt.Println("时:", s)
	fmt.Println("天:", ttt)

}

// 测试不同小时制的时间
func TestClearPointerSlice(t *testing.T) {
	// 12小时
	str := "2019-09-20 02:54:21"
	tt := TimeConvStringWithOtherWay(str)
	unix := tt.Unix()
	fmt.Println(tt.String())
	int2Time := FormatInt2Time(int(unix))
	fmt.Println(int2Time.String())
	i := 1575530288
	formatInt2Time := FormatInt2Time(i)
	fmt.Println(formatInt2Time.String())
}
