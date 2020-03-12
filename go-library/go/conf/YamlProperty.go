/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-06-26 13:25 
# @File : YamlProperty.go
# @Description : 
*/
package conf


type VessingSmsConfig struct {
	AppID  string `yaml:"AppID"`
	AppKey string `yaml:"AppKey"`
	SmsApiURI string `yaml:"smsApiURI"`
	TemplateID int `yaml:"templateID"`
}