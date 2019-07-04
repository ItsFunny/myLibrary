/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-05-04 08:49
# @File : validation.go
# @Description :
*/
package utils

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	Convey("正常手机号", t, func() {
		str := "+8618757883747"
		fmt.Println(ValidatePhone(str))
	})

	Convey("加86的手机号", t, func() {
		str := "86-010-88888888"
		fmt.Println(ValidatePhone(str))
	})

	Convey("8位", t, func() {
		str := "88888888"
		fmt.Println(ValidatePhone(str))
	})
	Convey("166,136正常手机号", t, func() {
		str := "16657883747"
		str2 := "13657883747"
		fmt.Println(ValidatePhone(str))
		fmt.Println(ValidatePhone(str2))
	})
	Convey("错误手机号", t, func() {
		str := "187783324123"
		fmt.Println(ValidatePhone(str))
	})

}
func TestClearEmpltyBlank(t *testing.T) {
	str:="^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	matched, _ := regexp.MatchString(str, "18757883747")
	fmt.Println(matched)
}

type BlockChainUser struct {
	ID int `json:"id,string"`
	// Password      string `json:"password"`
	IDCard        string `json:"id_card"`
	WalletAddress string `json:"wallet_address"`
	RealName      string `json:"real_name"`
	// TODO
	// 装载总次数
	TransfedTotalCount int `json:"transfed_total_count"`
	// 被转载的文章数
	TransfedTotalItemCount int `json:"transfed_total_item_count"`
}

func TestGenValidateCode(t *testing.T) {
	str:=`{"id":"2","id_card":"2","wallet_address":"e059ab67485408c843511abeb2336c76","real_name":"","transfed_total_count":0,"transfed_total_item_count":0}`
	var u BlockChainUser
	e := json.Unmarshal([]byte(str), &u)
	if nil!=e{
		panic(e)
	}else{
		fmt.Println(u)
	}
}