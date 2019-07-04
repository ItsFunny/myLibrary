/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 11:20 
# @File : time.go
# @Description : 
*/
package utils

import (
	"time"
)

const (
	BASE_TIME_FORMAT_TILL_SEC = "2006-01-02 03:04:05"
)

func GetBeiJingTimeZone()*time.Location{
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	return cstZone
}

func Int64ConvT2TimeStrTilSec(timeStamp int64) string {
	return time.Unix(timeStamp,0).Format(BASE_TIME_FORMAT_TILL_SEC)
}
