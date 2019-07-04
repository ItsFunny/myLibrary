package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/wumansgy/goEncrypt"
	"io/ioutil"
	"os"
)

type RSAEncryptModel struct {
	EncryptTime int64
	PublickKey  []byte
	PrivateKey  []byte
}

func Gen1024RSAKey() ([]byte, []byte, error) {
	return GenRsaKey(1024)
}

// RSA公钥私钥产生
func GenRsaKey(bits int) ([]byte, []byte, error) {
	var (
		privateBytes, publicBytes []byte
	)
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}
	if bytes, err := ioutil.ReadFile("private.pem"); nil != err {
		fmt.Println(err)
	} else {
		// fmt.Println(hex.EncodeToString(bytes))
		// privateK = string(bytes)
		privateBytes = bytes
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return nil, nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, nil, err
	}
	if bytes, err := ioutil.ReadFile("public.pem"); nil != err {
		return nil, nil, err
	} else {
		// fmt.Println(hex.EncodeToString(bytes))
		// publicK = string(bytes)
		publicBytes = bytes
	}
	go func() {
		os.Remove("public.pem")
		os.Remove("private.pem")
	}()
	return privateBytes, publicBytes, nil
}

func encrypt(datas, key []byte) ([]byte, error) {
	return goEncrypt.RsaEncrypt(datas, key)
}
func decrypt(datas, key []byte) ([]byte, error) {
	return goEncrypt.RsaDecrypt(datas, key)
}

func RSASign(item interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(item)
	if cryptText, err := goEncrypt.RsaSign(bytes, model.PrivateKey); nil != err {
		log.Errorf("[RSAEncrypt] sign failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSAVeriSign(item interface{}, signCodes []byte, model RSAEncryptModel) bool {
	bytes, _ := json.Marshal(item)
	return goEncrypt.RsaVerifySign(bytes, signCodes, model.PublickKey)
}

// rsa 加密
func RSAEncryptByPub(item interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(item)
	if cryptText, err := encrypt(bytes, model.PublickKey); nil != err {
		log.Errorf("[RSAEncrypt] encrypt failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSAEncryptByPrv(data interface{}, model RSAEncryptModel) ([]byte, error) {
	bytes, _ := json.Marshal(data)
	if cryptText, err := encrypt(bytes, model.PrivateKey); nil != err {
		log.Errorf("[RSAEncrypt] encrypt failed:%v", err.Error())
		return nil, err
	} else {
		return cryptText, nil
	}
}

func RSADecrypt(encryptStr string, key []byte) ([]byte, error) {
	// if plainText, err := goEncrypt.RsaDecrypt([]byte(encryptStr), property.PrivateBytes); nil != err {
	bytes, e := hex.DecodeString(encryptStr)
	if nil != e {
		return nil, e
	}
	if plainText, err := goEncrypt.RsaDecrypt(bytes, key); nil != err {
		log.Errorf("[RSADecrypt]faield:%v", err.Error())
		return nil, err
	} else {
		return plainText, nil
	}
}



func MD5Encrypt(bytes []byte) string {
	h := md5.New()
	h.Write(bytes) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
