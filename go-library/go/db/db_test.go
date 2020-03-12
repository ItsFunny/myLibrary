/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-14 16:47 
# @File : db_test.go
# @Description : 
*/
package db

import (
	"fmt"
	"testing"
)

func TestBuildBatchInsertExecSQL(t *testing.T) {
	var p [][]string
	p[1][2] = "1"
	fmt.Println(p)
}

func TestOrderedDeleted(t *testing.T) {
	// tableNames := []string{"VLINK_M1", "VLINK_M2"}
	// clounms := []string{"ID", "IDD"}
	// values := make([][]int, 0)
	// values = append(values, []int{1, 2, 3}, []int{4, 5, 6})
	// strings, l := OrderedDeleted(tableNames, clounms, values)
	//
	// for _, s := range strings {
	// 	fmt.Println(s)
	// }
	// for i := 0; i < len(l); i++ {
	// 	fmt.Println(l[i])
	// }
	var a []int
	fmt.Println(a==nil)
}
