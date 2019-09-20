/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-06 14:10 
# @File : sort_interface.go
# @Description :  sort 接口
*/
package sort

import (
	"sort"
	"time"
)


type ISortInterface interface {
	sort.Interface
	// 根据ID 排序
	GetID() int
	// 根据名称排序
	GetName() string
	// 根据时间排序
	GetTimezome() time.Time
}
