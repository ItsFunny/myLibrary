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
	"encoding/asn1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
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
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
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
	m["serialNo"] = "364129854503006208"
	m["salt"] = encrypt.MD5EncryptByBytes("201908261512362679AVU0X77JO4AAB8GY3QXW364129854503006208") + SALT
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
	reqs = append(reqs, NewEncryptoMachine("beijing", 3000),
		NewEncryptoMachine("bidsun", 3001),
		NewEncryptoMachine("dongguan", 3002),
		NewEncryptoMachine("foshan", 3003),
		NewEncryptoMachine("guangzhou", 3004),
		NewEncryptoMachine("guizhou", 3005),
		NewEncryptoMachine("haikou", 3006),
		NewEncryptoMachine("hangzhou", 3007),
		NewEncryptoMachine("lianzixin", 3008),
		NewEncryptoMachine("nanjing", 3009),
		NewEncryptoMachine("neimenggu", 3010),
		NewEncryptoMachine("nic", 3011),
		NewEncryptoMachine("qingdao", 3012),
		NewEncryptoMachine("tbi", 3013),
		NewEncryptoMachine("wuhan", 3014),
		NewEncryptoMachine("wuhangzhengshuju", 3015),
		NewEncryptoMachine("xiamen", 3016),
		NewEncryptoMachine("zhuhai", 3017), )
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
		TemplateFilePath:        "/Users/joker/go/src/myLibrary/go-library/go/static/config.json",
		StaticHtmlBaseStorePath: "/Users/joker/go/src/myLibrary/go-library/go/static/",
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
	reqs = append(reqs, NewWWStaticDynamicOrdererReq("docker-compose-orderer", "orderer", "Org1MSP", 5050, 11061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer1", "orderer1", "Org2MSP", 5051, 12061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer2", "orderer2", "Org3MSP", 5052, 13061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer3", "orderer3", "Org4MSP", 5053, 14061),
		NewWWStaticDynamicOrdererReq("docker-compose-orderer4", "orderer4", "Org5MSP", 5054, 15061))
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
		NewWWStaticDynamicHtmlReq("docker-compose-org5", "org5", "Org5MSP", 15051, 15061),
		NewWWStaticDynamicHtmlReq("docker-compose-org6", "org6", "Org6MSP", 16051, 16061),
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

type BidsunJsonKeyInfo struct {
	ServerId         string `json:"serverId"`
	BidsunTicketUrl  string `json:"bidsunTicketUrl"`
	DoeDecUrl        string `json:"doeDecUrl"`
	CryptKeyLabel    string `json:"cryptKeyLabel"`
	CryptLicToken    string `json:"cryptLicToken"`
	CryptPubKey      string `json:"cryptPubKey"`
	ServerUserKey    string `json:"serverUserKey"`
	SignKeyLabel     string `json:"signKeyLabel"`
	SignPubKey       string `json:"signPubKey"`
	TicketPubKey     string `json:"ticketPublicKey"`
	TicketPrivateKey string `json:"ticketPrivateKey"`
}

type BidsunKeyInfoConfiguration struct {
	ServerId         string `json:"serverId"`
	BidsunTicketUrl  string `json:"bidsunTicketUrl"`
	DoeDecUrl        string `json:"doeDecUrl"`
	CryptKeyLabel    string `json:"cryptKeyLabel"`
	CryptLicToken    string `json:"cryptLicToken"`
	CryptPubKey      *sm2.PublicKey
	ServerUserPrvKey *sm2.PrivateKey
	SignKeyLabel     string `json:"signKeyLabel"`
	SignPubKey       *sm2.PublicKey
	TicketPubKey     string
	TicketPrvKey     *sm2.PrivateKey
}

type DoeDecRespDTO struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data string `json:"data"`
}

type DoeDecRequest struct {
	Ticket        string   `json:"ticket"`
	CryptKeyLabel string   `json:"cryptKeyLabel"`
	LicToken      string   `json:"licToken"`
	Data          []string `json:"data"`
}

