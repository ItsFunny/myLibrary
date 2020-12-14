/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-14 16:25 
# @File : result.go
# @Description : 
# @Attention : 
*/
package dto

import "myLibrary/go-library/go/constants"

type ResultDTO struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func SuccessWithDetail(data interface{}, code int, msg string) ResultDTO {
	return ResultDTO{
		Data: data,
		Code: code,
		Msg:  msg,
	}
}

func Success(data interface{}) ResultDTO {
	return ResultDTO{
		Data: data,
		Code: constants.SUCCESS,
		Msg:  "success",
	}
}
