/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-26 09:41
# @File : test_lock.go
# @Description :
# @Attention :
*/
package test_data

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

type A struct {
	lock sync.RWMutex
}

func (this *A) sss() {
	this.lock.Lock()
	defer this.lock.Unlock()
	log.Println(1)
	this.lock.Lock()
	log.Println(2)
	this.lock.Unlock()
}
func Test_aaa(t *testing.T) {
	a := &A{}
	a.sss()
}

type B struct {
	BB string
}
type C struct {
	*B
	Name string
}
type D struct {
	C
	Age int
}

func Test_ccc(t *testing.T) {
	d := &D{
		C: C{
			B:    &B{"bb"},
			Name: "ccName",
		},
		Age: 1,
	}
	v, _ := json.Marshal(d)
	fmt.Println(string(v))
}

func Test_d(t *testing.T) {
	fmt.Println(len("E3028811AD7CE856054783D9A4F4CAEADDC12E01"))
}

func TestB64(t *testing.T) {
	crt := `-----BEGIN CERTIFICATE-----
MIICEjCCAbigAwIBAgIRAIRAeoP2cWJFBwhPc8mZ6B0wCgYIKoEcz1UBg3UwYzEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xETAPBgNVBAoTCG9yZzAuY29tMRQwEgYDVQQDEwtjYS5vcmcwLmNv
bTAeFw0yMTAyMDcwMjA2MDBaFw0zMTAyMDUwMjA2MDBaMGMxCzAJBgNVBAYTAlVT
MRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMQ4w
DAYDVQQLEwVhZG1pbjEXMBUGA1UEAwwOQWRtaW5Ab3JnMC5jb20wWTATBgcqhkjO
PQIBBggqgRzPVQGCLQNCAARBAuu48MkoHb8F9MOFXO5yUwWPqBuDo1L6POoNND3h
F4+mkSMgERjUc86x5AQ/6b9w4DGBaYarSAr6qz7kAmU7o00wSzAOBgNVHQ8BAf8E
BAMCB4AwDAYDVR0TAQH/BAIwADArBgNVHSMEJDAigCBaHvBi25QAbVw009KSsVoy
blNj39a4Hx8PmnXY8+/mmjAKBggqgRzPVQGDdQNIADBFAiEA8WvTRlmHtAiAFys0
YEkYAUsOp9iS8KTUQ6f5MxIGVCACIEWQLv47rR9vf6cpM6uOIUFNTn8fW5F3T3ui
GROQmQpS
-----END CERTIFICATE-----`
	toString := base64.StdEncoding.EncodeToString([]byte(crt))
	fmt.Println(toString)

	prv := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgxmL0qAJbhH7UYgbJ
4swkwBDqHwqEM7UOfXBp9Y2kr9OgCgYIKoEcz1UBgi2hRANCAARBAuu48MkoHb8F
9MOFXO5yUwWPqBuDo1L6POoNND3hF4+mkSMgERjUc86x5AQ/6b9w4DGBaYarSAr6
qz7kAmU7
-----END PRIVATE KEY-----
`
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(prv)))
}

func Test_Sync(t *testing.T) {
	yamlStr := `version: 1.0.0

client:
  organization: Org1MSP
  logging:
    level: INFO  # 为了方便调试，正式环境应该改为info
  # 用户的身份也就是对应的私钥和签名证书 通过代码指定
  BCCSP:
    security:
      enabled: true
      default:
        provider: "GM"
      hashAlgorithm: "GMSM3"
      softVerify: true
      level: 256

organizations:
  Org0MSP:
    mspid: Org0MSP
    cryptoPath: crypto-config/peerOrganizations/org1.com/users/{username}@org0.com/msp
    peers:
      - peer0.org0.com
  Org1MSP:
    mspid: Org1MSP
    cryptoPath: crypto-config/peerOrganizations/org1.com/users/{username}@org1.com/msp
    peers:
      - peer0.org1.com
    certificateAuthorities:
      - ca.org1.com
  Org2MSP:
    mspid: Org2MSP
    cryptoPath: crypto-config/peerOrganizations/org1.com/users/{username}@org2.com/msp
    peers:
      - peer0.org2.com
#  Org3:
#    mspid: Org3MSP
#    cryptoPath: crypto-config/peerOrganizations/org1.com/users/{username}@org3.com/msp
#    peers:
#      - peer0.org3.com
channels:
  # 采用了服务发现，如果单纯想要配置简单，只需要配置一个节点，获取channel最新配置后自然可以获取到其他节点的证书，然后只需要在hosts文件中添加peer节点ip和容器名的映射即可。
  demochannel0:
    orderers:
      - orderer.demo.com
      - orderer1.demo.com
      - orderer2.demo.com
    peers:
      peer0.org0.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true
      peer0.org1.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true
      peer0.org2.com:
        endorsingPeer: true
        chaincodeQuery: true
        ledgerQuery: true
        eventSource: true
#      peer0.org3.com:
#        endorsingPeer: true
#        chaincodeQuery: true
#        ledgerQuery: true
#        eventSource: true
#      peer0.org4.com:
#        endorsingPeer: true
#        chaincodeQuery: true
#        ledgerQuery: true
#        eventSource: true
    policies:
      discovery:
        #[Optional] discovery info will be retrieved for these number of random targets
        maxTargets: 2
        #[Optional] retry options for retriving discovery info
        retryOpts:
          #[Optional] number of retry attempts
          attempts: 4
          #[Optional] the back off interval for the first retry attempt
          initialBackoff: 500ms
          #[Optional] the maximum back off interval for any retry attempt
          maxBackoff: 5s
          #[Optional] he factor by which the initial back off period is exponentially incremented
          backoffFactor: 2.0
      queryChannelConfig:
        minResponses: 1
        maxTargets: 1
        retryOpts:
          attempts: 5
          initialBackoff: 500ms
          maxBackoff: 5s
          backoffFactor: 2.0
orderers:
  orderer.demo.com:
    url: orderer.demo.com:5050
    grpcOptions:
      ssl-target-name-override: orderer.demo.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/ordererOrganizations/demo.com/tlsca/tlsca.demo.com-cert.pem

  orderer1.demo.com:
    url: orderer1.demo.com:5051
    grpcOptions:
      ssl-target-name-override: orderer1.demo.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/ordererOrganizations/demo.com/tlsca/tlsca.demo.com-cert.pem

  orderer2.demo.com:
    url: orderer2.demo.com:5052
    grpcOptions:
      ssl-target-name-override: orderer2.demo.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/ordererOrganizations/demo.com/tlsca/tlsca.demo.com-cert.pem

peers:
  peer0.org0.com:
    url: grpcs://peer0.org0.com:10051
    grpcOptions:
      ssl-target-name-override: peer0.org0.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/peerOrganizations/org0.com/tlsca/tlsca.org0.com-cert.pem
  peer0.org1.com:
    url: grpcs://peer0.org1.com:11051
    grpcOptions:
      ssl-target-name-override: peer0.org1.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/peerOrganizations/org1.com/tlsca/tlsca.org1.com-cert.pem
  peer0.org2.com:
    url: grpcs://peer0.org2.com:12051
    grpcOptions:
      ssl-target-name-override: peer0.org2.com
      allow-insecure: false
    tlsCACerts:
      path: ./crypto-config/peerOrganizations/org2.com/tlsca/tlsca.org2.com-cert.pem
#  peer0.org3.com:
#    url: grpcs://peer0.org3.com:13051
#    grpcOptions:
#      ssl-target-name-override: peer0.org3.com
#      allow-insecure: false
#    tlsCACerts:
#      path: ./crypto-config/peerOrganizations/org3.com/tlsca/tlsca.org3.com-cert.pem

certificateAuthorities:
  ca.org1.com:
    url: http://ca.org1.com:7054
    tlsCACerts:
      path: ./crypto-config/peerOrganizations/org1.com/ca/ca.org1.com-cert.pem
    registrar:
      enrollId: admin
      enrollSecret: adminpw
    caName: ca.org1.com

entityMatchers:
  peer:
    - pattern: (\w*)peer0.org0.com(\w*)
      urlSubstitutionExp: grpcs://peer0.org0.com:10051
      #      eventUrlSubstitutionExp: peer0.org1.com:11053
      #      sslTargetOverrideUrlSubstitutionExp: peer0.org1.com
      mappedHost: peer0.org0.com

    - pattern: (\w*)peer0.org1.com(\w*)
      urlSubstitutionExp: grpcs://peer0.org1.com:11051
      mappedHost: peer0.org1.com

    - pattern: (\w*)peer0.org2.com(\w*)
      urlSubstitutionExp: grpcs://peer0.org2.com:12051
      mappedHost: peer0.org2.com

    - pattern: (\w*)peer0.org3.com(\w*)
      urlSubstitutionExp: grpcs://peer0.org3.com:13051
      mappedHost: peer0.org3.com

    - pattern: (\w*)peer0.org4.com(\w*)
      urlSubstitutionExp: grpcs://peer0.org4.com:14051
      mappedHost: peer0.org4.com


  orderer:
    - pattern: (\w*)orderer.demo.com(\w*)
      urlSubstitutionExp: grpcs://orderer.demo.com:5050
      #sslTargetOverrideUrlSubstitutionExp: orderer.org1.com
      mappedHost: orderer.demo.com

    - pattern: (\w*)orderer1.demo.com(\w*)
      urlSubstitutionExp: grpcs://orderer2.demo.com:5051
      #sslTargetOverrideUrlSubstitutionExp: orderer.org1.com
      mappedHost: orderer1.demo.com

    - pattern: (\w*)orderer2.demo.com(\w*)
      urlSubstitutionExp: grpcs://orderer3.demo.com:5052
      #sslTargetOverrideUrlSubstitutionExp: orderer.org1.com
      mappedHost: orderer2.demo.com

#  certificateAuthorities:
#    - pattern: (\w*)ca.org1.com(\w*)
#      urlSubstitutionExp: peer0.org2.com:7054
#      mappedHost: ca.org1.com`
	base64Str := base64.StdEncoding.EncodeToString([]byte(yamlStr))
	fmt.Println(base64Str)
}