var (
	DefaultBidsunHttpHandler func(response *http.Response) (result *HttpResult, b bool, e error) = func(response *http.Response) (result *HttpResult, b bool, e error) {
		dataBytes, e := ioutil.ReadAll(response.Body)
		if nil != e {
			return nil, true, errors.New("ioutil读取失败:" + e.Error())
		}
		if response.StatusCode != 200 {
			msg := "未知"
			if len(dataBytes) != 0 {
				msg = string(dataBytes)
			}
			return nil, true, errors.New("响应码状态不为200,状态信息为:" + response.Status + ",msg:" + msg)
		}
		errCode := response.Header.Get("errorCode")
		if len(errCode) != 0 && errCode != "0" {
			errorDesc := response.Header.Get("errorDesc")
			if len(errorDesc) == 0 {
				errorDesc = "未知错误"
			}
			var msg string
			s, e := url.QueryUnescape(errorDesc)
			// errorStrBytes, e := base64.URLEncoding.DecodeString(errorDesc)
			if nil != e {
				errorDesc = "url decode 未知错误:" + e.Error() + ",原先:" + errorDesc
			}
			msg = s
			return nil, false, errors.New("逻辑错误了,错误码为:" + errCode + ",错误信息为:" + msg)
		}
		var res TicketResp
		if e = json.Unmarshal(dataBytes, &res); nil != e {
			return nil, false, errors.New("json反序列化失败:" + e.Error())
		}
		return &HttpResult{
			Data: res,
			Code: 1,
			Msg:  "success",
		}, true, nil
	}
	DefaultDotHttpHandler = func(response *http.Response) (result *HttpResult, b bool, e error) {
		dataBytes, e := ioutil.ReadAll(response.Body)
		if nil != e {
			return nil, true, errors.New("ioutil读取失败:" + e.Error())
		}
		if response.StatusCode != 200 {
			msg := "未知"
			if len(dataBytes) != 0 {
				msg = string(dataBytes)
			}
			return nil, true, errors.New("响应码状态不为200,状态信息为:" + response.Status + ",msg:" + msg)
		}
		fmt.Println("doe返回值:" + string(dataBytes))
		var resp DoeDecRespDTO
		if e = json.Unmarshal(dataBytes, &resp); nil != e {
			return nil, false, errors.New("反序列化失败:" + e.Error())
		}
		fmt.Println(resp)
		if resp.Code != 0 {
			return nil, false, errors.New("调用深思失败:" + resp.Desc)
		}
		maps := make([]map[string]interface{}, 0)
		if e = json.Unmarshal([]byte(resp.Data), &maps); nil != e {
			return nil, false, errors.New("反序列化为map失败:" + e.Error())
		}
		return &HttpResult{
			Data: maps,
			Code: 1,
			Msg:  "success",
		}, true, nil
	}
)

