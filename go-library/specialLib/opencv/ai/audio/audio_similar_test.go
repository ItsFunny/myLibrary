/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-19 15:08 
# @File : audio_similar_test.go
# @Description : 
*/
package audio
//
// import (
// 	"fmt"
// 	"github.com/mfonda/simhash"
// 	"io/ioutil"
// 	"testing"
// )
//
// func TestConvTAudio2FreQuence(t *testing.T) {
// 	filePath := "/Users/joker/Music/网易云音乐/QT - Heatrate.mp3"
// 	// filePath := "/Users/joker/Music/网易云音乐/Nicole Ferris - Just Don't Matter.mp3"
// 	bytes, e := ioutil.ReadFile(filePath)
// 	if nil != e {
// 		panic(e)
// 	}
// 	ConvTAudio2FreQuence(bytes)
//
// }
//
// func TestGochroma(t *testing.T) {
// 	filePath := "/Users/joker/Music/网易云音乐/QT - Heatrate.mp3"
// 	info, e := GetAudioFinger(filePath)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(info)
//
// 	filePath2 := "/Users/joker/Music/网易云音乐/Matt Cab - Akashi.mp3"
// 	info2, e := GetAudioFinger(filePath2)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(info2)
// 	compare := simhash.Compare(info.SimilarCode, info2.SimilarCode)
// 	fmt.Println(compare)
// }
//
// // similar test
// func TestGetAudioFinger(t *testing.T) {
// 	filePath1 := "/Users/joker/Music/网易云音乐/Charlie Puth,Selena Gomez - We Don't Talk Anymore.mp3"
// 	// filePath2 := "/Users/joker/Music/网易云音乐/Alex G,TJ Brown - We Don‘t Talk Anymore.mp3"
// 	filePath2 := "/Users/joker/Music/网易云音乐/Matt Cab - Akashi.mp3"
// 	// filePath2=filePath1
// 	info1, e := GetAudioFinger(filePath1)
// 	if nil != e {
// 		panic(e)
// 	}
// 	info2, e := GetAudioFinger(filePath2)
// 	if nil != e {
// 		panic(e)
// 	}
// 	// compare := simhash.Compare(info1.SimilarCode, info2.SimilarCode)
// 	// fmt.Println(compare)
//
// 	f, e := AudioInt32Compare(info1.SimilarCodes, info2.SimilarCodes)
// 	if nil != e {
// 		panic(e)
// 	}
// 	fmt.Println(f)
// 	fmt.Println(info1.SimilarCode)
// 	fmt.Println(info2.SimilarCode)
// 	compare := simhash.Compare(info1.SimilarCode, info2.SimilarCode)
// 	fmt.Println(compare)
//
// }
// func TestAudioInt32Compare(t *testing.T) {
// 	filePath1 := "/Users/joker/Music/网易云音乐/Charlie Puth,Selena Gomez - We Don't Talk Anymore.mp3"
// 	// filePath2 := "/Users/joker/Music/网易云音乐/Alex G,TJ Brown - We Don‘t Talk Anymore.mp3"
// 	filePath2 := "/Users/joker/Music/网易云音乐/Matt Cab - Akashi.mp3"
// 	bytes1, _ := ioutil.ReadFile(filePath1)
// 	bytes2, _ := ioutil.ReadFile(filePath2)
// 	s1 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(bytes1)))
// 	s2 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(bytes2)))
// 	fmt.Println(s1)
// 	fmt.Println(s2)
// 	compare := simhash.Compare(s1, s2)
// 	fmt.Println(compare)
// }
