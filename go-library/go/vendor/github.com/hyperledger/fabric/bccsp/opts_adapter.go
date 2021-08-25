package bccsp

import (
	"github.com/hyperledger/fabric/base"
)

type CertificateImportOptsAdapter interface {
	base.CertificateBaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type CertificateGeneratorOpts interface {
	base.CertificateBaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type SignatureMarshalOptsAdapter interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type SignatureUnMarshalOptsAdapter interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type KeyGenOptsAdapter interface {
	KeyGenOpts
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type KeyDerivOptsAdapter interface {
	KeyDerivOpts
	base.BaseOpts
}

type KeyImportOptsAdapter interface {
	KeyImportOpts
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type OriginPublicKeyImportOptsAdapter interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type PublicKeyImportOptsAdapter interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type HashOptsAdapter interface {
	HashOpts
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type SignerOptsAdapter interface {
	SignerOpts
	GetDetailType() base.DetailTypeFunc
	GetBaseType() base.BaseTypeFunc

	GetHashOpts() HashOptsAdapter
	GetMinExcepted() base.BaseTypeFunc
}

type VerifierOptsAdapter interface {
	GetDetailType() base.DetailTypeFunc
	GetBaseType() base.BaseTypeFunc
	GetHashOpts() HashOptsAdapter
	// 用于与运算,期望的type 值
	GetMinExcepted() base.BaseTypeFunc
}

type EncrypterOptsAdapter interface {
	base.BaseOpts
}

type DecrypterOptsAdapter interface {
	base.BaseOpts
}

type CSRGenerateOpts interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}

type DefaultCSRGenerateOptsImpl struct {
	base.CommonBaseOpts
}
type CRTWithCsrGeneratorOpts interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type DefaultCRTWithCSRGenerateOptsImpl struct {
	base.CommonBaseOpts
}

type CSRParserOpts interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type DefaultCsrParserOptsImpl struct {
	base.CommonBaseOpts
}

type SignCertGenOpts interface {
	base.BaseOpts
	GetMinExcepted() base.BaseTypeFunc
}
type DefaultSignCertGenOptsImpl struct {
	base.CommonBaseOpts
}
