/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-23 13:11
# @File : date.go
# @Description :
# @Attention :
*/
package utils

import "time"

func FormatYYMMDDHHMMSS(t int64) string {
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
	//
	// timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	// loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	// theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	// sr := theTime.Unix()
}
