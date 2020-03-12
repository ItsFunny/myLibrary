/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-04 14:42 
# @File : spide_test.go
# @Description : 
*/
package plugin

import (
	"fmt"
	"myLibrary/go-library/go/page"
	"testing"
)

func TestSpide_SpideOne(t *testing.T) {
}

func TestSpide_SpideOne_ARTICLE(t *testing.T) {
	var a page.BasePageReq
	a.PageNum = 1
	tt := &a
	tt.Cc()
	fmt.Println(a.PageNum)
}
func c(req *page.BasePageReq) {
	req.PageNum = 111
}
