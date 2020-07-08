/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-30 11:26 
# @File : pic_test.go
# @Description : 
*/
package pic

import (
	"fmt"
	"myLibrary/go-library/go/utils"
	"testing"
	"time"
)

func TestBatchComparePicSimilarity(t *testing.T) {
	now := time.Now().Unix()
	dir1 := "/Users/joker/Desktop/漫画/彭波--魔力青春摇一摇/moli_ceshi1"
	dir2 := "/Users/joker/Desktop/漫画/彭波--魔力青春摇一摇/moli_ceshi2"
	filePathes1, e := utils.GetFilesBelownDirFilterBySuffix(dir1, ".jpg")
	if nil != e {
		// ExecutionConsumer
		panic(e)
	}
	filePathes2, e := utils.GetFilesBelownDirFilterBySuffix(dir2, ".jpg")
	if nil != e {
		panic(e)
	}
	i, f, e := BatchComparePicSimilarity("a", filePathes2, filePathes1, 6)
	if nil != e {
		panic(e)
	}
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println("消耗时间:", time.Now().Unix()-now)
}
