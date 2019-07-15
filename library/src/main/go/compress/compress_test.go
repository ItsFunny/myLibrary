/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-08 14:44 
# @File : compress_test.go
# @Description : 
*/
package compress

import (
	"fmt"
	"testing"
)

func TestGzipEncode(t *testing.T) {
	str:="joker"
	fmt.Println([]byte(str))
	bytes, e := GzipEncode([]byte(str))
	if nil!=e{
		fmt.Println(e)
	}else{
		fmt.Println(bytes)
	}
}

func TestGzipDecode(t *testing.T) {

}