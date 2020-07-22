/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-17 09:14
# @File : scraptch.go
# @Description :
# @Attention :
*/
package utils

import (
	"testing"
)

func Test_scrape(t *testing.T) {
	scrape()
}

func Test_scrapeDetail(t *testing.T) {
	scrapeDetail("https://tieba.baidu.com/p/6814856936")
}
