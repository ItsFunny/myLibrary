/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-30 11:11 
# @File : pic.go
# @Description : 
*/
package pic

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	"image/jpeg"
	"image/png"
	"myLibrary/go-libary/go/utils"
	"os"
	"path"
)

// batch比较图片的相似度
func BatchComparePicSimilarity(funName string, filePathSlice1 []string, filePathSlice2 []string, threshold int) (int, float64, error) {
	hash1 := make([]*goimagehash.ImageHash, 0)
	for _, p := range filePathSlice1 {
		imageHash, e := GetImgHash(funName, p)
		if nil != e {
			return 0, 0.0, e
		}
		hash1 = append(hash1, imageHash)
	}
	hash2 := make([]*goimagehash.ImageHash, 0)
	for _, p := range filePathSlice2 {
		imageHash, e := GetImgHash(funName, p)
		if nil != e {
			return 0, 0.0, e
		}
		hash2 = append(hash2, imageHash)
	}
	f, e := CompareImgHashes(hash1, hash2, threshold)
	return -1, f, e
}

// 比较单个图片的相似度
func CompareSingleImageSimilarity(hashFuncName string, filePath1, filePath2 string, threshold int) (int, float64, error) {
	return BatchComparePicSimilarity(hashFuncName, []string{filePath1}, []string{filePath2}, threshold)
}

func GetImgHash(funcName string, filePath string) (*goimagehash.ImageHash, error) {
	suffix := path.Ext(filePath)
	if exists := utils.IsFileOrDirExists(filePath); !exists {
		return nil, errors.New(fmt.Sprintf("[%s]文件不存在", filePath))
	}
	var img image.Image
	if suffix == ".jpeg" || suffix == ".jpg" {
		file, e := os.Open(filePath)
		if nil != e {
			return nil, e
		}
		image, e := jpeg.Decode(file)
		if nil != e {
			return nil, e
		}
		img = image
	} else if suffix == ".png" {
		file, e := os.Open(filePath)
		if nil != e {
			return nil, e
		}
		image, e := png.Decode(file)
		if nil != e {
			return nil, e
		}
		img = image
	} else {
		return nil, errors.New("图片格式错误,现暂仅支持jpeg,jpg,png结尾的图片")
	}

	imgHash := new(goimagehash.ImageHash)
	var e error
	switch funcName {
	case "a":
		imgHash, e = goimagehash.AverageHash(img)
	case "d":
		imgHash, e = goimagehash.DifferenceHash(img)
	}
	if nil != e {
		return nil, e
	}
	return imgHash, nil
}

// Deprecated: Use CompareImgHashes
func ComparePicSimilarity(hashFuncName string, prevBytes, newPicBytes []byte) (int, error) {
	var (
		prevHash *goimagehash.ImageHash
		newHash  *goimagehash.ImageHash
		e        error
		distance int
	)

	prev := bytes.NewReader(prevBytes)
	newPic := bytes.NewBuffer(newPicBytes)
	prevImg, e := jpeg.Decode(prev)
	if nil != e {
		return 0, e
	}
	newImg, e := jpeg.Decode(newPic)
	if nil != e {
		return 0, e
	}
	switch hashFuncName {
	case "a":
		prevHash, e = goimagehash.AverageHash(prevImg)
		newHash, e = goimagehash.AverageHash(newImg)
	}
	distance, e = prevHash.Distance(newHash)
	if nil != e {
		return 0, e
	}

	return distance, e
}

// srcHashes : 代表的原先存在着的hash
// newHashes: 代表的是新上传来匹配的
func CompareImgHashes(prevHashes, newHashes []*goimagehash.ImageHash, threshold int) (float64, error) {
	l1 := len(prevHashes)
	l2 := len(newHashes)

	count := 0
	for i := 0; i < l2; i++ {
		for j := 0; j < l1; j++ {
			distance, e := newHashes[i].Distance(prevHashes[j])
			if nil != e {
				return 0.0, e
			}
			if distance <= threshold {
				count++
				break
			}
		}
	}

	return float64(count) / float64(l2), nil
}

// func CompareImgHashes(hashes1, hashes2 []*goimagehash.ImageHash, threshold int) (float64, error) {
// 	l1 := len(hashes1)
// 	l2 := len(hashes2)
// 	count := 0
// 	less := l1
// 	lessesHashes := hashes1
// 	bigHashes := hashes2
// 	big := l2
// 	if l1 > l2 {
// 		less = l2
// 		lessesHashes = hashes2
// 		big = l1
// 		bigHashes = hashes1
// 	}
// 	for i := 0; i < less; i++ {
// 		for j := 0; j < big; j++ {
// 			distance, e := lessesHashes[i].Distance(bigHashes[j])
// 			if nil != e {
// 				return 0.0, e
// 			}
// 			if distance <= threshold {
// 				count++
// 				break
// 			}
// 		}
// 	}
//
// 	return float64(count) / float64(big), nil
// }
