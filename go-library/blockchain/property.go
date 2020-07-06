/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-15 14:24 
# @File : property.go
# @Description : 
# @Attention : 
*/
package config

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"myLibrary/go-library/blockchain/base"
)

// cc的配置
type ChainCodeProperties struct {
	ChainCodeID   base.ChainCodeID `yaml:"chainCodeId" json:"chainCodeId"`
	ChainCodePath string           `yaml:"chainCodePath" json:"chainCodePath"`
	NeedUpdate    bool             `yaml:"needUpdate" json:"needUpdate"`
	// 2020-06-16 是否是需要监听block事件的chaincode
	NeedListOnBlockEvent bool `yaml:"needListOnBlockEvent" json:"needListOnBlockEvent"`
}

type Peer struct {
	Address string `yaml:"address" json:"address"`
	Port    int    `yaml:"port" json:"port"`
}

type AnchorPeer struct {
	Address string `yaml:"address" json:"address"`
	Port    int    `yaml:"port" json:"port"`
	// 这个peer上要挂载哪些cc
	ChainCodes []ChainCodeProperties `yaml:"chainCodes" json:"chainCodes"`
}

type OrdererProperties struct {
	OrdererID      string `yaml:"ordererId" json:"ordererId"`
	OrdererAddress string `yaml:"ordererAddress" json:"ordererAddress"`
	OrdererPort    int    `yaml:"ordererPort" json:"ordererPort"`
}

// peer的配置
type PeerProperties struct {
	// anchorpeer的节点地址
	AnchorPeers []AnchorPeer `yaml:"anchorPeers" json:"anchorPeers"`
}

type UserProperties struct {
	UserName string `yaml:"userName" json:"userName"`
	// 是否是admin角色
	Admin bool `yaml:"admin" json:"admin"`
	// 是否需要enroll
}

type CaProperties struct {
	CaName string `yaml:"caName" json:"caName"`
}
type OrganizationProperties struct {
	// 组织id
	OrganizationID base.OrganizationID `yaml:"organizationId" json:"organizationId"`
	// 用户list,暂且只是一个string切片
	Users []UserProperties `yaml:"users" json:"users"`

	// 该组织下的peer节点
	Peers []PeerProperties `yaml:"peers" json:"peers"`

	// ca 信息
	Ca CaProperties `yaml:"ca" json:"ca"`
}

func (org OrganizationProperties) GetAdminUser() []string {
	users := make([]string, 0)

	for _, user := range org.Users {
		if user.Admin {
			users = append(users, user.UserName)
		}
	}

	return users
}

// 获取记录在该组织下的用户
func (this *OrganizationProperties) getEnrollUsers() []fabsdk.ContextOption {
	options := make([]fabsdk.ContextOption, 0)
	for _, user := range this.Users {
		options = append(options, fabsdk.WithUser(user.UserName))
	}
	return options
}
func (this OrganizationProperties) String() string {
	str := ""
	for _, u := range this.Users {
		str += "user:" + u.UserName + " "
	}
	str += "organizationId:" + string(this.OrganizationID) + " "

	return str
}

// channel的配置
type ChannelProperties struct {
	// channel名称
	ChannelID base.ChannelID `yaml:"channelId" json:"channelId"`
	// channel的定义路径: 既*.tx 路径
	ChannelConfigPath string `yaml:"channelConfigPath" json:"channelConfigPath"`
	// 2019-12-28 add
	// order 信息
	// 该channel下的参与的order节点 一个network网络只需要定义一个即可,其他的会自动gossip参与
	Orders []OrdererProperties `yaml:"orders" json:"orders"`
	// 该channel下的组织
	Organizations []OrganizationProperties `yaml:"organizations"json:"organizations"`

	// 2020-06-26 是否监听block事件
	NeedListOnBlockEvent bool `yaml:"needListOnBlockEvent" json:"needListOnBlockEvent"`
}

func (this ChannelProperties) GetInterestBlockEventChainCodes() []string {
	results := make([]string, 0)
	for _, org := range this.Organizations {
		for _, pp := range org.Peers {
			for _, anchorPeer := range pp.AnchorPeers {
				for _, chaincode := range anchorPeer.ChainCodes {
					results = append(results, string(chaincode.ChainCodeID))
				}
			}
		}
	}
	return results
}

type BlockChainProperties struct {
	// GOPATH路径
	GoPath string `yaml:"goPath" json:"goPath"`
	// config的路径,绝对路径
	ConfigPath string `yaml:"configPath" json:"configPath"`
	// 加密的版本号
	CryptVersion base.Version        `yaml:"cryptVersion" json:"version"`
	Channels     []ChannelProperties `yaml:"channels" json:"channels"`
}
