/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-02-10 07:50
# @File : certificate.go
# @Description :
# @Attention :
*/
package base

import (
	"crypto/x509"
	"encoding/base64"
)

type ICertificateHelper interface {
	GetX509Certificate() *x509.Certificate
	GetBaseOptsWrapper() BaseOptsWrapper
}

type CertificateHelperWrapper struct {
	Crt        *x509.Certificate
	DetailType DetailType
}

func (this CertificateHelperWrapper) GetX509Certificate() *x509.Certificate {
	return this.Crt
}

func (this CertificateHelperWrapper) GetBaseOptsWrapper() BaseOptsWrapper {
	res := BaseOptsWrapper{
		CommonBaseOpts: CommonBaseOpts{},
		FailStrategy:   FailOver,
	}
	if IsECDSASignedCert(this.Crt) {
		res.CommonBaseOpts.AlgorithmType = ECDSA
		res.CommonBaseOpts.DetailType, res.CommonBaseOpts.MinExcepted = DetailType(ECDSA), func() BaseType {
			return ECDSA
		}
	} else if IsRSACertificate(this.Crt) {
		res.CommonBaseOpts.AlgorithmType = RSA
		res.CommonBaseOpts.DetailType, res.CommonBaseOpts.MinExcepted = DetailType(RSA), func() BaseType {
			return RSA
		}
	} else if IsGmCertificate(this.Crt) {
		res.CommonBaseOpts.AlgorithmType = SM2
		res.CommonBaseOpts.DetailType, res.CommonBaseOpts.MinExcepted = DetailType(SM2), func() BaseType {
			return SM2
		}
	} else {
		panic("未知的证书:" + base64.StdEncoding.EncodeToString(this.Crt.Raw))
	}
	if this.DetailType > 0 {
		res.CommonBaseOpts.DetailType = this.DetailType
	}
	return res
}
