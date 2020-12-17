/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-17 23:24 
# @File : gmutils_test.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func Test_AABB(t *testing.T) {
	// 	certStr:=`-----BEGIN CERTIFICATE-----
	// MIICDzCCAbagAwIBAgIQb3eHVmSnKC+CW+CHJ4JsdzAKBggqhkjOPQQDAjBjMQsw
	// CQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy
	// YW5jaXNjbzERMA8GA1UEChMIb3JnMS5jb20xFDASBgNVBAMTC2NhLm9yZzEuY29t
	// MB4XDTIwMTIxNzExMDUwMFoXDTMwMTIxNTExMDUwMFowYjELMAkGA1UEBhMCVVMx
	// EzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDTAL
	// BgNVBAsTBHBlZXIxFzAVBgNVBAMTDnBlZXIwLm9yZzEuY29tMFkwEwYHKoZIzj0C
	// AQYIKoZIzj0DAQcDQgAEWoa3Yrr7M7idxZ5AYtrYYzVZ0TaPSiZjgr94iWFlg6m8
	// FictYi0BSbd7l2yYs3QxUzPCxpnJslwvZ93RsC7WVqNNMEswDgYDVR0PAQH/BAQD
	// AgeAMAwGA1UdEwEB/wQCMAAwKwYDVR0jBCQwIoAgfz5VjkDl1WHAgbUINyv+0Fm+
	// JEePd35r4y3TQNJutiwwCgYIKoZIzj0EAwIDRwAwRAIgO26srpCuMQ88Yi4iQ1em
	// 6sTNxRlwJO20OQ5IvsS6+WMCIBH3BNIC/SNIyAxg7s2hz0nrN6KFGFQW78NZcGhA
	// tvBa
	// -----END CERTIFICATE-----
	// `
}

func Test_BB(t *testing.T) {
	// prvStr := "ZCgxpkpJZ4+rcu3JK8SU+NXUBjzKdDgdEMx86vf8kJE="
	// pubStr := "BI+U3ScwpyxzrwnEKuLIpnv17hh/ycM8ZuFNzoljdl/6gIg2IZNJmtsFhKhcjYr6zsccyvhgTqusv4hqSo5+Dog="
	prvStr := "fjOgXrYnP1xkVYeCQrQyVseX6ZBbGBh/FI0Y7ezO1Ms="
	pubStr := "BL/xA/A5q0kK9eKR0oHplDCsPGpApa2UAbx1QTj5p8peuzBeBr9JUEmdKstmKDHlgvUFcCGZ2k5dWKqmlvr8xbs="

	prvK := parseSm2PrvKeyFromString(prvStr)
	pubK := parsePubKeyFromString(pubStr)
	bs := []byte("bidsun-3001")
	uid := []byte("1234567812345678")
	s, e := SM2Sign(prvK, bs, uid)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println("signature为:" + s)
	// pubK=&prvK.PublicKey
	bytesX := prvK.PublicKey.X.Bytes()
	bytesY := prvK.PublicKey.Y.Bytes()
	fmt.Println("bXX:" + Base64Encode(bytesX))
	fmt.Println("bYY:" + Base64Encode(bytesY))
	b, e := SM2Verify(pubK, bs, uid, s)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(b)

}
func Test_versi(t *testing.T){
	origin:=[]byte("123")
	uid := []byte("1234567812345678")
	sig:="ZDhjNDBjZTIwNjllMTVjNWFkZDc5OGJkNjcwOTc5M2I2M2JjMDQxYjI2ZmNmOTk1MTRhNTQxNWRhMGM4NjI5NDcwYjk1YWUzYjE2NjQ3Njc4YThkOTE4YmQ0MmFlMDc3ZmNhN2YzNzczOGEwYTRkNDViNDg0N2RmMmI2YjFmMGI="
	pubStr := "BI+U3ScwpyxzrwnEKuLIpnv17hh/ycM8ZuFNzoljdl/6gIg2IZNJmtsFhKhcjYr6zsccyvhgTqusv4hqSo5+Dog="
	pubK := parsePubKeyFromString(pubStr)
	bytes, _ := Base64Decode(sig)

	b, e := SM2Verify(pubK, origin, uid, hex.EncodeToString(bytes))
	if nil!=e{
		log.Fatal(e)
	}
	fmt.Println(b)
}

