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
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"golang.org/x/fmt"
	"myLibrary/go-library/go/blockchain/base"
	"myLibrary/go-library/go/converters"
	error3 "myLibrary/go-library/go/error"
	"myLibrary/go-library/go/log"
	"strconv"
	"strings"
	"time"
)

type organizationChannelClientInfo struct {
	Client *channel.Client
}

type channelClientInfo struct {
	clients map[base.OrganizationID]*organizationChannelClientInfo
}

// 选择channel->选择channel下的organization->获取client
type ChannelClientWrapper struct {
	clients map[base.ChannelID]*channelClientInfo
}

// 组织的资源admin,用于控制channel中的资源
type OrganizationResourceAdmin struct {
	Admin *resmgmt.Client
}

type ResourceAdminInfo struct {
	admins map[base.OrganizationID]*OrganizationResourceAdmin
}

type ResourceAdminWrapper struct {
	admins map[base.ChannelID]*ResourceAdminInfo
}

// //

type ChannelEventInfo struct {
	blockRegistration      fab.Registration
	blockStopEventFlagChan chan struct{}
	EventClient            *event.Client
}

type ChannelEventWrapper struct {
	Events map[base.ChannelID]*ChannelEventInfo
}

// ///////
type ChannelLedgerInfo struct {
	Ledger *ledger.Client
}
type ChannelLedgerWrapper struct {
	ledgers map[base.ChannelID]*ChannelLedgerInfo
}

//
type VlinkBlockChainConfiguration struct {
	Log                  log.Logger
	Version              base.Version
	sdk                  *fabsdk.FabricSDK
	clientWrapper        *ChannelClientWrapper
	adminResourceWrapper *ResourceAdminWrapper
	events               *ChannelEventWrapper
	ledgers              *ChannelLedgerWrapper
}

type ExecuteReq struct {
	MethodName     base.MethodName
	ChannelID      base.ChannelID
	OrganizationID base.OrganizationID
	ChainCodeID    base.ChainCodeID
	ReqData        interface{}
}

