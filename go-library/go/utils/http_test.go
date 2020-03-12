/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-27 13:39 
# @File : http_test.go
# @Description : 
*/
package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"myLibrary/go-library/go/crypt"
	"net/http"
	"testing"
)

func TestDoPostForm(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if nil != err {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type T struct {
	Success    bool   `json:"success"`
	ResultCode string `json:"resultCode"`
	Msg        string `json:"msg"`
	Data       string `json:"data"`
	Salt       string `json:"salt"`
}

func TestDoPostForm2(t *testing.T) {
	SALT := "AVU0X77JO4AAB8GY3QXW"
	m := make(map[string]interface{})
	m["partnerID"] = "201908261512362679"
	m["partnerKey"] = "AVU0X77JO4AAB8GY3QXW"
	m["serialNo"] = "364129854503006208"
	m["salt"] = encrypt.MD5EncryptByBytes("201908261512362679AVU0X77JO4AAB8GY3QXW364129854503006208") + SALT
	args := ConvtMap2FastHttpArgs(m)
	statusCode, body, err := fasthttp.Post(nil, "https://ipp.tsa.cn/v2/api/confirm/downloadOpusCertificate", args)
	if nil != err {
		panic(err)
	}
	if statusCode != http.StatusOK {
		fmt.Println("失败")
	} else {
		var t T
		if er := json.Unmarshal(body, &t); nil != er {
			panic(er)
		}
		bytes, err := base64.StdEncoding.DecodeString(t.Data)
		if nil != err {
			panic(err)
		}
		if err := Write2File("asd.pdf", bytes); nil != err {
			panic(err)
		}

	}
}
