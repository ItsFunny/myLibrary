/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-12-11 22:45 
# @File : intergration.go
# @Description :    汇总了hash,signer,verifer
# @Attention : 
*/
package bccsp

import (
	"myLibrary/go-library/go/internal/bccsp/base"
	"myLibrary/go-library/go/internal/bccsp/opts"
)
type BaseIntegrationOpts struct {
	// Ephemeral() bool
	EphemeralFunc func()bool
}
func(this BaseIntegrationOpts)Ephemeral()bool{
	return this.EphemeralFunc()
}
type DefaultECDSAIntegrationOpts struct {
	BaseIntegrationOpts
	base.ECDSASignerOpts
}

func ( DefaultECDSAIntegrationOpts) GetHashOpts() HashOptsAdapter {
	return opts.DefaultSHA256HashOptsImpl{}
}

func ( DefaultECDSAIntegrationOpts) Algorithm() string {
	return "ECDSA"
}

type DefaultSM2IntegrationOpts struct {
	BaseIntegrationOpts
	base.SM2SignerOpts
	// SM2VerifyOpts
}

func ( DefaultSM2IntegrationOpts) GetHashOpts() HashOptsAdapter {
	return opts.DefaultSM3HashOptsImpl{}
}

func ( DefaultSM2IntegrationOpts) Algorithm() string {
	return "SM2"
}

type DefaultRSAIntergrationOpts struct {
	BaseIntegrationOpts
	base.RSASignerOpts
	// RSAVerifyOpts
}

func ( DefaultRSAIntergrationOpts) GetHashOpts() HashOptsAdapter {
	panic("implement me")
}

func ( DefaultRSAIntergrationOpts) Algorithm() string {
	return "RSA"
}



