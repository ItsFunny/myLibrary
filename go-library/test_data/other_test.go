/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-27 13:42
# @File : test_other.go
# @Description :
# @Attention :
*/
package test_data

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/ipfs/go-cid"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func Test_aa(t *testing.T) {
	// 2500000000
	time.Sleep(2500000000)
	fmt.Println(1)
}

func Test_bb(t *testing.T) {
	crt := `-----BEGIN CERTIFICATE-----
                    MIICATCCAaegAwIBAgIRAP+9miYL6YmQeGBoLNy+nQwwCgYIKoZIzj0EAwIwWjEL
                    MAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxEDAOBgNVBAcTB0JlaWppbmcx
                    ETAPBgNVBAoTCG9yZzEuY29tMRQwEgYDVQQDEwtjYS5vcmcxLmNvbTAeFw0yMTAy
                    MDYwMzE1MTRaFw0zMTAyMDQwMzE1MTRaMFsxCzAJBgNVBAYTAkNOMRAwDgYDVQQI
                    EwdCZWlqaW5nMRAwDgYDVQQHEwdCZWlqaW5nMQ8wDQYDVQQLEwZjbGllbnQxFzAV
                    BgNVBAMMDkFkbWluQG9yZzEuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
                    wSsfLbTOVLF411+ZUMlBjKzZFMJl61UWitWKNGBzmBa8tPS68NEuGmFqpHw0cyTG
                    g+KkoDkH1zTRltUNU26PB6NNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQC
                    MAAwKwYDVR0jBCQwIoAgqXVyknzT2tVmuvXrIccsN36SwtjbwQHVgsV7npVVE6Aw
                    CgYIKoZIzj0EAwIDSAAwRQIhAJ6ncxnEoNYeTyRHcuylL+HHwWk8lATUoKzZZUjX
                    da0kAiAT9KbBdz9f0WyvpjZcsQo0q9f5fcrHySCkICKMDSg6sw==
                    -----END CERTIFICATE-----`
	p, rest := pem.Decode([]byte(crt))
	fmt.Println(rest)
	fmt.Println(p)
}

func Test_AAccc(t *testing.T) {
	fmt.Println(string(filepath.ListSeparator))
}

func Test_Cid(t *testing.T) {
	bu := cid.V0Builder{}
	// bytes, _ := ioutil.ReadFile("/Users/joker/Downloads/a.log")

	bytes := []byte("zxckasddasasdkjasld")
	fmt.Println(base64.StdEncoding.EncodeToString(bytes))
	sum, _ := bu.Sum(bytes)
	// QmVVYgGt3RuZQY3zjVsf21QXGJh3S6mSLALTEWAqBTb2PC
	// QmQLqvuv726B6HaUd1Zk1UgY7Hz4CeMGtRMSeFJ9jckUiW
	encoding, _ := cid.ExtractEncoding(sum.String())
	fmt.Println(encoding)
	fmt.Println(sum.String())
}

func TestMetrics(t *testing.T) {
	t1 := time.Now().Unix()
	fmt.Println(t1)
	// 1621496118
	// 1621496082879
	time.Sleep(time.Microsecond + 12333)
	since := time.Since(time.Unix(t1, 0))
	fmt.Println(since.Seconds())
	fmt.Println(since.Milliseconds())
}

func Test_MultiCtx(t *testing.T) {
	c := context.Background()

	c1 := context.WithValue(c, "1", 1)
	c2 := context.WithValue(c1, "2", 2)

	fmt.Println(c2.Value("1"))
	fmt.Println(c2.Value("2"))
}

const (
	JAVA_TIME = 1000000000000
)

func Test_parse(t *testing.T) {
	v := int64(1622172147578)
	if v >= JAVA_TIME {
		v = v / 1000
	}
	sendT := time.Unix(v, 0)
	sendTStr := sendT.Format("2006-01-02 15:04:05")
	since := time.Since(sendT)
	seconds := since.Seconds()
	milliseconds := since.Milliseconds()
	fmt.Println("传递了sendTime",
		"发送时间", sendTStr,
		"当前时间", time.Now().String(),
		"时间间隔(s)", seconds, "时间间隔(毫秒)", milliseconds)
}

func Test_aaab(t *testing.T) {
	fmt.Println(time.Now().Unix())
}

type SS struct {
	incTime uint64
	seqNum  uint64
}

func TestBBB(t *testing.T) {
	ss := SS{
		incTime: uint64(time.Now().UnixNano()),
		seqNum:  uint64(0),
	}
	go func() {
		tt := time.NewTicker(time.Second)
		for {
			select {
			case <-tt.C:
				ss.seqNum++
				fmt.Println(uint64(ss.incTime), ss.seqNum)
			}
		}
	}()
	for {
		select {}
	}
}

func TestCCC(t *testing.T) {

	tm := uint64(1622341189577409323)
	toTime := tsToTime(tm)
	fmt.Println(toTime.String())

}
func tsToTime(ts uint64) time.Time {
	return time.Unix(int64(0), int64(ts))
}
func TestArr(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, 0, 1, 2, 3, 4, 5)
	arr = arr[:len(arr)-1]
	for _, v := range arr {
		fmt.Println(v)
	}
}

func TestPrint(t *testing.T) {
	fmt.Println(byte('1'))
}

type ITt interface {
	SetData(d interface{})
}
type AA struct {
	data interface{}
}
type ttImpl struct {
	a AA
}

func (this *ttImpl) SetData(d interface{}) {
	this.a.data = d
}
func Test_123333(t *testing.T) {
	ttt := &ttImpl{
		a: AA{data: 123},
	}
	fmt.Println(ttt.a)
	ttt.SetData(456)
	fmt.Println(ttt.a)

}

func Test_Strings(t *testing.T) {
	var sts []string
	v := strings.Join(sts, ",")
	fmt.Println(v)
}
func Test_asdd(t *testing.T) {
	str := "1997-04-02 10:07:18"
	parse, err := time.Parse("2006-01-02 03:04:05", str)
	if nil != err {
		panic(err)
	}
	fmt.Println(parse.Unix())
}
func Test_Unix(t *testing.T) {
	unix := int64(859975638)
	tt := time.Unix(unix, 0)
	fmt.Println(tt.String())
}

func Test_c(t *testing.T) {
	c := make(chan int, 1)
	go func() {
		for {
			select {
			case v, ok := <-c:
				if ok {
					print(v, 1)
				} else {
					fmt.Println("close")
					return
				}
			}
		}
	}()
	// time.Sleep(time.Second * 2)

	// close(c)
	for {
		select {}
	}
}
func Test_TimeSub(t *testing.T) {
	now := time.Now()

	time.Sleep(time.Second * 5)

	fmt.Println(time.Now().Sub(now).Seconds())
}

func Test_Goto(t *testing.T) {

	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
	goto ready2
ready2:
	for i := 100; i < 110; i++ {
		fmt.Println(i)
	}
	goto ready
ready:
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}

	fmt.Println("end")
}

func Test_CloseChan(t *testing.T) {
	cc := make(chan int, 1)
	cc <- 1
	close(cc)
	v := <-cc
	fmt.Println(v)
}
