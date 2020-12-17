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
package sw

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	utils2 "github.com/hyperledger/fabric/utils"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"myLibrary/go-library/go/internal/bccsp"
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/internal/bccsp/cache"
	"myLibrary/go-library/go/constants"
	"myLibrary/go-library/go/utils"
	"reflect"
	"strings"
	"unsafe"
)

// 注册到bccspContainer中
func init() {
	// 默认256 位
	// config2.RegisterOpts(constants.ECDSA,base.DefaultECDSAIntegrationOpts{})
	// config2.RegisterOpts(constants.SM2,base.DefaultSM2IntegrationOpts{})
	// config2.RegisterOpts(constants.RSA,base.DefaultRSAIntergrationOpts{})

	// 私钥导入
	bccsp.RegisterKeyImporter(constants.ECDSA, ECDSA_KEY_IMPORTER)
	bccsp.RegisterKeyImporter(constants.SM2, SM2_KEY_IMPORTER)
	bccsp.RegisterKeyImporter(constants.RSA, nil)
	// 公私钥生成
	bccsp.RegisterKeyGenerator(constants.ECDSA, ECDSA_KEYGENERATOR)
	bccsp.RegisterKeyGenerator(constants.SM2, SM2_KEYGENERATOR)
	bccsp.RegisterKeyGenerator(constants.RSA, nil)

	// 签名

	bccsp.RegisterKeySigner(constants.ECDSA, ECDSA_SIGNER)
	bccsp.RegisterKeySigner(constants.SM2, SM2_SIGNER)
	bccsp.RegisterKeySigner(constants.RSA, nil)
	// 验签
	bccsp.RegisterKeyVerifier(constants.ECDSA, ECDSA_VERIFIER)
	bccsp.RegisterKeyVerifier(constants.SM2, SM2_VERIFIER)
	bccsp.RegisterKeyVerifier(constants.RSA, nil)
	// 导入
	bccsp.RegisterCertificateImporter(constants.ECDSA, ECDSA_CERTIFICATE_IMPORTER)
	bccsp.RegisterCertificateImporter(constants.SM2, SM2_CERTIFICATE_IMPORTER)
	bccsp.RegisterCertificateImporter(constants.RSA, nil)

	// 公钥导入
	bccsp.RegisterPublicKeyImporter(constants.ECDSA, ECDSA_PUBLICKEY_IMPORTER)
	bccsp.RegisterPublicKeyImporter(constants.SM2, SM2_PUBLICKEY_IMPORTER)
	bccsp.RegisterPublicKeyImporter(constants.RSA, RSA_PUBLICKEY_IMPORTER)

	// 原生公钥导入
	bccsp.RegisterOriginPublicKeyImporter(constants.ECDSA, ECDSA_ORIGIN_PUBLICKEY_IMPORTER)
	bccsp.RegisterOriginPublicKeyImporter(constants.SM2, SM2_ORIGIN_PUBLICKEY_IMPORTER)
	bccsp.RegisterOriginPublicKeyImporter(constants.RSA, nil)

	// 序列化
	bccsp.RegisterMarshal(constants.ECDSA, ECDSA_MARSHAL)
	bccsp.RegisterMarshal(constants.SM2, SM2_MARSHAL)

	// 反序列化
	bccsp.RegisterUnmarshal(constants.ECDSA, ECDSA_UNMARSHAL)
	bccsp.RegisterUnmarshal(constants.SM2, SM2_UNMARSHAL)

	// hash
	bccsp.RegisterHash(constants.HASH_SHA256, SHA256_HASH)
	bccsp.RegisterHash(constants.HASH_SM3_WITH_NO, SM3_HASH_WITH_NO)

	// privateKey generator
	bccsp.RegisterPrivateKeyGenerator(constants.ECDSA, ECDSA_PRIVATEKEY_GENERATOR)
	bccsp.RegisterPrivateKeyGenerator(constants.SM2, SM2_PRIVATEKEY_GENERATOR)

	// certificate generator
	bccsp.RegisterCertificateGenerator(constants.ECDSA,ECDSA_CERTIFICATE_GENERATOR)
	bccsp.RegisterCertificateGenerator(constants.SM2,SM2_CERTIFICATE_GENERATOR)

	// opts
	cache.RegisterSignerOpts(constants.ECDSA, bccsp.DefaultECDSAIntegrationOpts{})
	cache.RegisterSignerOpts(constants.SM2, bccsp.DefaultSM2IntegrationOpts{})
	cache.RegisterVerifierOpts(constants.ECDSA, bccsp.DefaultECDSAIntegrationOpts{})
	cache.RegisterVerifierOpts(constants.SM2, bccsp.DefaultSM2IntegrationOpts{})
	cache.RegisterHashOpts(constants.ECDSA, bccsp.DefaultECDSAIntegrationOpts{})
	cache.RegisterHashOpts(constants.SM2, bccsp.DefaultSM2IntegrationOpts{})
	cache.RegisterKeyGenOpts(constants.ECDSA,bccsp.DefaultECDSAIntegrationOpts{})
	cache.RegisterKeyGenOpts(constants.SM2,bccsp.DefaultSM2IntegrationOpts{})
	cache.RegisterCertificateOpts(constants.ECDSA,&base.ECDSACertificateOpts{})
	cache.RegisterCertificateOpts(constants.SM2,&base.SM2CertificateOpts{})
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 私钥导入
func SM2_KEY_IMPORTER(raw, pwd []byte) (interface{}, error) {
	if strings.Contains(string(raw), "BEGIN") {
		var block *pem.Block
		block, _ = pem.Decode(raw)
		if block == nil {
			return nil, errors.New("failed to decode private key")
		}
		raw = block.Bytes
	}
	key, e := sm2.ReadPrivateKeyFromMem(raw, pwd)
	if nil == e {
		return key, nil
	}
	if len(raw) != 32 {
		return nil, errors.New("标准私钥为32个字节")
	}
	var privateKey sm2.PrivateKey
	d := raw
	privateKey.D = new(big.Int).SetBytes(d)
	curve := sm2.P256Sm2()
	privateKey.Curve = curve
	x, y := curve.ScalarBaseMult(d)
	privateKey.X, privateKey.Y = x, y
	return &privateKey, nil
}

// ecdsa私钥读取
func ECDSA_KEY_IMPORTER(raw, pwd []byte) (interface{}, error) {
	if len(raw) == 0 {
		return nil, errors.New("Invalid PEM. It must be different from nil.")
	}

	if strings.Contains(string(raw), "BEGIN") {
		block, _ := pem.Decode(raw)
		if block == nil {
			return nil, fmt.Errorf("Failed decoding PEM. Block must be different from nil. [% x]", raw)
		}

		// TODO: derive from header the type of the key

		if x509.IsEncryptedPEMBlock(block) {
			if len(pwd) == 0 {
				return nil, errors.New("Encrypted Key. Need a password")
			}

			decrypted, err := x509.DecryptPEMBlock(block, pwd)
			if err != nil {
				return nil, fmt.Errorf("Failed PEM decryption [%s]", err)
			}

			key, err := DERToPrivateKey(decrypted)
			if err != nil {
				return nil, err
			}
			return key, err
		}
		raw = block.Bytes
	}
	cert, err := DERToPrivateKey(raw)
	if err != nil {
		return nil, err
	}
	return cert, err
}

// DERToPrivateKey unmarshals a der to private key
func DERToPrivateKey(der []byte) (key interface{}, err error) {
	if key, err = x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}

	if key, err = x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return
		default:
			return nil, errors.New("Found unknown private key type in PKCS#8 wrapping")
		}
	}

	if key, err = x509.ParseECPrivateKey(der); err == nil {
		return
	}
	// if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
	// 	return key, nil
	// } else {
	// 	fmt.Printf("error!!!!! %s", err.Error())
	// }

	return nil, errors.New("Invalid key type. The DER must contain an rsa.PrivateKey or ecdsa.PrivateKey")
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 生成公私钥
func SM2_KEYGENERATOR() (bccsp.Key, error) {
	// 调用 SM2的注册证书方法
	// privKey, err := sm2.GenerateKey()
	// modified by charlie : 调用 Java jar包生成私钥
	privKey, err := utils.GenStandardPrvKeyByJar()
	if err != nil {
		return nil, fmt.Errorf("Failed generating GMSM2 key  [%s]", err)
	}

	return &sm2PrivateKey{privKey}, nil
}

