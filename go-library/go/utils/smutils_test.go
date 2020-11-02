/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-05 12:07
# @File : smutils.go
# @Description :
# @Attention :
*/
package utils

import (
	"crypto/ecdsa"
	"encoding/asn1"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"myLibrary/go-library/go/tjfoc/gmsm/sm2"
	"testing"
)

// MIICNDCCAdigAwIBAgIPBwACICAQBhAAAAIAGEJiMAwGCCqBHM9VAYN1BQAwYDELMAkGA1UEBhMCQ04xMjAwBgNVBAoMKUdVQU5HIERPTkcgQ0VSVElGSUNBVEUgQVVUSE9SSVRZIENPLixMVEQuMR0wGwYDVQQDDBRHRENBIFRydXN0QVVUSCBFMSBDQTAeFw0yMDEwMDUxNjAwMDBaFw0yMTEwMDYwODU5MTNaMEcxFTATBgNVBAoeDABiAGkAZABzAHUAbjEVMBMGA1UECx4MAGIAaQBkAHMAdQBuMRcwFQYDVQQDHg4AbwByAGcATgBhAG0AZTBZMBMGByqGSM49AgEGCCqBHM9VAYItA0IABKm1My9iSigG
func TestReadSM2CertFromFile(t *testing.T) {
	path := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-fabric-ca/src/main/resources/crt/demoOrg.crt"
	certificate, e := ReadSM2CertFromFile(path)
	if nil != e {
		log.Fatal(e)
	} else {
		fmt.Println(certificate)
	}
}

func TestGM(t *testing.T) {
	GM()
}

func TestParseSMPrvFromFile(t *testing.T) {
	prvPath := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-fabric-ca/src/main/resources/prv/demoOrg.pem"
	crtPath := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-fabric-ca/src/main/resources/crt/demoOrg.crt"
	prvK, pubK, e := ParseSMKeyPairFromFile(prvPath, crtPath)
	if nil != e {
		log.Fatal(e)
	}
	bytes := []byte("123")
	r, s, e := sm2.Sign(prvK, bytes)
	if nil != e {
		log.Fatal(e)
	}
	verify := sm2.Verify(pubK, bytes, r, s)
	assert.Equal(t, true, verify)
}

func TestParseSMKeyPairFromFile(t *testing.T) {
	str := "xJhcWYZ26/xRUJHc8dwxfLoKeCB3BgVND/122GMdFD4="
	bytes, _ := Base64Decode(str)
	key, e := sm2.ParseSm2PrivateKey(bytes)
	if nil != e {
		fmt.Println("der 解析失败")
	}
	fmt.Println(key)
	privateKey, e := sm2.ParsePKCS8UnecryptedPrivateKey(bytes)
	if nil != e {
		fmt.Println("pkcs8 解析失败:" + e.Error())
	}
	fmt.Println(privateKey)
	pkcs1PrivateKey, e := sm2.ParsePKCS1PrivateKey(bytes)
	if nil != e {
		fmt.Println("pkcs1解析失败:" + e.Error())
	}
	fmt.Println(pkcs1PrivateKey)
}

func TestParseSM2PrvKBytesWithPubKey(t *testing.T) {
	pvStr := `MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg+/7WIrdWtr+85blX
4EQLCF8yN4CQYshSvBxIUFfInWWgCgYIKoEcz1UBgi2hRANCAAQlU4+tM2dAb5km
S72uRsGADj8707fVyA/m4YxJ/puwU57SMQDIn2TV8WDOhQfgKVdjiWRu12YvvPT1
anvgxLGn`
	bytes, _ := Base64Decode(pvStr)
	key, e := sm2.ParsePKCS8UnecryptedPrivateKey(bytes)
	if nil != e {
		panic(e)
	}
	bs := []byte("123")
	r, s, e := sm2.Sign(key, bs)
	if nil != e {
		panic(e)
	}

	type Sm2Sing struct {
		R *big.Int
		S *big.Int
	}
	unmarshal, err := asn1.Marshal(&Sm2Sing{
		R: r,
		S: s,
	})
	if nil != err {
		panic(err)
	}
	fmt.Println(unmarshal)
}

