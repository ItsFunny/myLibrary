package encrypt

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"myLibrary/library/src/main/go/converters"
	"os"
	"testing"
)

const (
	pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDYQJZ+kaWOlW9RD7B3PfFc/OOa
ShEcPhXs8ZW1gb9KGFlkLpslVjvqsbfZXU07i2zR6WjzxhW68l0DkKfMo6IuOZw7
0HxLLtiNYF9CXL08Oc+NIGiUk6buoMQ3xMHqv69mN13dTMGxgdItYFULaL4r5uAj
CHXJuizxsxFWadpnqQIDAQAB
-----END PUBLIC KEY-----`
)

func TestRSAEncryptByPrv(t *testing.T) {
	bytes, e := RSAEncryptByPub("123", RSAEncryptModel{
		PublickKey: []byte(pubKey),
	})
	if nil != e {
		panic(e)
	}
	str := hex.EncodeToString(bytes)
	fmt.Println(str)
	rsaDecrypt, e := RSADecrypt(str, []byte(privateKey))
	if nil != e {
		panic(e)
	}
	fmt.Println(string(rsaDecrypt))
}

const (
	value      = `ceaJqFbjTblbVhlMXgyXj0SVLzVC7Z60N0yZt650H1dI973fSfJf777mwye3tDjFXlxsmgA/tDYGUxB0xYPz8I/o4nWuH2xKuQxwvYdnNL0aNa5yxEEiOuKm7f4d489q2IhjjyNaL3GQfb5QVuRRJYZYHWhUiiJPR5lGvYiSQfI=`
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDYQJZ+kaWOlW9RD7B3PfFc/OOaShEcPhXs8ZW1gb9KGFlkLpsl
VjvqsbfZXU07i2zR6WjzxhW68l0DkKfMo6IuOZw70HxLLtiNYF9CXL08Oc+NIGiU
k6buoMQ3xMHqv69mN13dTMGxgdItYFULaL4r5uAjCHXJuizxsxFWadpnqQIDAQAB
AoGBAMmTQvUODHWhI9c6ubIc4jxcWkg6nEJoOQXZl0ROgiGuGr1AmEuwWz+Ekywh
RVmouDHe0X7R+PV+72bgUxl0ujmGoZLNCjrtt9KHOZkX4xryZRKcBOogxHuno9fE
6m/FWbuzlc1LaTCD5XUGoFaCHj2teIB8431gXhuIQHjEPJUBAkEA/Mqe0qPE5OhT
zCHZWmBtGmUWOPuOcYI9g7M6PsSNfp0H8s7ehBSGFqfYJ6nrXhOSsGTA91IInMb2
mVb/9yK+oQJBANr/PjKOHLGx1YlYAbNKBOLYQiPCkgV3+FLydcUKbGu/LRO+iLRK
ksRaqu1kHE5RARDOwnMD/QqGqfCDY+9JNAkCQQCgwgYAdE9JNwnbPgdoNvwLFg/s
yuTKAIY4E6lNs1c8Foawfaf6HMcs7y5CAwf/+riFXn1siomZkdnOqAn9UuYhAkEA
vmV/G9D2HX9xGXGMOOY3jUlbZ1+4OEzvdp4Zye+gB6U0eaADlkvnghMZ3D5XZbeD
Z6t63ygujUI66UIyVk2ckQJAM9toHYn5aHCwSLmzMgQSBSNfQApP4NptkNsR2UEM
0AsibeVgoflh82Ss7Ns5JsA9mX7Zhj+o75PW8mc06Nra9g==
-----END RSA PRIVATE KEY-----
`
)

func TestRSADecrypt(t *testing.T) {
	pub, _ := RSAEncryptByPub("123", RSAEncryptModel{
		PublickKey: []byte(pubKey),
	})

	s := base64.StdEncoding.EncodeToString(pub)
	fmt.Println(s)

	decodeString, e := base64.StdEncoding.DecodeString(s)
	if nil != e {
		panic(e)
	} else {
		fmt.Println(string(decodeString))
	}

	bytes, e := rsaDecrypt(decodeString, []byte(privateKey))
	if nil != e {
		panic(e)
	} else {
		fmt.Println(string(bytes))
	}
}
func rsaDecrypt(value []byte, key []byte) ([]byte, error) {
	if plainText, err := goEncrypt.RsaDecrypt(value, key); nil != err {
		return nil, err
	} else {
		return plainText, nil
	}
}

