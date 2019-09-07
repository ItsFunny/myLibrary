/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-09-03 14:12 
# @File : sms.go
# @Description : 
*/
package interfaces

type ISmsConfiguration interface {
	GetAppID() string
	GetAppKey() string
	GetSmsApiURL() string

	SetAppID(id string) ISmsConfiguration
	SetAppKey(k string) ISmsConfiguration
	SetSmsApiURL(u string) ISmsConfiguration

	BuildValues(req ISmsSendReq) ([]string, []string)
}
type ISmsSendReq interface {
	GetTemplateID() int
	SetTemplateID(id int) ISmsSendReq
	SetPhone(p string) ISmsSendReq
	GetPhone() string
	GetMsg() []string
	SetMsg(m []string) ISmsSendReq
}
