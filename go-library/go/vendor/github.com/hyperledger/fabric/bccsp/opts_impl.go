/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-09 15:32
# @File : opts_impl.go
# @Description :
# @Attention :
*/
package bccsp

import (
	"crypto"
	"github.com/hyperledger/fabric/base"
)

type DefaultEcdsaOriginPublicKeyImportOptsAdapterImpl struct {
	base.EcdsaSameOpts
}
type DefaultSM2OriginPublicKeyImportOptsAdapterImpl struct {
	base.SM2SameOpts
}

type DefaultEcdsaPublicKeyImportOptsAdapterImpl struct {
	base.EcdsaSameOpts
}
type DefaultSM2PublicKeyImportOptsAdapterImpl struct {
	base.SM2SameOpts
}
type DefaultRSAPublicKeyImportOptsAdapterImpl struct {
	base.RSASameOpts
}

type DefaultSignerOptsAdapterImpl struct {
	base.CommonBaseOpts
	HashOpts  func() HashOptsAdapter
	Algorithm func() string
}

func (DefaultSignerOptsAdapterImpl) HashFunc() crypto.Hash {
	panic("implement me")
}

func NewDefaultSignerOptsAdapterImpl(baseType base.BaseType, detailType base.DetailType, getHashOpts func() HashOptsAdapter, algorithm func() string) *DefaultSignerOptsAdapterImpl {
	return &DefaultSignerOptsAdapterImpl{
		CommonBaseOpts: base.CommonBaseOpts{
			AlgorithmType: baseType,
			DetailType:    detailType,
			MinExcepted: func() base.BaseType {
				return baseType
			},
		},
		HashOpts:  getHashOpts,
		Algorithm: algorithm,
	}
}
func (c DefaultSignerOptsAdapterImpl) GetHashOpts() HashOptsAdapter {
	return c.HashOpts()
}

// verifier

type DefaultVerifierOptsAdapterImpl struct {
	base.CommonBaseOpts
	HashOpts func() HashOptsAdapter
}

func (this DefaultVerifierOptsAdapterImpl) GetHashOpts() HashOptsAdapter {
	return this.HashOpts()
}

// hash
type DefaultHashOptsAdapterImpl struct {
	base.CommonBaseOpts
	AlgorithmFunc func() string
}

func (d DefaultHashOptsAdapterImpl) Algorithm() string {
	return d.AlgorithmFunc()
}

// KeyGenOptsAdapter
type DefaultKeyGenOptsAdapterImpl struct {
	base.CommonBaseOpts
	AlgorithmFunc func() string
	EphemeralFunc func() bool
}

func (d DefaultKeyGenOptsAdapterImpl) Algorithm() string {
	return d.AlgorithmFunc()
}

func (d DefaultKeyGenOptsAdapterImpl) Ephemeral() bool {
	return d.EphemeralFunc()
}

// CertificateBaseOpts
type DefaultCertificateBaseOptsImpl struct {
	base.CommonBaseOpts
}
