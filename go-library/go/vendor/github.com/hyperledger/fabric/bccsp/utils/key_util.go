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
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const DEFAULT_JAVA_JAR_PATH = "/usr/local/bin"
const ENV_FABRIC_SM2_KEY_GENERATOR_PATH = "FABRIC_SM2_KEY_GENERATOR_PATH"

// 通过环境变量获取生成私钥的jar包路径
func GenStandardPrvKeyByJar() (*sm2.PrivateKey, error) {
	path := DEFAULT_JAVA_JAR_PATH
	tmpPath := os.Getenv(ENV_FABRIC_SM2_KEY_GENERATOR_PATH)
	if tmpPath != "" {
		path = tmpPath
	}
	jarPath := path
	if !strings.ContainsAny(path, "jar") {
		jarPath = filepath.Join(path, "sm2KeyGenerator.jar")
		if _, err := os.Stat(jarPath); os.IsNotExist(err) {
			return nil, fmt.Errorf("sm2Generator jar file[%s] not exists", jarPath)
		}
	}
	if _, err := os.Stat(jarPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("sm2Generator jar file[%s] not exists", jarPath)
	}

	cmd := exec.Command("java", "-jar", jarPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New("调用命令行失败:" + err.Error() + "," + string(output))
	}

	var privateKey sm2.PrivateKey
	d, err := base64.StdEncoding.DecodeString(string(output))
	if err != nil {
		return nil, err
	}
	privateKey.D = new(big.Int).SetBytes(d)
	curve := sm2.P256Sm2()
	privateKey.Curve = curve
	x, y := curve.ScalarBaseMult(d)
	privateKey.X, privateKey.Y = x, y
	return &privateKey, nil
}

func ParseInterface2SM2PubKey(pubInterface interface{}) (*sm2.PublicKey, error) {
	if sm2Pub, ok := pubInterface.(*sm2.PublicKey); ok {
		return sm2Pub, nil
	} else if ecdsaPub, ok := pubInterface.(*ecdsa.PublicKey); ok {
		return ParseEcdsaPubKey2SM2PubK(ecdsaPub), nil
	}
	return nil, errors.New("无法解析的类型,该key既不是sm2的公钥,也不是ecdsa的公钥")
}
func ParseEcdsaPub2SM2PubKey(key *ecdsa.PublicKey) *sm2.PublicKey {
	return &sm2.PublicKey{
		Curve: key.Curve,
		X:     key.X,
		Y:     key.Y,
	}
}

// 三种方式读取私钥   block解析丢里面去了,所以可能其他地方会有问题?
func ThreeWaysReadPrvKey(raw []byte, pwd []byte) (interface{}, error) {
	// var block *pem.Block
	// block, _ = pem.Decode(raw)
	// if block == nil {
	// 	return nil, errors.New("failed to decode private key")
	// }
	pipeline := []PRV_READER{STANDARD_PRV_READER, ECDSA_PRV_READER, COMMON_SM2_PRV_READER}
	var key interface{}
	var e error
	// raw = block.Bytes
	for _, reader := range pipeline {
		key, e = reader(raw, pwd)
		if nil != e {
			continue
		}
		return key, nil
	}
	return nil, errors.New("无法解析私钥:" + e.Error())
}

type PRV_READER = func(raw, pwd []byte) (interface{}, error)

func STANDARD_PRV_READER(raw, pwd []byte) (interface{}, error) {
	if strings.Contains(string(raw), "BEGIN") {
		var block *pem.Block
		block, _ = pem.Decode(raw)
		if block == nil {
			return nil, errors.New("failed to decode private key")
		}
		raw = block.Bytes
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

// 常见的私钥读取
func COMMON_SM2_PRV_READER(raw, pwd []byte) (interface{}, error) {
	return sm2.ReadPrivateKeyFromMem(raw, pwd)
}

// ecdsa私钥读取
func ECDSA_PRV_READER(raw, pwd []byte) (interface{}, error) {
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

	if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return key, nil
	} else {
		fmt.Printf("error!!!!! %s", err.Error())
	}

	return nil, errors.New("Invalid key type. The DER must contain an rsa.PrivateKey or ecdsa.PrivateKey")
}
