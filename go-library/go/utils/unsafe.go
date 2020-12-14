/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-11 11:46 
# @File : unsafe.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"reflect"
	"unsafe"
)

// 用法:
/*

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
 */
 // 就是在那个内存地址中new 一个变量
func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}