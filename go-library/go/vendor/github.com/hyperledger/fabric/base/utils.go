/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-02-10 07:52
# @File : utils.go
# @Description :
# @Attention :
*/
package base

import (
	"crypto/x509"
	"fmt"
)

func GetAlgorithmTypeFromCert(crt *x509.Certificate) BaseType {
	if crt.SignatureAlgorithm.String() == GM_SIGNATURE_STRING {
		return SM2
	} else if IsECDSASignedCert(crt) {
		return ECDSA
	} else if IsRSACertificate(crt) {
		return RSA
	} else {
		panic(fmt.Sprintf("未知的crt:%v", crt))
	}
}

func IsGmCertificate(cert *x509.Certificate) bool {
	return cert.SignatureAlgorithm.String() == GM_SIGNATURE_STRING
}

func IsECDSASignedCert(cert *x509.Certificate) bool {
	return cert.SignatureAlgorithm == x509.ECDSAWithSHA1 ||
		cert.SignatureAlgorithm == x509.ECDSAWithSHA256 ||
		cert.SignatureAlgorithm == x509.ECDSAWithSHA384 ||
		cert.SignatureAlgorithm == x509.ECDSAWithSHA512
}

func IsRSACertificate(cert *x509.Certificate) bool {
	return cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA256WithRSAPSS ||
		cert.SignatureAlgorithm == x509.SHA512WithRSAPSS || cert.SignatureAlgorithm == x509.SHA512WithRSA
}
