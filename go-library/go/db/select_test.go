/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-11-13 16:33 
# @File : select_test.go
# @Description : 
*/
package db

import (
	"fmt"
	"testing"
)

func TestSpliceInSQL(t *testing.T) {
	s, i := SpliceInSQL("ID", []int{1, 2, 3, 4, 5})
	fmt.Println(s)
	fmt.Println(i)

}
