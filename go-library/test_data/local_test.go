/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-26 09:41 
# @File : test_lock.go
# @Description : 
# @Attention : 
*/
package test_data

import (
	"log"
	"sync"
	"testing"
)
type A struct{
	lock sync.RWMutex
}
func(this *A)sss(){
	this.lock.Lock()
	defer this.lock.Unlock()
	log.Println(1)
	this.lock.Lock()
	log.Println(2)
	this.lock.Unlock()
}
func Test_aaa(t *testing.T) {
	a:=&A{}
	a.sss()
}
