/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-23 13:52 
# @File : distributelock_test.go
# @Description : 
*/
package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestNewDistributeLock(t *testing.T) {
	conn, e := redis.Dial("tcp", "127.0.0.1:6379")
	if nil != e {
		panic(e)
	}
	defer conn.Close()

	l := NewDistributeLock(conn, 5)
	b, e := l.TryLock(conn, "asd", []byte("123"), 3, 4, false)
	if nil != e {
		panic(e)
	}
	fmt.Println(b)
	defer l.TryReleaseLock(conn, "asd")
}
func TestNewDistributeLock2(t *testing.T) {
	fmt.Println(time.Now().Unix())
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now().Unix())
}
