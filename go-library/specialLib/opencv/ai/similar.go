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
	"github.com/corona10/goimagehash"
	"github.com/mfonda/simhash"
	"github.com/syyongx/php2go"
	"image/jpeg"
)

// 校验字符串的相似度,使用simhash进行判断
func CompareTextSimilarity(prev, newUpload string) (float64, int) {
	hash1 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(prev)))
	hash2 := simhash.Simhash(simhash.NewWordFeatureSet([]byte(newUpload)))
	compare := simhash.Compare(hash1, hash2)
	return 0.0, int(compare)
}

// 简单字符串匹配,适用于当长度小的情况
func SimpleCompareTextSimilarity(prev, newUpload string) (float64, int) {
	per := 0.0
	i := php2go.SimilarText(prev, newUpload, &per)
	return per, i
}

// Deprecated: Use GetImgHash insteaded
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
