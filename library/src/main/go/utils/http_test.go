/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-27 13:39 
# @File : http_test.go
# @Description : 
*/
package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDoPostForm(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if nil!=err{
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil!=err{
		panic(err)
	}
	fmt.Println(string(bytes))
}

