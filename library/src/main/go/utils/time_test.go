/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-28 11:24 
# @File : time_test.go
# @Description : 
*/
package utils

import (
	"fmt"
	"github.com/akkagao/citizens/utils"
	"testing"
)

func TestInt64ConvT2TimeStrTilSec(t *testing.T) {
	// unix := time.Now().Unix()
	// timeunix:=156233873
	time := utils.GetBeiJingTime()
	sec := Int64ConvT2TimeStrTilSec(int64(time))
	fmt.Println(sec)
}
