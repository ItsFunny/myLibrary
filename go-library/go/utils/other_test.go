/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-13 05:38 
# @File : other_test.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_AAA(t *testing.T) {
	c := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 1)
		close(c)
	}()
	<-c
	fmt.Print(1)
}

func Test_bbb(t *testing.T) {
	logrus.WithField("",1).Warn()
	logrus.Warn("123")
	bb()
}
func bb() {
	defer func() {
		if rec := recover(); nil != rec {
			msg := "%s"
			fmt.Println(fmt.Sprintf(msg, rec))
		}
	}()
    type A struct {
    	Name string
    }
	panic(A{
		Name: "asd",
	})
}
