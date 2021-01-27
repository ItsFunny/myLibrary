/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-13 21:43 
# @File : rsa_test.go
# @Description : 
# @Attention : 
*/
package algorithm

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/helper"
	"math/big"
	"time"

	"crypto/rsa"
	"testing"
)

// 生成rsa的私钥
func Test_GenRsa(t *testing.T) {
	privateKey, e := rsa.GenerateKey(rand.Reader, 2048)
	helper.CheckError(e)
	fmt.Println(privateKey)
}

// rsa 签名 验签
func Test_RsaSignVerify(t *testing.T) {
	privateKey, e := rsa.GenerateKey(rand.Reader, 2048)
	msg := []byte("123")
	helper.CheckError(e)
	myhash := crypto.SHA256
	hashInstance := myhash.New()
	hashInstance.Write(msg)
	hashed := hashInstance.Sum(nil)
	signature, e := rsa.SignPKCS1v15(rand.Reader, privateKey, myhash, hashed)
	helper.CheckError(e)
	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	// 校验
	verifyHash := crypto.SHA256
	verifyHashInstance := myhash.New()
	verifyHashInstance.Write(msg)
	verifyHashed := verifyHashInstance.Sum(nil)
	e = rsa.VerifyPKCS1v15(&privateKey.PublicKey, verifyHash, verifyHashed, signature)
	if nil == e {
		fmt.Println("签名验签通过")
	} else {
		fmt.Println("失败")
	}
}

// rsa 生成证书
func Test_Rsa_GenCertificate(t *testing.T) {
	key, e := rsa.GenerateKey(rand.Reader, 2048)
	helper.CheckError(e)

	// Generate a pem block with the private key
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	fmt.Println(string(keyPem))

	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		// you have to generate a different serial number each execution
		SerialNumber: big.NewInt(123123),
		Subject: pkix.Name{
			CommonName:   "New Name",
			Organization: []string{"New Org."},
		},
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	helper.CheckError(err)

	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	fmt.Println(string(certPem))
}
