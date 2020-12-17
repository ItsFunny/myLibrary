package base

import (
	"fmt"
	"myLibrary/go-library/go/constants"
)

var (
	// defaultOptsMap = map[SerializableInterface]IntegrationOpts{
	// 	constants.ECDSA: DefaultECDSAIntegrationOpts{},
	// 	constants.SM2:   DefaultSM2IntegrationOpts{},
	// 	constants.RSA:   DefaultRSAIntergrationOpts{},
	// }
)

// func RegisterOpts(algInterface SerializableInterface,opts IntegrationOpts){
// 	defaultOptsMap[algInterface]=opts
// }


type BaseOptsWrapperBuilder func() BaseOptsWrapper
type SerializableInterface interface{}
type SerializableFunc func() SerializableInterface
type FailStrategy func() bool

type Template interface {
	fmt.Stringer
	ValidIsMine(serializableInterface SerializableInterface) bool
	LinkLast(template Template)
	GetNext() Template
	SetNext(template Template)
}

type BaseTemplate struct {
	Template
	next     Template
	toString string
}

func (this *BaseTemplate) GetNext() Template {
	return this.next
}

func (this *BaseTemplate) SetNext(template Template) {
	this.next = template
}

func (this *BaseTemplate) LinkLast(template Template) {
	if nil == this.next {
		this.next = template
		return
	}
	tmp := this.next
	for ; nil != tmp.GetNext(); tmp = tmp.GetNext() {
	}
	tmp.SetNext(template)
}

type AlgorithmBaseTemplate struct {
	*BaseTemplate
	Algorithm constants.Algorithm
}

func NewAlgorithmBaseTemplate(algorithm constants.Algorithm) *AlgorithmBaseTemplate {
	return &AlgorithmBaseTemplate{
		BaseTemplate: &BaseTemplate{
			toString: "算法为:[" + constants.GetAlgorithmDescription(algorithm) + "]",
		},
		Algorithm: algorithm,
	}
}

func (this *AlgorithmBaseTemplate) ValidIsMine(serializableInterface SerializableInterface) bool {
	return this.Algorithm == serializableInterface
}

func (this *AlgorithmBaseTemplate) String() string {
	return this.toString
}

// func GetOptsByAlgorithm(alg constants.Algorithm) IntegrationOpts {
// 	return defaultOptsMap[alg]
// }

type BaseOptsWrapper struct {
	CommonBaseOpts CommonBaseOpts
	FailStrategy   FailStrategy
}

type CommonBaseOpts struct {
	Algorithm constants.Algorithm
}

func (c CommonBaseOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return c.Algorithm
	}
}

func (c CommonBaseOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return c.Algorithm
	}
}

type BaseOpts interface {
	GetTypeFunc() SerializableFunc
	GetBaseType()SerializableFunc
}

type CertificateBaseOpts interface {
	BaseOpts
}
type ECDSACertificateOpts struct {
}

func (this *ECDSACertificateOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.ECDSA
	}
}


func (this *ECDSACertificateOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.ECDSA
	}
}

type SM2CertificateOpts struct {
}

func (this *SM2CertificateOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.SM2
	}
}

func (this *SM2CertificateOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.SM2
	}
}

type EcdsaOpts struct {
}

func (EcdsaOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.ECDSA
	}
}





func (EcdsaOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.ECDSA
	}
}

type SM2Opts struct {
}

func (SM2Opts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.SM2
	}
}

func (SM2Opts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.SM2
	}
}

type RSAOpts struct {
}

func (RSAOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.RSA
	}
}


func (RSAOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.RSA
	}
}

type AESOpts struct {
}

func (AESOpts) GetTypeFunc() SerializableFunc {
	return func() SerializableInterface {
		return constants.AES
	}
}

func (AESOpts) GetBaseType() SerializableFunc {
	return func() SerializableInterface {
		return constants.AES
	}
}

var FailFast = func() bool {
	return true
}
var FailOver = func() bool {
	return false
}




// 即包含了签名opts,也包含了校验opts ,
// type IntegrationOpts interface {
// 	SignerOptsAdapter
	// crypto.SignerOpts
	// BaseOpts
// }
//
// type DefaultIntegrationOpts struct {
//
// }
//
// func (this DefaultIntegrationOpts) HashFunc() crypto.Hash {
// 	panic("implement me")
// }
//
// func (this DefaultIntegrationOpts) GetBaseType() SerializableFunc {
// 	return func() SerializableInterface {
// 		return constants.ECDSA
// 	}
// }
//
// func (this DefaultIntegrationOpts) GetTypeFunc() SerializableFunc {
// 	return func() SerializableInterface {
// 		return constants.ECDSA
// 	}
// }

