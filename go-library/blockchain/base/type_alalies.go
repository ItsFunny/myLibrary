/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 14:06 
# @File : type.go
# @Description : 
# @Attention : 
*/
package base

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"myLibrary/go-library/go/authentication"
	error2 "myLibrary/go-library/go/error"
)

type TransBaseType int
type ChannelID string
type OrganizationID string
type ChainCodeID string
type MethodName string
// 代表区块链上的key
type Key string
// fromWalletAddress 从哪个钱包过来的
type From string
// toWalletAddress 去往哪个钱包的
type To string
// token 交易coin
type Token float64
type Version uint64
type ObjectType string
type KeyGenerater func(stub shim.ChaincodeStubInterface, objectType ObjectType, param ...interface{}) (string,   error2.IBaseError)
type TransBaseTypeV2 authentication.Authority
type TransBaseTypeV2Value authentication.AuthValue