var (
	keyConfiguration *BidsunKeyInfoConfiguration
)
// "doeDecUrl":"http://101.133.144.179:2159/doe/client/decData",
// "bidsunTicketUrl":"https://api.ebidsun.com/cloudshieldca/clientServerGenTicket",
var jsonT = `
		{
    "appId":"88888",
    "bidsunTicketUrl":"http://api.ebidsun.com/cloudshieldca/clientServerGenTicket",
    "cryptKeyLabel":"AAE40Hsy/eNMlZ4FbcK0L3R5FeDFxEXuRku3BD9w4qv2jg==",
    "cryptLicToken":"AAEV4MXERe5GS7cEP3Diq/aO9vUx93y8QzqZ/RPbVnS+sDjQezL940yVngVtwrQvdHk40Hsy/eNMlZ4FbcK0L3R5",
    "cryptPubKey":"ti2gksu8u0TZgPxZ1MgkR533zW4e+vA8pgJu/9BrCzSlI2wqRwtt11rZlJwQd1uZaMsDaQsrQONYf54jbNgFAw==",
    "doeDecUrl":"http://101.133.144.179:2159/doe/client/decData",
    "serverId":"bidsun-3001",
    "serverUserKey":"AAA40Hsy/eNMlZ4FbcK0L3R5APYsTzZuMK3ckxbHsAkpEXK+b/lCroomLRQ74Fh46si0",
    "signKeyLabel":"AAE40Hsy/eNMlZ4FbcK0L3R5hXfhu7VwSZG+xw/tMUy3+w==",
    "signLicToken":"AAGFd+G7tXBJkb7HD+0xTLf7ndCDjryTQnWuqCIZCNXTjzjQezL940yVngVtwrQvdHk40Hsy/eNMlZ4FbcK0L3R5",
    "signPubKey":"Z3raSFvhAsinYS+/vrN5L/JtKvMfBUxExTznB1MDuy1y2KRKRcpgT9P62O+5HxPnyu1x4J/qsqHWOQnKMNvRog==",
    "ticket":"CmcKIQKShRKPJ6hvdMR4fJuancArAZ1Mxz+JsRS4ctHFzNKTbBIgmf1uEdWOGseeuStBdG53GTNw20d9pt8lCpiif2doyT4aIDNlUcrvbNeUTvARv2VHXY4NNbk5c/OmkYiDZFUB5UFcEnBBlR8pG/b+yN/PXUzvsoywRrwvYy6UeHsVMyD+tZjHJ+7djDs4iWzVWJcd+qgHZ8+Vmyb6i72nxIgRLd45xXcQJHnUBdcmyg58s5M8IaPqskxoBUcXb0VDOq1H1pFUHgJXWWLjRM5ll6iY/Yvz8cQc",
    "ticketPrivateKey":"fjOgXrYnP1xkVYeCQrQyVseX6ZBbGBh/FI0Y7ezO1Ms=",
    "ticketPublicKey":"BL/xA/A5q0kK9eKR0oHplDCsPGpApa2UAbx1QTj5p8peuzBeBr9JUEmdKstmKDHlgvUFcCGZ2k5dWKqmlvr8xbs="
}
`

func init() {
	var keyInfo BidsunJsonKeyInfo
	if e := json.Unmarshal([]byte(jsonT), &keyInfo); nil != e {
		log.Fatal(e)
	}
	keyConfiguration = &BidsunKeyInfoConfiguration{
		ServerId:         keyInfo.ServerId,
		BidsunTicketUrl:  keyInfo.BidsunTicketUrl,
		CryptKeyLabel:    keyInfo.CryptKeyLabel,
		CryptLicToken:    keyInfo.CryptLicToken,
		CryptPubKey:      parsePubKeyFromString(keyInfo.CryptPubKey),
		ServerUserPrvKey: nil,
		// ServerUserPrvKey: parseSm2PrvKeyFromString(keyInfo.ServerUserKey),
		SignKeyLabel: keyInfo.SignKeyLabel,
		SignPubKey:   parsePubKeyFromString(keyInfo.SignPubKey),
		TicketPubKey: keyInfo.TicketPubKey,
		TicketPrvKey: parseSm2PrvKeyFromString(keyInfo.TicketPrivateKey),
		DoeDecUrl:    keyInfo.DoeDecUrl,
	}
}

type TicketReq struct {
	ServerId string `json:"serverId"`
}

type TicketResp struct {
	AppId     string `json:"appId"`
	Ticket    string `json:"ticket"`
	AppSecret string `json:"appSecret"`
}

func getTicket() (TicketResp, error) {
	defaultHttpReq := HttpReq{
		TimeOutSeconds:  10,
		MaxFailCount:    5,
		HttpUrl:         keyConfiguration.BidsunTicketUrl,
		BytesData:       nil,
		Headers:         nil,
		HandlerResponse: DefaultBidsunHttpHandler,
	}
	headers := make(map[string]string, 0)
	privateKey := keyConfiguration.TicketPrvKey
	originBytes := []byte(keyConfiguration.ServerId)
	userId := []byte("1234567812345678")
	r, s, err := sm2.Sm2Sign(privateKey, originBytes, userId)
	if nil != err {
		return TicketResp{}, err
	}
	signature := encodeSignatureRS(r, s)
	headers["signStr"] = string(signature)
	headers["pubKey"] = keyConfiguration.TicketPubKey
	headers["Content-Type"] = "application/json"
	defaultHttpReq.Headers = headers
	req := TicketReq{
		ServerId: keyConfiguration.ServerId,
	}
	reqBytes, _ := json.Marshal(req)
	defaultHttpReq.BytesData = reqBytes
	resString, err := PostPerClient(defaultHttpReq)
	if nil != err {
		panic("获取ticket失败:" + err.Error())
	}
	fmt.Println(resString)
	return resString.(TicketResp), err
}