func TestDERToPrivateKey(t *testing.T) {
	prvP := `MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg+/7WIrdWtr+85blX
4EQLCF8yN4CQYshSvBxIUFfInWWgCgYIKoEcz1UBgi2hRANCAAQlU4+tM2dAb5km
S72uRsGADj8707fVyA/m4YxJ/puwU57SMQDIn2TV8WDOhQfgKVdjiWRu12YvvPT1
anvgxLGn`
	bytes, _ := Base64Decode(prvP)
	privateKey, e := sm2.ParsePKCS8UnecryptedPrivateKey(bytes)
	if nil != e {
		log.Fatal(e)
	}
	certP := `MIICJzCCAc2gAwIBAgIQe47KVEbNbZaykHFiOAEb9zAKBggqgRzPVQGDdTBeMQsw
CQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzET
MBEGA1UEChMKYmlkc3VuLmNvbTEWMBQGA1UEAxMNY2EuYmlkc3VuLmNvbTAeFw0y
MDEwMTAwODE4MDBaFw0zMDEwMDgwODE4MDBaMF4xCzAJBgNVBAYTAkNOMRAwDgYD
VQQIEwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMRMwEQYDVQQKEwpiaWRzdW4u
Y29tMRYwFAYDVQQDEw1jYS5iaWRzdW4uY29tMFkwEwYHKoZIzj0CAQYIKoEcz1UB
gi0DQgAEdkhJvbM+4FEK81+eKeBYW8pfyRNOYxjs0R7Slu1Zt8bM5MvkOaLAFPsk
TmC/oQO4rR2NRoDt26DUCzDI5H334qNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1Ud
JQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1Ud
DgQiBCDeu/pnTAdiOq/Zxwu8eDnpPPq78VbGOgg6EVT4W3w/kDAKBggqgRzPVQGD
dQNIADBFAiEAtuLroXjUYcV4887huh6xbhM8Aw5VgoTYsfKlH0lv8BQCICH0e/iz
UBEtNyL3b73UozVhvw5TPVLSKDJXiEtnKceK`
	bytes, _ = Base64Decode(certP)
	certificate, e := sm2.ParseCertificate(bytes)
	if nil != e {
		log.Fatal(e)
	}

	// key, publicKey, e := ParseSMKeyPairFromFile(prvP, certP)
	// if nil != e {
	// 	log.Fatal(e)
	// }
	bs := []byte("123")
	r, s, e := sm2.Sign(privateKey, bs)
	if nil != e {
		log.Fatal(e)
	}
	key := certificate.PublicKey.(*ecdsa.PublicKey)
	pubKey := &sm2.PublicKey{
		Curve: key.Curve,
		X:     key.X,
		Y:     key.Y,
	}
	verify := sm2.Verify(pubKey, bs, r, s)
	fmt.Println(verify)

}

func TestDERToPrivateKey2ByFile(t *testing.T) {
	prvP := `/Users/joker/Java/ebidsun/bidsun-sdk/bidsun-fabric-client/src/main/testresources/prv_sk`
	// bytes, _ := ioutil.ReadFile(prvP)
	privateKey, e := sm2.ReadPrivateKeyFromPem(prvP, nil)
	if nil != e {
		log.Fatal(e)
	}

	certP := "/Users/joker/Java/ebidsun/bidsun-sdk/bidsun-fabric-client/src/main/testresources/Admin@bidsun.com-cert.pem"
	// bytes, _ = ioutil.ReadFile(certP)
	certificate, e := sm2.ReadCertificateFromPem(certP)
	if nil != e {
		log.Fatal(e)
	}
	bs := []byte("123")
	r, s, e := sm2.Sign(privateKey, bs)
	if nil != e {
		log.Fatal(e)
	}
	key := certificate.PublicKey.(*ecdsa.PublicKey)
	pubKey := &sm2.PublicKey{
		Curve: key.Curve,
		X:     key.X,
		Y:     key.Y,
	}
	verify := sm2.Verify(pubKey, bs, r, s)
	fmt.Println(verify)

}

func TestJavaGo(t *testing.T) {
	str := "MEYCIQDaOoTzhN3fjBAGU4PbZeBHKvPPTO9NQeTBov3N+4mDjgIhAIYK0UBNnjfsVZWqQ/83DZIQw8YiQlAvSRTclAO9nVBt"
	rawSig, _ := Base64Decode(str)
	R, S, e := UnmarshalSM2Signature(rawSig)
	if nil != e {
		log.Fatal(e)
	}

	// prvP := `/Users/joker/Java/ebidsun/bidsun-sdk/bidsun-fabric-client/src/main/testresources/prv_sk`
	// // bytes, _ := ioutil.ReadFile(prvP)
	// privateKey, e := sm2.ReadPrivateKeyFromPem(prvP,nil)
	// if nil != e {
	// 	log.Fatal(e)
	// }
	// bs:=[]byte("123")
	// r, s, e := sm2.Sign(privateKey, bs)
	// if nil!=e{
	// 	log.Fatal(e)
	// }
	// R=r
	// S=s
	certP := "/Users/joker/Java/ebidsun/bidsun-sdk/bidsun-fabric-client/src/main/testresources/Admin@bidsun.com-cert.pem"
	certificate, e := sm2.ReadCertificateFromPem(certP)
	if nil != e {
		log.Fatal(e)
	}

	key := certificate.PublicKey.(*ecdsa.PublicKey)
	pubKey := &sm2.PublicKey{
		Curve: key.Curve,
		X:     key.X,
		Y:     key.Y,
	}
	bytes, e := sm2.MarshalSm2PublicKey(pubKey)
	if nil != e {
		log.Fatal(e)
	}
	// Java编译的公钥:MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEKuaMXFbkFSTPhdyiwK26TfkjlQyLaxfrqKgZeVzmPz4kyiC+Z/z4bfCK7m7ZGHat+hlzQL/suB9VapqjyzKskg==
	// Go解析的公钥: MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEKuaMXFbkFSTPhdyiwK26TfkjlQyLaxfrqKgZeVzmPz4kyiC+Z/z4bfCK7m7ZGHat+hlzQL/suB9VapqjyzKskg==
	fmt.Println(Base64Encode(bytes))
	bs := []byte("123456")
	verify := sm2.Sm2Verify(pubKey, bs, nil, R, S)
	fmt.Println(verify)
	// utils.UnmarshalSM2Signature(signature)
}

