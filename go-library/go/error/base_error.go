/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-14 13:44 
# @File : base_error.go
# @Description : 
# @Attention : 
*/
package error

import (
	constants2 "myLibrary/go-library/go/constants"
)

type BaseError struct {
	Cause error
	Code  ErrorCode
	Msg   string
}

func NewBaseError(ee error, c ErrorCode, msg string) *BaseError {
	e := new(BaseError)
	e.Code = c
	e.Msg = msg
	e.Cause = ee

	return e
}
func NewJSONSerializeError(e error, msg string) *BaseError {
	return NewBaseError(e, constants2.JSON_SERIALIZE_ERROR_CODE, msg)
}
func NewArguError(e error, msg string) IBaseError {
		return  NewBaseError(e, constants2.ARGUMENT_ERROR_CODE, msg)
}

func NewFabricError(e error, msg string) IBaseError {
	return  NewBaseError(e, constants2.FABRIC_ERROR_CODE, msg)
}
func NewRecordNotExistError(msg string) IBaseError {
		return NewBaseError(nil, constants2.RECORD_NOT_EXIST_ERROR, msg)
}
func (e *BaseError) Error() string {
	if nil != e.Cause {
		return e.Msg + ":" + e.Cause.Error()
	}

	return e.Msg
}

func (e *BaseError) GetCode() ErrorCode {
	return e.Code
}

func (e *BaseError) SetCode(code ErrorCode) {
	e.Code = code
}

func (e *BaseError) GetMsg() string {
	return e.Msg
}

func (e *BaseError) SetMsg(m string) {
	e.Msg = m
}