func BuildSignDataHeader(appSecret, appId, timeStamp, originData string) string {
	sb := strings.Builder{}
	sb.WriteString(appId)
	sb.WriteString(timeStamp)
	sb.WriteString(originData)
	msg := sb.String()
	fmt.Println("hmac加密的原数据为:" + msg)
	fmt.Println("hex加密的原数据为:" + hex.EncodeToString([]byte(msg)))
	return base64.StdEncoding.EncodeToString(encodeHmacSHA256([]byte(appSecret), []byte(msg)))
}
func encodeHmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func Test_GetTicket(t *testing.T) {
	s, e := getTicket()
	fmt.Println(e)
	fmt.Println(s)
}

/*
*sm2密文转asn.1编码格式
*sm2密文结构如下:
*  x
*  y
*  hash
*  CipherText
 */
type sm2Cipher struct {
	XCoordinate *big.Int
	YCoordinate *big.Int
	HASH        []byte
	CipherText  []byte
}

func CipherMarshal(data []byte) ([]byte, error) {
	data = data[1:]
	x := new(big.Int).SetBytes(data[:32])
	y := new(big.Int).SetBytes(data[32:64])
	hash := data[64:96]
	cipherText := data[96:]
	return asn1.Marshal(sm2Cipher{x, y, hash, cipherText})
}

/*
sm2密文asn.1编码格式转C1|C3|C2拼接格式
*/
func CipherUnmarshal(data []byte) ([]byte, error) {
	var cipher sm2Cipher
	_, err := asn1.Unmarshal(data, &cipher)
	if err != nil {
		return nil, err
	}
	x := cipher.XCoordinate.Bytes()
	y := cipher.YCoordinate.Bytes()
	hash := cipher.HASH
	if err != nil {
		return nil, err
	}
	cipherText := cipher.CipherText
	if err != nil {
		return nil, err
	}
	c := []byte{}
	c = append(c, x...)          // x分量
	c = append(c, y...)          // y分
	c = append(c, hash...)       // x分量
	c = append(c, cipherText...) // y分
	return append([]byte{0x04}, c...), nil
}
func TestHttp(t *testing.T) {
	// decUrl := "http://101.133.144.179:2159/doe/client/decData"
	// decUrl := "http://doe-bx-client.senseyun.com/doe/client/decData"
	ticketResp, e := getTicket()
	if nil != e {
		panic("获取ticket失败:" + e.Error())
	}

	req := DoeDecRequest{
		Ticket:        ticketResp.Ticket,
		CryptKeyLabel: keyConfiguration.CryptKeyLabel,
		LicToken:      keyConfiguration.CryptLicToken,
	}
	req.Data = []string{"atQ4gL0zpgT4JGLfybmKjoUFhSnZday5LwnpDhXKMqux5Ld4uEpvJyPUcBQKuxwBnxojBR2a2JzGoqnq4FTLr1368SRFPNGzxhjLRM/0LBL4S/dh1Hy/QElPrL7xt8Hx8ctXgZJMUmtkjw=="}
	fmt.Println("加密数据为:" + req.Data[0])
	marshal, _ := json.Marshal(req)
	fmt.Println("请求参数数据为:" + string(marshal))
	originData := string(marshal)
	// originData="123"
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	headers["appId"] = ticketResp.AppId
	headers["timestamp"] = timeStamp
	signDataHeader := BuildSignDataHeader(ticketResp.AppSecret, ticketResp.AppId, timeStamp, originData)
	fmt.Println("timeStamp:" + timeStamp)
	fmt.Println("signDataHeader:" + signDataHeader)
	headers["signData"] = signDataHeader
	s, e := PostPerClient(HttpReq{
		TimeOutSeconds:  10,
		MaxFailCount:    5,
		HttpUrl:         keyConfiguration.DoeDecUrl,
		BytesData:       marshal,
		Headers:         headers,
		HandlerResponse: DefaultDotHttpHandler,
	})
	if nil != e {
		panic("调用深思失败:" + e.Error())
	}
	resp := s.([]map[string]interface{})
	if len(req.Data) != len(resp) {
		panic("返回值不一致,请求了:" + strconv.Itoa(len(req.Data)) + ",但是实际对象只返回了:" + strconv.Itoa(len(resp)))
	}
	for i := 0; i < len(req.Data); i++ {
		v := resp[i]
		if code, exist := v["code"]; !exist {
			panic("调用深思失败,没有code码返回")
		} else {
			switch code.(type) {
			case float64:
				codeInt := int(code.(float64))
				if codeInt != 0 {
					msg := ""
					desc := v["desc"]
					if desc != nil {
						msg = desc.(string)
					}
					panic("调用深思失败,对内部的数据解密失败:" + msg)
				}
			}
		}
		pureSec := v[req.Data[i]]
		if pureSec == nil {
			panic("返回的密钥为空")
		}
		bytes, e := base64.StdEncoding.DecodeString(pureSec.(string))
		if nil != e {
			panic("返回的密钥base64 解码失败:" + e.Error())
		}
		fmt.Println("密钥为:" + string(bytes))
	}
	fmt.Println(resp)
}