// 32,64 test
func TestX509ReadCertificate(t *testing.T) {
	certPath := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-fabric-ca/src/main/resources/crt/demoOrg.crt"
	prvPath := "/Users/joker/Java/ebidsun/bidsun-tool/bidsun-tool-fabric-ca/src/main/resources/prv/demoOrg.pem"
	key, publicKey, e := ParseSMKeyPairFromFile(prvPath, certPath)
	if nil != e {
		log.Fatal(e)
	}
	bs := []byte("123")
	r, s, e := sm2.Sm2Sign(key, bs, nil)
	if nil != e {
		log.Fatal(e)
	}
	verify := sm2.Sm2Verify(publicKey, bs, nil, r, s)
	fmt.Println(verify)
}

func TestUnmarshalSM2Signature(t *testing.T) {
	pubStr := "BI8Ru4Dap5++jjuR2BSCvWDlZ9zgX92glUQ7Zpmow++lMUAdgFmlWyNvAsLqLTzb9NfOLohNhInQ0jWHB+lusbE="
	prvStr := `MIICBQIBADCB7AYHKoZIzj0CATCB4AIBATAsBgcqhkjOPQEBAiEA/////v////////////////////8AAAAA//////////8wRAQg/////v////////////////////8AAAAA//////////wEICjp+p6dn140TVqeS89lCafzl4n1FauPkt28vUFNlA6TBEEEMsSuLB8ZgRlfmQRGajnJlI/jC7/yZgvhcVpFiTNMdMe8Nzai9PZ3nFm9zuNraSFT0KmHfMYqR0AC3zLlITnwoAIhAP////7///////////////9yA99rIcYFK1O79Ak51UEjAgEBBIIBDzCCAQsCAQEEIPN0xmLKCNMJhpefHmjFHJH2zxjBOyJpa2CavH5A5l8moIHjMIHgAgEBMCwGByqGSM49AQECIQD////+/////////////////////wAAAAD//////////zBEBCD////+/////////////////////wAAAAD//////////AQgKOn6np2fXjRNWp5Lz2UJp/OXifUVq4+S3by9QU2UDpMEQQQyxK4sHxmBGV+ZBEZqOcmUj+MLv/JmC+FxWkWJM0x0x7w3NqL09necWb3O42tpIVPQqYd8xipHQALfMuUhOfCgAiEA/////v///////////////3ID32shxgUrU7v0CTnVQSMCAQE=`
	bytes, e := Base64Decode(prvStr)
	if nil != e {
		panic(e)
	}
	key, e := sm2.ParseSm2PrivateKey(bytes)
	if nil != e {
		log.Fatal(e)
	}
	pubKeyBytes, e := Base64Decode(pubStr)
	if nil != e {
		log.Fatal(e)
	}
	publicKey, e := sm2.ParseSm2PublicKey(pubKeyBytes)
	if nil != e {
		log.Fatal(e)
	}
	bs := []byte("123")
	r, s, e := sm2.Sm2Sign(key, bs, nil)
	verify := sm2.Verify(publicKey, bs, r, s)
	fmt.Println(verify)
}

func TestSM2ReadCertificate(t *testing.T) {
	str := "ME0CAQAwEwYHKoZIzj0CAQYIKoEcz1UBg3UEMzAxAgEBBCDzdMZiygjTCYaXnx5oxRyR9s8YwTsiaWtgmrx+QOZfJqAKBggqgRzPVQGDdQ=="
	bytes, _ := Base64Decode(str)
	key, e := sm2.ParsePKCS8UnecryptedPrivateKey(bytes)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(key)

	pubStr := "BI8Ru4Dap5++jjuR2BSCvWDlZ9zgX92glUQ7Zpmow++lMUAdgFmlWyNvAsLqLTzb9NfOLohNhInQ0jWHB+lusbE="
	pubByets, _ := Base64Decode(pubStr)
	publicKey, e := sm2.ParseSm2PublicKey(pubByets)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(publicKey)
}