func TestSemaphore(t *testing.T) {
	// notify := make(chan struct{})
	// q := make([]chan int, 5)
	// consume := func() {
	// 	for {
	// 		select {
	// 		case r := <-q:
	// 			fmt.Println(r)
	// 		}
	// 	}
	// }
	// go func() {
	//
	// }()
}

func TestA(t *testing.T) {
	// StateChannel       = p2p.ChannelID(0x20) NewRoundStepMessage
	// DataChannel        = p2p.ChannelID(0x21)
	// VoteChannel        = p2p.ChannelID(0x22)
	// VoteSetBitsChannel = p2p.ChannelID(0x23)
	fmt.Println(0x20)
	fmt.Println(0x21)
	fmt.Println(0x22)
	fmt.Println(0x23)
}
func TestC(t *testing.T) {
	erc := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		close(erc)
	}()
	select {
	case <-erc:
		fmt.Println("close")
	}
}

// test selectN
// 通过反射的形式获取到selectN
func Test_ReflectSelectN(t *testing.T) {
	size := 20
	chans := make([]chan int, size)
	errChs := make([]chan error, size)
	for i := 0; i < size; i++ {
		chans[i] = make(chan int, size)
		errChs[i] = make(chan error, size)
	}

	go func() {
		count := 0
		ticker := time.NewTicker(time.Second * 3)
		for {
			select {
			case <-ticker.C:
				randIndex := rand.Intn(size)
				chans[randIndex] <- count
				fmt.Println("下标为", randIndex, "添加元素", count)
				count++
			}
		}
	}()
	go func() {
		count := 0
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				randIndex := rand.Intn(size)
				errChs[randIndex] <- errors.New("123")
				fmt.Println("下标为", randIndex, "添加error", count)
				count++
			}
		}
	}()
	go func() {
		cases := make([]reflect.SelectCase, len(chans))
		errCases := make([]reflect.SelectCase, len(chans))
		for i, ch := range chans {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
			errCases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(errChs[i])}
		}
		for {
			errIndex, errValue, ok := reflect.Select(errCases)
			if ok {
				fmt.Println("收到错误", errValue, errIndex)
			} else {
				fmt.Println("没有收到错误")
			}

			chosen, value, _ := reflect.Select(cases)
			// ok will be true if the channel has not been closed.
			// ch := chans[chosen]
			msg := value.Int()
			fmt.Println("accquire msg", msg, "index", chosen)
		}
	}()

	for {
		select {}
	}
}

