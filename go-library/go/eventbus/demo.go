/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/3/8 4:51 下午
# @File : demo.go
# @Description :
# @Attention :
*/
package main

import (
	"encoding/json"
	"fmt"
)

// func main() {
// 	c := make(chan int, 10)
// 	c <- 1
// 	cc := c
// 	cc = nil
// 	count := 0
// 	for {
// 		select {
// 		case v := <-cc:
// 			fmt.Println(v)
// 			if count%2 == 0 {
// 				cc = nil
// 			}
// 		default:
// 			time.Sleep(time.Second * 1)
// 			count++
// 			if count%2 != 0 {
// 				cc = c
// 			}
// 			c <- 1
// 		}
// 	}
// }

type A struct {
	Name string
}
func main() {
	var bytes []byte
	bytes=nil
	var a A
	json.Unmarshal(bytes,&a)

	fmt.Println(&a)
}