/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-13 15:53 
# @File : similar_test.go
# @Description : 
*/
package ai

import (
	"fmt"
	"github.com/mfonda/simhash"
	"testing"
)

var (
	str1 = `中文 格式`
	str2 = `中文                                  格式时间戳系统需要遵循国际和国家标准有关对时间戳格式的规范要求。在时间戳系统在生成时间戳时，并不需要用户的原始信息（数据），而是只对用户的原始信息（数据）的某些关键特征（HASH值）进行时间戳签名，从而保证了用户原始信息（数据）的保密性和安全性。不同地区建立的时间戳，它们采用的数字签名证书是不同的，这主要用于区别不同的时间戳服务中心。时间戳服务中心用于时间戳签名的数字证书采用树形交叉认证体系，利用交叉认证技术实现不同时间戳服务中心所产生的时间戳文件的验证。我国的时间戳系统由北京联合信任技术有限公司与中国科学院国家授时中心联合建成，采用授时中心的时间源，并由时间守时系统与时间监控系统保证时间源准确性，从而保证时间戳系统产生时间戳的权威性。`
)

var (
	str3 = str1
)

func TestCompareTextSimilarity(t *testing.T) {
	hash1 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(str1)))
	hash2 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(str2)))
	fmt.Println(hash1)
	fmt.Println(hash2)
	compare := simhash.Compare(hash1, hash2)
	fmt.Println(compare)
	percent := calcPercent(float64(hash1), float64(hash2))
	fmt.Println(percent)
}
func TestSimpleCompareTextSimilarity(t *testing.T) {
	f, i := SimpleCompareTextSimilarity(str1, str2)
	fmt.Println(f)
	fmt.Println(i)
}
//
// func TestComparePicSimilarity(t *testing.T) {
// 	pic1 := "/Users/joker/Desktop/图片/小丑女.jpg"
// 	pic2 := "/Users/joker/Desktop/图片/刀塔.jpg"
// 	file1, _ := os.Open(pic1)
// 	file2, _ := os.Open(pic1)
// 	fil3, _ := os.Open(pic2)
// 	defer file1.Close()
// 	defer file2.Close()
// 	defer fil3.Close()
//
// 	bytes1, _ := ioutil.ReadAll(file1)
// 	bytes2, _ := ioutil.ReadAll(file2)
// 	bytes3, _ := ioutil.ReadAll(fil3)
// 	d1, _ := ComparePicSimilarity("a", bytes1, bytes2)
// 	fmt.Println(d1)
// 	d2, _ := ComparePicSimilarity("a", bytes1, bytes3)
// 	fmt.Println(d2)
// }
//
// func TestGetPicSimHash(t *testing.T) {
// 	filePath := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_2.png"
// 	// filePath := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_1.jpeg"
// 	bbs, e := ioutil.ReadFile(filePath)
// 	if nil != e {
// 		panic(e)
// 	}
// 	pic := bytes.NewBuffer(bbs)
// 	img, e := png.Decode(pic)
// 	imgHash, e := goimagehash.AverageHash(img)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(imgHash.GetHash())
//
// 	u, e := GetPicSimHash("a", bbs)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(u)
// }
//
// func TestGetImgHash(t *testing.T) {
// 	filePath := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_3.jpg"
// 	u, e := GetImgHash("a", filePath)
// 	if nil != e {
// 		panic(e)
// 	}
// 	filePath2 := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_4.jpg"
// 	hash2, e := GetImgHash("a", filePath2)
// 	if nil != e {
// 		panic(e)
// 	}
// 	compare := simhash.Compare(u.GetHash(), hash2.GetHash())
// 	fmt.Println(compare)
// 	// fmt.Println(u)
// }
// func TestCompareImageSimilarity(t *testing.T) {
// 	filePath1 := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_1.jpg"
// 	filePath2 := "/Users/joker/Desktop/图片/测试/图片相似度测试/sim_2.jpg"
// 	i, e := CompareImageSimilarity("a", filePath1, filePath2)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(i)
// }

func calcPercent(a, b float64) float64 {
	if a > b {
		return b / a
	} else {
		return a / b
	}
}
