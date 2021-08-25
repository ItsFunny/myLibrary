/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 23:22
# @File : BaseType.go
# @Description :
# @Attention :
*/
package base

const GM_SIGNATURE_STRING = "Ed25519"
const GM_SIGNATURE_SM2WITHSM3 = "SM2-SM3"

var (
	algorithDescMap = map[BaseType]string{
		ECDSA: "ecdsa",
		SM2:   "sm2",
		RSA:   "rsa",
		SHA:   "sha",
	}
)

// base 算法
const (
	ECDSA BaseType = 1 << 0
	SM2   BaseType = 1 << 1
	RSA   BaseType = 1 << 2
	AES   BaseType = 1 << 3
	SHA   BaseType = 1 << 4
	NONE  BaseType = 1 << 7
)
const (
	ECDSA_256 DetailType = 1 << 0
	SM2_256   DetailType = 1 << 1
	RSA_2048  DetailType = 1 << 2
)

// hash算法
const (
	HASH_SHA256      BaseType = 100
	HASH_SM3_WITH_NO BaseType = 101
)

// 序列化

const (
	SERIALIZE_ECDSA BaseType = 30
	SERIALIZE_SM2   BaseType = 31
	SERIALIZE_ALL   BaseType = 32
)

// 前开后闭
const (
	// 15个字节为magicWord
	MAGIC_WORD = "magicWord"
	// magicWord起始下标
	EXTENSION_MAGICWORD_START_INDEX BaseType = 0
	// magicWord截止下标
	EXTENSION_MAGICWORD_END_INDEX BaseType = 10
	// 算法下标
	EXTENSION_BaseType_BASE_INDEX BaseType = 10

	EXTENSION_MIN_LENGTH BaseType = 12
)

func GetAlgorithmDescription(alg BaseType) string {
	return algorithDescMap[alg]
}
