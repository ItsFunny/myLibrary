package base

import (
	"crypto"
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/base/helper"
	"reflect"
	"strconv"
)

var (
// defaultOptsMap = map[SerializableInterface]IntegrationOpts{
// 	ECDSA: DefaultECDSAIntegrationOpts{},
// 	base.SM2:   DefaultSM2IntegrationOpts{},
// 	constants.RSA:   DefaultRSAIntergrationOpts{},
// }
)

// func RegisterOpts(algInterface SerializableInterface,opts IntegrationOpts){
// 	defaultOptsMap[algInterface]=opts
// }

type Type int

type BaseType Type

func (t BaseType) EqualsBaseType(alg BaseType) bool {
	return t == BaseType(alg)
}
func (t BaseType) Contains(value BaseType, minExcepted BaseType) bool {
	if t&value >= minExcepted {
		return true
	}
	return false
}

type DetailType Type

type BaseOptsWrapperBuilder func() BaseOptsWrapper

type BaseTypeFunc func() BaseType

type DetailTypeFunc func() DetailType

type FailStrategy func() bool

type TypeValue int64

type TypeFul interface {
	GetType() TypeValue
}

type StateFul interface {
}

type AlgorithmBaseTemplate struct {
	*BaseTemplate
	BaseType BaseType
}

// 分级template
type AlgorithmHieraTemplate struct {
	*AlgorithmBaseTemplate
	detailHandler IDetailTemplate
}

type AlgorithmHieraManager struct {
	AlgorithmHieraTemplate IHieraTemplate
}

func (this *AlgorithmHieraManager) GetHieraHandlerByBaseType(baseType, minConpareType BaseType) IHieraTemplate {
	for tmp := this.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(IHieraTemplate) {
		if tmp.ValidIsMine(baseType, minConpareType) {
			return tmp
		}
		if tmp.GetNext() == nil {
			break
		}
	}
	panic("找不到对应的template")
}

func NewAlgorithmHieraManager(baseType BaseType) *AlgorithmHieraManager {
	return &AlgorithmHieraManager{
		AlgorithmHieraTemplate: newAlgorithmHieraTemplate(baseType),
	}
}

func (this *AlgorithmHieraManager) LinkHiera(baseType BaseType, detailTemplate IDetailTemplate) {
	if this.AlgorithmHieraTemplate == nil {
		this.AlgorithmHieraTemplate = newAlgorithmHieraTemplate(baseType)
		this.AlgorithmHieraTemplate.SetDetailTemplate(detailTemplate)
		return
	}
	tmp := this.AlgorithmHieraTemplate
	for {
		if nil == tmp.GetNext() || tmp.ValidIsMine(baseType, baseType) {
			break
		} else if nil != tmp.GetNext() {
			tmp = tmp.GetNext().(IHieraTemplate)
		}
	}
	if tmp.ValidIsMine(baseType, baseType) {
		if nil == tmp.GetDetailTemplate() {
			tmp.SetDetailTemplate(detailTemplate)
		} else {
			tmp.GetDetailTemplate().LinkLast(detailTemplate)
		}
	} else {
		newBaseTemplate := newAlgorithmHieraTemplate(baseType)
		newBaseTemplate.detailHandler = detailTemplate
		tmp.LinkLast(newBaseTemplate)
	}
}

func (this *AlgorithmHieraTemplate) SetDetailTemplate(detailTemplate IDetailTemplate) {
	this.detailHandler = detailTemplate
}

func (this *AlgorithmHieraTemplate) GetDetailTemplate() IDetailTemplate {
	return this.detailHandler
}

type AlgorithmDetailTemplate struct {
	*BaseDetailTemplate
}

func newAlgorithmBaseTemplate(BaseType BaseType) *AlgorithmBaseTemplate {
	return &AlgorithmBaseTemplate{
		BaseTemplate: &BaseTemplate{
			Template: nil,
			next:     nil,
			toString: "算法为:[" + GetAlgorithmDescription(BaseType) + "]",
		},
		BaseType: BaseType,
	}
}

func NewAlgorithmDetailTemplate(detailType DetailType) *AlgorithmDetailTemplate {
	return &AlgorithmDetailTemplate{
		BaseDetailTemplate: &BaseDetailTemplate{
			IDetailTemplate: nil,
			next:            nil,
			toString:        "detailType:[" + strconv.Itoa(int(detailType)) + "]",
			DetailType:      detailType,
		},
	}
}

func (this *AlgorithmBaseTemplate) ValidIsMine(compareValue BaseType, minExcepted BaseType) bool {
	return BaseType(this.BaseType).Contains(compareValue, minExcepted)
}

func (this *AlgorithmBaseTemplate) String() string {
	return this.toString
}