func ECDSA_KEYGENERATOR() (bccsp.Key, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", elliptic.P256(), err)
	}

	return &ecdsaPrivateKey{privateKey}, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 签名
var ECDSA_SIGNER = func(prvKInterface interface{}, msg []byte) ([]byte, error) {
	prvK := prvKInterface.(*ecdsaPrivateKey)
	r, s, err := ecdsa.Sign(rand.Reader, prvK.privKey, msg)
	if nil != err {
		return nil, errors.New("ecdsa 签名失败:" + err.Error())
	}
	signature, err := utils.MarshalECDSASignature(r, s)
	if nil != err {
		return nil, errors.New("ecdsa 签名后序列化失败:" + err.Error())
	}
	// utils2.Warning(nil, fmt.Sprintf("ecdsa签名,元数据长度:[%d],内容为:[ "+base64.StdEncoding.EncodeToString(msg)+"],签名后的数据长度:[%d],具体内容为: [ "+base64.StdEncoding.EncodeToString(signature)+" ]", len(msg), len(signature)))
	return signature, err
}
var SM2_SIGNER = func(prvKInterface interface{}, msg []byte) ([]byte, error) {
	prvK := prvKInterface.(*sm2PrivateKey)
	r, s, err := sm2.Sm2Sign(prvK.privKey, msg, nil)
	if nil != err {
		return nil, errors.New("sm2 签名失败:" + err.Error())
	}
	signature := encodeSignature(r, s)

	// utils2.Warning(nil, fmt.Sprintf("sm2签名,元数据长度:[%d],内容为:[ "+base64.StdEncoding.EncodeToString(msg)+"],签名后的数据长度:[%d],具体内容为: [ "+base64.StdEncoding.EncodeToString(signature)+" ],私钥为:[ %s ]", len(msg), len(signature), base64.StdEncoding.EncodeToString(prvK.privKey.D.Bytes())))
	return signature, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 验签
var ECDSA_VERIFIER = func(pubKInterface interface{}, signature, msg []byte) (bool, error) {
	r, s, err := utils.UnmarshalECDSASignature(signature)
	if nil != err {
		return false, errors.New("ecdsa反序列化失败:" + err.Error())
	}
	pubK := pubKInterface.(*ecdsaPublicKey)
	verify := ecdsa.Verify(pubK.pubKey, msg, r, s)
	// utils2.Warning(nil, fmt.Sprintf("ecdsa校验,结果为:[%v],signature=[%s],msg=[%s]", verify, base64.StdEncoding.EncodeToString(signature), base64.StdEncoding.EncodeToString(msg)))
	return verify, nil
}
var SM2_VERIFIER = func(pubKInterface interface{}, signature, msg []byte) (bool, error) {
	// 签名是由2个32个字节的大整数拼接而成
	r, s, err := decodeSignature(signature)
	if nil != err {
		return false, errors.New("sm2反序列化失败:" + err.Error())
	}

	pubK := pubKInterface.(*sm2PublicKey)
	verify := sm2.Sm2Verify(pubK.pubKey, msg, nil, r, s)
	// fmt.Println(r.Text(16), s.Text(16))
	// utils2.Warning(nil, fmt.Sprintf("国密校验,结果为:[%v],signature=[%s],msg=[%s]", verify, base64.StdEncoding.EncodeToString(signature), base64.StdEncoding.EncodeToString(msg)))
	return verify, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 证书导入
var SM2_CERTIFICATE_IMPORTER = func(bytes []byte) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if strings.Contains(string(bytes), "BEGIN") {
		block, _ := pem.Decode(bytes)
		if block == nil {
			return nil, nil, errors.New("sm2 解析block certificate失败")
		}
		bytes = block.Bytes
	}
	certificate, e := sm2.ParseCertificate(bytes)
	if nil != e {
		return nil, nil, errors.New("sm2 解析证书失败")
	}
	// 确保sm2证书 返回sm2 的opts
	if certificate.SignatureAlgorithm != sm2.SM2WithSM3 {
		return nil, nil, errors.New("该证书不是sm2,需要换一种方式读取")
	}
	return ParseSm2Certificate2X509(certificate), &base.SM2CertificateOpts{}, nil
}

// ecdsa 证书导入
var ECDSA_CERTIFICATE_IMPORTER = func(bytes []byte) (*x509.Certificate, base.CertificateBaseOpts, error) {
	if strings.Contains(string(bytes), "BEGIN") {
		block, _ := pem.Decode(bytes)
		if block == nil {
			return nil, nil, errors.New("ECDSA 解析block certificate失败")
		}
		bytes = block.Bytes
	}
	certificate, e := x509.ParseCertificate(bytes)
	if nil != e {
		return nil, nil, errors.New("ecdsa证书解析失败:" + e.Error())
	}
	// 确保ecdsa证书 返回ecdsa 的opts
	if !utils2.IsECDSASignedCert(certificate) {
		return nil, nil, errors.New("该证书不是ecdsa,需要换一种方式读取")
	}
	return certificate, &base.ECDSACertificateOpts{}, e
}

// X509证书格式转换为 SM2证书格式
func ParseX509Certificate2Sm2(x509Cert *x509.Certificate) *sm2.Certificate {
	sm2cert := &sm2.Certificate{
		Raw:                     x509Cert.Raw,
		RawTBSCertificate:       x509Cert.RawTBSCertificate,
		RawSubjectPublicKeyInfo: x509Cert.RawSubjectPublicKeyInfo,
		RawSubject:              x509Cert.RawSubject,
		RawIssuer:               x509Cert.RawIssuer,

		Signature:          x509Cert.Signature,
		SignatureAlgorithm: sm2.SignatureAlgorithm(x509Cert.SignatureAlgorithm),

		PublicKeyAlgorithm: sm2.PublicKeyAlgorithm(x509Cert.PublicKeyAlgorithm),
		PublicKey:          x509Cert.PublicKey,

		Version:      x509Cert.Version,
		SerialNumber: x509Cert.SerialNumber,
		Issuer:       x509Cert.Issuer,
		Subject:      x509Cert.Subject,
		NotBefore:    x509Cert.NotBefore,
		NotAfter:     x509Cert.NotAfter,
		KeyUsage:     sm2.KeyUsage(x509Cert.KeyUsage),

		Extensions: x509Cert.Extensions,

		ExtraExtensions: x509Cert.ExtraExtensions,

		UnhandledCriticalExtensions: x509Cert.UnhandledCriticalExtensions,

		// ExtKeyUsage:	[]x509.ExtKeyUsage(x509Cert.ExtKeyUsage) ,
		UnknownExtKeyUsage: x509Cert.UnknownExtKeyUsage,

		BasicConstraintsValid: x509Cert.BasicConstraintsValid,
		IsCA:                  x509Cert.IsCA,
		MaxPathLen:            x509Cert.MaxPathLen,
		// MaxPathLenZero indicates that BasicConstraintsValid==true and
		// MaxPathLen==0 should be interpreted as an actual maximum path length
		// of zero. Otherwise, that combination is interpreted as MaxPathLen
		// not being set.
		MaxPathLenZero: x509Cert.MaxPathLenZero,

		SubjectKeyId:   x509Cert.SubjectKeyId,
		AuthorityKeyId: x509Cert.AuthorityKeyId,

		// RFC 5280, 4.2.2.1 (Authority Information Access)
		OCSPServer:            x509Cert.OCSPServer,
		IssuingCertificateURL: x509Cert.IssuingCertificateURL,

		// Subject Alternate Name values
		DNSNames:       x509Cert.DNSNames,
		EmailAddresses: x509Cert.EmailAddresses,
		IPAddresses:    x509Cert.IPAddresses,

		// Name constraints
		PermittedDNSDomainsCritical: x509Cert.PermittedDNSDomainsCritical,
		PermittedDNSDomains:         x509Cert.PermittedDNSDomains,

		// CRL Distribution Points
		CRLDistributionPoints: x509Cert.CRLDistributionPoints,

		PolicyIdentifiers: x509Cert.PolicyIdentifiers,
	}
	for _, val := range x509Cert.ExtKeyUsage {
		sm2cert.ExtKeyUsage = append(sm2cert.ExtKeyUsage, sm2.ExtKeyUsage(val))
	}

	return sm2cert
}

func ParseX509Options2SM2Options(ops *x509.VerifyOptions) *sm2.VerifyOptions {
	intermediatesPool := sm2.NewCertPool()
	rootPool := sm2.NewCertPool()
	if ops.Intermediates != nil {
		rf := reflect.ValueOf(ops.Intermediates)
		certs := GetUnexportedField(rf.Elem().FieldByName("certs"))
		if nil != certs {
			certificates := certs.([]*x509.Certificate)
			for _, crt := range certificates {
				intermediatesPool.AddCert(ParseX509Certificate2Sm2(crt))
			}
		}
	}
	if ops.Roots != nil {
		rf := reflect.ValueOf(ops.Roots)
		certs := GetUnexportedField(rf.Elem().FieldByName("certs"))
		if nil != certs {
			certificates := certs.([]*x509.Certificate)
			for _, crt := range certificates {
				rootPool.AddCert(ParseX509Certificate2Sm2(crt))
			}
		}
	}

	ks := make([]sm2.ExtKeyUsage, 0)
	for _, v := range ops.KeyUsages {
		ks = append(ks, sm2.ExtKeyUsage(v))
	}
	res := &sm2.VerifyOptions{
		DNSName:       ops.DNSName,
		Intermediates: intermediatesPool,
		Roots:         rootPool,
		CurrentTime:   ops.CurrentTime,
		KeyUsages:     ks,
	}
	return res
}

func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

// sm2 证书转换 x509 证书
func ParseSm2Certificate2X509(sm2Cert *sm2.Certificate) *x509.Certificate {
	x509cert := &x509.Certificate{
		Raw:                     sm2Cert.Raw,
		RawTBSCertificate:       sm2Cert.RawTBSCertificate,
		RawSubjectPublicKeyInfo: sm2Cert.RawSubjectPublicKeyInfo,
		RawSubject:              sm2Cert.RawSubject,
		RawIssuer:               sm2Cert.RawIssuer,

		Signature:          sm2Cert.Signature,
		SignatureAlgorithm: x509.SignatureAlgorithm(sm2Cert.SignatureAlgorithm),

		PublicKeyAlgorithm: x509.PublicKeyAlgorithm(sm2Cert.PublicKeyAlgorithm),
		PublicKey:          sm2Cert.PublicKey,

		Version:      sm2Cert.Version,
		SerialNumber: sm2Cert.SerialNumber,
		Issuer:       sm2Cert.Issuer,
		Subject:      sm2Cert.Subject,
		NotBefore:    sm2Cert.NotBefore,
		NotAfter:     sm2Cert.NotAfter,
		KeyUsage:     x509.KeyUsage(sm2Cert.KeyUsage),

		Extensions: sm2Cert.Extensions,

		ExtraExtensions: sm2Cert.ExtraExtensions,

		UnhandledCriticalExtensions: sm2Cert.UnhandledCriticalExtensions,

		// ExtKeyUsage:	[]x509.ExtKeyUsage(sm2Cert.ExtKeyUsage) ,
		UnknownExtKeyUsage: sm2Cert.UnknownExtKeyUsage,

		BasicConstraintsValid: sm2Cert.BasicConstraintsValid,
		IsCA:                  sm2Cert.IsCA,
		MaxPathLen:            sm2Cert.MaxPathLen,
		// MaxPathLenZero indicates that BasicConstraintsValid==true and
		// MaxPathLen==0 should be interpreted as an actual maximum path length
		// of zero. Otherwise, that combination is interpreted as MaxPathLen
		// not being set.
		MaxPathLenZero: sm2Cert.MaxPathLenZero,

		SubjectKeyId:   sm2Cert.SubjectKeyId,
		AuthorityKeyId: sm2Cert.AuthorityKeyId,

		// RFC 5280, 4.2.2.1 (Authority Information Access)
		OCSPServer:            sm2Cert.OCSPServer,
		IssuingCertificateURL: sm2Cert.IssuingCertificateURL,

		// Subject Alternate Name values
		DNSNames:       sm2Cert.DNSNames,
		EmailAddresses: sm2Cert.EmailAddresses,
		IPAddresses:    sm2Cert.IPAddresses,

		// Name constraints
		PermittedDNSDomainsCritical: sm2Cert.PermittedDNSDomainsCritical,
		PermittedDNSDomains:         sm2Cert.PermittedDNSDomains,

		// CRL Distribution Points
		CRLDistributionPoints: sm2Cert.CRLDistributionPoints,

		PolicyIdentifiers: sm2Cert.PolicyIdentifiers,
	}
	for _, val := range sm2Cert.ExtKeyUsage {
		x509cert.ExtKeyUsage = append(x509cert.ExtKeyUsage, x509.ExtKeyUsage(val))
	}

	return x509cert
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 公钥导入
var ECDSA_PUBLICKEY_IMPORTER = func(raw interface{}) (bccsp.Key, error) {
	var (
		lowLevelKey *ecdsa.PublicKey
	)
	if bytes, ok := raw.([]byte); ok {
		if res, err := bccsp.OriginPublicKeyImport(bytes, base.FailFast, base.EcdsaOpts{}); nil != err {
			return nil, errors.New("ecdsa 原生公钥解析失败:" + err.Error())
		} else {
			lowLevelKey = res.(*ecdsa.PublicKey)
		}
	} else {
		lowLevelKey, ok = raw.(*ecdsa.PublicKey)
		if !ok {
			return nil, errors.New("Invalid raw material. Expected *ecdsa.PublicKey.")
		}
	}

	return &ecdsaPublicKey{lowLevelKey}, nil
}
var SM2_PUBLICKEY_IMPORTER = func(raw interface{}) (bccsp.Key, error) {
	var (
		lowLevelKey *sm2.PublicKey
	)

	if bytes, ok := raw.([]byte); ok {
		if res, err := bccsp.OriginPublicKeyImport(bytes, base.FailFast, base.SM2Opts{}); nil != err {
			return nil, errors.New("ecdsa 原生公钥解析失败:" + err.Error())
		} else {
			lowLevelKey = res.(*sm2.PublicKey)
		}
	} else {
		ecdsapubk, okk := raw.(*ecdsa.PublicKey)
		if !okk {
			return nil, errors.New("Invalid raw material. Expected *ecdsa.PublicKey.")
		}
		lowLevelKey = &sm2.PublicKey{
			Curve: ecdsapubk.Curve,
			X:     ecdsapubk.X,
			Y:     ecdsapubk.Y,
		}
	}

	return &sm2PublicKey{lowLevelKey}, nil
}

var RSA_PUBLICKEY_IMPORTER = func(raw interface{}) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *rsa.PublicKey.")
	}

	return &rsaPublicKey{lowLevelKey}, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 原生公钥导入
var ECDSA_ORIGIN_PUBLICKEY_IMPORTER = func(bytes []byte) (interface{}, error) {
	// if len(bytes) == 0 {
	// 	return nil, errors.New("[ECDSA_PUBLICKEY_IMPORTER] Invalid raw. It must not be nil.")
	// }
	return x509.ParsePKIXPublicKey(bytes)
}

// 原生sm2 公钥导入
var SM2_ORIGIN_PUBLICKEY_IMPORTER = func(bytes []byte) (interface{}, error) {
	k, err := sm2.ParseSm2PublicKey(bytes)
	if nil != err {
		return nil, errors.New("sm2 公钥导入失败:" + err.Error())
	}
	return k, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 序列化,反序列化
var ECDSA_UNMARSHAL = func(sig []byte) (*big.Int, *big.Int, error) {
	return utils.UnmarshalECDSASignature(sig)
}
var SM2_UNMARSHAL = func(sig []byte) (*big.Int, *big.Int, error) {
	signature := string(sig)
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

// 序列化
var ECDSA_MARSHAL = func(r, s *big.Int) ([]byte, error) {
	return utils.MarshalECDSASignature(r, s)
}
var SM2_MARSHAL = func(r, s *big.Int) ([]byte, error) {
	fmt.Printf("siganture r=%v, s=%v\n", r.Text(16), s.Text(16))
	return []byte(fmt.Sprintf("%s%s", r.Text(16), s.Text(16))), nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// hash
var SHA256_HASH = func(msg []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(msg)
	out2 := h.Sum(nil)
	return out2, nil
}

var SM3_HASH_WITH_NO = func(msg []byte) ([]byte, error) {
	return msg, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 公私钥生成
var ECDSA_PRIVATEKEY_GENERATOR = func() (bccsp.IPrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to generate private key")
	}

	// pkcs8Encoded, err := x509.MarshalPKCS8PrivateKey(priv)
	// if err != nil {
	// 	return nil, errors.WithMessage(err, "failed to marshal private key")
	// }
	return &ecdsaPrivateKey{
		privKey: priv,
	},nil
	// pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: pkcs8Encoded})
	// return pemEncoded,nil
}

var SM2_PRIVATEKEY_GENERATOR = func() (bccsp.IPrivateKey, error) {
	// GenStandardPrvKeyByJar
	key, e := utils.GenStandardPrvKeyByJar()
	if nil != e {
		return nil, errors.New("生成sm2私钥失败:" + e.Error())
	}
	return &sm2PrivateKey{
		privKey: key,
	},nil
	// return key.D.Bytes(), nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////


// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
// certificate generator
var ECDSA_CERTIFICATE_GENERATOR=func(builder base.ParamBuilder)(*x509.Certificate,[]byte,error){
	params:=builder()
	parent:=params["parent"].(*x509.Certificate)
	template:=params["template"].(*x509.Certificate)
	pub:=params["pub"]
	priv:=params["priv"]
	// create the x509 public cert
	certBytes, err := x509.CreateCertificate(rand.Reader, template, parent, pub, priv)
	if err != nil {
		return nil,nil, err
	}
	x509Cert, err := x509.ParseCertificate(certBytes)
	return x509Cert,certBytes,err
}

var SM2_CERTIFICATE_GENERATOR= func(builder base.ParamBuilder) (*x509.Certificate,[]byte,error){
	params:=builder()
	parent:=params["parent"].(*x509.Certificate)
	template:=params["template"].(*x509.Certificate)
	pub:=params["pub"]
	priv:=params["priv"]
	sm2Template := ParseX509Certificate2Sm2(template)
	parentTemplate:=ParseX509Certificate2Sm2(parent)
	cert, err := sm2.CreateCertificate(rand.Reader, sm2Template, parentTemplate, pub, priv)
	if nil!=err{
		return nil,nil,err
	}
	certificate, err := sm2.ParseCertificate(cert)
	if nil!=err{
		return nil,nil,err
	}
	return ParseSm2Certificate2X509(certificate),cert,nil
}



// //////////////////////////////////////////////////////////////////////////////////////////////////////////////