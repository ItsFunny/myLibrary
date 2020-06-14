/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-24 12:43
# @File : blockchian.go
# @Description :
# @Attention :
*/
package config

import (
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"golang.org/x/fmt"
	"myLibrary/go-library/go/converters"
	error3 "myLibrary/go-library/go/error"
	"myLibrary/go-library/go/log"
	"strconv"
	"time"
)



// // FIXME 命名规范
type TClientMap map[interface{}]*channel.Client
type TAdminMap map[OrganizationID]*resmgmt.Client
type TEventMap map[interface{}]*event.Client

//
type VlinkEventAdapter struct {
	client                 *event.Client
	blockRegistration      fab.Registration
	blockStopEventFlagChan chan struct{}
}

//
type VlinkBlockChainConfiguration struct {
	Log     log.Logger
	Version Version
	sdk     *fabsdk.FabricSDK
	clients map[ChannelID]*channel.Client
	admins  map[ChannelID]TAdminMap
	events  map[ChannelID]*VlinkEventAdapter
	ledgers map[ChannelID]*ledger.Client
}



func (this *VlinkBlockChainConfiguration) Execute(name MethodName, id ChannelID, codeID ChainCodeID, req interface{}) (ServiceLogicBaseResp, []byte,  error3.IBaseError) {
	var (
		logicRes ServiceLogicBaseResp
		e        error3.IBaseError
	)
	response, vlinkError := this.defaultExecute(ChainBaseReq{
		MethodName:  name,
		ChannelID:   id,
		ChainCodeID: codeID,
	}, req)
	if nil != vlinkError {
		return logicRes, nil, vlinkError
	}
	resp, _ := HandleResponse(response)
	logicRes.LogicCode = int(converter.BigEndianBytes2Int64(resp.CodeBytes))
	logicRes.LogicMsg = string(resp.MsgBytes)
	logicRes.LogBytes = resp.LogBytes
	if e := json.Unmarshal(resp.OtherBytes, logicRes.CommAttribute); nil != e {
		return logicRes, nil, error3.NewJSONSerializeError(e, "CommAttribute反序列化失败")
	}

	return logicRes, resp.DataBytes, e
}

func (this *VlinkBlockChainConfiguration) defaultExecute(b ChainBaseReq, req interface{}) (channel.Response, error3.IBaseError) {
	var d interface{}
	switch req.(type) {
	case ICrypter:
		encrypt, e := req.(ICrypter).Encrypt(this.Version)
		if nil != e {
			return channel.Response{}, error3.NewArguError(e, "参数加密失败")
		}
		d = encrypt
	default:
		d = req
	}
	// response, vlinkError := this.execute(b, d)
	// if nil!=vlinkError{
	// 	return channel.Response{},vlinkError
	// }
	// return response,nil

	return this.execute(b, d)
}

func (this *VlinkBlockChainConfiguration) execute(b ChainBaseReq, data interface{}) (channel.Response, error3.IBaseError) {
	var args []string
	args = append(args, string(b.MethodName))

	bytes, e := json.Marshal(data)
	if e != nil {
		return channel.Response{}, error3.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", data))
	}
	args = append(args, string(bytes))
	response, err := this.clients[b.ChannelID].Execute(channel.Request{
		ChaincodeID: string(b.ChainCodeID),
		Fcn:         args[0],
		// 2020-01-08 update 为了与invokechaincode 一致,此处补齐一个string
		Args: [][]byte{[]byte(args[1]), []byte(strconv.Itoa(int(this.Version)))},
	}, channel.WithTimeout(fab.Execute, time.Second*60))

	if nil != err {
		return response, error3.NewFabricError(err, fmt.Sprintf("调用fabric失败,方法名称为:%s", b.MethodName))
	}

	return response, nil
}

func (this *VlinkBlockChainConfiguration) Close() error {
	for _, eveAdapter := range this.events {
		if nil != eveAdapter.client {
			eveAdapter.blockStopEventFlagChan <- struct{}{}
			eveAdapter.client.Unregister(eveAdapter.blockRegistration)
		}
	}
	return nil
}

func HandleResponse(response channel.Response) (BaseFabricResp, error3.IBaseError) {
	bytes := response.Payload
	var resp BaseFabricResp

	e := json.Unmarshal(bytes, &resp)
	if nil != e {
		return resp, error3.NewJSONSerializeError(e, "反序列化结构体失败")
	}

	return resp, nil
}
