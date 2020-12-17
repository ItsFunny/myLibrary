/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package bccsp

import (
	"crypto/x509"
	"errors"
	"log"
	"math/big"
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/constants"
)

var (
	bccspContainer  *BccspContainer
)

type IKeyGenerator func() (Key, error)
type IKeyImporter func(raw, pwd []byte) (interface{}, error)
type IPublicKeyImporter func(raw interface{}) (Key, error)
type IOriginPublicKeyImporter func(bytes []byte) (interface{}, error)
type ISigner func(prvKInterface interface{}, msg []byte) ([]byte, error)
type IVerifier func(pubKeyInterface interface{}, signature, msg []byte) (bool, error)
type ICertificateImporter func(bytes []byte) (*x509.Certificate, base.CertificateBaseOpts, error)
type IUnmarshal func(sig []byte) (r, s *big.Int, err error)
type IMarshal func(r, s *big.Int) ([]byte, error)
type IHash func(msg []byte)(hash []byte, err error)
type IPrivateKeyGenerator func()(IPrivateKey,error)
// FIXME 添加一个生成证书的opts,返回值最好也还是一个接口 ,兼容ecdsa和sm2
type ICertificateGenerator func(params base.ParamBuilder)(*x509.Certificate,[]byte,error)

type CertificateGeneratorTemplate struct {
	*base.AlgorithmBaseTemplate
	 ICertificateGenerator ICertificateGenerator
}

type MarshalTemplate struct {
	*base.AlgorithmBaseTemplate
	IMarshal IMarshal
}

type UnmarshalTemplate struct {
	*base.AlgorithmBaseTemplate
	IUnmarshal IUnmarshal
}

type KeyImporterTemplate struct {
	*base.AlgorithmBaseTemplate
	IKeyImporter IKeyImporter
}

type KeyGeneratorTemplate struct {
	*base.AlgorithmBaseTemplate
	IKeyGenerator IKeyGenerator
}

type SignerTemplate struct {
	*base.AlgorithmBaseTemplate
	ISigner ISigner
}
type VerifierTemplate struct {
	*base.AlgorithmBaseTemplate
	IVerifier IVerifier
}

type CertificateImporterTemplate struct {
	*base.AlgorithmBaseTemplate
	ICertificateImporter ICertificateImporter
}

type PublicKeyImportTemplate struct {
	*base.AlgorithmBaseTemplate
	IPublicKeyImporter IPublicKeyImporter
}

type OriginPublicKeyImporter struct {
	*base.AlgorithmBaseTemplate
	IOriginPublicKeyImporter IOriginPublicKeyImporter
}

type HashTemplate struct {
	*base.AlgorithmBaseTemplate
	IHash IHash
}
type PrivateKeyGeneratorTemplate struct {
	*base.AlgorithmBaseTemplate
	IPrivateKeyGenerator IPrivateKeyGenerator
}

type BccspContainer struct {
	IKeyImporter             *KeyImporterTemplate
	IKeyGenerator            *KeyGeneratorTemplate
	ISigner                  *SignerTemplate
	IVerifier                *VerifierTemplate
	ICertificateImporter     *CertificateImporterTemplate
	IPublicKeyImporter       *PublicKeyImportTemplate
	IOriginPublicKeyImporter *OriginPublicKeyImporter
	IUnmarshal               *UnmarshalTemplate
	IMarshal                 *MarshalTemplate
	IHash *HashTemplate
	IPrivateKeyGenerator *PrivateKeyGeneratorTemplate
	ICertificateGenerator *CertificateGeneratorTemplate
}

