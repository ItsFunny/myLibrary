/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 17:09 
# @File : converter_test.go
# @Description : 
*/
package converter

import (
	"fmt"
	"testing"
)

func TestBytes2Int64(t *testing.T) {
	a := int64(1)
	bytes := BigEndianInt642Bytes(int64(a))
	fmt.Println(bytes)
	bytes2Int64 := BigEndianBytes2Int64(bytes)
	fmt.Println(bytes2Int64)
}

func TestInt64ToBytes(t *testing.T) {

}