func TestHttp2(t *testing.T) {
	// decUrl := "http://101.133.144.179:2159/doe/client/decData"
	// decUrl := "http://doe-bx-client.senseyun.com/doe/client/decData"
	// ticketResp, e := getTicket()
	// if nil != e {
	// 	panic("获取ticket失败:" + e.Error())
	// }
	ticketResp := TicketResp{
		AppId:     "88888",
		Ticket:    "CmcKIQIQbWPMz6Ka9LjSPt5jR78kBbkf5O0oG3Zi3zJHsp4KnRIgcAFKyvyKgv7L7AhKSrnF0UFTiQvlVWvZ9VzsB6uwHAkaIPAaoAXZM0MrbAfFBgpJ+Bb1AoGXdpWRobn/vXp2z3V0EnA8znjnmeq8SaMiBNT4otViEiruIzi3uPq6jnnShHPHXX5sEdi2WuP5gqHW1ConNaUJgBX4fswY2BvjVewtYri2HWCTFdbTMO0pz01v3sR3/yy3yaDtLVW+HXzC6ZrpKlGlsAWEe4Smc/ZIVG0wqYQF",
		AppSecret: "f86781b9ffbb8f1a84cf74da5f9fe425",
	}
	origin := []byte("123")
	pub := keyConfiguration.CryptPubKey
	encBytes, e := sm2.Encrypt(pub, origin)
	if nil != e {
		log.Fatal(e)
	}

	req := DoeDecRequest{
		Ticket:        ticketResp.Ticket,
		CryptKeyLabel: keyConfiguration.CryptKeyLabel,
		LicToken:      keyConfiguration.CryptLicToken,
	}
	req.Data = []string{base64.StdEncoding.EncodeToString(encBytes)}
	// req.Data = []string{"k7nvMkQ1eyr+t3npg5pXl8Z0BiAb/y4TtXHd+ct7gLp/OR8n6c2I5yvaJHJ1bhG8N6rBB5RjzZc4FsHZQiaUB01HJK1mmk4qOnAHWAPEWZPnNg5qkBAgKkGNmBJk2fqMTDNQ4aUTDgAS3S9/"}
	marshal, _ := json.Marshal(req)
	fmt.Println("请求参数数据为:" + string(marshal))
	originData := string(marshal)
	// originData="123"
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	headers["appId"] = ticketResp.AppId
	headers["timestamp"] = timeStamp
	signDataHeader := BuildSignDataHeader(ticketResp.AppSecret, ticketResp.AppId, timeStamp, originData)
	fmt.Println("timeStamp:" + timeStamp)
	fmt.Println("signDataHeader:" + signDataHeader)
	headers["signData"] = signDataHeader
	s, e := PostPerClient(HttpReq{
		TimeOutSeconds:  10,
		MaxFailCount:    5,
		HttpUrl:         keyConfiguration.DoeDecUrl,
		BytesData:       marshal,
		Headers:         headers,
		HandlerResponse: DefaultDotHttpHandler,
	})
	if nil != e {
		panic("调用深思失败:" + e.Error())
	}
	resp := s.([]map[string]interface{})
	if len(req.Data) != len(resp) {
		panic("返回值不一致,请求了:" + strconv.Itoa(len(req.Data)) + ",但是实际对象只返回了:" + strconv.Itoa(len(resp)))
	}
	for i := 0; i < len(req.Data); i++ {
		v := resp[i]
		if code, exist := v["code"]; !exist {
			panic("调用深思失败,没有code码返回")
		} else {
			switch code.(type) {
			case float64:
				codeInt := int(code.(float64))
				if codeInt != 0 {
					msg := ""
					desc := v["desc"]
					if desc != nil {
						msg = desc.(string)
					}
					panic("调用深思失败,对内部的数据解密失败:" + msg)
				}
			}
		}
		pureSec := v[req.Data[i]]
		if pureSec == nil {
			panic("返回的密钥为空")
		}
		bytes, e := base64.StdEncoding.DecodeString(pureSec.(string))
		if nil != e {
			panic("返回的密钥base64 解码失败:" + e.Error())
		}
		fmt.Println("密钥为:" + string(bytes))
	}
	fmt.Println(resp)
}

