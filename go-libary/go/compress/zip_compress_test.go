/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-10-21 09:43 
# @File : zip_test.go
# @Description : 
*/
package compress

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	src := "/Users/joker/Desktop/个人/区块链"
	dst := "/Users/joker/Desktop/个人/a.zip"
	if e := ZipCompress(src, dst); nil != e {
		fmt.Println(e)
	}
}
