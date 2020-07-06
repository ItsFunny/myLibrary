/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 14:05 
# @File : args.go
# @Description : 
# @Attention : 
*/
package base

import (

	"github.com/hyperledger/fabric-chaincode-go/shim"
	error2 "myLibrary/go-library/blockchain/error"
	error3 "myLibrary/go-library/go/error"
)

var (
	COMMON_STRING_OR_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error3.IBaseError) {
		switch param[0].(type) {
		case string:
			return COMMON_STRING_KEY_GENERATOR(stub, objectType, param[0])
		case []string:
			return COMMON_STRING_ARRAY_KEY_GENERATOR(stub, objectType, param[0])
		}
		return "", error2.NewFabricError(nil, "找不到匹配的处理")
	}
	COMMON_STRING_ARRAY_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error3.IBaseError) {
		strings := param[0].([]string)
		s, e := stub.CreateCompositeKey(string(objectType), strings)
		if nil != e {
			return "", error2.NewFabricError(e, "创建组合键失败")
		}
		return s, nil
	}

	COMMON_STRING_KEY_GENERATOR = func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string, error3.IBaseError) {
		strings := param[0].(string)
		s, e := stub.CreateCompositeKey(string(objectType), []string{strings})
		if nil != e {
			return "", error2.NewFabricError(e, "创建组合键失败")
		}
		return s, nil
	}
)



type ArgsChecker = func(args []string) error3.IBaseError
type ArgsConverter = func(args []string) (interface{}, error3.IBaseError)
type ArgsDecrypter = func(data interface{}, version string) (interface{}, error3.IBaseError)

type ArgsParameter struct {
	ArgsChecker   ArgsChecker
	ArgsConverter ArgsConverter
}

type TransBaseDescription struct {
	TransBaseType TransBaseTypeV2
	Description   string
}