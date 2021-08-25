/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 16:18
# @File : bccsp_register.go
# @Description :    bccsp 静态方法
# @Attention :
*/
package bccsp

import (
	"crypto/x509"
	"errors"
	"github.com/cloudflare/cfssl/certdb"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
	"github.com/hyperledger/fabric/base"
	"log"
	"math/big"
)

var _ BccspManager = &DefaultChainAbleBccspManager{}

type BccspManager interface {
	SignatureMarshal(r, s *big.Int, opts SignatureMarshalOptsAdapter) ([]byte, error)
	SignatureUnmarshal(sig []byte, opts SignatureUnMarshalOptsAdapter) (r, s *big.Int, e error)
	Sign(key Key, digest []byte, opts SignerOptsAdapter) (signature []byte, err error)
	Verify(key Key, digest, signature []byte, failStrategy base.FailStrategy, opts VerifierOptsAdapter) error
	InterfaceKeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts KeyImportOptsAdapter) (PrivateKeyWrapper, error)
	PublicKeyImport(raw interface{}, failStrategy base.FailStrategy, opts PublicKeyImportOptsAdapter) (KeyAdapter, error)
	OriginPublicKeyImport(raw []byte, failStrategy base.FailStrategy, opts OriginPublicKeyImportOptsAdapter) (interface{}, error)
	ParseCertificate(raw []byte, failStrategy base.FailStrategy, opts CertificateImportOptsAdapter) (*x509.Certificate, base.CertificateBaseOpts, error)
	Hash(msg []byte, opts HashOptsAdapter) ([]byte, error)
	GenPrivateKey(opts KeyGenOptsAdapter) (KeyAdapter, error)
	GenCertificate(opts CertificateGeneratorOpts, params base.ParamBuilder) (*x509.Certificate, []byte, error)
	GenCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CSRGenerateOpts) ([]byte, error)
	GenCRTWithCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CRTWithCsrGeneratorOpts, wrapper CrtCsrWrapper) ([]byte, []byte, error)
	ParseCSR(raw interface{}, failStrategy base.FailStrategy, opts CSRParserOpts) (*x509.CertificateRequest, error)
	GenSignCert(keyAdapter KeyAdapter, rootCa *x509.Certificate, req signer.SignRequest, opts SignCertGenOpts, others ...interface{}) (*certdb.CertificateRecord, []byte, error)
}

type DefaultChainAbleBccspManager struct {
	IKeyImporter             *DefaultMediator
	IKeyGenerator            *DefaultMediator
	ISigner                  *DefaultMediator
	IVerifier                *DefaultMediator
	ICertificateImporter     *DefaultMediator
	IPublicKeyImporter       *DefaultMediator
	IOriginPublicKeyImporter *DefaultMediator
	IUnmarshal               *DefaultMediator
	IMarshal                 *DefaultMediator
	IHash                    *DefaultMediator
	IPrivateKeyGenerator     *DefaultMediator
	ICertificateGenerator    *DefaultMediator
	ICSRGenerator            *DefaultMediator
	ICRTWithCsrGenerator     *DefaultMediator
	ICSRParser               *DefaultMediator
	ISignCertGenerator       *DefaultMediator
}

func (this *DefaultChainAbleBccspManager) GenSignCert(keyAdapter KeyAdapter, rootCa *x509.Certificate, req signer.SignRequest, opts SignCertGenOpts, others ...interface{}) (*certdb.CertificateRecord, []byte, error) {
	handler := this.ISignCertGenerator.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).GenSignCert(keyAdapter, rootCa, req, opts, others...)
}

func (this *DefaultChainAbleBccspManager) ParseCSR(raw interface{}, failStrategy base.FailStrategy, opts CSRParserOpts) (*x509.CertificateRequest, error) {
	if failStrategy() {
		handler := this.ICSRParser.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).ParseCSR(raw, failStrategy, opts)
	}
	for tmp := this.ICSRParser.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if k, e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).ParseCSR(raw, failStrategy, opts); nil != e {
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法parse,失败")
}
func (this *DefaultChainAbleBccspManager) GenCRTWithCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CRTWithCsrGeneratorOpts, wrapper CrtCsrWrapper) ([]byte, []byte, error) {
	handler := this.ICRTWithCsrGenerator.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).GenCRTWithCSR(adapter, req, opts, wrapper)
}

func (this *DefaultChainAbleBccspManager) GenCSR(adapter KeyAdapter, req *csr.CertificateRequest, opts CSRGenerateOpts) ([]byte, error) {
	handler := this.ICSRGenerator.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).GenCSR(adapter, req, opts)
}

func (this *DefaultChainAbleBccspManager) SignatureMarshal(r, s *big.Int, opts SignatureMarshalOptsAdapter) ([]byte, error) {
	handler := this.IMarshal.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).SignatureMarshal(r, s, opts)
}

