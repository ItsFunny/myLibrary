/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 15:56
# @File : template.go
# @Description :
# @Attention :
*/
package bccsp

import (
	"crypto/x509"
	"github.com/hyperledger/fabric/base"
	"math/big"
)

type ICertificateGeneratorTemplate interface {
	base.Template
	CertificateGen(opts CertificateGeneratorOpts, params base.ParamBuilder) (*x509.Certificate, []byte, error)
}

type ISignatureMarshalTemplate interface {
	base.Template
	SignatureMarshal(r *big.Int, s *big.Int, opts SignatureMarshalOptsAdapter) ([]byte, error)
}

type ISignatureUnmarshalTemplate interface {
	base.Template
	SignatureUnmarshal(sig []byte, opts SignatureUnMarshalOptsAdapter) (*big.Int, *big.Int, error)
}

type IKeyImporterTemplate interface {
	base.Template
	KeyImport(raw, pwd []byte, failStrategy base.FailStrategy, opts KeyImportOptsAdapter) (interface{}, error)
}

type IKeyGeneratorTemplate interface {
	base.Template
	KeyGen(opts KeyGenOptsAdapter) (KeyAdapter, error)
}

type ISignerTemplate interface {
	base.Template
	Sign(key Key, digest []byte, opts SignerOptsAdapter) (signature []byte, err error)
}

type IVerifierTemplate interface {
	base.Template
	Verify(key Key, digest, signature []byte, opts VerifierOptsAdapter) error
}

type ICertificateImporterTemplate interface {
	base.Template
	ParseCertificate(bytes []byte, opts CertificateImportOptsAdapter) (*x509.Certificate, base.CertificateBaseOpts, error)
}

type IPublicKeyImportTemplate interface {
	base.Template
	PublicKeyImport(raw interface{}, opts PublicKeyImportOptsAdapter) (Key, error)
}

type IOriginPublicKeyImporter interface {
	base.Template
	OriginPublicKeyImport(raw []byte, opts OriginPublicKeyImportOptsAdapter) (interface{}, error)
}

type IHashTemplate interface {
	base.Template
	Hash(msg []byte, opts HashOptsAdapter) ([]byte, error)
}

type IPrivateKeyGeneratorTemplate interface {
	base.Template
}