// func LinkHiera(baseType BaseType, cur IHieraTemplate, detailTemplate IDetailTemplate) IHieraTemplate {
// 	if cur == nil {
// 		cur = NewAlgorithmHieraTemplate(baseType)
// 		cur.SetDetailTemplate(detailTemplate)
// 		return cur
// 	}
// 	if cur.ValidIsMine(baseType, baseType) {
// 		cur.GetDetailTemplate().LinkLast(detailTemplate)
// 	} else {
// 		tmp := cur.GetNext()
// 		for ; nil != tmp.GetNext() || tmp.ValidIsMine(baseType, baseType); tmp = tmp.GetNext() {
// 		}
// 		if tmp.GetNext() == nil {
// 			newBaseTemplate := NewAlgorithmHieraTemplate(baseType)
// 			newBaseTemplate.DetailHandler = detailTemplate
// 			tmp.SetNext(newBaseTemplate)
// 		} else {
// 			tmp.(HieraTemplate).GetDetailTemplate().LinkLast(detailTemplate)
// 		}
// 	}
// 	return cur
// }

func newAlgorithmHieraTemplate(baseType BaseType) *AlgorithmHieraTemplate {
	return &AlgorithmHieraTemplate{
		AlgorithmBaseTemplate: newAlgorithmBaseTemplate(baseType),
	}
}

// func GetOptsByBaseType(alg BaseType) IntegrationOpts {
// 	return defaultOptsMap[alg]
// }

type BaseOptsWrapper struct {
	CommonBaseOpts CommonBaseOpts
	FailStrategy   FailStrategy
}

type CommonBaseOpts struct {
	AlgorithmType BaseType
	DetailType    DetailType
	MinExcepted   BaseTypeFunc
}

func (c CommonBaseOpts) GetMinExcepted() BaseTypeFunc {
	return c.MinExcepted
}

func (c CommonBaseOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(c.AlgorithmType)
	}
}

func (c CommonBaseOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return c.DetailType
	}
}

type BaseOpts interface {
	// parent type 类型 , 返回的是 GM,ECDSA,RSA
	GetBaseType() BaseTypeFunc
	// 详情类型 ,返回的是 SM2_256/SM2_512 ECDSA_256 等具体的算法
	GetDetailType() DetailTypeFunc
}

type CertificateBaseOpts interface {
	BaseOpts
	GetMinExcepted() BaseTypeFunc
}

type DefaultCertificateBaseOptsImpl struct {
	DetailType DetailType
	BaseType   BaseType
}

func (d DefaultCertificateBaseOptsImpl) GetMinExcepted() BaseTypeFunc {
	return d.GetBaseType()
}

func (d DefaultCertificateBaseOptsImpl) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return d.BaseType
	}
}

func (d DefaultCertificateBaseOptsImpl) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return d.DetailType
	}
}

type ECDSACertificateOpts struct {
}

func (this *ECDSACertificateOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(ECDSA)
	}
}

func (this *ECDSACertificateOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(ECDSA)
	}
}

type SM2CertificateOpts struct {
}

func (this *SM2CertificateOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(SM2)
	}
}

func (this *SM2CertificateOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(SM2)
	}
}

// same: minExcepted 和 typeFunc都是相同的
type EcdsaSameOpts struct {
}

func (EcdsaSameOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(ECDSA)
	}
}

func (c EcdsaSameOpts) GetMinExcepted() BaseTypeFunc {
	return c.GetBaseType()
}
func (EcdsaSameOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(ECDSA)
	}
}

type SM2SameOpts struct {
}

func (c SM2SameOpts) GetMinExcepted() BaseTypeFunc {
	return c.GetBaseType()
}
func (SM2SameOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(SM2)
	}
}

func (SM2SameOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(SM2)
	}
}

type RSASameOpts struct {
}

func (c RSASameOpts) GetMinExcepted() BaseTypeFunc {
	return c.GetBaseType()
}
func (RSASameOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(RSA)
	}
}

func (RSASameOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(RSA)
	}
}

type AESOpts struct {
}

func (c AESOpts) GetMinExcepted() DetailTypeFunc {
	return c.GetDetailType()
}

func (AESOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return DetailType(AES)
	}
}

func (AESOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return BaseType(AES)
	}
}

type DefaultOpts struct {
	BaseType   BaseType
	DetailType DetailType
	Hash       crypto.Hash
}

func (this DefaultOpts) HashFunc() crypto.Hash {
	return this.Hash
}

func (this DefaultOpts) GetBaseType() BaseTypeFunc {
	return func() BaseType {
		return this.BaseType
	}
}

func (this DefaultOpts) GetDetailType() DetailTypeFunc {
	return func() DetailType {
		return this.DetailType
	}
}

var FailFast = func() bool {
	return true
}
var FailOver = func() bool {
	return false
}

func CycledToJsonBytes(list interface{}) ([]byte, error) {
	value, err := helper.ToValue(list)
	if nil != err {
		return nil, err
	}
	return json.Marshal(value)
}

func JsonBytes2Cycled(bytes []byte, res interface{}) error {
	of := reflect.ValueOf(res)
	if of.Kind() != reflect.Ptr {
		return errors.New("结果必须为指针")
	}
	value := &helper.Value{}
	err := json.Unmarshal(bytes, value)
	if nil != err {
		return err
	}
	return helper.FromValue(value, res)
}