func Test_SwithchSelectN(t *testing.T) {

}

type envelope struct {
	out chan int
	in  chan int
}

/*
	动态监听多个channel
	场景: 用于listener或者说是consumer端

*/

type ChannelWatcher struct {
}

// 通过长度自动分配selectN
func autoSelectN(t *testing.T, allChans []chan envelope) {
	l := len(allChans)
	i := 0
	for i < l {
		l = l - i

	}
}
func select4(t *testing.T, fourChans []chan int) {

}

func Test_HashValie(t *testing.T) {
	bytes := []byte("123333333")
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
}

func Test_Ch(t *testing.T) {
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		time.Sleep(time.Second * 2)
		close(ch)
	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println(v)
				} else {
					fmt.Println("close")
				}
			}
		}
	}()
	for {
		select {}
	}
}

func Test_For_select(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c <- 1
	}()
	cc, cancel := context.WithTimeout(context.Background(), time.Second*10)
LOOP:
	for {
		select {
		case v := <-c:
			cancel()
			fmt.Println("结果:", v)
			break LOOP
		case <-cc.Done():
			cancel()
			fmt.Println("触发了error:", cc.Err())
			return
		}
	}
	fmt.Println("end")
}

func Test_Parse(t *testing.T) {
	// time=1624520890,
	// v := time.Now().Unix()
	tt := time.Unix(1624524817, 0)

	fmt.Println(tt.String())
}

