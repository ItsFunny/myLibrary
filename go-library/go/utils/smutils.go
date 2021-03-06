// /*
// # -*- coding: utf-8 -*-
// # @Author : joker
// # @Time : 2020-10-05 12:07
// # @File : smutils.go
// # @Description :
// # @Attention :
// */
package utils
//
// import (
// 	"crypto/ecdsa"
// 	"crypto/rsa"
// 	"crypto/x509"
// 	"encoding/asn1"
// 	"encoding/base64"
// 	"encoding/pem"
// 	"fmt"
// 	"github.com/pkg/errors"
// 	"io/ioutil"
// 	"math/big"
// 	"myLibrary/go-library/go/tjfoc/gmsm/sm2"
// )
//
// func GM() {
// 	prvStr := `o3Vl18oUETpni/2D/t634jzvyIn08zEka9ZpAMRjeZUEdAWYLO32qLsVDudvaT6Y
// tqw57FmEFs5eHqpfqrEjneR4mT44M+RGD1ehDqqP7FltZlqENDkgrtYp56h/Qa8L
// LA==`
// 	bytes, _ := Base64Decode(prvStr)
//
// 	prv := bytes[:32]
// 	pub := bytes[32:]
// 	xBs := pub[0:32]
// 	yBs := pub[32:]
// 	X := &big.Int{}
// 	X = X.SetBytes(xBs)
// 	Y := &big.Int{}
// 	Y = Y.SetBytes(yBs)
// 	pubK := sm2.PublicKey{
// 		Curve: sm2.P256Sm2(),
// 		X:     X,
// 		Y:     Y,
// 	}
// 	privateKey, e := ParseSM2PrvKBytesWithPubKey(prv, pubK)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(privateKey)
// 	bs := []byte("123")
// 	r, s, e := sm2.Sign(privateKey, bs)
// 	if nil != e {
// 		panic(e)
// 	}
// 	verify := sm2.Verify(&pubK, bs, r, s)
// 	fmt.Println(verify)
// }
//
// func buildPrv(pub sm2.PublicKey) *sm2.PrivateKey {
// 	encPrv := `O51oWDipjyoG08MASHraVDTjDNWophmUFgkVbH/CbYY=`
// 	// 	signPrv:=`MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgWt9z9iL/aeASq1na
// 	// 3DZKM+t1O758F6YOxBKWEN0+pVegCgYIKoEcz1UBgi2hRANCAAT3spXBAjGcPp0y
// 	// Rwa2OeFQEQvxZ6c4i7VKBXJjTDOU3YmUFOhIGziw9jrLINmoH+M8QXF+S2s715PF
// 	// bPA0ZtM3`
//
// 	prvBytes, e := base64.StdEncoding.DecodeString(encPrv)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(len(prvBytes))
//
// 	d := &big.Int{}
// 	d = d.SetBytes(prvBytes)
//
// 	pk := &sm2.PrivateKey{
// 		PublicKey: pub,
// 		D:         d,
// 	}
// 	return pk
// }
//
// // prvKey, e := sm2.ParsePKCS8UnecryptedPrivateKey(prvBytes)
// // if nil != e {
// // 	panic(e)
// // }
//
// // pkcs8 的形式读取标准的 sm2 国密的公私钥
// // 私钥:32位 公钥: 64位
// func ParseSMKeyPairFromFile(prvPath string, crtPath string) (*sm2.PrivateKey, *sm2.PublicKey, error) {
// 	if len(prvPath) == 0 || len(crtPath) == 0 {
// 		return nil, nil, errors.New("私钥或者证书地址不可为空")
// 	}
// 	certificate, e := ReadSM2CertFromFile(crtPath)
// 	if nil != e {
// 		return nil, nil, e
// 	}
// 	ecdsaPubK := certificate.PublicKey.(*ecdsa.PublicKey)
// 	sm2PubK := &sm2.PublicKey{
// 		Curve: ecdsaPubK.Curve,
// 		X:     ecdsaPubK.X,
// 		Y:     ecdsaPubK.Y,
// 	}
// 	prvK, e := ParseSM2PrvKFileWithPubKey(prvPath, *sm2PubK)
// 	if nil != e {
// 		return nil, nil, errors.New("解密私钥失败:" + e.Error())
// 	}
//
// 	return prvK, sm2PubK, nil
// }
//
// func ParseSM2PrvKFileWithPubKey(prvPath string, pubK sm2.PublicKey) (*sm2.PrivateKey, error) {
// 	if len(prvPath) == 0 {
// 		return nil, errors.New("私钥不可为空")
// 	}
// 	prvBytes, e := ioutil.ReadFile(prvPath)
// 	if nil != e {
// 		return nil, errors.New("读取私钥文件失败:" + e.Error())
// 	}
//
// 	str := string(prvBytes)
// 	str = replace(str, "PRIVATE KEY")
// 	prvBytes, e = Base64Decode(str)
// 	// if nil != e {
// 	// 	return nil, errors.New("base64解码失败:" + e.Error())
// 	// }
// 	// d := &big.Int{}
// 	// d = d.SetBytes(prvBytes)
// 	//
// 	// prvK := &sm2.PrivateKey{
// 	// 	PublicKey: pubK,
// 	// 	D:         d,
// 	// }
// 	return ParseSM2PrvKBytesWithPubKey(prvBytes, pubK)
// }
//
// func ParseSM2PrvKBytesWithPubKey(origin []byte, pub sm2.PublicKey) (*sm2.PrivateKey, error) {
// 	d := &big.Int{}
// 	d = d.SetBytes(origin)
//
// 	prvK := &sm2.PrivateKey{
// 		PublicKey: pub,
// 		D:         d,
// 	}
// 	return prvK, nil
// }
//
// func ReadSM2CertFromFile(path string) (*sm2.Certificate, error) {
// 	bytes, e := ioutil.ReadFile(path)
// 	if nil != e {
// 		return nil, e
// 	}
// 	str := string(bytes)
// 	str = replace(str, "CERTIFICATE")
// 	cerBytes, e := Base64Decode(str)
// 	if nil != e {
// 		return nil, e
// 	}
// 	return sm2.ParseCertificate(cerBytes)
// }
//
// type CertificateReader func(encoded []byte) (interface{}, error)
//
// // 2种方式读取证书
// func ReadCertificateTwoWays(encoded []byte) (interface{}, error) {
// 	pipeline := []CertificateReader{SM2ReadCertificate, X509ReadCertificate}
// 	for _, reader := range pipeline {
// 		result, e := reader(encoded)
// 		if nil != e {
// 			return result, nil
// 		}
// 	}
// 	return nil, errors.New("无法解析")
// }
//
// // 国密读取
// func SM2ReadCertificate(encoded []byte) (interface{}, error) {
// 	str := string(encoded)
// 	str = replace(str, "CERTIFICATE")
// 	cerBytes, e := Base64Decode(str)
// 	if nil != e {
// 		return nil, e
// 	}
// 	return sm2.ParseCertificate(cerBytes)
// }
//
// // x509 原生读取
// func X509ReadCertificate(encoded []byte) (interface{}, error) {
// 	block, _ := pem.Decode(encoded)
// 	if block == nil {
// 		return nil, errors.New("Failed to PEM decode certificate")
// 	}
// 	result, err := x509.ParseCertificate(block.Bytes)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Error parsing certificate")
// 	}
// 	return result, nil
// }
//
// type PrvKReader func(raw, pwd []byte) (interface{}, error)
//
// // 三种方式读取私钥
// func ReadPrvKThreeWays(raw []byte) (interface{}, error) {
// 	pipeLine := []PrvKReader{ECDSA_READER, SMPRV_READER}
// 	//
// 	var result interface{}
// 	var e error
// 	for _, reader := range pipeLine {
// 		result, e = reader(raw, nil)
// 		if nil != e {
// 			return result, e
// 		}
// 	}
// 	return nil, errors.New("无法解析")
// }
//
// // 标准公私钥读取
//
// // 国密读取
// var SMPRV_READER = func(raw, pwd []byte) (interface{}, error) {
// 	return sm2.ParseSm2PrivateKey(raw)
// }
//
// // 读取ecdsa的私钥
// var ECDSA_READER = func(raw, pwd []byte) (interface{}, error) {
// 	if len(raw) == 0 {
// 		return nil, errors.New("Invalid PEM. It must be different from nil.")
// 	}
// 	block, _ := pem.Decode(raw)
// 	if block == nil {
// 		return nil, fmt.Errorf("Failed decoding PEM. Block must be different from nil. [% x]", raw)
// 	}
//
// 	// TODO: derive from header the type of the key
//
// 	if x509.IsEncryptedPEMBlock(block) {
// 		if len(pwd) == 0 {
// 			return nil, errors.New("Encrypted Key. Need a password")
// 		}
//
// 		decrypted, err := x509.DecryptPEMBlock(block, pwd)
// 		if err != nil {
// 			return nil, fmt.Errorf("Failed PEM decryption [%s]", err)
// 		}
//
// 		key, err := DERToPrivateKey(decrypted)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return key, err
// 	}
//
// 	cert, err := DERToPrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return cert, err
// }
//
// // DERToPrivateKey unmarshals a der to private key
// func DERToPrivateKey(der []byte) (key interface{}, err error) {
//
// 	if key, err = x509.ParsePKCS1PrivateKey(der); err == nil {
// 		return key, nil
// 	}
//
// 	if key, err = x509.ParsePKCS8PrivateKey(der); err == nil {
// 		switch key.(type) {
// 		case *rsa.PrivateKey, *ecdsa.PrivateKey:
// 			return
// 		default:
// 			return nil, errors.New("Found unknown private key type in PKCS#8 wrapping")
// 		}
// 	}
//
// 	if key, err = x509.ParseECPrivateKey(der); err == nil {
// 		return
// 	}
//
// 	return nil, errors.New("Invalid key type. The DER must contain an rsa.PrivateKey or ecdsa.PrivateKey")
// }
//
//
// type SM2Signature struct {
// 	R, S *big.Int
// }
// func UnmarshalSM2Signature(raw []byte) (*big.Int, *big.Int, error) {
// 	// Unmarshal
// 	sig := new(SM2Signature)
// 	_, err := asn1.Unmarshal(raw, sig)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
// 	}
//
// 	// Validate sig
// 	if sig.R == nil {
// 		return nil, nil, errors.New("invalid signature, R must be different from nil")
// 	}
// 	if sig.S == nil {
// 		return nil, nil, errors.New("invalid signature, S must be different from nil")
// 	}
//
// 	if sig.R.Sign() != 1 {
// 		return nil, nil, errors.New("invalid signature, R must be larger than zero")
// 	}
// 	if sig.S.Sign() != 1 {
// 		return nil, nil, errors.New("invalid signature, S must be larger than zero")
// 	}
//
// 	return sig.R, sig.S, nil
// }
//
//
// ///////////////////////////////////////////////
// //////////////////////////
// //////////////////////////
// ///////////////////////////////////////////////
//
// // 从文件中加载私钥
// func ReadPrivateKeyFromPem(fileName string) (*sm2.PrivateKey, error) {
// 	data, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	block, _ := pem.Decode(data)
// 	if block == nil {
// 		return nil, errors.New("failed to decode private key")
// 	}
// 	var privateKey sm2.PrivateKey
// 	d := block.Bytes
// 	privateKey.D = new(big.Int).SetBytes(d)
// 	curve := sm2.P256Sm2()
// 	privateKey.Curve = curve
// 	x, y := curve.ScalarBaseMult(d)
// 	privateKey.X, privateKey.Y = x, y
// 	return &privateKey, nil
// }
//
// // 从私钥文件中获取公钥
// func ReadPublicKeyFromPrivateKeyPem(privateKeyFileName string) (*sm2.PublicKey, error) {
// 	privateKey, err := ReadPrivateKeyFromPem(privateKeyFileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &privateKey.PublicKey, nil
//
// }
//
// // 从证书中加载公钥
// func ReadPublicKeyFromCertPem(certFileName string) (*sm2.PublicKey, error) {
// 	cert, err := sm2.ReadCertificateFromPem(certFileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	publicKey, ok := cert.PublicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("failed to parse ecdsaPublicKey from sm2 cert")
// 	}
// 	curve := sm2.P256Sm2()
// 	sm2PublicKey := &sm2.PublicKey{
// 		Curve: curve,
// 		X:     publicKey.X,
// 		Y:     publicKey.Y,
// 	}
//
// 	return sm2PublicKey, nil
// }
//
// // 签名(兼容java客户端，世纪数码签名验签接口)
// func SM2Sign(privateKey *sm2.PrivateKey, rawMessage []byte) (string, error) {
// 	r, s, err := sm2.Sm2Sign(privateKey, rawMessage, nil)
// 	if err != nil {
// 		return "", err
// 	}
// 	return encodeSignature(r, s), nil
// }
//
// // 验签
// func SM2Verify(publicKey *sm2.PublicKey, rawMessage []byte, signature string) (bool, error) {
// 	// 签名是由2个32个字节的大整数拼接而成
// 	r, s, err := decodeSignature(signature)
// 	if err != nil {
// 		return false, err
// 	}
// 	verify := sm2.Sm2Verify(publicKey, rawMessage, nil, r, s)
// 	fmt.Println(r.Text(16), s.Text(16))
// 	return verify, nil
// }
//
// // 将两个大整数拼接成字符串
// func encodeSignature(r, s *big.Int) string {
// 	fmt.Printf("siganture r=%v, s=%v\n", r.Text(16), s.Text(16))
// 	return fmt.Sprintf("%s%s", r.Text(16), s.Text(16))
// }
//
// // 将签名值转换成2个大整数
// func decodeSignature(signature string) (*big.Int, *big.Int, error) {
// 	signatureR := signature[:64]
// 	signatureS := signature[64:]
// 	var r, s big.Int
// 	_, err := fmt.Sscanf(signatureR, "%x", &r)
// 	if err != nil {
// 		return nil, nil, errors.Wrap(err, "failed to parse big.Integer R from signature")
// 	}
// 	_, err = fmt.Sscanf(signatureS, "%x", &s)
// 	if err != nil {
// 		return nil, nil, errors.Wrap(err, "failed to parse big.Integer S from signature")
// 	}
// 	return &r, &s, nil
// }
//
// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
//
// func main() {
// 	prvStr := "rWlQn95FhPLuVabXhCbYHa8HiGDV32t2C4QBwIute74="
// 	key, err := readPrvFromStr2(prvStr)
// 	checkError(err)
//
// 	// 	pubKey, err := ReadPublicKeyFromCertPem("cert.pem")
// 	// 	checkError(err)
// 	pubKey := &key.PublicKey
//
// 	msg := []byte("hello world")
// 	signature, err := SM2Sign(key, msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	ok, err := SM2Verify(pubKey, msg, signature)
// 	checkError(err)
// 	fmt.Println("ok:", ok)
//
// }
//
// // 字节码转换为私钥
// func readPrvFromStr2(str string) (*sm2.PrivateKey, error) {
// 	bytes, e := base64.StdEncoding.DecodeString(str)
// 	if nil != e {
// 		return nil, e
// 	}
// 	// block, _ := pem.Decode(bytes)
// 	// if block == nil {
// 	// 	return nil, errors.New("failed to decode private key")
// 	// }
// 	var privateKey sm2.PrivateKey
// 	d := bytes
// 	privateKey.D = new(big.Int).SetBytes(d)
// 	curve := sm2.P256Sm2()
// 	privateKey.Curve = curve
// 	x, y := curve.ScalarBaseMult(d)
// 	privateKey.X, privateKey.Y = x, y
// 	return &privateKey, nil
// }
