/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-19 13:34 
# @File : audio.go
# @Description : 
*/
package audio
//
// import (
// 	"errors"
// 	"fmt"
// 	"github.com/go-fingerprint/fingerprint"
// 	"github.com/go-fingerprint/gochroma"
// 	"github.com/mfonda/simhash"
// 	"io/ioutil"
// 	"math"
// 	"math/cmplx"
// 	"myLibrary/library/src/main/go/converters"
// 	. "myLibrary/library/src/main/go/utils"
// 	"os"
// 	"scientificgo.org/fft"
// 	"strconv"
// )
//
// const (
// 	UPPER_LIMIT = 300
// 	LOWER_LIMIT = 40
// 	FUZ_FACTOR  = 2
// )
//
// var (
// 	RANGE []int
// )
//
// func init() {
// 	RANGE = []int{40, 80, 0, 180, UPPER_LIMIT + 1}
// }
//
// // 将字节码转换为频域并且获取其hash值
// func ConvTAudio2FreQuence(data []byte) {
//
// 	totalSize := len(data)
// 	chunkSize := 4096
// 	amountPossible := totalSize / chunkSize
// 	cs := make([][]complex128, amountPossible)
// 	zero := float64(0.0)
// 	for times := 0; times < amountPossible; times++ {
// 		css := make([]complex128, chunkSize)
// 		for i := 0; i < 4096; i++ {
// 			b := int(data[times*chunkSize]) + i
// 			css[i] = complex(float64(b), zero)
// 		}
// 		cs[times] = fft.Fft(css, false)
// 	}
// 	calcFrequence(cs)
// }
//
// func calcFrequence(cs [][]complex128) {
// 	l := len(cs)
// 	highscores := make([][]float64, l)
// 	for i := 0; i < l; i++ {
// 		highscores[i] = make([]float64, 5)
// 		for j := 0; j < 5; j++ {
// 			highscores[i][j] = 0.0
// 		}
// 	}
// 	recordPoints := make([][]float64, l)
// 	for i := 0; i < l; i++ {
// 		recordPoints[i] = make([]float64, UPPER_LIMIT)
// 		for j := 0; j < UPPER_LIMIT; j++ {
// 			recordPoints[i][j] = 0.0
// 		}
// 	}
// 	points := make([][]int, l)
// 	for i := 0; i < l; i++ {
// 		points[i] = make([]int, 5)
// 		for j := 0; j < 5; j++ {
// 			points[i][j] = 0
// 		}
// 	}
// 	newFIle := "test2.txt"
// 	str := ""
// 	for t := 0; t < l; t++ {
// 		for freq := LOWER_LIMIT; freq < UPPER_LIMIT-1; freq++ {
// 			// Math.hypot(re, im)
// 			rf := real(cs[t][freq])
// 			ifloat := imag(cs[t][freq])
// 			mag := math.Log(math.Hypot(rf, ifloat))
// 			index := getIndex(freq)
// 			if mag > highscores[t][index] {
// 				highscores[t][index] = mag
// 				recordPoints[t][freq] = 1
// 				points[t][index] = freq
// 			}
// 		}
// 		for k := 0; k < 5; k++ {
// 			// strconv.FormatFloat(float64,'E',-1,64)
// 			str += "" + strconv.FormatFloat(highscores[t][k], 'E', -1, 64) + ";" +
// 				"" + strconv.FormatFloat(recordPoints[t][k], 'E', -1, 64) + "\t \n"
// 		}
// 		h := hash(points[t][0], points[t][1], points[t][2], points[t][3])
// 		// fmt.Println(h)
// 		if h != 8000004040 {
// 			fmt.Println(h)
// 		}
//
// 	}
// 	fmt.Println(len(str))
// 	err := ioutil.WriteFile(newFIle, []byte(str), 0777)
// 	if nil != err {
// 		panic(err)
// 	}
// }
//
// func getIndex(freq int) int {
// 	i := 0
// 	for RANGE[i] < freq {
// 		i++
// 	}
// 	return i
// }
// func hash(p1, p2, p3, p4 int) int {
// 	return (p4-(p4%FUZ_FACTOR))*100000000 + (p3-(p3%FUZ_FACTOR))*100000 + (p2-(p2%FUZ_FACTOR))*100 + (p1 - (p1 % FUZ_FACTOR))
// }
//
// func sqrt(x float64) string {
// 	if x < 0 {
// 		return fmt.Sprint(cmplx.Sqrt(complex(x, 0)))
// 	}
//
// 	return fmt.Sprint(math.Sqrt(x))
//
// }
//
// type AudioFingerPrintInfo struct {
// 	FingerPrintStr string
// 	SimilarCode    uint64
// 	SimilarCodes   []int32
// }
//
// func GetAudioFinger(filePath string) (AudioFingerPrintInfo, error) {
//
// 	var res AudioFingerPrintInfo
// 	if !IsFileOrDirExists(filePath) {
// 		return res, errors.New("文件不存在")
// 	}
// 	file, e := os.Open(filePath)
// 	if nil != e {
// 		return res, e
// 	}
// 	defer file.Close()
// 	info := fingerprint.RawInfo{
// 		Src:        file,
// 		Channels:   3,
// 		Rate:       44100,
// 		MaxSeconds: 300,
// 	}
// 	p2 := gochroma.New(gochroma.AlgorithmDefault)
// 	defer p2.Close()
// 	fprint, e := p2.RawFingerprint(info)
// 	if nil != e {
// 		return res, e
// 	}
// 	res.SimilarCodes = fprint
//
// 	bytes := make([]byte, 0)
// 	for _, i64 := range res.SimilarCodes {
// 		bytes = append(bytes, converter.LittelEndianInt642Bytes(int64(i64))...)
// 	}
// 	res.SimilarCode = simhash.Simhash(simhash.NewWordFeatureSet(bytes))
//
// 	return res, nil
// }
//
// func AudioInt32Compare(comp1 []int32, comp2 []int32) (float64, error) {
// 	l1 := len(comp1)
// 	l2 := len(comp2)
// 	if l1 < l2 {
// 		for i := 0; i < l2-l1; i++ {
// 			comp1 = append(comp1, 0)
// 		}
// 		// comp2 = comp2[:l1]
// 	} else {
// 		for i := 0; i < l1-l2; i++ {
// 			comp2 = append(comp2, 0)
// 		}
// 		// comp1 = comp1[:l2]
// 	}
// 	return fingerprint.Compare(comp1, comp2)
//
// }
//
// func Gochroma() {
// 	p := gochroma.New(1)
// 	filePath := "/Users/joker/Music/网易云音乐/QT - Heatrate.mp3"
// 	file, e := os.Open(filePath)
// 	if nil != e {
// 		panic(e)
// 	}
// 	// fprint, e := p.Fingerprint(fingerprint.RawInfo{
// 	// 	Src:        file,
// 	// 	Channels:   2,
// 	// 	Rate:       44100,
// 	// 	MaxSeconds: 3,
// 	// })
// 	// if nil != e {
// 	// 	panic(e)
// 	// }
// 	// fmt.Println(fprint)
// 	fprints, e := p.RawFingerprint(fingerprint.RawInfo{
// 		Src:        file,
// 		Channels:   2,
// 		Rate:       44100,
// 		MaxSeconds: 3,
// 	})
// 	if nil != e {
// 		panic(e)
// 	} else {
// 		fmt.Println(fprints)
// 	}
// }
