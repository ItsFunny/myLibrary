/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-09-07 09:12 
# @File : fabric.go
# @Description : 
# @Attention : 
*/
package internal

import (
	"errors"
	"fmt"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protoutil"
	"myLibrary/go-library/blockchain/internal/genesisconfig"
	error2 "myLibrary/go-library/common/error"
)

func NewConsortiumOrgGroup(conf *genesisconfig.Organization) (*cb.ConfigGroup, error2.IBaseError) {
	consortiumsOrgGroup := NewConfigGroup()
	consortiumsOrgGroup.ModPolicy = channelconfig.AdminsPolicyKey

	if conf.SkipAsForeign {
		return consortiumsOrgGroup, nil
	}

	mspConfig, err := msp.GetVerifyingMspConfig(conf.MSPDir, conf.ID, conf.MSPType)
	if err != nil {
		return nil, error2.ErrorsWithMessage(err, fmt.Sprintf("1 - Error loading MSP configuration for org: %s", conf.Name))
	}

	if err := AddPolicies(consortiumsOrgGroup, conf.Policies, channelconfig.AdminsPolicyKey); err != nil {
		return nil, error2.ErrorsWithMessage(err, fmt.Sprintf("error adding policies to consortiums org group '%s'", conf.Name))
	}

	addValue(consortiumsOrgGroup, channelconfig.MSPValue(mspConfig), channelconfig.AdminsPolicyKey)

	return consortiumsOrgGroup, nil
}

func NewConfigGroup() *cb.ConfigGroup {
	return &cb.ConfigGroup{
		Groups:   make(map[string]*cb.ConfigGroup),
		Values:   make(map[string]*cb.ConfigValue),
		Policies: make(map[string]*cb.ConfigPolicy),
	}
}

const (
	// ConsensusTypeSolo identifies the solo consensus implementation.
	ConsensusTypeSolo = "solo"
	// ConsensusTypeKafka identifies the Kafka-based consensus implementation.
	ConsensusTypeKafka = "kafka"
	// ConsensusTypeKafka identifies the Kafka-based consensus implementation.
	ConsensusTypeEtcdRaft = "etcdraft"

	// BlockValidationPolicyKey TODO
	BlockValidationPolicyKey = "BlockValidation"

	// OrdererAdminsPolicy is the absolute path to the orderer admins policy
	OrdererAdminsPolicy = "/Channel/Orderer/Admins"

	// SignaturePolicyType is the 'Type' string for signature policies
	SignaturePolicyType = "Signature"

	// ImplicitMetaPolicyType is the 'Type' string for implicit meta policies
	ImplicitMetaPolicyType = "ImplicitMeta"
)

func AddPolicies(cg *cb.ConfigGroup, policyMap map[string]*genesisconfig.Policy, modPolicy string) error {
	switch {
	case policyMap == nil:
		return errors.New("no policies defined")
	case policyMap[channelconfig.AdminsPolicyKey] == nil:
		return errors.New("no Admins policy defined")
	case policyMap[channelconfig.ReadersPolicyKey] == nil:
		return errors.New("no Readers policy defined")
	case policyMap[channelconfig.WritersPolicyKey] == nil:
		return errors.New("no Writers policy defined")
	}

	for policyName, policy := range policyMap {
		switch policy.Type {
		case ImplicitMetaPolicyType:
			imp, err := policies.ImplicitMetaFromString(policy.Rule)
			if err != nil {
				return error2.ErrorsWithMessage(err, fmt.Sprintf("invalid implicit meta policy rule '%s'", policy.Rule))
			}
			cg.Policies[policyName] = &cb.ConfigPolicy{
				ModPolicy: modPolicy,
				Policy: &cb.Policy{
					Type:  int32(cb.Policy_IMPLICIT_META),
					Value: protoutil.MarshalOrPanic(imp),
				},
			}
		case SignaturePolicyType:
			sp, err := cauthdsl.FromString(policy.Rule)
			if err != nil {
				return error2.ErrorsWithMessage(err, fmt.Sprintf("invalid signature policy rule '%s'", policy.Rule))
			}
			cg.Policies[policyName] = &cb.ConfigPolicy{
				ModPolicy: modPolicy,
				Policy: &cb.Policy{
					Type:  int32(cb.Policy_SIGNATURE),
					Value: protoutil.MarshalOrPanic(sp),
				},
			}
		default:
			return error2.ErrorsWithMessage(nil,fmt.Sprintf("unknown policy type: %s", policy.Type))
		}
	}
	return nil
}

func addValue(cg *cb.ConfigGroup, value channelconfig.ConfigValue, modPolicy string) {
	cg.Values[value.Key()] = &cb.ConfigValue{
		Value:     protoutil.MarshalOrPanic(value.Value()),
		ModPolicy: modPolicy,
	}
}
