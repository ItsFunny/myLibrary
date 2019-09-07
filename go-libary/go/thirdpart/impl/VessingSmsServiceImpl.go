/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 13:23 
# @File : SmsServiceImpl.go
# @Description :  TODO 统一整理
*/
package impl

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"myLibrary/go-libary/go/base/services/impl"
	"myLibrary/go-libary/go/thirdpart/interfaces"
	"myLibrary/go-libary/go/thirdpart/models"
	"net/http"
)

type VessingSmsServiceImpl struct {
	baseImpl.BaseServiceImpl
}

func (this *VessingSmsServiceImpl) SendSMS(config interfaces.ISmsConfiguration, req *models.SmsSendReq) (models.SmsSendResp, error) {
	var (
		result models.SmsSendResp
	)
	this.BeforeStart("SendSMS")
	defer this.AfterEnd()

	keys, values := config.BuildValues(req)
	l := len(keys)
	args := fasthttp.Args{}
	for i := 0; i < l; i++ {
		args.Set(keys[i], values[i])
	}
	statusCode, body, err := fasthttp.Post(nil, config.GetSmsApiURL(), &args)
	if nil != err {
		return result, err
	} else if statusCode != http.StatusOK {
		this.GetLogger().Error("rest调用失败")
		return result, errors.New("未知错误")
	} else if nil != body && len(body) != 0 {
		var restResp models.SmsSendRestResp
		if err := json.Unmarshal(body, &restResp); nil != err {
			this.GetLogger().Error("反序列化失败:%s,原始数据为:[%s]", err.Error(), string(body))
			return result, err
		}
		if restResp.ReturnStatus != 1 {
			result.Msg = restResp.ReturnMessage
		} else {
			this.GetLogger().Info("成功向用户[%v]发送消息:[%v]", req.GetPhone(), req.GetMsg())
			result.Msg = "发送成功"
		}
	}

	return result, nil
}