// func TestHttp(t *testing.T) {
// 	// decUrl := "http://101.133.144.179:2159/doe/client/decData"
// 	decUrl := "http://doe-bx-client.senseyun.com/doe/client/decData"
// 	origin := []byte("123")
// 	pub := keyConfiguration.CryptPubKey
// 	encBytes, e := sm2.Encrypt(pub, origin)
// 	if nil != e {
// 		log.Fatal(e)
// 	}
// 	req := DoeDecRequest{
// 		CryptKeyLabel: keyConfiguration.CryptKeyLabel,
// 		LicToken:      keyConfiguration.CryptLicToken,
// 	}
// 	req.Data = []string{base64.StdEncoding.EncodeToString(encBytes)}
// 	marshal, _ := json.Marshal(req)
// 	fmt.Println("请求参数数据为:" + string(marshal))
// 	originData := string(marshal)
// 	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
// 	headers := make(map[string]string, 0)
// 	headers["Content-Type"] = "application/json"
// 	headers["appId"] = keyConfiguration.AppId
// 	headers["timestamp"], timeStamp)
// 	signDataHeader := BuildSignDataHeader(keyInfo.AppSecret, keyInfo.AppId, timeStamp, originData)
// 	fmt.Println("timeStamp:" + timeStamp)
// 	fmt.Println("signDataHeader:" + signDataHeader)
// 	request.Header.Set("signData", signDataHeader)
// 	httpResp, e := (&http.Client{}).Do(request)
// 	if nil != e {
// 		log.Fatal(e)
// 	}
// 	if httpResp.StatusCode != 200 {
// 		log.Fatal("请求失败")
// 	}
// 	respBytes, e := ioutil.ReadAll(httpResp.Body)
// 	if nil != e {
// 		log.Fatal("读取返回值失败")
// 	}
// 	var resp DoeDecResp
// 	if e = json.Unmarshal(respBytes, &resp); nil != e {
// 		log.Fatal("反序列化返回值失败")
// 	}
// 	fmt.Println(resp)
// }

func parsePubKeyFromString(b64Str string) *sm2.PublicKey {
	pubBytes, e := base64.StdEncoding.DecodeString(b64Str)
	if nil != e {
		log.Fatal(e)
	}
	if len(pubBytes) == 65 {
		pubBytes = pubBytes[1:]
	}
	x := pubBytes[:32]
	y := pubBytes[32:64]
	fmt.Println("x:" + base64.StdEncoding.EncodeToString(x))
	fmt.Println("y:" + base64.StdEncoding.EncodeToString(y))
	pubX := new(big.Int).SetBytes(x)
	pubY := new(big.Int).SetBytes(y)
	pub := sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     pubX,
		Y:     pubY,
	}
	return &pub
}
func parseSm2PrvKeyFromString(b64String string) *sm2.PrivateKey {
	decodeString, e := base64.StdEncoding.DecodeString(b64String)
	if nil != e {
		panic("解析私钥失败")
	}
	if len(decodeString) == 32 {
		var (
			raw []byte
		)
		if strings.Contains(string(decodeString), "BEGIN") {
			var block *pem.Block
			block, _ = pem.Decode(decodeString)
			if block == nil {
				panic("failed to decode private key")
			}
			raw = block.Bytes
		}
		raw = decodeString
		if len(raw) != 32 {
			panic("标准私钥为32个字节")
		}
		var privateKey sm2.PrivateKey
		d := raw
		privateKey.D = new(big.Int).SetBytes(d)
		curve := sm2.P256Sm2()
		privateKey.Curve = curve
		x, y := curve.ScalarBaseMult(d)
		privateKey.X, privateKey.Y = x, y
		return &privateKey
	}
	key, e := sm2.ParseSm2PrivateKey(decodeString)
	if nil != e {
		panic("ParseSm2PrivateKey#解析私钥失败")
	}
	return key
}

