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
	"encoding/json"
	"fmt"
	"testing"
	"time"
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

func TestFormyAsyyyyMMddHHmmssSSS(t *testing.T) {
	now := time.Now()
	sss := FormtAsyyyyMMddHHmmssSSS(now)
	fmt.Println(sss)
}

type AStruct struct {
	Name string
	AlgT int
}

type PAStruct struct {
	*AStruct
	VVV string
}

func Test_A(t *testing.T) {
	pa := &PAStruct{
		AStruct: &AStruct{
			Name: "1111",
			AlgT: 111,
		},
		VVV: "vvv",
	}
	bytes, _ := json.Marshal(pa)
	fmt.Println(string(bytes))

	paa := PAStruct{}

	json.Unmarshal(bytes, &paa)
	fmt.Println(paa)

}

func calcTimes(send, recv int64) float64 {
	return float64((387) / 1000.0)
}
func TestCalc(t *testing.T) {
	sendUnix:=1616129917491
	//           1616130134766182000  纳秒
	//           1616130174076 毫秒
	nowUnix := time.Now().UnixNano() / 1e6
	fmt.Println(nowUnix)
	times := calcTimes(int64(sendUnix), nowUnix)
	fmt.Println(times)
}