func (this *VlinkBlockChainConfiguration) Execute(executeReq ExecuteReq) (base.ServiceLogicBaseResp, []byte, error3.IBaseError) {
	// name base.MethodName, id base.ChannelID, codeID ChainCodeID, req interface{}
	name := executeReq.MethodName
	id := executeReq.ChannelID
	codeID := executeReq.ChainCodeID
	req := executeReq.ReqData
	var (
		logicRes base.ServiceLogicBaseResp
		e        error3.IBaseError
	)
	response, vlinkError := this.defaultExecute(base.ChainBaseReq{
		MethodName:     name,
		ChannelID:      id,
		OrganizationID: executeReq.OrganizationID,
		ChainCodeID:    codeID,
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

func (this *VlinkBlockChainConfiguration) defaultExecute(b base.ChainBaseReq, req interface{}) (channel.Response, error3.IBaseError) {
	var d interface{}
	switch req.(type) {
	case base.ICrypter:
		encrypt, e := req.(base.ICrypter).Encrypt(this.Version)
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

func (this *VlinkBlockChainConfiguration) execute(b base.ChainBaseReq, data interface{}) (channel.Response, error3.IBaseError) {
	var args []string
	args = append(args, string(b.MethodName))

	bytes, e := json.Marshal(data)
	if e != nil {
		return channel.Response{}, error3.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", data))
	}
	args = append(args, string(bytes))
	adminClient := this.clientWrapper.clients[b.ChannelID].clients[b.OrganizationID].Client
	response, err := adminClient.Execute(channel.Request{
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
	for _, eveAdapter := range this.events.Events {
		if nil != eveAdapter.EventClient {
			eveAdapter.blockStopEventFlagChan <- struct{}{}
			eveAdapter.EventClient.Unregister(eveAdapter.blockRegistration)
		}
	}
	return nil
}

func HandleResponse(response channel.Response) (base.BaseFabricResp, error3.IBaseError) {
	bytes := response.Payload
	var resp base.BaseFabricResp

	e := json.Unmarshal(bytes, &resp)
	if nil != e {
		return resp, error3.NewJSONSerializeError(e, "反序列化结构体失败")
	}

	return resp, nil
}

func (setUp *VlinkBlockChainConfiguration) Config(p VlinkBlockChainProperties) error3.IBaseError {
	if e := setUp.initialize(p); nil != e {
		return e
	}

	if e := setUp.InstallAndInstantiateCC(p); nil != e {
		return e
	}

	// 启动监听block
	RunTasks()

	return nil
}

func (setUp *VlinkBlockChainConfiguration) initialize(p VlinkBlockChainProperties) error3.IBaseError {
	fmt.Println("begin 初始化SDK")
	c := config.FromFile(p.ConfigPath)
	sdk, e := fabsdk.New(c)
	if nil != e {
		panic(e)
	}
	// defer sdk.Close()
	setUp.sdk = sdk

	fmt.Println("end 初始化SDK")

	fmt.Println("begin 初始化资源管理器")
	// 多个组织有多个不同的organization ,所以需要for 遍历多次初始化
	channelCreated := false
	for _, channel := range p.Channels {
		for _, organization := range channel.Organizations {
			fmt.Println(fmt.Sprintf("begin 初始化组织为[%s]的资源管理器", organization.OrganizationID))
			fmt.Println("信息为:", organization.String())
			resourceManagerClientContext := sdk.Context(fabsdk.WithOrg(string(organization.OrganizationID)), fabsdk.WithUser(organization.OrganizationAdmin))
			admin, e := resmgmt.New(resourceManagerClientContext)
			if nil != e {
				s := fmt.Sprintf("初始化组织为[%s]的资源管理器失败:%s", organization.OrganizationID, e.Error())
				return error3.NewConfigError(e, s)
			}
			if setUp.adminResourceWrapper == nil {
				w := new(ResourceAdminWrapper)
				w.admins = make(map[base.ChannelID]*ResourceAdminInfo)
				setUp.adminResourceWrapper = w
			}
			m := setUp.adminResourceWrapper.admins[channel.ChannelID]
			if m == nil {
				resourceMap := new(ResourceAdminInfo)
				resourceMap.admins = make(map[base.OrganizationID]*OrganizationResourceAdmin)
				setUp.adminResourceWrapper.admins[channel.ChannelID] = resourceMap
				m = resourceMap
			}
			m.admins[organization.OrganizationID] = &OrganizationResourceAdmin{
				Admin: admin,
			}
			fmt.Println(fmt.Sprintf("end 初始化组织为[%s]的资源管理器", organization.OrganizationID))

			fmt.Println("begin 开始初始化admin-mspclient")
			mspClient, e := mspclient.New(sdk.Context(), mspclient.WithOrg(string(organization.OrganizationID)))
			if nil != e {
				panic(e)
			}
			fmt.Println("begin 组装identites")
			identites := make([]msp.SigningIdentity, 0)
			identity, e := mspClient.GetSigningIdentity(organization.OrganizationAdmin)
			if nil != e {
				panic(e)
			}
			identites = append(identites, identity)
			fmt.Println("end 组装identites")
			fmt.Println("end 初始化msp-client")

			for _, peer := range organization.Peers {
				fmt.Println("begin 查询已经存在的channel")
				channelResp, e := admin.QueryChannels(resmgmt.WithTargetEndpoints(peer.AnchorPeers[0].Address))
				if nil != e {
					return error3.NewConfigError(e, fmt.Sprintf("查询anchorpeer=[%s]上的channel失败:%s", peer.AnchorPeers[0].Address, e.Error()))
				}
				fmt.Println("end 查询已经存在的channel")

				fmt.Println("begin 判断channel是否已经存在")
				if nil != channelResp {
					for _, c := range channelResp.Channels {
						if strings.EqualFold(c.ChannelId, string(channel.ChannelID)) {
							channelCreated = true
							break
						}
					}
				}

				// p := models.VlinkPeer{
				// 	ChannelName: string(channel.base.ChannelID),
				// 	Domain:      peer.AnchorPeers[0].Address,
				// 	Port:        peer.AnchorPeers[0].Port,
				// }
				if channelCreated {
					fmt.Println(fmt.Sprintf("channel:[%s]已经存在\n", string(channel.ChannelID)))
				} else {
					fmt.Println("begin 创建channel")
					saveChanReq := resmgmt.SaveChannelRequest{ChannelID: string(channel.ChannelID), ChannelConfigPath: channel.ChannelConfigPath, SigningIdentities: identites}
					// 获取某个order的keys 即可
					endPoints := make([]resmgmt.RequestOption, 0)
					orderP := channel.Orders[0]
					endPoints = append(endPoints, resmgmt.WithOrdererEndpoint(orderP.OrdererAddress))
					saveChanResp, e := admin.SaveChannel(saveChanReq, endPoints...)
					if nil != e || saveChanResp.TransactionID == "" {
						panic(e)
					}
					fmt.Println("end 创建channel,channel创建成功")

					fmt.Println("begin 将节点加入channel")
					if e = admin.JoinChannel(string(channel.ChannelID), endPoints...); nil != e {
						s := fmt.Sprintf("channelId=[%s]加入通道失败:%s", channel.ChannelID, e.Error())
						return error3.NewConfigError(e, s)
					}
					fmt.Println("end 将节点加入channel")
				}
			}

			fmt.Println("begin 创建区块链账本相关")

			cCtx := make([]fabsdk.ContextOption, 0)
			cCtx = append(cCtx, fabsdk.WithOrg(string(organization.OrganizationID)))
			cCtx = append(cCtx, fabsdk.WithUser(organization.OrganizationAdmin))
			ledgerContext := sdk.ChannelContext(string(channel.ChannelID), cCtx...)
			ledgerClient, e := ledger.New(ledgerContext)
			if nil != e {
				panic(e)
			}
			if setUp.ledgers == nil {
				ledgerWrapper := new(ChannelLedgerWrapper)
				ledgerWrapper.ledgers = make(map[base.ChannelID]*ChannelLedgerInfo)
				setUp.ledgers = ledgerWrapper
			}
			setUp.ledgers.ledgers[channel.ChannelID] = &ChannelLedgerInfo{
				Ledger: ledgerClient,
			}
			fmt.Println("end 创建区块链账本相关")
		}
	}

	return nil
}

func (setUp *VlinkBlockChainConfiguration) InstallAndInstantiateCC(p VlinkBlockChainProperties) error3.IBaseError {
	fmt.Println("begin InstallAndInstantiateCC")
	for _, ccan := range p.Channels {
		for _, org := range ccan.Organizations {
			for _, peer := range org.Peers {
				for _, anchorPeer := range peer.AnchorPeers {
					for _, chaincode := range anchorPeer.ChainCodes {
						fmt.Println(fmt.Sprintf("begin 创建chaincode package,chaincodeId=[%s]"), chaincode.ChainCodeID)
						ccPackage, e := packager.NewCCPackage(chaincode.ChainCodePath, p.GoPath)
						if nil != e {
							panic(e)
						}
						fmt.Println("end 创建chaincodepackage")
						ccIsInstall := false

						fmt.Println("begin 查询已经安装的chaincode")
						admin := setUp.adminResourceWrapper.admins[ccan.ChannelID].admins[org.OrganizationID].Admin
						anchorAddress := peer.AnchorPeers[0].Address
						queryInstallCcResp, e := admin.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(anchorAddress))
						if nil != e {
							panic(e)
						}

						fmt.Printf("begin 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)
						Versions := make([]int, 0)
						Versions = append(Versions, 0)
						for _, c := range queryInstallCcResp.Chaincodes {
							if strings.EqualFold(string(chaincode.ChainCodeID), c.Name) {
								ccIsInstall = true
								i, _ := strconv.Atoi(c.Version)
								Versions = append(Versions, i+1)
							}
						}

						fmt.Printf("end 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)
						if ccIsInstall {
							fmt.Printf("该cc[%s]已经安装\n", chaincode.ChainCodeID)
							if chaincode.NeedUpdate {
								fmt.Println(fmt.Sprintf("链码[%s]需要升级,版本号为:%d", chaincode.ChainCodeID, Versions[len(Versions)-1]))
								newInstallCCReq := resmgmt.InstallCCRequest{Name: string(chaincode.ChainCodeID), Path: chaincode.ChainCodePath, Version: strconv.Itoa(Versions[len(Versions)-1]), Package: ccPackage}
								_, err := admin.InstallCC(newInstallCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
								if err != nil {
									return error3.NewConfigError(err, "failed to install chaincode")
								}
								fmt.Println("Chaincode install successfully ,begin upgrade")
								request := resmgmt.UpgradeCCRequest{
									Name:    string(chaincode.ChainCodeID),
									Path:    chaincode.ChainCodePath,
									Version: strconv.Itoa(Versions[len(Versions)-1]),
								}
								member := cauthdsl.SignedByMspMember(string(org.OrganizationID))
								request.Policy = member
								response, err := admin.UpgradeCC(string(ccan.ChannelID), request)
								if nil != err {
									fmt.Println("更新链码失败:", err.Error())
									return error3.NewConfigError(err, "更新链码失败")
								} else {
									fmt.Println(response.TransactionID)
								}
								fmt.Println(fmt.Sprintf("更新链码[%s]成功"), chaincode.ChainCodeID)
							}
						} else {
							fmt.Printf("begin 安装[%s]链码\n", chaincode.ChainCodeID)
							installCcReq := resmgmt.InstallCCRequest{
								Name:    string(chaincode.ChainCodeID),
								Path:    chaincode.ChainCodePath,
								Version: "0",
								Package: ccPackage,
							}
							responses, e := admin.InstallCC(installCcReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
							if nil != e || responses == nil {
								panic(e)
							}
							fmt.Printf("end 安装[%s]链码\n", chaincode)
						}
						ccIsInstance := false
						fmt.Println("begin 查询已经实例化的链码")
						queryInstanReps, e := admin.QueryInstantiatedChaincodes(string(ccan.ChannelID), resmgmt.WithTargetEndpoints(anchorAddress))
						if nil != e {
							s := fmt.Sprintf("查询通道=[%s]上实例化的链码失败:%s", ccan.ChannelID)
							return error3.NewConfigError(e, s)
						}

						for _, c := range queryInstanReps.Chaincodes {
							if strings.EqualFold(c.Name, string(chaincode.ChainCodeID)) {
								ccIsInstance = true
								break
							}
						}
						fmt.Println("end 查询已经实例化的链码")

						if ccIsInstance {
							fmt.Printf("该[%s]已经实例化\n", chaincode.ChainCodeID)
						} else {
							fmt.Printf("begin cc[%s]实例化\n", chaincode.ChainCodeID)
							ccPolicy := cauthdsl.SignedByMspMember(string(org.OrganizationID))
							instanReq := resmgmt.InstantiateCCRequest{
								Name:       string(chaincode.ChainCodeID),
								Path:       chaincode.ChainCodePath,
								Version:    "0",
								Args:       [][]byte{[]byte("init"), []byte("init")},
								Policy:     ccPolicy,
								CollConfig: nil,
							}
							instanResp, e := admin.InstantiateCC(string(ccan.ChannelID), instanReq, resmgmt.WithTargetEndpoints(anchorAddress))
							if nil != e || instanResp.TransactionID == "" {
								panic(e)
							}
							fmt.Printf("end cc[%s]实例化\n", chaincode.ChainCodeID)
						}
					}

				}
			}

			fmt.Println("begin 创建用于execute和query的channel client,基于channel->organization")
			// FIXME 这里需要确定,是通过channelId,还是org还是user 获取client信息
			// 若是org ,则 需要有一个 通过channelId获取org的map
			channelContext := setUp.sdk.ChannelContext(string(ccan.ChannelID), org.getEnrollUsers()...)
			channelClient, e := channel.New(channelContext)
			if nil != e {
				return error3.NewConfigError(e, "创建通道:"+string(ccan.ChannelID)+"失败")
			}

			if setUp.clientWrapper == nil {
				clientWrapper := new(ChannelClientWrapper)
				clientWrapper.clients = make(map[base.ChannelID]*channelClientInfo)
				setUp.clientWrapper = clientWrapper
			}
			clientOrganzationInfo := setUp.clientWrapper.clients[ccan.ChannelID]
			if clientOrganzationInfo.clients == nil {
				clientOrganzationInfo = &channelClientInfo{
					clients: map[base.OrganizationID]*organizationChannelClientInfo{org.OrganizationID: &organizationChannelClientInfo{
						Client: channelClient,
					}},
				}
				setUp.clientWrapper.clients[ccan.ChannelID] = clientOrganzationInfo
			} else if _, exist := clientOrganzationInfo.clients[org.OrganizationID]; !exist {
				clientOrganzationInfo.clients[org.OrganizationID]= &organizationChannelClientInfo{
					Client: channelClient,
				}
			}else{
				continue
			}

			fmt.Println("end 创建用于execute和query的channel client")

			fmt.Println("begin 创建event事件客户端")
			eveClient, e := event.New(channelContext, event.WithBlockEvents())
			if nil != e {
				return error3.NewConfigError(e, "创建通道事件:"+string(ccan.ChannelID)+"失败")
			}
			if setUp.events == nil {
				eventWrapper:=&ChannelEventWrapper{
					Events: make( map[base.ChannelID]*ChannelEventInfo),
				}
				setUp.events =eventWrapper
			}
			eveAdapter := new(ChannelEventInfo)
			blockStopEventFlagChan := make(chan struct{}, 1)
			eveAdapter.blockStopEventFlagChan = blockStopEventFlagChan
			eveAdapter.EventClient = eveClient
			setUp.events.Events[ccan.ChannelID] = eveAdapter
			fmt.Println("begin 创建event事件客户端")
			registration, events, e := eveClient.RegisterBlockEvent()
			if nil != e {
				return error3.NewSystemError(e, "监听block事件失败")
			}
			eveAdapter.blockRegistration = registration
			RegisterBlockEvent(ccan.ChannelID, ccan.GetInterestBlockEventChainCodes(), events, blockStopEventFlagChan)

			fmt.Println("end InstallAndInstantiateCC")
		}

	}

	return nil
}
