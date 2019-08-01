/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-24 09:26 
# @File : file_test.go
# @Description : 
*/
package utils

import (
	"fmt"
	"testing"
)

func TestGetFileSize(t *testing.T) {
	path := "/Users/joker/Downloads/fengkuangwaixingren.mp4"
	size, e := GetFileSize(path)
	fmt.Println(e)
	fmt.Println(size)
}