type SS struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMD5Encrypt(t *testing.T) {
	baseBytes := make([]byte, 0)

	transType := 1
	int642Bytes := converter.BigEndianInt642Bytes(int64(transType))
	fmt.Println(int642Bytes)
	baseBytes = int642Bytes[4:]
	fmt.Println(baseBytes)

	str := "vmV/G9D2HX9xGXGMOOY3jUlbZ1+4OEzvdp4Zye+gB6U0eaADlkvnghMZ3D5XZbeD"
	walletAddress := MD5Encrypt([]byte(str))
	bytes := []byte(walletAddress)

	baseBytes = append(baseBytes, bytes...)
	fmt.Println(baseBytes)
	fmt.Println(len(baseBytes))

	m := SS{
		Name: "joker",
		Age:  23,
	}
	marshal, _ := json.Marshal(m)
	baseBytes = append(baseBytes, marshal...)
	fmt.Println(baseBytes)
	fmt.Println(len(baseBytes))

	// 反序列化

	typeBytes := make([]byte, 4)
	typeBytes = append(typeBytes, baseBytes[:4]...)
	fmt.Println(typeBytes)
	fmt.Println(converter.BigEndianBytes2Int64(typeBytes))
	walletBytes := baseBytes[4:36]
	fmt.Println(string(walletBytes))
	fmt.Println(string(walletBytes) == walletAddress)

	modelBytes := baseBytes[36:]
	var s SS
	unmarshal := json.Unmarshal(modelBytes, &s)
	if nil != unmarshal {
		panic(unmarshal)
	}
	fmt.Println(s)

}

func TestMD5EncryptFile(t *testing.T) {
	path := "/Users/joker/Downloads/fengkuangwaixingren.mp4"
	file, _ := os.Open(path)
	defer file.Close()

	encryptFile := MD5EncryptFile(file)
	fmt.Println(encryptFile)
}

func TestAesEncrypt(t *testing.T) {
	str := ""
	aseKey := []byte("321423u9y8d2fwfl")
	bytes, e := AesEncrypt([]byte(str), aseKey)
	if nil != e {
		fmt.Println(e)
	} else {
		fmt.Println(hex.EncodeToString(bytes))
	}
}
func TestAesDecrypt(t *testing.T) {
	str := "aaa"
	aseKey := []byte("9f07debc1dde4ae58f41fef3b6aca13c")
	fmt.Println(len(aseKey))
	bytes, e := AesEncrypt([]byte(str), aseKey)
	if nil != e {
		fmt.Println(e)
	} else {
		fmt.Println(hex.EncodeToString(bytes))
	}
	aesDecrypt, e := AesDecrypt(bytes, []byte(aseKey))
	if nil != e {
		fmt.Println(e)
	} else {
		fmt.Println(string(aesDecrypt))
	}
}

func TestMD5Encrypt2(t *testing.T) {
	str := "FDSW$t34tregt5tO&$(#RHuyoyiUYE*&OI$HRLuy87odlfh)"
	key := "12346z"
	str = key + str

	bytes1 := []byte(str)
	bytes2, _ := hex.DecodeString(str)

	md5Encrypt := MD5Encrypt(bytes1)
	fmt.Println(md5Encrypt)

	md52 := MD5Encrypt(bytes2)
	fmt.Println(md52)
}

func TestBytes2ECDSAPrv(t *testing.T) {

}

func TestECCSign(t *testing.T) {
	prvStr := "c402a9081e4736760518271ed7652ff6cf2b688359a5f328215863c67a550a8f"
	prvBytes, _ := hex.DecodeString(prvStr)
	bytes, e := ECCSign("joker", prvBytes)
	if nil != e {
		panic(e)
	}
	pubStr := "0454fb4158e1ec31286e2e825967380a1c75f31d4cf438430b24e84387e2e3fc92eb0242babdd3bfbc3a31edd3090fc32e20428b89753d6c2b0f8419af6c484d42"
	pubBytes, _ := hex.DecodeString(pubStr)
	signature := hex.EncodeToString(bytes)
	b := ECCVerifySignWithHex("joker", signature, pubBytes)
	fmt.Println(b)
}

func TestECCVerifySignWithHex(t *testing.T) {
	signature := "969409a0fbbf1aea899554299d673bef8efe99183bb0dd2a37332b8430bb85a3b8f89bab7668e7edd9ca4008515272b6fd13f3b383da71092625fb361505d610"
	md5Codes := "ecffa640755d2778355e80786a2b689a"
	pubStr := "0454fb4158e1ec31286e2e825967380a1c75f31d4cf438430b24e84387e2e3fc92eb0242babdd3bfbc3a31edd3090fc32e20428b89753d6c2b0f8419af6c484d42"
	pubBytes, _ := hex.DecodeString(pubStr)
	b := ECCVerifySignWithHex(md5Codes, signature, pubBytes)
	fmt.Println(b)

	m := make(map[string]string)
	if v, exist := m["key"]; exist {
		fmt.Println(v)
	}

	delete(m, "key")
	a := interface{}("123")
	switch a.(type) {
	case map[string]string:
		println("map")
	case string:
		fmt.Println("string")
	}
}
