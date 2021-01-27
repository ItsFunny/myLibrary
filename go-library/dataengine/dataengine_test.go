/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-26 14:44 
# @File : dataengine_test.go
# @Description : 
# @Attention : 
*/
package dataengine

import (
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"testing"
)

func TestDataEngine(t *testing.T) {
	// jsonStr:=`{"_objectId":"haha_BFB2A85200000000165DB56561D2516E","_schema":"haha","_zoneId":"default","a":"t1","b":123}`
	bytes, e := ioutil.ReadFile("/Users/joker/go/src/myLibrary/go-library/test_data/json.json")
	if nil != e {
		panic(e)
	}
	jsonStr := string(bytes)
	bytes, e = ioutil.ReadFile("/Users/joker/go/src/myLibrary/go-library/test_data/json.schemal")
	if nil != e {
		panic(e)
	}
	schemalLoader := gojsonschema.NewStringLoader(string(bytes))
	// schema, e := gojsonschema.NewSchema(schemalLoader)
	// if nil!=e{
	// 	panic(e)
	// }
	dataLoader := gojsonschema.NewStringLoader(jsonStr)
	_, e = gojsonschema.Validate(schemalLoader, dataLoader)
	if nil != e {
		panic(e)
	}
}
