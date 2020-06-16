/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:41 
# @File : base.go
# @Description : 
# @Attention : 
*/
package cc

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"myLibrary/go-library/go/base/service"
	error3 "myLibrary/go-library/go/error"
)

type IBaseChainCode interface {
	shim.Chaincode
	Config() error3.IBaseError
}

type IConcreteChainCode interface {
	InitDetail(stub shim.ChaincodeStubInterface) error3.IBaseError
	InvokeDetail(stub shim.ChaincodeStubInterface)BasePeerResponse
	ConfigDetail() error3.IBaseError
}


type BaseChainCodeServiceImpl struct {
	*service.BaseServiceImpl
	ConcreteChainCode IConcreteChainCode
}

func NewBaseChainCodeServiceImpl(ConcreteChainCode IConcreteChainCode) *BaseChainCodeServiceImpl {
	v := new(BaseChainCodeServiceImpl)
	v.BaseServiceImpl = service.NewBaseServiceImplWithLog4goLogger()
	v.ConcreteChainCode = ConcreteChainCode
	return v
}

// 启动时候的调用,发生在init之前,当跨链
// func (v *BaseChainCodeServiceImpl) BootStrapConfig() peer.Response {
// }

func (v *BaseChainCodeServiceImpl) Init(stub shim.ChaincodeStubInterface) peer.Response {
	e := v.ConcreteChainCode.InitDetail(stub)
	if nil != e {
		return shim.Error(e.Error())
	}
	e = v.ConcreteChainCode.ConfigDetail()
	if nil != e {
		return shim.Error(e.Error())
	}

	return shim.Success(nil)
}

func (v *BaseChainCodeServiceImpl) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	resp := v.ConcreteChainCode.InvokeDetail(stub)
	return resp.Response
}

func (v *BaseChainCodeServiceImpl) Config() error3.IBaseError {
	return v.ConcreteChainCode.ConfigDetail()
}

// type IChaincodeFacadedService interface {
// 	log.Logger
// 	ValidateArguAndReturn(method base.MethodName, args []string) (models.BaseFabricAfterValidModel, error3.IBaseError)
// }

// 参数加解密

// type IArgumentDecrypt interface {
// 	Decrypt(argu interface{}, version string) (interface{}, error3.IBaseError)
// 	SetParent(c IChaincodeFacadedService)
// }


// version,from
type BasePeerResponse struct {
	peer.Response
	// Version             uint64
	// From                From
	// To                  To
	// Token               Token
	// BaseTransactionType TransBaseType
}