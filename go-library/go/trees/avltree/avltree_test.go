/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-05 14:13 
# @File : avltree_test.go
# @Description : 
# @Attention : 
*/
package avltree

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/utils"
	"testing"
)

func TestNewWithIntComparator(t *testing.T) {
	tre := new(Tree)
	tre.Put(utils.StringComparator, "1", 2)
	value, _ := tre.Get(utils.StringComparator, "1")
	fmt.Println(value)
	bytes, e := json.Marshal(tre)
	if nil != e {
		panic(e)
	}
	fmt.Println(string(bytes))

	var tre2 Tree
	e = json.Unmarshal(bytes, &tre2)
	if nil != e {
		panic(e)
	}
	fmt.Println(tre2.Get(utils.StringComparator, "1"))
	tre2.Put(utils.StringComparator, "3", 4)
	fmt.Println(tre2.Get(utils.StringComparator, "3"))
}
