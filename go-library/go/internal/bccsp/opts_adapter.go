package bccsp

import (
	"myLibrary/go-library/go/internal/bccsp/base"
)

type CertificateGeneratorOpts interface {
	base.CertificateBaseOpts
}

type KeyGenOptsAdapter interface {
	KeyGenOpts
	base.BaseOpts
}

type KeyDerivOptsAdapter interface {
	KeyDerivOpts
	base.BaseOpts
}

type KeyImportOptsAdapter interface {
	KeyImportOpts
	base.BaseOpts
}

type HashOptsAdapter interface {
	HashOpts
	base.BaseOpts
}

type SignerOptsAdapter interface {
	SignerOpts
	GetTypeFunc() base.SerializableFunc
	GetBaseType()base.SerializableFunc
	GetHashOpts() HashOptsAdapter
}

type VerifierOptsAdapter interface {
	GetTypeFunc() base.SerializableFunc
	GetBaseType()base.SerializableFunc
	GetHashOpts() HashOptsAdapter
}

type EncrypterOptsAdapter interface {
	base.BaseOpts
}

type DecrypterOptsAdapter interface {
	base.BaseOpts
}
