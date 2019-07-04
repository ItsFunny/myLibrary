/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 13:18 
# @File : ISmsService.go
# @Description : https://www.veesing.com/ 短信接口
*/
package thirdpart

import (
	"github.com/akkagao/citizens/models"
	"github.com/akkagao/citizens/webbase"
)

type (
	IVessingSmsService interface {
		SendSMS(req models.SmsVerifySendModel) (models.SmsSendResp, error)
	}
)

func NewUserService(init webbase.IWebBaseServiceInit) ISmsService {
	return nil
}
