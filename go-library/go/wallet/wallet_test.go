/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 15:03 
# @File : wallet_test.go
# @Description : 
# @Attention : 
*/
package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/hbakhtiyor/schnorr"
	"github.com/stretchr/testify/assert"
	"myLibrary/go-library/go/crypt"
	"testing"
)

func TestHDWallet_GetAddress(t *testing.T) {
	prevPwd := "123"
	cryptPrevpwd := encrypt.MD5EncryptByBytes(prevPwd)
	wallet := NewHDWallet(1, 1, cryptPrevpwd)
	address := wallet.GetAddress()
	fmt.Println(wallet.Position)
	fmt.Println(address)
	fmt.Println(wallet.Path)

	childWallet, e := wallet.NewChildWallet(2)
	if nil != e {
		panic(e)
	}
	fmt.Println("======")
	fmt.Println(childWallet.Position)
	fmt.Println(childWallet.GetAddress())
	fmt.Println(childWallet.Path)
}
func TestHDWallet_ECCSign(t *testing.T) {
	str := "joker"
	wallet := NewHDWallet(1, 1, "qwe")
	bytes, e := wallet.SchnorrSign(str)
	assert.NoError(t, e)
	fmt.Println(hex.EncodeToString(wallet.PrvKey.Key))
	fmt.Println(hex.EncodeToString(wallet.PubKey.Key))
	fmt.Println(hex.EncodeToString(bytes))
	// hexPrvKey=44bb11264f10e8489418c8753f6c037df1839d5cc59bd330c402c575a4d86964
	// hexPubKey=027faef15cd558673516f03c9391d6d2da0e9cc65773bf8304888ebaafa7cbf09e
	// signature=7ae511c71a735b6063384c2005e6deaa47ce7b9875736996f998b7c2f5498ffdfc698f12ded7dcdea7e6d47da71d6a5362bc924132e76d9a3f7d3e05a0229cd6
}
func TestHDWallet_ECCVerisign(t *testing.T) {
	hexPubKey := "027faef15cd558673516f03c9391d6d2da0e9cc65773bf8304888ebaafa7cbf09e"
	str := "7ae511c71a735b6063384c2005e6deaa47ce7b9875736996f998b7c2f5498ffdfc698f12ded7dcdea7e6d47da71d6a5362bc924132e76d9a3f7d3e05a0229cd6"
	k, _ := hex.DecodeString(hexPubKey)
	verysign := SchnorrVerysign("joker", str, k)

	fmt.Println(verysign)
}

func SchnorrVerysign(exceptStr string, signatureStr string, key []byte) bool {
	var (
		publicKey [33]byte
		message   [32]byte
		signature [64]byte
	)
	copy(publicKey[:], key)
	copy(message[:], []byte(exceptStr))
	bytes, e := hex.DecodeString(signatureStr)
	if nil != e {
		return false
	}
	copy(signature[:], bytes)

	b, e := schnorr.Verify(publicKey, message, signature)
	if nil != e {
		return false
	}
	return b
}

func TestNewHDWallet(t *testing.T) {
	prevPwd := "123"
	cryptPrevpwd := encrypt.MD5EncryptByBytes(prevPwd)
	w1 := NewHDWallet(1, 1, cryptPrevpwd)
	nowpwd := "1234"
	nowcryptPrevpwd := encrypt.MD5EncryptByBytes(nowpwd)
	w2 := NewHDWallet(1, 1, nowcryptPrevpwd)
	fmt.Println(w1.GetAddress())
	fmt.Println(w2.GetAddress())

}
