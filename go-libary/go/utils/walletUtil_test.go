/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-17 16:30 
# @File : walletUtilTest.go
# @Description : 
*/
package utils

import (
	"encoding/binary"
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"math"
	"myLibrary/go-libary/go/crypt"
	"testing"
)

func TestNewWallet(t *testing.T) {
	wallet, e := NewWallet()
	if nil != e {
		panic(e)
	}
	fmt.Println(len(wallet.PrivateKeyBytes))
	fmt.Println(len(wallet.PublicKeyBytes))
	str := "joker"
	strs, e := encrypt.ECCSignWithHex(str, wallet.PrivateKeyBytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(strs)
	hex := encrypt.ECCVerifySignWithHex(str, strs, wallet.PublicKeyBytes)
	fmt.Println(hex)
}

func TestWalletEncrypt(t *testing.T) {
	wallet, e := NewWallet()
	if nil != e {
		panic(e)
	}
	str := "joker"
	s, e := encrypt.ECCEncryptWithHex(str, wallet.PublicKeyBytes)
	if nil != e {
		panic(e)
	}
	hex, e := encrypt.ECCDecryptWithHex(s, wallet.PrivateKeyBytes)
	if nil != e {
		panic(e)
	}
	fmt.Println(hex)

}

func TestNewWallet2(t *testing.T) {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()
	key, e := publicKey.NewChildKey(2)
	if nil != e {
		panic(e)
	}
	newChildKey, _ := key.NewChildKey(3)
	bytes, _ := newChildKey.Serialize()

	fmt.Println(len(bytes))
	fmt.Println(key)
	childKey, e := publicKey.NewChildKey(2)
	if nil != e {
		panic(e)
	}
	fmt.Println(childKey)

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master private keyLength: ", len(masterKey.Key))
	fmt.Println("Master public key: ", publicKey)
	fmt.Println("Master public keyLength: ", len(publicKey.Key))

	fmt.Println(masterKey.Depth)
}
func TestCombineOrConditionSQL2(t *testing.T) {

	f := 12.666512
	result := float64ToByte(f)
	fmt.Printf("result:%x", result)

}
func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func TestCompareAndSwap(t *testing.T) {

}