func init() {
	bccspContainer = new(BccspContainer)
}
func RegisterCertificateGenerator(algInterface base.SerializableInterface,im ICertificateGenerator){
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &CertificateGeneratorTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			ICertificateGenerator:          im,
		}
		if bccspContainer.ICertificateGenerator == nil {
			bccspContainer.ICertificateGenerator = newTemplate
		} else {
			bccspContainer.ICertificateGenerator.LinkLast(newTemplate)
		}
	}
}
func RegisterPrivateKeyGenerator(algInterface base.SerializableInterface, importer IPrivateKeyGenerator) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &PrivateKeyGeneratorTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IPrivateKeyGenerator:          importer,
		}
		if bccspContainer.IPrivateKeyGenerator == nil {
			bccspContainer.IPrivateKeyGenerator = newTemplate
		} else {
			bccspContainer.IPrivateKeyGenerator.LinkLast(newTemplate)
		}
	}
}
func RegisterKeyImporter(algInterface base.SerializableInterface, importer IKeyImporter) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &KeyImporterTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IKeyImporter:          importer,
		}
		if bccspContainer.IKeyImporter == nil {
			bccspContainer.IKeyImporter = newTemplate
		} else {
			bccspContainer.IKeyImporter.LinkLast(newTemplate)
		}
	}
}
func RegisterKeyGenerator(algInterface base.SerializableInterface, generator IKeyGenerator) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &KeyGeneratorTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IKeyGenerator:         generator,
		}
		if bccspContainer.IKeyGenerator == nil {
			bccspContainer.IKeyGenerator = newTemplate
		} else {
			bccspContainer.IKeyGenerator.LinkLast(newTemplate)
		}
	}
}
func RegisterKeySigner(algInterface base.SerializableInterface, signer ISigner) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &SignerTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			ISigner:               signer,
		}
		if bccspContainer.ISigner == nil {
			bccspContainer.ISigner = newTemplate
		} else {
			bccspContainer.ISigner.LinkLast(newTemplate)
		}
	}

}
func RegisterKeyVerifier(algInterface base.SerializableInterface, verifier IVerifier) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &VerifierTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IVerifier:             verifier,
		}
		if bccspContainer.IVerifier == nil {
			bccspContainer.IVerifier = newTemplate
		} else {
			bccspContainer.IVerifier.LinkLast(newTemplate)
		}
	}
}
func RegisterCertificateImporter(algInterface base.SerializableInterface, certificateImporter ICertificateImporter) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &CertificateImporterTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			ICertificateImporter:  certificateImporter,
		}
		if bccspContainer.ICertificateImporter == nil {
			bccspContainer.ICertificateImporter = newTemplate
		} else {
			bccspContainer.ICertificateImporter.LinkLast(newTemplate)
		}
	}
}
func RegisterPublicKeyImporter(algInterface base.SerializableInterface, im IPublicKeyImporter) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &PublicKeyImportTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IPublicKeyImporter:    im,
		}
		if bccspContainer.IPublicKeyImporter == nil {
			bccspContainer.IPublicKeyImporter = newTemplate
		} else {
			bccspContainer.IPublicKeyImporter.LinkLast(newTemplate)
		}
	}
}
func RegisterOriginPublicKeyImporter(algInterface base.SerializableInterface, im IOriginPublicKeyImporter) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &OriginPublicKeyImporter{
			AlgorithmBaseTemplate:    base.NewAlgorithmBaseTemplate(alg),
			IOriginPublicKeyImporter: im,
		}
		if bccspContainer.IOriginPublicKeyImporter == nil {
			bccspContainer.IOriginPublicKeyImporter = newTemplate
		} else {
			bccspContainer.IOriginPublicKeyImporter.LinkLast(newTemplate)
		}
	}
}
func RegisterUnmarshal(algInterface base.SerializableInterface, im IUnmarshal) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &UnmarshalTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IUnmarshal:            im,
		}
		if bccspContainer.IUnmarshal == nil {
			bccspContainer.IUnmarshal = newTemplate
		} else {
			bccspContainer.IUnmarshal.LinkLast(newTemplate)
		}
	}
}
func RegisterMarshal(algInterface base.SerializableInterface, im IMarshal) {
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &MarshalTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IMarshal:              im,
		}
		if bccspContainer.IMarshal == nil {
			bccspContainer.IMarshal = newTemplate
		} else {
			bccspContainer.IMarshal.LinkLast(newTemplate)
		}
	}
}
func RegisterHash(algInterface base.SerializableInterface,im IHash){
	if alg, ok := algInterface.(constants.Algorithm); !ok {
		log.Fatal("register的时候必须为algorithm常量")
	} else {
		newTemplate := &HashTemplate{
			AlgorithmBaseTemplate: base.NewAlgorithmBaseTemplate(alg),
			IHash:              im,
		}
		if bccspContainer.IHash == nil {
			bccspContainer.IHash = newTemplate
		} else {
			bccspContainer.IHash.LinkLast(newTemplate)
		}
	}
}

