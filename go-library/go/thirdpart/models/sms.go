/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-09-03 13:54 
# @File : sms.go
# @Description : 
*/
package models

import (
	"myLibrary/go-library/go/thirdpart/interfaces"
	"strconv"
	"strings"
)

type SmsConfiguration struct {
	AppID     string `yaml:"AppID"`
	AppKey    string `yaml:"AppKey"`
	SmsApiURL string `yaml:"smsApiURI"`
}

func (this *SmsConfiguration) GetAppID() string {
	return this.AppID
}

func (this *SmsConfiguration) GetAppKey() string {
	return this.AppKey
}

func (this *SmsConfiguration) GetSmsApiURL() string {
	return this.SmsApiURL
}

func (this *SmsConfiguration) SetAppID(id string) interfaces.ISmsConfiguration {
	this.AppID = id
	return this
}

func (this *SmsConfiguration) SetAppKey(k string) interfaces.ISmsConfiguration {
	this.AppKey = k
	return this
}

func (this *SmsConfiguration) SetSmsApiURL(u string) interfaces.ISmsConfiguration {
	this.SmsApiURL = u
	return this
}
func (this *SmsConfiguration) BuildValues(req interfaces.ISmsSendReq) ([]string, []string) {
	keys, values := []string{"appId", "appKey", "templateId"}, []string{this.AppID, this.AppKey, strconv.Itoa(req.GetTemplateID())}
	keys = append(keys, "Variables", "phone")
	values = append(values, strings.Join(req.GetMsg(), ";"), req.GetPhone())
	return keys, values
}

// 发送短信的模板
type SmsSendReq struct {
	Phone string
	Msgs  []string
	// 模板ID
	TemplateID int
}

func (this *SmsSendReq) SetPhone(p string) interfaces.ISmsSendReq {
	this.Phone = p
	return this
}

func (this *SmsSendReq) GetPhone() string {
	return this.Phone
}

func (this *SmsSendReq) GetTemplateID() int {
	return this.TemplateID
}

func (this *SmsSendReq) SetTemplateID(id int) interfaces.ISmsSendReq {
	this.TemplateID = id
	return this
}

func (this *SmsSendReq) GetMsg() []string {
	return this.Msgs
}

func (this *SmsSendReq) SetMsg(m []string) interfaces.ISmsSendReq {
	this.Msgs = m
	return this
}

type SmsSendRestResp struct {
	// 返回状态
	ReturnStatus int `json:"returnStatus,string"`
	// 返回信息
	ReturnMessage string `json:"message"`
	// 返回余额
	RemainPoint int `json:"remainpoint,string"`
	// 任务ID
	TaskId string `json:"taskId"`
	// 成功条数
	SuccessCounts int `json:"successCounts,string"`
}

type SmsSendResp struct {
	Msg            string `json:"msg"`
}
