/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-10 07:56
# @File : template.go
# @Description :
# @Attention :
*/
package base

import "fmt"

type Template interface {
	fmt.Stringer
	ValidIsMine(typeValue BaseType, excepted BaseType) bool
	LinkLast(template Template)
	GetNext() Template
	SetNext(template Template)
}

type IDetailTemplate interface {
	fmt.Stringer
	ValidIsMine(detailType DetailType) bool
	LinkLast(template IDetailTemplate)
	GetNext() IDetailTemplate
	SetNext(template IDetailTemplate)
}

type IHieraTemplateManager interface {
	LinkHiera(baseType BaseType, detailTemplate IDetailTemplate)
	GetHieraHandlerByBaseType(baseType, minCompareType BaseType) IHieraTemplate
}

type IHieraTemplate interface {
	Template
	GetDetailTemplate() IDetailTemplate
	SetDetailTemplate(detailTemplate IDetailTemplate)
}

type BaseTemplate struct {
	Template
	next     Template
	toString string
}

func (this *BaseTemplate) GetNext() Template {
	return this.next
}

func (this *BaseTemplate) SetNext(template Template) {
	this.next = template
}

func (this *BaseTemplate) LinkLast(template Template) {
	if nil == this.next {
		this.next = template
		return
	}
	tmp := this.next
	for ; nil != tmp.GetNext(); tmp = tmp.GetNext() {
	}
	tmp.SetNext(template)
}

type BaseDetailTemplate struct {
	IDetailTemplate
	next       IDetailTemplate
	toString   string
	DetailType DetailType
}

func (this *BaseDetailTemplate) String() string {
	return this.toString
}

func (this *BaseDetailTemplate) ValidIsMine(detailType DetailType) bool {
	return this.DetailType == detailType
}

func (this *BaseDetailTemplate) GetNext() IDetailTemplate {
	return this.next
}

func (this *BaseDetailTemplate) SetNext(template IDetailTemplate) {
	this.next = template
}

func (this *BaseDetailTemplate) LinkLast(template IDetailTemplate) {
	if nil == this.next {
		this.next = template
		return
	}
	tmp := this.next
	for ; nil != tmp.GetNext(); tmp = tmp.GetNext() {
	}
	tmp.SetNext(template)
}
