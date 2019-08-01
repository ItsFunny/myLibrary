/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-07-25 10:36 
# @File : sha.go
# @Description : 
*/
package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}
