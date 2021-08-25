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
package utils

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperledger/fabric/base"
)

func GetAlgorithmTypeFromCert(crt *x509.Certificate) base.BaseType {
	if crt.SignatureAlgorithm.String() == base.GM_SIGNATURE_STRING {
		return base.SM2
	} else if IsECDSASignedCert(crt) {
		return base.ECDSA
	} else if IsRSACertificate(crt) {
		return base.RSA
	} else {
		panic(fmt.Sprintf("未知的crt:%v", crt))
	}
}

func IsGmCertificate(cert *x509.Certificate) bool {
	// var cert x509.Certificate
	// rest, err := asn1.Unmarshal(ans1Bytes, &cert)
	// if err != nil {
	// 	return false
	// }
	// if len(rest) > 0 {
	// 	return false
	// }
	return cert.SignatureAlgorithm.String() == base.GM_SIGNATURE_STRING
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
