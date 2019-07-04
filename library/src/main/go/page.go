/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 15:40 
# @File : page.go
# @Description : 
*/
package main

type PageInfo struct {
	PageNum    int
	PageSize   int
	Data       interface{}
	TotalCount int
}