func (this *DefaultChainAbleBccspManager) SignatureUnmarshal(sig []byte, opts SignatureUnMarshalOptsAdapter) (r, s *big.Int, e error) {
	handler := this.IUnmarshal.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).SignatureUnmarshal(sig, opts)
}

func (this *DefaultChainAbleBccspManager) Sign(key Key, digest []byte, opts SignerOptsAdapter) (signature []byte, err error) {
	handler := this.ISigner.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).Sign(key, digest, opts)
}

func (this *DefaultChainAbleBccspManager) Verify(key Key, digest, signature []byte, failStrategy base.FailStrategy, opts VerifierOptsAdapter) error {
	if failStrategy() {
		handler := this.IVerifier.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		if e := handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).Verify(key, digest, signature, failStrategy, opts); nil != e {
			goto loopVerify
		}
		return nil
	}
	goto loopVerify

loopVerify:
	for tmp := this.IVerifier.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).Verify(key, digest, signature, failStrategy, opts); nil != e {
			log.Printf("[Verify]:alg=%s,校验失败,下一个继续,失败原因:[%s]", tmp.String(), e.Error())
			if tmp.GetNext() != nil {
				continue
			} else {
				return errors.New("遍历也无法校验成功:" + e.Error())
			}
		} else {
			return nil
		}
	}
	return errors.New("遍历也无法校验成功")

}

func (this *DefaultChainAbleBccspManager) InterfaceKeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts KeyImportOptsAdapter) (PrivateKeyWrapper, error) {
	if failStrategy() {
		handler := this.IKeyImporter.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).InterfaceKeyImport(raw, pwd, failStrategy, opts)
	}
	for tmp := this.IKeyImporter.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if k, e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).InterfaceKeyImport(raw, pwd, failStrategy, opts); nil != e {
			continue
		} else {
			return k, nil
		}
	}
	return PrivateKeyWrapper{}, errors.New("遍历也无法import key,失败")
}
func (this *DefaultChainAbleBccspManager) PublicKeyImport(raw interface{}, failStrategy base.FailStrategy, opts PublicKeyImportOptsAdapter) (KeyAdapter, error) {
	if failStrategy() {
		handler := this.IPublicKeyImporter.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).PublicKeyImport(raw, failStrategy, opts)
	}
	for tmp := this.IPublicKeyImporter.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if k, e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).PublicKeyImport(raw, failStrategy, opts); nil != e {
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("遍历也无法PublicKeyImport,失败")
}

func (this *DefaultChainAbleBccspManager) OriginPublicKeyImport(raw []byte, failStrategy base.FailStrategy, opts OriginPublicKeyImportOptsAdapter) (interface{}, error) {
	if failStrategy() {
		handler := this.IOriginPublicKeyImporter.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).OriginPublicKeyImport(raw, failStrategy, opts)
	}
	for tmp := this.IOriginPublicKeyImporter.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if k, e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).OriginPublicKeyImport(raw, failStrategy, opts); nil != e {
			continue
		} else {
			return k, nil
		}
	}
	return nil, errors.New("解析失败,都无法解析原生公钥")
}

func (this *DefaultChainAbleBccspManager) ParseCertificate(raw []byte, failStrategy base.FailStrategy, opts CertificateImportOptsAdapter) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if failStrategy() {
		handler := this.ICertificateImporter.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
		return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).ParseCertificate(raw, failStrategy, opts)
	}
	for tmp := this.ICertificateImporter.AlgorithmHieraTemplate; nil != tmp; tmp = tmp.GetNext().(*base.AlgorithmHieraTemplate) {
		if k, opts, e := tmp.GetDetailTemplate().(*DefaultDetailTemplateMediator).ParseCertificate(raw, failStrategy, opts); nil != e {
			if tmp.GetNext() == nil {
				return k, opts, e
			}
			continue
		} else {
			return k, opts, nil
		}
	}
	return nil, nil, errors.New("解析失败,都无法解析证书")
}

func (this *DefaultChainAbleBccspManager) Hash(msg []byte, opts HashOptsAdapter) ([]byte, error) {
	handler := this.IHash.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).Hash(msg, opts)
}

func (this *DefaultChainAbleBccspManager) GenPrivateKey(opts KeyGenOptsAdapter) (KeyAdapter, error) {
	handler := this.IPrivateKeyGenerator.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).KeyGen(opts)
}

func (this *DefaultChainAbleBccspManager) GenCertificate(opts CertificateGeneratorOpts, params base.ParamBuilder) (*x509.Certificate, []byte, error) {
	handler := this.ICertificateGenerator.GetHieraHandlerByBaseType(opts.GetBaseType()(), opts.GetMinExcepted()())
	return handler.GetDetailTemplate().(*DefaultDetailTemplateMediator).CertificateGen(opts, params)
}
