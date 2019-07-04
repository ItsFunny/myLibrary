/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 13:23 
# @File : SmsServiceImpl.go
# @Description :  TODO 统一整理
*/
package impl
//
// import (
// 	"encoding/json"
// 	"github.com/akkagao/citizens/conf"
// 	"github.com/akkagao/citizens/models"
// 	"github.com/akkagao/citizens/models/dto"
// 	"github.com/akkagao/citizens/webbase/impl"
// 	"io/ioutil"
// 	utils2 "myLibrary/library/src/main/go/utils"
// )
//
// type VessingSmsServiceImpl struct {
// 	webImpl.WebBaseServiceImpl
// }
//
// func (this *VessingSmsServiceImpl) SendSMS(req models.SmsVerifySendModel) (models.SmsSendResp, error) {
// 	var (
// 		result models.SmsSendResp
// 	)
// 	this.BeforeStart("SendSMS")
// 	defer this.AfterEnd()
//
// 	property := conf.GetSmsProperty()
//
// 	// 验证码
// 	validateCode := utils2.GenValidateCode(6)
// 	result.ValidationCode=validateCode
// 	keys, values := property.BuildSmsVerifyValues([]string{validateCode})
// 	response, e := utils2.DoPostForm(property.SmsApiURI, keys, values)
// 	if nil != e {
// 		this.GetLogger().Error("[SendSMS] 发送的时候发生错误:%v", e.Error())
// 		return result, e
// 	}
// 	bytes, e := ioutil.ReadAll(response.Body)
// 	if nil != e {
// 		this.GetLogger().Error("[SendSMS]读取http response body 失败:%s", e.Error())
// 		return result, e
// 	}
// 	var apiResp models.SmsAPIResp
// 	if e := json.Unmarshal(bytes, &apiResp); nil != e {
// 		this.GetLogger().Error("[SendSMS] 反序列化body 失败:%s", e.Error())
// 		return result, e
// 	}
// 	if apiResp.ReturnStatus != 2000 {
// 		result.Status = dto.FAIL
// 		result.Msg = apiResp.GetResultInfo()
// 	} else {
// 		result.Status = dto.SUCCESS
// 	}
//
// 	return result,nil
// }
