/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-27 07:44 
# @File : error.go
# @Description : 
# @Attention : 
*/
package error

import (
	error2 "myLibrary/go-library/go/error"
)

type BlockChainError struct {
	*error2.BaseError
}

func NewBlockChainError(e error, code error2.ErrorCode, msg string) *BlockChainError {
	berr := new(BlockChainError)
	berr.BaseError = error2.NewBaseError(e, code, msg)
	return berr
}

func OrganizationNotExistError(e error, msg string) error2.IBaseError {
	return NewBlockChainError(e, RECORD_NOT_EXIST_ERROR, msg)
}

func UserRegistrationError(e error, msg string) error2.IBaseError {
	return NewBlockChainError(e, USER_REGISTRATION_ERROR, msg)
}

func NewJSONSerializeError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, JSON_SERIALIZE_ERROR_CODE, msg)
}
func NewArguError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, ARGUMENT_ERROR_CODE, msg)
}

func NewFabricError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, FABRIC_ERROR_CODE, msg)
}
func NewCryptError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, CRYPT_ERROR, msg)
}
func NewRecordNotExistError(msg string) error2.IBaseError {
	return error2.NewBaseError(nil, RECORD_NOT_EXIST_ERROR, msg)
}

func NewConfigError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, CONFIG_ERROR_CODE, msg)
}

func NewSystemError(e error, msg string) error2.IBaseError {
	return error2.NewBaseError(e, SYSTEM_ERROR_CODE, msg)
}