// 将两个大整数拼接成字符串
func encodeSignatureRS(r, s *big.Int) []byte {
	// 缺位补足
	r1 := r.Text(16)
	if len(r1) < 64 {
		for i := len(r1); i < 64; i++ {
			r1 = "0" + r1
		}
	}

	// 缺位补足
	s1 := s.Text(16)
	if len(s1) < 64 {
		for i := len(s1); i < 64; i++ {
			s1 = "0" + s1
		}
	}

	return []byte(fmt.Sprintf("%s%s", r1, s1))
}

func Test_ACSD(t *testing.T) {
	// {{57896044551258231062740198220913455226441901632205615997740090104278065086466 0xc000021140 [536870905 268435455 895 268428288 536870911 268435455 536870911 150994943 268435455] [394377860 220399154 355969936 163370829 236861671 88177300 303341152 24396229 75627569] [137364797 52992271 113266657 202339045 31563580 107393171 24488059 247693942 35835723] [408558522 55895443 311818945 254526569 75260154 203012265 258167614 151236203 209300666]} 78062733960036975656793595378943889026827322737440942312410347508863577709861 25173883307730307121934226407653351122022746290457055470940048998556846361257}
	// {"ticket":"CmcKIQO438f2L6sEH0RYk3wj7hDeijIQGNWLdmGHXGyQHRJ29xIgiuqlRsPK8nJE4qTmnfITPLOGt7NEygIIUfvyDY8xmNIaIBTiBSFDy5Cwv+1+Kh/jAf3zcF5NarMYGWDUlMqJBshcEnAvvpHAXhiB5EYj0r1Owx9tg0CxEwUpRQW4lQu8IyG5dVmID/nMeo1/BmGQuuYKzvg3JvHfc0QNOGkM5v4vChs0OtQaiv6LDph5isuODJ9kgRA1wxAshrIeiMP1mAD/DC1RG7LcSZ4aDv3gkokpMtZ7","cryptKeyLabel":"AAHdB+LwMm5HhphSDT9vAwLoIoXGEhs0RICEBS7T16s52A==","licToken":"AAEihcYSGzREgIQFLtPXqznY6fpf4CfHQI+zeIomo08K2t0H4vAybkeGmFINP28DAujdB+LwMm5HhphSDT9vAwLo","data":["BOYxDp+rGVay1TFri4+eLomSH7x1NiWhVk14YQu4szNWBgKp5O+UPMNxZwBIkt8mU1smRHmiifIevJaFtyBJL1rglcB0mlJeiQaY2U8S7wX0qSFCeIXnEJdxepktAg3QwP9eFQ=="]}
	// hmac中的数据为:888881609839731917123
	// timeStamp:1609839731917
	// signDataHeader:JW/fZXhszpo374jhKlwMK7ETLD4qRiUiBfeDAuhtDXw=
	// {98 认证失败 }
	type A struct {
		Name string `json:"name"`
	}
	a := A{
		Name: "123",
	}
	key := "f86781b9ffbb8f1a84cf74da5f9fe425"
	timeStamp := "1609839731917"
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
		n, _ := new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
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
	pubStr := "JzowzBHdTXITJY4qNd60F7s05Tianox+/EaeoL9wI/gkrX+0pMoV/Y4Q2pVPYFtudiI1LXuei8rdraWd/PJNog=="
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
