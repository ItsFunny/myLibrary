/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-07-05 13:02 
# @File : register.go
# @Description : 
# @Attention : 
*/
package model

import "myLibrary/go-library/common/blockchain/base"

type UserRegisterReq struct {
	Oid base.OrganizationID
	Name string
	Secret string
	Type string
}

type UserRegistrationResp struct {
	EnrollSecret string
}
