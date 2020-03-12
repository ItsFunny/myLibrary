/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-09-03 14:35 
# @File : sms_test.go
# @Description : 
*/
package impl

import (
	"fmt"
	"myLibrary/go-library/go/base/services/impl"
	"myLibrary/go-library/go/thirdpart/models"
	"testing"
)

func TestVessingSmsServiceImpl_SendSMS(t *testing.T) {
	config := new(models.SmsConfiguration)
	config.SetAppID("").SetAppKey("").SetSmsApiURL("")
	req := new(models.SmsSendReq)
	req.SetTemplateID(411).SetPhone("18757883747").SetMsg([]string{"12345s"})
	smsImpl := new(VessingSmsServiceImpl)

	smsImpl.BaseServiceImpl = *baseImpl.NewBaseServiceImpl()

	if resp, e := smsImpl.SendSMS(config, req); nil != e {
		panic(e)
	} else {
		fmt.Println(resp)
	}
}
