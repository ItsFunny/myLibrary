/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-10-06 16:18 
# @File : base64utils.go
# @Description : 
# @Attention : 
*/
package utils

import "encoding/base64"

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
func Base64Decode(str string)([]byte,error){
	return base64.StdEncoding.DecodeString(str)
}
