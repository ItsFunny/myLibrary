/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-11 11:46 
# @File : unsafe_test.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func Test_AA(t *testing.T) {
	crt := `-----BEGIN CERTIFICATE-----
MIICMTCCAdigAwIBAgIRANuluz10nyq+k+QQZcbP//AwCgYIKoZIzj0EAwIwYzEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xETAPBgNVBAoTCG9yZzEuY29tMRQwEgYDVQQDEwtjYS5vcmcxLmNv
bTAeFw0yMDEwMjExNDU0MDBaFw0zMDEwMTkxNDU0MDBaMGMxCzAJBgNVBAYTAlVT
MRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMREw
DwYDVQQKEwhvcmcxLmNvbTEUMBIGA1UEAxMLY2Eub3JnMS5jb20wWTATBgcqhkjO
PQIBBggqhkjOPQMBBwNCAARSZvnc67KPnTLmgFNO9mDAtgGi6vf+Gc5ey/4WOmiM
jpBxCgDMjqQLJz8m2FTKnNgeg45d5Ykp7s4NZdlffwuTo20wazAOBgNVHQ8BAf8E
BAMCAaYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMA8GA1UdEwEB/wQF
MAMBAf8wKQYDVR0OBCIEIIT5a/d0OcH7l0sppubpR3eUuGt5k3ZGFkKne3GTsQcb
MAoGCCqGSM49BAMCA0cAMEQCIHh1Y2eNRu39qjngcUUQc1WBEBiGTzSfYTpuBCl0
UoORAiAte8U1yyA2B9qKIKVJI7Z0/ufkiA+QUWt7qMgHQI4S7Q==
-----END CERTIFICATE-----
`

	var c *x509.Certificate


	pool := x509.NewCertPool()
	p, _ := pem.Decode([]byte(crt))
	certificate, e := x509.ParseCertificate(p.Bytes)
	if nil != e {
		log.Fatal(e)
	}
	pool.AddCert(certificate)

	rs := reflect.ValueOf(pool).Elem()
	rf := rs.Field(2)
	ri:=reflect.ValueOf(&c).Elem()

	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()


	rf.Bytes()
	ri.Bytes()
	}

func Test_ABS(t *testing.T){
	crt := `-----BEGIN CERTIFICATE-----
MIICMTCCAdigAwIBAgIRANuluz10nyq+k+QQZcbP//AwCgYIKoZIzj0EAwIwYzEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xETAPBgNVBAoTCG9yZzEuY29tMRQwEgYDVQQDEwtjYS5vcmcxLmNv
bTAeFw0yMDEwMjExNDU0MDBaFw0zMDEwMTkxNDU0MDBaMGMxCzAJBgNVBAYTAlVT
MRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMREw
DwYDVQQKEwhvcmcxLmNvbTEUMBIGA1UEAxMLY2Eub3JnMS5jb20wWTATBgcqhkjO
PQIBBggqhkjOPQMBBwNCAARSZvnc67KPnTLmgFNO9mDAtgGi6vf+Gc5ey/4WOmiM
jpBxCgDMjqQLJz8m2FTKnNgeg45d5Ykp7s4NZdlffwuTo20wazAOBgNVHQ8BAf8E
BAMCAaYwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMA8GA1UdEwEB/wQF
MAMBAf8wKQYDVR0OBCIEIIT5a/d0OcH7l0sppubpR3eUuGt5k3ZGFkKne3GTsQcb
MAoGCCqGSM49BAMCA0cAMEQCIHh1Y2eNRu39qjngcUUQc1WBEBiGTzSfYTpuBCl0
UoORAiAte8U1yyA2B9qKIKVJI7Z0/ufkiA+QUWt7qMgHQI4S7Q==
-----END CERTIFICATE-----
`



	pool := x509.NewCertPool()
	p, _ := pem.Decode([]byte(crt))
	certificate, e := x509.ParseCertificate(p.Bytes)
	if nil != e {
		log.Fatal(e)
	}
	pool.AddCert(certificate)

	rf:=reflect.ValueOf(pool)
	field := GetUnexportedField(rf.Elem().FieldByName("certs"))
	fmt.Println(field)
}