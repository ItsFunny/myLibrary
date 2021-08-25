package bccsp

import "github.com/hyperledger/fabric/base"

type DefaultMediator struct {
	*base.AlgorithmHieraManager
}

func NewDefaultMediator(baseType base.BaseType) *DefaultMediator {
	return &DefaultMediator{
		AlgorithmHieraManager: base.NewAlgorithmHieraManager(baseType),
	}
}

type SignatureMarshalMediator struct {
	*base.AlgorithmHieraManager
}
