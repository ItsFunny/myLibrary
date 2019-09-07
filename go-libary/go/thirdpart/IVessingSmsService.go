/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 13:18 
# @File : ISmsService.go
# @Description : https://www.veesing.com/ 短信接口
*/
package thirdpart

import (
	"myLibrary/go-libary/go/thirdpart/interfaces"
	"myLibrary/go-libary/go/thirdpart/models"
)

type (
	IVessingSmsService interface {
		SendSMS(config interfaces.ISmsConfiguration,req models.SmsSendReq) (models.SmsSendResp, error)
	}
)

