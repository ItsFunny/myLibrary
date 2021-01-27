/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-27 13:39 
# @File : http_test.go
# @Description : 
*/
package utils

import (
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/valyala/fasthttp"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"myLibrary/go-library/go/crypt"
	"net/http"
	"strconv"
	"testing"
)

func TestDoPostForm(t *testing.T) {
	resp, err := http.Get("https://www.baidu.com")
	if nil != err {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type T struct {
	Success    bool   `json:"success"`
	ResultCode string `json:"resultCode"`
	Msg        string `json:"msg"`
	Data       string `json:"data"`
	Salt       string `json:"salt"`
}

func TestDoPostForm2(t *testing.T) {
	SALT := "AVU0X77JO4AAB8GY3QXW"
	m := make(map[string]interface{})
	m["partnerID"] = "201908261512362679"
	m["partnerKey"] = "AVU0X77JO4AAB8GY3QXW"
	m["serialNo"] = "364129854703006208"
	m["salt"] = encrypt.MD5EncryptByBytes("201908261512362679AVU0X77JO4AAB8GY3QXW364129854703006208") + SALT
	args := ConvtMap2FastHttpArgs(m)
	statusCode, body, err := fasthttp.Post(nil, "https://ipp.tsa.cn/v2/api/confirm/downloadOpusCertificate", args)
	if nil != err {
		panic(err)
	}
	if statusCode != http.StatusOK {
		fmt.Println("失败")
	} else {
		var t T
		if er := json.Unmarshal(body, &t); nil != er {
			panic(er)
		}
		bytes, err := base64.StdEncoding.DecodeString(t.Data)
		if nil != err {
			panic(err)
		}
		if err := Write2File("asd.pdf", bytes); nil != err {
			panic(err)
		}

	}
}

type PeerOrganization struct {
	OrgMsp  string
	OrgName string
	OrgPort int
}

// 生成加密机
func TestGenerateEncryptMachine(t *testing.T) {
	reqs := make([]StaticDynamicHtmlReq, 0)
	reqs = append(reqs, NewEncryptoMachine("beijing", 7000),
		NewEncryptoMachine("bidsun", 7001),
		NewEncryptoMachine("dongguan", 7002),
		NewEncryptoMachine("foshan", 7003),
		NewEncryptoMachine("guangzhou", 7004),
		NewEncryptoMachine("guizhou", 7005),
		NewEncryptoMachine("haikou", 7006),
		NewEncryptoMachine("hangzhou", 7007),
		NewEncryptoMachine("lianzixin", 7008),
		NewEncryptoMachine("nanjing", 7009),
		NewEncryptoMachine("neimenggu", 7010),
		NewEncryptoMachine("nic", 7011),
		NewEncryptoMachine("qingdao", 7012),
		NewEncryptoMachine("tbi", 7013),
		NewEncryptoMachine("wuhan", 7014),
		NewEncryptoMachine("wuhangzhengshuju", 7015),
		NewEncryptoMachine("xiamen", 7016),
		NewEncryptoMachine("zhuhai", 7017), )
	for _, req := range reqs {
		html, e := StaticDynamicHtml(req)
		if nil != e {
			log.Fatal(e)
		}
		fmt.Println(html)
	}
}

func NewEncryptoMachine(serverName string, index int) StaticDynamicHtmlReq {
	return StaticDynamicHtmlReq{
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/production/config.json",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/production/",
		NewName:                 serverName,
		Suffix:                  ".json",
		Data: map[string]interface{}{
			"serverId": serverName + "-" + strconv.Itoa(index),
			"serverO":  serverName,
			"serverOu": serverName,},
	}
}

func TestStaticWithOrderer(t *testing.T) {
	reqs := make([]StaticDynamicHtmlReq, 0)
	reqs = append(reqs, NewWWStaticDynamicOrdererReq("docker-compose-orderer", "orderer", "Org1MSP", 7070, 11061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer1", "orderer1", "Org2MSP", 7051, 12061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer2", "orderer2", "Org3MSP", 7052, 13061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer3", "orderer3", "Org4MSP", 7053, 14061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer4", "orderer4", "Org5MSP", 7054, 17061))
	for _, req := range reqs {
		html, e := StaticDynamicHtml(req)
		if nil != e {
			log.Fatal(e)
		}
		fmt.Println(html)
	}
}
func NewWWStaticDynamicOrdererReq(newName, ordererName, orgMsp string, orgPort, chainCodeport int) StaticDynamicHtmlReq {
	return StaticDynamicHtmlReq{
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/docker-compose-orderer-template.yaml",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/",
		NewName:                 newName,
		Suffix:                  ".yaml",
		Data: map[string]interface{}{
			"ordererName": ordererName,
			"OrdererPort": orgPort,
		},
	}
}
func TestStaticDynamicHtmlWithBatch(t *testing.T) {
	reqs := make([]StaticDynamicHtmlReq, 0)
	reqs = append(reqs, NewWWStaticDynamicHtmlReq("docker-compose-org1", "org1", "Org1MSP", 11051, 11061),
		NewWWStaticDynamicHtmlReq("docker-compose-org2", "org2", "Org2MSP", 12051, 12061),
		NewWWStaticDynamicHtmlReq("docker-compose-org3", "org3", "Org3MSP", 13051, 13061),
		NewWWStaticDynamicHtmlReq("docker-compose-org4", "org4", "Org4MSP", 14051, 14061),
		NewWWStaticDynamicHtmlReq("docker-compose-org5", "org5", "Org5MSP", 17051, 17061),
		NewWWStaticDynamicHtmlReq("docker-compose-org6", "org6", "Org6MSP", 17051, 17061),
		NewWWStaticDynamicHtmlReq("docker-compose-org7", "org7", "Org7MSP", 17051, 17061), )
	for _, req := range reqs {
		html, e := StaticDynamicHtml(req)
		if nil != e {
			log.Fatal(e)
		}
		fmt.Println(html)
	}
}
func NewWWStaticDynamicHtmlReq(newName, orgName, orgMsp string, orgPort, chainCodeport int) StaticDynamicHtmlReq {
	return StaticDynamicHtmlReq{
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/docker-compose.yaml",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/",
		NewName:                 newName,
		Suffix:                  ".yaml",
		Data: map[string]interface{}{
			"OrgName":             orgName,
			"OrgMSP":              orgMsp,
			"OrgPort":             orgPort,
			"PEER_CHAINCODE_PORT": chainCodeport,
		},
	}
}
func TestStaticDynamicHtml2(t *testing.T) {
	req := StaticDynamicHtmlReq{
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/docker-compose.yaml",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/",
		NewName:                 "new",
		Suffix:                  ".yaml",
		Data: map[string]interface{}{
			"OrgName": "org1",
			"OrgMSP":  "Org1MSP",
			"OrgPort": 11051,},
	}
	html, e := StaticDynamicHtml(req)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(html)
}
func TestStaticDynamicHtml(t *testing.T) {
	req := StaticDynamicHtmlReq{
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/test.html",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/",
		NewName:                 "new",
		Data:                    map[string]interface{}{"name": "fffff"},
	}
	html, e := StaticDynamicHtml(req)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(html)
}


func Test_ACSD(t *testing.T) {
	// {{57897044551258231062740198220913455226441901632205615997740090104278067086466 0xc000021140 [536870905 268435455 895 268428288 536870911 268435455 536870911 170994943 268435455] [394377870 220399154 355969936 163370829 236861671 88177300 303341152 24396229 75627569] [137364797 52992271 113266657 202339045 31563580 107393171 24488059 247693942 35835723] [408558522 55895443 311818945 254526569 75270154 203012265 258167614 151236203 209300666]} 78062733970036975656793595378943889026827322737440942312410347708863577709861 25173883307730307121934226407653351122022746290457055470940048998556846361257}
	// {"ticket":"CmcKIQO438f2L6sEH0RYk3wj7hDeijIQGNWLdmGHXGyQHRJ29xIgiuqlRsPK8nJE4qTmnfITPLOGt7NEygIIUfvyDY8xmNIaIBTiBSFDy5Cwv+1+Kh/jAf3zcF5NarMYGWDUlMqJBshcEnAvvpHAXhiB5EYj0r1Owx9tg0CxEwUpRQW4lQu8IyG5dVmID/nMeo1/BmGQuuYKzvg3JvHfc0QNOGkM5v4vChs0OtQaiv6LDph5isuODJ9kgRA1wxAshrIeiMP1mAD/DC1RG7LcSZ4aDv3gkokpMtZ7","cryptKeyLabel":"AAHdB+LwMm5HhphSDT9vAwLoIoXGEhs0RICEBS7T16s52A==","licToken":"AAEihcYSGzREgIQFLtPXqznY6fpf4CfHQI+zeIomo08K2t0H4vAybkeGmFINP28DAujdB+LwMm5HhphSDT9vAwLo","data":["BOYxDp+rGVay1TFri4+eLomSH7x1NiWhVk14YQu4szNWBgKp5O+UPMNxZwBIkt8mU1smRHmiifIevJaFtyBJL1rglcB0mlJeiQaY2U8S7wX0qSFCeIXnEJdxepktAg3QwP9eFQ=="]}
	// hmac中的数据为:888881709839731917123
	// timeStamp:1709839731917
	// signDataHeader:JW/fZXhszpo374jhKlwMK7ETLD4qRiUiBfeDAuhtDXw=
	// {98 认证失败 }
	type A struct {
		Name string `json:"name"`
	}
	a := A{
		Name: "123",
	}
	key := "f86781b9ffbb8f1a84cf74da5f9fe425"
	timeStamp := "1709839731917"
	appId := "88888"
	bs, _ := json.Marshal(a)
	reqStr := string(bs)
	reqStr = `{"ticket":"CmcKIQO438f2L6sEH0RYk3wj7hDeijIQGNWLdmGHXGyQHRJ29xIgiuqlRsPK8nJE4qTmnfITPLOGt7NEygIIUfvyDY8xmNIaIBTiBSFDy5Cwv+1+Kh/jAf3zcF5NarMYGWDUlMqJBshcEnAvvpHAXhiB5EYj0r1Owx9tg0CxEwUpRQW4lQu8IyG5dVmID/nMeo1/BmGQuuYKzvg3JvHfc0QNOGkM5v4vChs0OtQaiv6LDph5isuODJ9kgRA1wxAshrIeiMP1mAD/DC1RG7LcSZ4aDv3gkokpMtZ7","cryptKeyLable":"AAHdB+LwMm5HhphSDT9vAwLoIoXGEhs0RICEBS7T16s52A==","licToken":"AAEihcYSGzREgIQFLtPXqznY6fpf4CfHQI+zeIomo08K2t0H4vAybkeGmFINP28DAujdB+LwMm5HhphSDT9vAwLo","data":["BLMoGY9RVAdKoKiq4bfAThFBOrElA7aNWWLhtg9RaetWFQo8tcwEq8fWFlAaiOCeIkfA/fW5l2aUZBERRs8mnXcn3Yc7X7DC9CZ1yzbCwH0qSH0hm4PGhfN3i+T+rdU+cMo6uw=="]}`
	res := BuildSignDataHeader(key, appId, timeStamp, reqStr)
	// Wd5uuj9UiOTkyiIvbVuneKeb+Kg8C4IYBzWicN82Flw=
	fmt.Println(res)
}

// Encrypt 为SM2加密函数:
// (1) 输入参数为: 公钥PB点(pub.X, pub.Y), 明文消息字节数组 in[], 密文类别标识 cipherTextType
// (2) 生成随机数k, k属于区间[1,N-1]
// (3) 利用标准包elliptic的方法CurveParams.ScalarBaseMult()生成倍点C1=kG=(c1x, c1y)
// (4) 由于SM2推荐曲线为素数域椭圆曲线，其余因子h=1，此时，点S=[h]PB就是公钥PB点，不可能为无穷远点O，
// 所以，国标4-6.1.A3被省略
// (5) 利用标准包elliptic的方法CurveParams.ScalarBaseMult()生成倍点kPB=(kPBx, kPBy)
// (6) 调用改进后的秘钥派生函数kdf(), 生成C2

type Sm2CipherTextType int32

const (
	// 旧标准的密文顺序
	C1C2C3 Sm2CipherTextType = 1
	// [GM/T 0009-2012]标准规定的顺序
	C1C3C2 Sm2CipherTextType = 2
)

func Encrypt(pub *sm2.PublicKey, in []byte, cipherTextType Sm2CipherTextType) ([]byte, error) {
	c2 := make([]byte, len(in))
	copy(c2, in)
	var c1 []byte
	digest := sm3.New()
	var kPBx, kPBy *big.Int
	for {
		// 利用标准库crypto/rand获取随机数k
		n, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C7052B53BBF40939D54123", 16)
		k, err := nextK(rand.Reader, n)
		if err != nil {
			return nil, err
		}
		kBytes := k.Bytes()
		// 利用标准库elliptic的方法CurveParams.ScalarBaseMult()计算倍点C1=kG=(c1x, c1y)
		c1x, c1y := pub.Curve.ScalarBaseMult(kBytes)

		// 将公钥曲线与C1点的坐标参数序列化。
		c1 = elliptic.Marshal(pub.Curve, c1x, c1y)

		// 利用标准库elliptic的方法CurveParams.ScalarMult()计算倍点kPB=(kPBx, kPBy)
		kPBx, kPBy = pub.Curve.ScalarMult(pub.X, pub.Y, kBytes)

		// 利用改造后的秘钥派生函数推算C2
		kdf(digest, kPBx, kPBy, c2)

		// 若中间变量t全部字节均为0则重启加密运算(详见国标4-6.1.A5)
		if !notEncrypted(c2, in) {
			break
		}
	}

	// 推算C3=Hash(kPBx || M || kPBy)，详见国标4-6.1.A7
	digest.Reset()
	digest.Write(kPBx.Bytes())
	digest.Write(in)
	digest.Write(kPBy.Bytes())
	c3 := digest.Sum(nil)

	// 根据密文格式标识的选择输出密文(C1C3C2新国准，或C1C2C3旧国标)
	c1Len := len(c1)
	c2Len := len(c2)
	c3Len := len(c3)
	result := make([]byte, c1Len+c2Len+c3Len)
	if cipherTextType == C1C2C3 {
		copy(result[:c1Len], c1)
		copy(result[c1Len:c1Len+c2Len], c2)
		copy(result[c1Len+c2Len:], c3)
	} else if cipherTextType == C1C3C2 {
		copy(result[:c1Len], c1)
		copy(result[c1Len:c1Len+c3Len], c3)
		copy(result[c1Len+c3Len:], c2)
	} else {
		return nil, errors.New("unknown cipherTextType:" + string(cipherTextType))
	}
	return result, nil
}

func kdf(digest hash.Hash, c1x *big.Int, c1y *big.Int, encData []byte) {
	bufSize := 4
	if bufSize < digest.Size() {
		bufSize = digest.Size()
	}
	buf := make([]byte, bufSize)

	encDataLen := len(encData)
	c1xBytes := c1x.Bytes()
	c1yBytes := c1y.Bytes()
	off := 0
	ct := uint32(0)
	for off < encDataLen {
		digest.Reset()
		digest.Write(c1xBytes)
		digest.Write(c1yBytes)
		ct++
		binary.BigEndian.PutUint32(buf, ct)
		digest.Write(buf[:4])
		tmp := digest.Sum(nil)
		copy(buf[:bufSize], tmp[:bufSize])

		xorLen := encDataLen - off
		if xorLen > digest.Size() {
			xorLen = digest.Size()
		}
		xor(encData[off:], buf, xorLen)
		off += xorLen
	}
}
func xor(data []byte, kdfOut []byte, dRemaining int) {
	for i := 0; i != dRemaining; i++ {
		data[i] ^= kdfOut[i]
	}
}
func notEncrypted(encData []byte, in []byte) bool {
	encDataLen := len(encData)
	for i := 0; i != encDataLen; i++ {
		if encData[i] != in[i] {
			return false
		}
	}
	return true
}

func nextK(rnd io.Reader, max *big.Int) (*big.Int, error) {
	intOne := new(big.Int).SetInt64(1)
	var k *big.Int
	var err error
	for {
		k, err = rand.Int(rnd, max)
		if err != nil {
			return nil, err
		}
		if k.Cmp(intOne) >= 0 {
			return k, err
		}
	}
}

func TestSM2Pub(t *testing.T) {
	prvStr := "Olafw9bAuWBj+qlwRCOeMSEcEp8kAB1MOovl1NaNFxA="
	pubStr := "JzowzBHdTXITJY4qNd70F7s05Tianox+/EaeoL9wI/gkrX+0pMoV/Y4Q2pVPYFtudiI1LXuei8rdraWd/PJNog=="
	bytes, _ := base64.StdEncoding.DecodeString(prvStr)

	var privateKey sm2.PrivateKey
	d := bytes
	privateKey.D = new(big.Int).SetBytes(d)
	curve := sm2.P256Sm2()
	privateKey.Curve = curve
	x, y := curve.ScalarBaseMult(d)
	privateKey.X, privateKey.Y = x, y
	// key, e := sm2.ParseSm2PrivateKey(bytes)
	// if nil != e {
	// 	log.Fatal(e)
	// }
	origin := []byte("123")
	r, s, e := sm2.Sm2Sign(&privateKey, origin, []byte("1234567812345678"))
	if nil != e {
		log.Fatal(e)
	}
	pub := parsePubKeyFromString(pubStr)
	verify := sm2.Sm2Verify(pub, origin, []byte("1234567812345678"), r, s)
	fmt.Println(verify)
	encrypt, e := sm2.Encrypt(pub, origin)
	if nil != e {
		log.Fatal(e)
	}
	decrypt, e := sm2.Decrypt(&privateKey, encrypt)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(string(decrypt))

}

func Test_AAAA(t *testing.T) {
	key := []byte("f86781b9ffbb8f1a84cf74da5f9fe425")
	h := hmac.New(sha256.New, key)
	h.Write([]byte("123"))
	sum := h.Sum(nil)
	fmt.Println(hex.EncodeToString(sum))
	fmt.Println(base64.StdEncoding.EncodeToString(sum))
}

type CC struct {
	VV string
}

func (*CC) String() string {
	return "ccc"
}

type B struct {
	Str  fmt.Stringer
	Name string
}

func Test_BBB(t *testing.T) {
	b := B{
		Str: &CC{
			VV: "cccccccc",
		},
		Name: "asd",
	}
	bytes, e := json.Marshal(b)
	if nil != e {
		log.Fatal(e)
	}
	fmt.Println(string(bytes))
}

func Test_ccc(t *testing.T) {
	str := `SGVsbG8gd29ybGQh`
	bytes, _ := Base64Decode(str)
	fmt.Println(string(bytes))
}
