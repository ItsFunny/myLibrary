/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-08 13:59 
# @File : chaincode.go
# @Description : 
# @Attention : 
*/
package error

import error2 "myLibrary/go-library/common/error"

type ChainCodeError struct {
	*error2.BaseError
}