type IsItemExistsInCatalogResp struct {
	Ok      bool
	ErrDesc string
}

func Test_json(t *testing.T) {
	v := IsItemExistsInCatalogResp{}
	marshal, err := json.Marshal(v)
	fmt.Println(err)
	fmt.Println(string(marshal))
}
func Test_ctx(t *testing.T) {
	c, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()
	go func() {
		for {
			select {
			case <-c.Done():
				fmt.Println("done")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-c.Done():
				fmt.Println("done2")
				return
			}
		}
	}()
	for {
		select {}
	}
}

func Test_Break(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	close(c)
	// }()
wait:
	for {
		select {
		case <-c:
			break wait
		default:
			fmt.Println("wait")
		}
	}

	fmt.Println(1)
}

func Test_Change(t *testing.T) {
	type A struct {
		ch chan interface{}
		bh chan interface{}
	}
	a := &A{
		ch: make(chan interface{}),
		bh: make(chan interface{}),
	}
	notify := make(chan struct{})
	go func() {
		for {
			select {
			case v, ok := <-a.ch:
				if !ok {
					a.ch = nil
					continue
				}
				fmt.Println(v)
			case v := <-a.bh:
				fmt.Println("bh", v)
			case <-notify:
				fmt.Println("notify")
			}
		}
	}()
	newCh := make(chan interface{})
	go func() {
		for i := 0; i < 200; i++ {
			a.ch <- i
		}
		close(a.ch)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		a.ch = newCh
		notify <- struct{}{}
	}()
	go func() {
		for i := 1000; i < 1010; i++ {
			newCh <- i
		}
	}()
	for {
		select {}
	}
}

func Test_Change2(t *testing.T) {
	type A struct {
		ch <-chan interface{}
		bh <-chan interface{}
	}
	ch1 := make(chan interface{}, 2)
	ch2 := make(chan interface{}, 2)
	a := &A{
		ch: ch1,
		bh: ch2,
	}
	notify := make(chan struct{})
	go func() {
		for {
			select {
			case v, ok := <-a.ch:
				if !ok {
					a.ch = nil
					continue
				}
				fmt.Println(v)
			case v := <-a.bh:
				fmt.Println("bh", v)
			case <-notify:
				fmt.Println("notify")
			}
		}
	}()
	newCh := make(chan interface{})
	go func() {
		for i := 0; i < 200; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		a.ch = newCh
		notify <- struct{}{}
	}()
	go func() {
		for i := 1000; i < 1010; i++ {
			newCh <- i
		}
	}()
	r := sync.Mutex{}
	r.Lock()
	select {}
}

func Test_Use(t *testing.T) {
	c := make(chan int, 1)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		default:
			return
		}
	}
	fmt.Println("bread")
}

func Test_Bit(t *testing.T) {
	index := 3
	v := 0
	v += 1 << index
	fmt.Println(v)
	v += 1 << index
	fmt.Println(v)
}
func Test_Cccah(t *testing.T) {
	c := make(chan int, 5)
	go func() {
		c <- 1
	}()
	go func() {
		for nil != c {
			select {
			case vv, ok := <-c:
				if !ok {
					return
				}
				fmt.Println(vv, ok)
			}
		}
	}()
	select {}
}

func TestBit(t *testing.T) {
	shif := 1
	v := 1
	v += 1 << shif
	log.Println(v)
	log.Println(v - 1>>shif)
}
func TestRwLock(t *testing.T) {
	r := sync.RWMutex{}
	r.RLock()
}

func TestDDD(t *testing.T) {
	c1 := make(chan int, 1)
	go func() {
		select {
		case v := <-c1:
			fmt.Println(v)
		}
	}()
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- 1
	}()
	select {}
}

func TestForSelect(t *testing.T) {
	c := make(chan int)
	c2 := make(chan int)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case v2 := <-c2:
				fmt.Println(v2)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second)
			c <- 1
			c2 <- 2
		}
	}()
	select {}
}
