/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 14:14 
# @File : factory.go
# @Description : 
# @Attention : 
*/
package error

import error2 "myLibrary/go-library/common/error"

func NewChainCodeError(e error, msg string)error2.IBaseError {
	return error2.NewBaseError(e, CHAINCODE_ERROR, msg)
}