func SignatureMarshal(r, s *big.Int, opts base.BaseOpts) ([]byte, error) {
	return bccspContainer.IMarshal.SignatureMarshal(r, s, opts)
}
// 反序列化
func SignatureUnmarshal(sig []byte, opts base.BaseOpts) (r, s *big.Int, e error) {
	return bccspContainer.IUnmarshal.SignatureUnmarshal(sig, opts)
}
func Sign(key Key, digest []byte, opts base.BaseOpts) (signature []byte, err error) {
	return bccspContainer.ISigner.Sign(key, digest, opts)
}
func Verify(key Key, digest, signature []byte, failStrategy base.FailStrategy, opts base.BaseOpts) error {
	if failStrategy() {
		return bccspContainer.IVerifier.Verify(key, digest, signature, opts);
	}

	for tmp := bccspContainer.IVerifier; nil != tmp; tmp = tmp.GetNext().(*VerifierTemplate) {
		if k, e := tmp.IVerifier(key, digest, signature); nil != e || !k {
			log.Printf("[Verify]:alg=%s,校验失败,下一个继续,失败原因:[%s]", tmp.String(),e.Error())
			continue
		} else {
			return nil
		}
	}
	return errors.New("遍历也无法校验成功")
}
func KeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts base.BaseOpts) (interface{}, error) {
	if failStrategy() {
		return bccspContainer.IKeyImporter.KeyImport(raw, pwd, opts)
	}
	for tmp := bccspContainer.IKeyImporter; nil != tmp; tmp = tmp.GetNext().(*KeyImporterTemplate) {
		if k, e := tmp.IKeyImporter(raw, pwd); nil != e {
			log.Printf("[KeyImport]alg=%s,导入失败,下一个继续,失败原因:[%s]", tmp.String(),e.Error())
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法import key,失败")
}
func PublicKeyImport(raw interface{}, failStrategy base.FailStrategy, opts base.BaseOpts) (Key, error) {
	if failStrategy() {
		return bccspContainer.IPublicKeyImporter.PublicKeyImport(raw, opts)
	}
	for tmp := bccspContainer.IPublicKeyImporter; nil != tmp; tmp = tmp.GetNext().(*PublicKeyImportTemplate) {
		if k, e := tmp.IPublicKeyImporter(raw); nil != e {
			log.Printf("[PublicKeyImport]alg=%s,导入失败,下一个继续,失败原因:[%s]", tmp.String(),e.Error())
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法PublicKeyImport,失败")
}
// 导入原先的key,如ecdsa.Publickey 或者是sm2.PublicKey等
func OriginPublicKeyImport(raw []byte, failStrategy base.FailStrategy, opts base.BaseOpts) (interface{}, error) {
	if failStrategy() {
		return bccspContainer.IOriginPublicKeyImporter.OriginPublicKeyImport(raw, opts)
	}
	for tmp := bccspContainer.IOriginPublicKeyImporter; nil != tmp; tmp = tmp.GetNext().(*OriginPublicKeyImporter) {
		if k, e := tmp.IOriginPublicKeyImporter(raw); nil != e {
			log.Printf("[OriginPublicKeyImport]alg=%s,导入失败,下一个继续,失败原因:[%s]", tmp.String(),e.Error())
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("解析失败,都无法解析原生公钥")
}
func ParseCertificate(raw []byte, failStrategy base.FailStrategy, opts base.BaseOpts) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if failStrategy() {
		return bccspContainer.ICertificateImporter.ParseCertificate(raw, opts)
	}
	for tmp := bccspContainer.ICertificateImporter; nil != tmp; tmp = tmp.GetNext().(*CertificateImporterTemplate) {
		if k, opts, e := tmp.ICertificateImporter(raw); nil != e {
			log.Printf("[ParseCertificate]alg=%s,导入失败,下一个继续,失败原因:[%s]", tmp.String(),e.Error())
			continue
		} else {
			return k, opts, nil
		}
	}
	return nil, nil, errors.New("解析失败,都无法解析证书")
}
func Hash(msg []byte,opts HashOptsAdapter)([]byte,error){
	return bccspContainer.IHash.Hash(msg,opts)
}
func GenPrivateKey(opts KeyGenOptsAdapter)(IPrivateKey,error){
	return bccspContainer.IPrivateKeyGenerator.KeyGen(opts)
}
func GenCertificate(opts base.CertificateBaseOpts,params base.ParamBuilder )(*x509.Certificate,[]byte,error){
	return bccspContainer.ICertificateGenerator.CertificateGen(opts,params)
}

func (this *KeyImporterTemplate) KeyImport(raw, pwd []byte, opts base.BaseOpts) (interface{}, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.IKeyImporter(raw, pwd)
	} else if nil != this.GetNext() {
		return this.GetNext().(*KeyImporterTemplate).KeyImport(raw, pwd, opts)
	} else {
		return nil, errors.New("找不到匹配的处理者")
	}
}
func (this *PublicKeyImportTemplate) PublicKeyImport(raw interface{}, opts base.BaseOpts) (Key, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.IPublicKeyImporter(raw)
	} else if nil != this.GetNext() {
		return this.GetNext().(*PublicKeyImportTemplate).PublicKeyImport(raw, opts)
	} else {
		return nil, errors.New("找不到匹配的处理者")
	}
}
func (this *OriginPublicKeyImporter) OriginPublicKeyImport(raw []byte, opts base.BaseOpts) (interface{}, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.IOriginPublicKeyImporter(raw)
	} else if nil != this.GetNext() {
		return this.GetNext().(*OriginPublicKeyImporter).OriginPublicKeyImport(raw, opts)
	} else {
		return nil, errors.New("OriginPublicKeyImporter:找不到匹配的处理者")
	}
}
func (this *CertificateImporterTemplate) ParseCertificate(bytes []byte, opts base.BaseOpts) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.ICertificateImporter(bytes)
	} else if nil != this.GetNext() {
		return this.GetNext().(*CertificateImporterTemplate).ParseCertificate(bytes, opts)
	} else {
		return nil, nil, errors.New("OriginPublicKeyImporter:找不到匹配的处理者")
	}
}
func (this *MarshalTemplate) SignatureMarshal(r *big.Int, s *big.Int, opts base.BaseOpts) ([]byte, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.IMarshal(r, s)
	} else if nil != this.GetNext() {
		return this.GetNext().(*MarshalTemplate).SignatureMarshal(r, s, opts)
	} else {
		return nil, errors.New("UnmarshalTemplate:找不到匹配的处理者")
	}
}
func (this *UnmarshalTemplate) SignatureUnmarshal(sig []byte, opts base.BaseOpts) (*big.Int, *big.Int, error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.IUnmarshal(sig)
	} else if nil != this.GetNext() {
		return this.GetNext().(*UnmarshalTemplate).SignatureUnmarshal(sig, opts)
	} else {
		return nil, nil, errors.New("UnmarshalTemplate:找不到匹配的处理者")
	}
}
func (this *SignerTemplate) Sign(key Key, digest []byte, opts base.BaseOpts) (signature []byte, err error) {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return this.ISigner(key, digest)
	} else if nil != this.GetNext() {
		return this.GetNext().(*SignerTemplate).Sign(key, digest, opts)
	} else {
		return nil, errors.New("SignerTemplate:找不到匹配的处理者")
	}
}
func (this *VerifierTemplate) Verify(key Key, digest, signature []byte, opts base.BaseOpts) error {
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		b, e := this.IVerifier(key, signature, digest)
		if nil != e {
			return e
		} else if !b {
			return errors.New("校验失败")
		}
		return nil
	} else if nil != this.GetNext() {
		return this.GetNext().(*VerifierTemplate).Verify(key, digest, signature, opts)
	} else {
		return errors.New("VerifierTemplate:找不到匹配的处理者")
	}
}
func (this *HashTemplate) Hash(msg []byte,opts HashOptsAdapter) ([]byte,error ){
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return  this.IHash(msg)
	} else if nil != this.GetNext() {
		return this.GetNext().(*HashTemplate).Hash(msg, opts)
	} else {
		return nil,errors.New("HashTemplate:找不到匹配的处理者")
	}
}
func(this *PrivateKeyGeneratorTemplate)KeyGen(opts KeyGenOptsAdapter)(IPrivateKey,error){
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return  this.IPrivateKeyGenerator()
	} else if nil != this.GetNext() {
		return this.GetNext().(*PrivateKeyGeneratorTemplate).KeyGen(opts)
	} else {
		return nil,errors.New("HashTemplate:找不到匹配的处理者")
	}
}

func(this *CertificateGeneratorTemplate)CertificateGen(opts base.CertificateBaseOpts,params base.ParamBuilder)(*x509.Certificate,[]byte,error){
	if this.ValidIsMine(opts.GetTypeFunc()()) {
		return  this.ICertificateGenerator(params)
	} else if nil != this.GetNext() {
		return this.GetNext().(*CertificateGeneratorTemplate).CertificateGen(opts,params)
	} else {
		return nil,nil,errors.New("CertificateGeneratorTemplate:找不到匹配的处理者")
	}
}
// func KeyGenerate() (Key, error) {
// 	alg := adapter.GetCurrentGoroutineAlg()
// 	if alg != constants.NONE {
// 		if gen, exist := bccspContainer.IKeyGenerator[alg]; !exist {
// 			return nil, errors.New("找不到匹配的生成器")
// 		} else {
// 			return gen()
// 		}
// 	}
// 	return nil, errors.New("找不到匹配的key生成器")
// }
