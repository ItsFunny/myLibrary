/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-17 23:38 
# @File : bccsp_container_test.go.go
# @Description : 
# @Attention : 
*/
package testdata

import (
	"encoding/hex"
	"log"
	"myLibrary/go-library/go/internal/bccsp"
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/internal/bccsp/opts"
	"testing"
)

func TestSKI(t *testing.T) {
	cert := `-----BEGIN CERTIFICATE-----
MIICDzCCAbagAwIBAgIQb3eHVmSnKC+CW+CHJ4JsdzAKBggqhkjOPQQDAjBjMQsw
CQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy
YW5jaXNjbzERMA8GA1UEChMIb3JnMS5jb20xFDASBgNVBAMTC2NhLm9yZzEuY29t
MB4XDTIwMTIxNzExMDUwMFoXDTMwMTIxNTExMDUwMFowYjELMAkGA1UEBhMCVVMx
EzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBGcmFuY2lzY28xDTAL
BgNVBAsTBHBlZXIxFzAVBgNVBAMTDnBlZXIwLm9yZzEuY29tMFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAEWoa3Yrr7M7idxZ5AYtrYYzVZ0TaPSiZjgr94iWFlg6m8
FictYi0BSbd7l2yYs3QxUzPCxpnJslwvZ93RsC7WVqNNMEswDgYDVR0PAQH/BAQD
AgeAMAwGA1UdEwEB/wQCMAAwKwYDVR0jBCQwIoAgfz5VjkDl1WHAgbUINyv+0Fm+
JEePd35r4y3TQNJutiwwCgYIKoZIzj0EAwIDRwAwRAIgO26srpCuMQ88Yi4iQ1em
6sTNxRlwJO20OQ5IvsS6+WMCIBH3BNIC/SNIyAxg7s2hz0nrN6KFGFQW78NZcGhA
tvBa
-----END CERTIFICATE-----`
	raw := []byte(cert)
	crt, _, e := bccsp.ParseCertificate(raw, base.FailOver, nil)
	if nil != e {
		log.Fatal(e)
	}
	bytes, e := bccsp.Hash(crt.Raw, opts.DefaultSHA256HashOptsImpl{})
	if nil != e {
		log.Fatal(e)
	}
	s := hex.EncodeToString(bytes)
	log.Println(s)
}

func TestAAA(t *testing.T) {
	// prvKeyStr := "FEYPhWXO8r+SuTGHH1sAHfUi7bWwZkbN5jtlDILQN3k="
	// // pubKeyStr := "公钥 205NGcj2F2/cB4YeOeMZUKURUJIOrLRoJOP/Yi88kHsAY6hcpmCz3mEzLLtNQJUTdv0mXaFVPDg0hk8j37i4/A=="
	// bytes, e := base64.StdEncoding.DecodeString(prvKeyStr)
	// if nil != e {
	// 	log.Fatal(e)
	// }
	// keyImport, e := bccsp.KeyImport(bytes, nil, base.FailFast, base.SM2Opts{})
	// if nil != e {
	// 	panic(e)
	// }
	// fmt.Println(keyImport)
	// pubKeyStr := "205NGcj2F2/cB4YeOeMZUKURUJIOrLRoJOP/Yi88kHsAY6hcpmCz3mEzLLtNQJUTdv0mXaFVPDg0hk8j37i4/A=="
	// pubBytes, _ := base64.StdEncoding.DecodeString(pubKeyStr)
	// msg := []byte("123")
	// k := keyImport.(*sm2.PrivateKey)
	// pubk, e := bccsp.PublicKeyImport(pubBytes, base.FailFast, base.SM2Opts{})
	// if nil != e {
	// 	log.Fatal(e)
	// }
	// fmt.Println(pubk)
}
