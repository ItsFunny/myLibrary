/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/3/11 4:40 下午
# @File : proto_test.go
# @Description :
# @Attention :
*/
package test_data

import (
	"encoding/hex"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// 08011203010203
func Test_proto(t *testing.T) {
	v := BitArrayProto{
		Bits:  1,
		Elems: []uint64{1, 2, 3},
	}
	marshal, err := proto.Marshal(&v)
	if nil != err {
		log.Fatalln(err)
	}
	fmt.Println(hex.EncodeToString(marshal))
}

// 测试 proto 追加字段之后,旧数据能否成功转换
func Test_Proto2(t *testing.T) {
	hexData := "08011203010203"
	{
		prevData, _ := hex.DecodeString(hexData)
		v := BitArrayProto2{}
		if err := proto.Unmarshal(prevData, &v); nil != err {
			log.Fatal(err)
		}
		assert.Equal(t, v.Bits, int64(1))
		fmt.Println(v.Bits)
	}

	// 测试新的proto,生成新的数据,能不能用老的解析
	{
		v := BitArrayProto2{
			Bits:  1,
			Elems: []uint64{1, 2, 3},
			Name:  "joker",
		}
		marshal, err := proto.Marshal(&v)
		if nil != err {
			log.Fatal(err)
		}
		oldV := BitArrayProto{}
		if err = proto.Unmarshal(marshal, &oldV); nil != err {
			log.Fatal(err)
		}
		assert.Equal(t, int64(1), oldV.Bits)
		fmt.Println("老数据结构能兼容新数据解析")
	}
}

func Test_Channel(t *testing.T) {
	c := make(chan int, 10)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()
	go func() {
		c <- 1
		c <- 2
	}()

	time.Sleep(time.Second * 10)
}
