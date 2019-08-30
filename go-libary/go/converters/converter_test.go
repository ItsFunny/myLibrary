/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 17:09 
# @File : converter_test.go
# @Description : 
*/
package converter

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestBytes2Int64(t *testing.T) {
	b := 9223372036854775807
	a := 18446744073709551615
	fmt.Println(a < b)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, a)
	fmt.Println(buf)
	// bytes := BigEndianInt642Bytes(int64(a))
	// fmt.Println(bytes)
	// bytes2Int64 := BigEndianBytes2Int64(bytes)
	// fmt.Println(bytes2Int64)
}
func TestBigEndianInt642Bytes(t *testing.T) {
	a := int64(0)
	bytes := BigEndianInt642Bytes(int64(a))
	fmt.Println(bytes)

}

func TestInt64ToBytes(t *testing.T) {
	m := uint64(3361437389424231001)
	bytes := BigEndianInt642Bytes(int64(m))
	fmt.Println(bytes)
	fmt.Println(len(bytes))
	a := make([]byte, 8)
	b := make([]byte, 8)
	c := make([]byte, 8)
	d := make([]byte, 8)
	a[0], a[1] = bytes[0], bytes[1]
	b[0], b[1] = bytes[2], bytes[3]
	c[0], c[1] = bytes[4], bytes[5]
	d[0], d[1] = bytes[6], bytes[7]
	fmt.Println(a)
	fmt.Println(BigEndianBytes2Int64(a))
	fmt.Println(b)
	fmt.Println(BigEndianBytes2Int64(b))
	fmt.Println(c)
	fmt.Println(BigEndianBytes2Int64(c))
	fmt.Println(d)
	fmt.Println(BigEndianBytes2Int64(d))
}

// 钱包记录相关交易信息
type BlockChainWalletTransRecordBO struct {
	WalletAddress string                            `json:"wallet_address"`
	SendCount     int                               `json:"send_count"`
	ReceiveCount  int                               `json:"receive_count"`
	Records       []BlockChainWalletTransRecordNode `json:"records"`
}

type BlockChainWalletTransRecordNode struct {
	TransactionID string `json:"transaction_id"`
	Type          int    `json:"type"` // 1 接收 2 发送
}

func TestFloat64ToByte(t *testing.T) {
	a := 0
	toByte := BigEndianFloat64ToByte(float64(a))
	fmt.Println(toByte)
	fmt.Println(BigEndianBytesToFloat64(toByte))
	str := `{"wallet_address":"fd8d298c1a5a3669562495670ffdd16d","send_count":1,"receive_count":0,"records":[{"transaction_id":"4588b8193aec7519dee114df654d5d7e4e5d2ba3c0fd653e075874abc81566e3","type":5}]}`
	bytes := make([]byte, 0)
	bytes = append(bytes, toByte...)
	bytes = append(bytes, []byte(str)...)

	modelBytes := bytes[8:]
	var aa BlockChainWalletTransRecordBO
	unmarshal := json.Unmarshal(modelBytes, &aa)
	if nil != unmarshal {
		panic(unmarshal)
	}
	fmt.Println(aa)

}
func TestBigEndianBytes2Int64(t *testing.T) {
	str := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDRt/aU7tAEHYjrjS4hLxCghMPs
MoLF5jcb5nCfOwQhrZbm3Ffkr2NDs6O7h/sZ5dS5gH51HMr7F6URCeHxZhe9bbjw
5thrY3wU6FJDkW1bLoSKEQGMO1GPDr9ks6IKmELF7b9bF0RKSUbpbzm57Dnl5d1Z
8jPmmY9lgffk7nqUSwIDAQAB
-----END PUBLIC KEY-----`
	fmt.Println(hex.EncodeToString([]byte(str)))
}
