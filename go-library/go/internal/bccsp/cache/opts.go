/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-12 18:24 
# @File : opts.go
# @Description : 
# @Attention : 
*/
package cache

import (
	"myLibrary/go-library/go/internal/bccsp"
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/constants"
)

var (
	hashOptsMap = map[base.SerializableInterface]bccsp.HashOptsAdapter{
	}
	signerOptsMap = map[base.SerializableInterface]bccsp.SignerOptsAdapter{

	}
	verifyOptsMap = map[base.SerializableInterface]bccsp.VerifierOptsAdapter{
	}
	keyGenOptsMap = map[base.SerializableInterface]bccsp.KeyGenOptsAdapter{

	}
	certificateOpts=map[base.SerializableInterface]base.CertificateBaseOpts{

	}
)
func RegisterCertificateOpts(serializableInterface base.SerializableInterface,adapter base.CertificateBaseOpts){
	certificateOpts[serializableInterface]=adapter
}
func RegisterKeyGenOpts(serializableInterface base.SerializableInterface, adapter bccsp.KeyGenOptsAdapter) {
	keyGenOptsMap[serializableInterface] = adapter
}

func RegisterHashOpts(serializableInterface base.SerializableInterface, adapter bccsp.HashOptsAdapter) {
	hashOptsMap[serializableInterface] = adapter
}

func RegisterSignerOpts(serializableInterface base.SerializableInterface, adapter bccsp.SignerOptsAdapter) {
	signerOptsMap[serializableInterface] = adapter
}

func RegisterVerifierOpts(serializableInterface base.SerializableInterface, adapter bccsp.VerifierOptsAdapter) {
	verifyOptsMap[serializableInterface] = adapter
}

func GetHashOptsByAlgorithm(alg constants.Algorithm) bccsp.HashOptsAdapter {
	return hashOptsMap[alg]
}
func GetSignerOptsByAlgorithm(alg constants.Algorithm) bccsp.SignerOptsAdapter {
	return signerOptsMap[alg]
}

func GetVerifierOptsByAlgorithm(alg constants.Algorithm) bccsp.VerifierOptsAdapter {
	return verifyOptsMap[alg]
}

func GetPrvKeyGenOptsByAlgorithm(alg constants.Algorithm)bccsp.KeyGenOptsAdapter{
	return keyGenOptsMap[alg]
}

func GetCertificateGenOptsByAlgorithm(alg constants.Algorithm)base.CertificateBaseOpts{
	return certificateOpts[alg]
}
