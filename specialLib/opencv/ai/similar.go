/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-08-13 15:15 
# @File : similar.go
# @Description : 
*/
package ai

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/syyongx/php2go"
	"image"
	"image/jpeg"
	"image/png"
	"myLibrary/library/src/main/go/utils"
	"os"
	"path"
)

// 校验字符串的相似度,使用simhash进行判断
func CompareTextSimilarity(prev, newUpload string) (float64, int) {
	// hash1 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(prev)))
	// 	// hash2:=simhash.Simhash(simhash.NewWordFeatureSet([]byte(newUpload)))
	// 	//
	// 	//
	// 	// per := 0.0
	// 	// text := php2go.SimilarText(prev, newUpload, &per)
	// 	// return per, text
	return 0.0, 1
}

// 简单字符串匹配,适用于当长度小的情况
func SimpleCompareTextSimilarity(prev, newUpload string) (float64, int) {
	per := 0.0
	i := php2go.SimilarText(prev, newUpload, &per)
	return per, i
}

func CompareImageSimilarity(hashFuncName string, filePath1, filePath2 string) (int, error) {
	imageHash, e := GetImgHash(hashFuncName, filePath1)
	if nil != e {
		panic(e)
	}
	hash, e := GetImgHash(hashFuncName, filePath2)
	if nil != e {
		panic(e)
	}
	return imageHash.Distance(hash)
}

// 图片相似度匹配
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
		return nil, errors.New("图片格式错误,现暂仅支持jpeg,png结尾的图片")
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

func GetPicSimHash(hashFuncName string, bbs []byte) (uint64, error) {
	imgHash := new(goimagehash.ImageHash)
	pic := bytes.NewBuffer(bbs)
	img, e := jpeg.Decode(pic)
	if nil != e {
		return 0, e
	}
	switch hashFuncName {
	case "a":
		imgHash, e = goimagehash.AverageHash(img)
	}

	return imgHash.GetHash(), nil
}

func CompareSimilarityByHash(prevHash int, newHash int) {
	// simhash.Simhash()
}
