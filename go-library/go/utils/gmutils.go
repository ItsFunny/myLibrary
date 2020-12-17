/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-13 08:36 
# @File : gmutils.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
)

// 签名(兼容java客户端，世纪数码签名验签接口)
func SM2Sign(privateKey *sm2.PrivateKey, rawMessage ,uid[]byte) (string, error) {
	r, s, err := sm2.Sm2Sign(privateKey, rawMessage, uid)
	if err != nil {
		return "", err
	}
	return encodeSignature(r, s), nil
}

// 验签
func SM2Verify(publicKey *sm2.PublicKey, rawMessage ,uid[]byte, signature string) (bool, error) {
	// fmt.Println("xx signature len:", len(signature))
	// fmt.Println("xx signature:", signature)
	// if len(signature) < 128 {
	// 	return false, errors.New("signature's length less than 128")
	// }
	// 签名是由2个32个字节的大整数拼接而成
	r, s, err := decodeSignature(signature)
	if err != nil {
		return false, err
	}
	verify := sm2.Sm2Verify(publicKey, rawMessage, uid, r, s)
	// fmt.Println(r.Text(16), s.Text(16))
	return verify, nil
}

// 将两个大整数拼接成字符串
func encodeSignature(r, s *big.Int) string {
	fmt.Printf("siganture r=%v, s=%v\n", r.Text(16), s.Text(16))
	return fmt.Sprintf("%s%s", r.Text(16), s.Text(16))
}

// 将签名值转换成2个大整数
func decodeSignature(signature string) (*big.Int, *big.Int, error) {
	signatureR := signature[:64]
	signatureS := signature[64:]
	var r, s big.Int
	_, err := fmt.Sscanf(signatureR, "%x", &r)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse big.Integer R from signature")
	}
	_, err = fmt.Sscanf(signatureS, "%x", &s)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse big.Integer S from signature")
	}
	return &r, &s, nil
}

const (
	BitSize    = 256
	KeyBytes   = (BitSize + 7) / 8
	UnCompress = 0x04
)

// 公钥序列化
func MarshalPublicKey(pub *sm2.PublicKey) []byte {
	data := getUnCompressPublicKeyBytes(pub)
	return data[1:]
}

func getUnCompressPublicKeyBytes(pub *sm2.PublicKey) []byte {
	xBytes := pub.X.Bytes()
	yBytes := pub.Y.Bytes()
	xl := len(xBytes)
	yl := len(yBytes)

	raw := make([]byte, 1+KeyBytes*2)
	raw[0] = UnCompress
	if xl > KeyBytes {
		copy(raw[1:1+KeyBytes], xBytes[xl-KeyBytes:])
	} else if xl < KeyBytes {
		copy(raw[1+(KeyBytes-xl):1+KeyBytes], xBytes)
	} else {
		copy(raw[1:1+KeyBytes], xBytes)
	}

	if yl > KeyBytes {
		copy(raw[1+KeyBytes:], yBytes[yl-KeyBytes:])
	} else if yl < KeyBytes {
		copy(raw[1+KeyBytes+(KeyBytes-yl):], yBytes)
	} else {
		copy(raw[1+KeyBytes:], yBytes)
	}
	return raw
}