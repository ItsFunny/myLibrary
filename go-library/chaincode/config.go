/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-16 13:50 
# @File : base.go
# @Description : 
# @Attention : 
*/
package cc

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"myLibrary/go-library/common"
	"myLibrary/go-library/common/blockchain"
	"myLibrary/go-library/common/blockchain/base"
	"myLibrary/go-library/go/base/service"

	error2 "myLibrary/go-library/common/error"
	"myLibrary/go-library/go/crypt"
	"myLibrary/go-library/go/services"
	"strconv"
)

type BaseConfigServiceImpl struct {
	*service.BaseServiceImpl
}

func NewBaseConfigServiceImpl() *BaseConfigServiceImpl {
	l := new(BaseConfigServiceImpl)
	l.BaseServiceImpl = service.NewBaseServiceImplWithLog4goLogger()
	return l
}

var (
	methods = []base.MethodName{
	}
	transDesc = []TransBaseDescription{
	}
	argsParameter = []ArgsParameter{
	}
)

var (
	keyConstants = []base.ObjectType{
	}
	keyGens = []base.KeyGenerater{
	}
	commonDecrypter = func(data interface{}, version string) (interface{}, error2.IBaseError) {
		return nil, nil
	}
)

func AddConfiguration(methodName []base.MethodName, methodDesc []TransBaseDescription, argParamter []ArgsParameter) {
	methods = append(methods, methodName...)
	transDesc = append(transDesc, methodDesc...)
	argsParameter = append(argsParameter, argParamter...)
}
func AddKey(ot []base.ObjectType, kg []base.KeyGenerater) {
	keyConstants = append(keyConstants, ot...)
	keyGens = append(keyGens, kg...)
}

func SetDecrypter(dec func(data interface{}, version string) (interface{}, error2.IBaseError)) {
	commonDecrypter = dec
}

// common 配置类,通过map存储参数等转换
type InMemoryBlockChainBaseConfiguration struct {
	// methods        []base.MethodName
	// argsParameters []ArgsParameter
	// transactionDescription    []TransBaseDescription
	//
	// keyConstants  []base.ObjectType
	// keyGenerators []base.KeyGenerater

	ArgsDecrypter       ArgsDecrypter
	ArgsCheckMap        map[base.MethodName]*ArgsParameter
	LogicDescriptionMap map[base.MethodName]TransBaseDescription

	// 多版本密钥管理
	MultiAESSecrets map[uint64]string

	// 用于生成key,是一个map结构,key是常量,value为生成方式
	BlockChainKeyContainer map[base.ObjectType]base.KeyGenerater

	*BaseConfigServiceImpl
}

func NewInMemoryBlockChainBaseConfiguration() *InMemoryBlockChainBaseConfiguration {
	c := new(InMemoryBlockChainBaseConfiguration)
	// argsParameter = argsParameter
	// methods = methods
	// transDesc = transDesc
	c.ArgsDecrypter = commonDecrypter
	// c.keyConstants = keyConstants
	// c.keyGenerators = keyGens
	c.BaseConfigServiceImpl = NewBaseConfigServiceImpl()

	// argsParameter, methods, transDesc, commonDecrypter, keyConstants, keyGens = nil, nil, nil, nil, nil, nil

	return c
}

func (c *InMemoryBlockChainBaseConfiguration) Config() error2.IBaseError {
	// l1, l2, l3 := len(methods), len(argsParameter), len(transDesc)
	l1, l2, l3 := len(methods), len(argsParameter), len(transDesc)
	if l1 == 0 {
		return error2.NewConfigError(nil, "配置参数method,argsParameter,transDesc不可为空")
	}
	if l1 != l2 || l2 != l3 || l1 != l3 {
		return error2.NewConfigError(nil, "method和参数转换以及config描述长度必须一致")
	}

	c.ArgsCheckMap = make(map[base.MethodName]*ArgsParameter)
	c.LogicDescriptionMap = make(map[base.MethodName]TransBaseDescription)
	for i := 0; i < l1; i++ {
		c.ArgsCheckMap[methods[i]] = &argsParameter[i]
		c.LogicDescriptionMap[methods[i]] = transDesc[i]
	}

	l4, l5 := len(keyConstants), len(keyGens)
	if l4 != l5 {
		return error2.NewConfigError(nil, "keyConstants和keyGenerators必须一致")
	}
	c.BlockChainKeyContainer = make(map[base.ObjectType]base.KeyGenerater)
	for i := 0; i < l4; i++ {
		fmt.Println(fmt.Sprintf("为key=[%v]的注册生成key函数 \n", keyConstants[i]))
		c.BlockChainKeyContainer[keyConstants[i]] = keyGens[i]
	}

	if c.MultiAESSecrets == nil {
		return error2.NewConfigError(nil, "多版本密钥不可为空")
	}

	argsParameter, methods, transDesc, commonDecrypter, keyConstants, keyGens = nil, nil, nil, nil, nil, nil

	return nil
}

