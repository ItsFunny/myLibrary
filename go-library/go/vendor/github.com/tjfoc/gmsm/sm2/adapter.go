/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-11-17 13:33 
# @File : adapter.go
# @Description : 
# @Attention : 
*/
package sm2

import (
	"encoding/asn1"
	"encoding/pem"
	"strings"
)

var (
	GM_SIGNATURE_STRING = "1.2.156.10197.1.501"
)

func IsGmCertificate(ans1Bytes []byte) bool {
	if strings.Contains(string(ans1Bytes), "BEGIN") {
		var block *pem.Block
		block, _ = pem.Decode(ans1Bytes)
		if block == nil {
			return false
		}
		ans1Bytes = block.Bytes
	}
	var cert certificate
	rest, err := asn1.Unmarshal(ans1Bytes, &cert)
	if err != nil {
		return false
	}
	if len(rest) > 0 {
		return false
	}
	return cert.SignatureAlgorithm.Algorithm.String() == GM_SIGNATURE_STRING
}