func (b *InMemoryBlockChainBaseConfiguration) CheckAndConvt(method base.MethodName, args []string) (blockchain.BaseFabricAfterValidModel, error2.IBaseError) {
	var (
		result blockchain.BaseFabricAfterValidModel
	)

	fmt.Println("检查参数是否安全")

	if p, exist := b.ArgsCheckMap[method]; !exist {
		return result, error2.NewConfigError(nil, "配置错误")
	} else {
		if err := p.ArgsChecker(args); nil != err {
			return result, error2.NewArguError(err, "参数checker无法通过")
		}
		if res, baseError := p.ArgsConverter(args); nil != baseError {
			return result, error2.NewArguError(baseError, "参数转换无法通过")
		} else {
			// 判断参数是否实现了某个接口
			switch res.(type) {
			case services.IValidater:
				if e := res.(services.IValidater).Validate(); nil != e {
					return result, error2.NewArguError(e, "参数校验错误")
				}
			default:
				b.Error("最好事先Validater接口,该参数尚未实现该接口,method=[%s]", method)
			}
			// 参数判断是否需要转换
			switch res.(type) {
			case common.ICrypter:
				//  0 为参数 1 为version
				d, e := res.(common.ICrypter).Decrypt(args[1])
				if nil != e {
					return result, error2.NewArguError(e, "参数转换失败")
				}
				result.Req = d
			default:
				result.Req = res
			}
			// 获取configType
			if configType, exist := b.LogicDescriptionMap[method]; !exist {
				return result, error2.NewConfigError(nil, "配置错误,configType未配置")
			} else {
				result.BaseTransType = configType.TransBaseType
				result.BaseTransDescription = configType.Description
			}
			v, err := strconv.ParseInt(args[1], 10, 64)
			if nil != err {
				return result, error2.NewArguError(nil, "参数错误,版本号错误:"+args[1])
			}
			result.Version = uint64(v)

			return result, nil
		}
	}

}
func (b *InMemoryBlockChainBaseConfiguration) Encrypt(valueBytes []byte, version uint64) ([]byte, error2.IBaseError) {
	// TODO
	if key, exist := b.MultiAESSecrets[version]; !exist {
		return nil, error2.NewConfigError(nil, "配置错误,无法找到匹配的密钥")
	} else {
		bytes, e := encrypt.AesEncrypt(valueBytes, []byte(key))
		if nil != e {
			return nil, error2.NewArguError(e, "加密错误")
		}
		return bytes, nil
	}
}
func (b *InMemoryBlockChainBaseConfiguration) Decrypt(encData []byte, version uint64) ([]byte, error2.IBaseError) {
	if key, exist := b.MultiAESSecrets[version]; !exist {
		return nil, error2.NewConfigError(nil, "配置错误,无法找到匹配的密钥")
	} else {
		bytes, e := encrypt.AesDecrypt(encData, []byte(key))
		if nil != e {
			return nil, error2.NewArguError(e, "解密错误")
		}
		return bytes, nil
	}
}

func (c *InMemoryBlockChainBaseConfiguration) GetKey(stub shim.ChaincodeStubInterface, key base.ObjectType, req ...interface{}) (string, error2.IBaseError) {
	if g, exist := c.BlockChainKeyContainer[key]; !exist {
		return "", error2.NewConfigError(nil, fmt.Sprintf("key=[%v]的生成key函数不存在", key))
	} else {
		return g(stub, key, req...)
	}
}

// 
// 
// func Encrypt(data []byte, version uint64) ([]byte, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return nil, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.Encrypt(data, version)
// }
// 
// func Decrypt(req []byte, version uint64) ([]byte, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return nil, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.Decrypt(req, version)
// }
// 
// func GetKey(stub shim.ChaincodeStubInterface, key base.ObjectType, args ...interface{}) (string, error2.IBaseError) {
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return "", error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.GetKey(stub, key, args...)
// }
// 
// func ValidateArguAndReturn(name base.MethodName, args []string) (base.BaseFabricAfterValidModel, error2.IBaseError) {
// 	fmt.Println("开始调用config的参数检查")
// 	utils.DebugPrintDetail("=", "配置", InMemoryBlockChainBaseConfiguration)
// 	if nil == InMemoryBlockChainBaseConfiguration {
// 		if e := Config(); nil != e {
// 			return base.BaseFabricAfterValidModel{}, error2.NewConfigError(e, "初始化config错误")
// 		}
// 	}
// 	return InMemoryBlockChainBaseConfiguration.CheckAndConvt(name, args)
// }
