// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180416

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-04-16"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyUserCertRequest() (request *ApplyUserCertRequest) {
    request = &ApplyUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "ApplyUserCert")
    return
}

func NewApplyUserCertResponse() (response *ApplyUserCertResponse) {
    response = &ApplyUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 申请用户证书
func (c *Client) ApplyUserCert(request *ApplyUserCertRequest) (response *ApplyUserCertResponse, err error) {
    if request == nil {
        request = NewApplyUserCertRequest()
    }
    response = NewApplyUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewBlockByNumberHandlerRequest() (request *BlockByNumberHandlerRequest) {
    request = &BlockByNumberHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "BlockByNumberHandler")
    return
}

func NewBlockByNumberHandlerResponse() (response *BlockByNumberHandlerResponse) {
    response = &BlockByNumberHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// Bcos根据块高查询区块信息
func (c *Client) BlockByNumberHandler(request *BlockByNumberHandlerRequest) (response *BlockByNumberHandlerResponse, err error) {
    if request == nil {
        request = NewBlockByNumberHandlerRequest()
    }
    response = NewBlockByNumberHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewCreateChaincodeAndInstallForUserRequest() (request *CreateChaincodeAndInstallForUserRequest) {
    request = &CreateChaincodeAndInstallForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "CreateChaincodeAndInstallForUser")
    return
}

func NewCreateChaincodeAndInstallForUserResponse() (response *CreateChaincodeAndInstallForUserResponse) {
    response = &CreateChaincodeAndInstallForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建并安装合约
func (c *Client) CreateChaincodeAndInstallForUser(request *CreateChaincodeAndInstallForUserRequest) (response *CreateChaincodeAndInstallForUserResponse, err error) {
    if request == nil {
        request = NewCreateChaincodeAndInstallForUserRequest()
    }
    response = NewCreateChaincodeAndInstallForUserResponse()
    err = c.Send(request, response)
    return
}

func NewDeployDynamicBcosContractRequest() (request *DeployDynamicBcosContractRequest) {
    request = &DeployDynamicBcosContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DeployDynamicBcosContract")
    return
}

func NewDeployDynamicBcosContractResponse() (response *DeployDynamicBcosContractResponse) {
    response = &DeployDynamicBcosContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 动态部署并发布Bcos合约
func (c *Client) DeployDynamicBcosContract(request *DeployDynamicBcosContractRequest) (response *DeployDynamicBcosContractResponse, err error) {
    if request == nil {
        request = NewDeployDynamicBcosContractRequest()
    }
    response = NewDeployDynamicBcosContractResponse()
    err = c.Send(request, response)
    return
}

func NewDeployDynamicContractHandlerRequest() (request *DeployDynamicContractHandlerRequest) {
    request = &DeployDynamicContractHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DeployDynamicContractHandler")
    return
}

func NewDeployDynamicContractHandlerResponse() (response *DeployDynamicContractHandlerResponse) {
    response = &DeployDynamicContractHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// 动态部署合约
func (c *Client) DeployDynamicContractHandler(request *DeployDynamicContractHandlerRequest) (response *DeployDynamicContractHandlerResponse, err error) {
    if request == nil {
        request = NewDeployDynamicContractHandlerRequest()
    }
    response = NewDeployDynamicContractHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadUserCertRequest() (request *DownloadUserCertRequest) {
    request = &DownloadUserCertRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "DownloadUserCert")
    return
}

func NewDownloadUserCertResponse() (response *DownloadUserCertResponse) {
    response = &DownloadUserCertResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下载用户证书
func (c *Client) DownloadUserCert(request *DownloadUserCertRequest) (response *DownloadUserCertResponse, err error) {
    if request == nil {
        request = NewDownloadUserCertRequest()
    }
    response = NewDownloadUserCertResponse()
    err = c.Send(request, response)
    return
}

func NewGetBcosBlockByNumberRequest() (request *GetBcosBlockByNumberRequest) {
    request = &GetBcosBlockByNumberRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBcosBlockByNumber")
    return
}

func NewGetBcosBlockByNumberResponse() (response *GetBcosBlockByNumberResponse) {
    response = &GetBcosBlockByNumberResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 使用块高查询Bcos区块信息
func (c *Client) GetBcosBlockByNumber(request *GetBcosBlockByNumberRequest) (response *GetBcosBlockByNumberResponse, err error) {
    if request == nil {
        request = NewGetBcosBlockByNumberRequest()
    }
    response = NewGetBcosBlockByNumberResponse()
    err = c.Send(request, response)
    return
}

func NewGetBcosBlockListRequest() (request *GetBcosBlockListRequest) {
    request = &GetBcosBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBcosBlockList")
    return
}

func NewGetBcosBlockListResponse() (response *GetBcosBlockListResponse) {
    response = &GetBcosBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos分页查询当前群组下的区块列表
func (c *Client) GetBcosBlockList(request *GetBcosBlockListRequest) (response *GetBcosBlockListResponse, err error) {
    if request == nil {
        request = NewGetBcosBlockListRequest()
    }
    response = NewGetBcosBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBcosTransByHashRequest() (request *GetBcosTransByHashRequest) {
    request = &GetBcosTransByHashRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBcosTransByHash")
    return
}

func NewGetBcosTransByHashResponse() (response *GetBcosTransByHashResponse) {
    response = &GetBcosTransByHashResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos根据交易哈希查看交易详细信息
func (c *Client) GetBcosTransByHash(request *GetBcosTransByHashRequest) (response *GetBcosTransByHashResponse, err error) {
    if request == nil {
        request = NewGetBcosTransByHashRequest()
    }
    response = NewGetBcosTransByHashResponse()
    err = c.Send(request, response)
    return
}

func NewGetBcosTransListRequest() (request *GetBcosTransListRequest) {
    request = &GetBcosTransListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBcosTransList")
    return
}

func NewGetBcosTransListResponse() (response *GetBcosTransListResponse) {
    response = &GetBcosTransListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Bcos分页查询当前群组的交易信息列表
func (c *Client) GetBcosTransList(request *GetBcosTransListRequest) (response *GetBcosTransListResponse, err error) {
    if request == nil {
        request = NewGetBcosTransListRequest()
    }
    response = NewGetBcosTransListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListRequest() (request *GetBlockListRequest) {
    request = &GetBlockListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockList")
    return
}

func NewGetBlockListResponse() (response *GetBlockListResponse) {
    response = &GetBlockListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看当前网络下的所有区块列表，分页展示
func (c *Client) GetBlockList(request *GetBlockListRequest) (response *GetBlockListResponse, err error) {
    if request == nil {
        request = NewGetBlockListRequest()
    }
    response = NewGetBlockListResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockListHandlerRequest() (request *GetBlockListHandlerRequest) {
    request = &GetBlockListHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockListHandler")
    return
}

func NewGetBlockListHandlerResponse() (response *GetBlockListHandlerResponse) {
    response = &GetBlockListHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// Bcos分页查询当前群组下的区块列表
func (c *Client) GetBlockListHandler(request *GetBlockListHandlerRequest) (response *GetBlockListHandlerResponse, err error) {
    if request == nil {
        request = NewGetBlockListHandlerRequest()
    }
    response = NewGetBlockListHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetBlockTransactionListForUserRequest() (request *GetBlockTransactionListForUserRequest) {
    request = &GetBlockTransactionListForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetBlockTransactionListForUser")
    return
}

func NewGetBlockTransactionListForUserResponse() (response *GetBlockTransactionListForUserResponse) {
    response = &GetBlockTransactionListForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块内的交易列表
func (c *Client) GetBlockTransactionListForUser(request *GetBlockTransactionListForUserRequest) (response *GetBlockTransactionListForUserResponse, err error) {
    if request == nil {
        request = NewGetBlockTransactionListForUserRequest()
    }
    response = NewGetBlockTransactionListForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeCompileLogForUserRequest() (request *GetChaincodeCompileLogForUserRequest) {
    request = &GetChaincodeCompileLogForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeCompileLogForUser")
    return
}

func NewGetChaincodeCompileLogForUserResponse() (response *GetChaincodeCompileLogForUserResponse) {
    response = &GetChaincodeCompileLogForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取合约编译日志
func (c *Client) GetChaincodeCompileLogForUser(request *GetChaincodeCompileLogForUserRequest) (response *GetChaincodeCompileLogForUserResponse, err error) {
    if request == nil {
        request = NewGetChaincodeCompileLogForUserRequest()
    }
    response = NewGetChaincodeCompileLogForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeInitializeResultForUserRequest() (request *GetChaincodeInitializeResultForUserRequest) {
    request = &GetChaincodeInitializeResultForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeInitializeResultForUser")
    return
}

func NewGetChaincodeInitializeResultForUserResponse() (response *GetChaincodeInitializeResultForUserResponse) {
    response = &GetChaincodeInitializeResultForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 实例化结果查询
func (c *Client) GetChaincodeInitializeResultForUser(request *GetChaincodeInitializeResultForUserRequest) (response *GetChaincodeInitializeResultForUserResponse, err error) {
    if request == nil {
        request = NewGetChaincodeInitializeResultForUserRequest()
    }
    response = NewGetChaincodeInitializeResultForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetChaincodeLogForUserRequest() (request *GetChaincodeLogForUserRequest) {
    request = &GetChaincodeLogForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChaincodeLogForUser")
    return
}

func NewGetChaincodeLogForUserResponse() (response *GetChaincodeLogForUserResponse) {
    response = &GetChaincodeLogForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取合约容器日志
func (c *Client) GetChaincodeLogForUser(request *GetChaincodeLogForUserRequest) (response *GetChaincodeLogForUserResponse, err error) {
    if request == nil {
        request = NewGetChaincodeLogForUserRequest()
    }
    response = NewGetChaincodeLogForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetChannelListForUserRequest() (request *GetChannelListForUserRequest) {
    request = &GetChannelListForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetChannelListForUser")
    return
}

func NewGetChannelListForUserResponse() (response *GetChannelListForUserResponse) {
    response = &GetChannelListForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取通道列表
func (c *Client) GetChannelListForUser(request *GetChannelListForUserRequest) (response *GetChannelListForUserResponse, err error) {
    if request == nil {
        request = NewGetChannelListForUserRequest()
    }
    response = NewGetChannelListForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterListForUserRequest() (request *GetClusterListForUserRequest) {
    request = &GetClusterListForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterListForUser")
    return
}

func NewGetClusterListForUserResponse() (response *GetClusterListForUserResponse) {
    response = &GetClusterListForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取该用户的网络列表。网络信息中包含组织信息，但仅包含该用户所在组织的信息。
func (c *Client) GetClusterListForUser(request *GetClusterListForUserRequest) (response *GetClusterListForUserResponse, err error) {
    if request == nil {
        request = NewGetClusterListForUserRequest()
    }
    response = NewGetClusterListForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetClusterSummaryRequest() (request *GetClusterSummaryRequest) {
    request = &GetClusterSummaryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetClusterSummary")
    return
}

func NewGetClusterSummaryResponse() (response *GetClusterSummaryResponse) {
    response = &GetClusterSummaryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取区块链网络概要
func (c *Client) GetClusterSummary(request *GetClusterSummaryRequest) (response *GetClusterSummaryResponse, err error) {
    if request == nil {
        request = NewGetClusterSummaryRequest()
    }
    response = NewGetClusterSummaryResponse()
    err = c.Send(request, response)
    return
}

func NewGetInvokeTxRequest() (request *GetInvokeTxRequest) {
    request = &GetInvokeTxRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetInvokeTx")
    return
}

func NewGetInvokeTxResponse() (response *GetInvokeTxResponse) {
    response = &GetInvokeTxResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// Invoke异步调用结果查询
func (c *Client) GetInvokeTx(request *GetInvokeTxRequest) (response *GetInvokeTxResponse, err error) {
    if request == nil {
        request = NewGetInvokeTxRequest()
    }
    response = NewGetInvokeTxResponse()
    err = c.Send(request, response)
    return
}

func NewGetLatesdTransactionListRequest() (request *GetLatesdTransactionListRequest) {
    request = &GetLatesdTransactionListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetLatesdTransactionList")
    return
}

func NewGetLatesdTransactionListResponse() (response *GetLatesdTransactionListResponse) {
    response = &GetLatesdTransactionListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取最新交易列表
func (c *Client) GetLatesdTransactionList(request *GetLatesdTransactionListRequest) (response *GetLatesdTransactionListResponse, err error) {
    if request == nil {
        request = NewGetLatesdTransactionListRequest()
    }
    response = NewGetLatesdTransactionListResponse()
    err = c.Send(request, response)
    return
}

func NewGetPeerLogForUserRequest() (request *GetPeerLogForUserRequest) {
    request = &GetPeerLogForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetPeerLogForUser")
    return
}

func NewGetPeerLogForUserResponse() (response *GetPeerLogForUserResponse) {
    response = &GetPeerLogForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取节点日志
func (c *Client) GetPeerLogForUser(request *GetPeerLogForUserRequest) (response *GetPeerLogForUserResponse, err error) {
    if request == nil {
        request = NewGetPeerLogForUserRequest()
    }
    response = NewGetPeerLogForUserResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransByHashHandlerRequest() (request *GetTransByHashHandlerRequest) {
    request = &GetTransByHashHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransByHashHandler")
    return
}

func NewGetTransByHashHandlerResponse() (response *GetTransByHashHandlerResponse) {
    response = &GetTransByHashHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// Bcos根据交易哈希查看交易详细信息
func (c *Client) GetTransByHashHandler(request *GetTransByHashHandlerRequest) (response *GetTransByHashHandlerResponse, err error) {
    if request == nil {
        request = NewGetTransByHashHandlerRequest()
    }
    response = NewGetTransByHashHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransListHandlerRequest() (request *GetTransListHandlerRequest) {
    request = &GetTransListHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransListHandler")
    return
}

func NewGetTransListHandlerResponse() (response *GetTransListHandlerResponse) {
    response = &GetTransListHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// Bcos分页查询当前群组的交易信息列表
func (c *Client) GetTransListHandler(request *GetTransListHandlerRequest) (response *GetTransListHandlerResponse, err error) {
    if request == nil {
        request = NewGetTransListHandlerRequest()
    }
    response = NewGetTransListHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewGetTransactionDetailForUserRequest() (request *GetTransactionDetailForUserRequest) {
    request = &GetTransactionDetailForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "GetTransactionDetailForUser")
    return
}

func NewGetTransactionDetailForUserResponse() (response *GetTransactionDetailForUserResponse) {
    response = &GetTransactionDetailForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取交易详情
func (c *Client) GetTransactionDetailForUser(request *GetTransactionDetailForUserRequest) (response *GetTransactionDetailForUserResponse, err error) {
    if request == nil {
        request = NewGetTransactionDetailForUserRequest()
    }
    response = NewGetTransactionDetailForUserResponse()
    err = c.Send(request, response)
    return
}

func NewInitializeChaincodeForUserRequest() (request *InitializeChaincodeForUserRequest) {
    request = &InitializeChaincodeForUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InitializeChaincodeForUser")
    return
}

func NewInitializeChaincodeForUserResponse() (response *InitializeChaincodeForUserResponse) {
    response = &InitializeChaincodeForUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 实例化合约
func (c *Client) InitializeChaincodeForUser(request *InitializeChaincodeForUserRequest) (response *InitializeChaincodeForUserResponse, err error) {
    if request == nil {
        request = NewInitializeChaincodeForUserRequest()
    }
    response = NewInitializeChaincodeForUserResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeRequest() (request *InvokeRequest) {
    request = &InvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Invoke")
    return
}

func NewInvokeResponse() (response *InvokeResponse) {
    response = &InvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新增交易
func (c *Client) Invoke(request *InvokeRequest) (response *InvokeResponse, err error) {
    if request == nil {
        request = NewInvokeRequest()
    }
    response = NewInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewInvokeBcosTransRequest() (request *InvokeBcosTransRequest) {
    request = &InvokeBcosTransRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "InvokeBcosTrans")
    return
}

func NewInvokeBcosTransResponse() (response *InvokeBcosTransResponse) {
    response = &InvokeBcosTransResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 执行Bcos交易，支持动态部署的合约
func (c *Client) InvokeBcosTrans(request *InvokeBcosTransRequest) (response *InvokeBcosTransResponse, err error) {
    if request == nil {
        request = NewInvokeBcosTransRequest()
    }
    response = NewInvokeBcosTransResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRequest() (request *QueryRequest) {
    request = &QueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "Query")
    return
}

func NewQueryResponse() (response *QueryResponse) {
    response = &QueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询交易
func (c *Client) Query(request *QueryRequest) (response *QueryResponse, err error) {
    if request == nil {
        request = NewQueryRequest()
    }
    response = NewQueryResponse()
    err = c.Send(request, response)
    return
}

func NewSendTransactionHandlerRequest() (request *SendTransactionHandlerRequest) {
    request = &SendTransactionHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "SendTransactionHandler")
    return
}

func NewSendTransactionHandlerResponse() (response *SendTransactionHandlerResponse) {
    response = &SendTransactionHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// Bcos发送交易
func (c *Client) SendTransactionHandler(request *SendTransactionHandlerRequest) (response *SendTransactionHandlerResponse, err error) {
    if request == nil {
        request = NewSendTransactionHandlerRequest()
    }
    response = NewSendTransactionHandlerResponse()
    err = c.Send(request, response)
    return
}

func NewSrvInvokeRequest() (request *SrvInvokeRequest) {
    request = &SrvInvokeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "SrvInvoke")
    return
}

func NewSrvInvokeResponse() (response *SrvInvokeResponse) {
    response = &SrvInvokeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// trustsql服务统一接口
func (c *Client) SrvInvoke(request *SrvInvokeRequest) (response *SrvInvokeResponse, err error) {
    if request == nil {
        request = NewSrvInvokeRequest()
    }
    response = NewSrvInvokeResponse()
    err = c.Send(request, response)
    return
}

func NewTransByDynamicContractHandlerRequest() (request *TransByDynamicContractHandlerRequest) {
    request = &TransByDynamicContractHandlerRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tbaas", APIVersion, "TransByDynamicContractHandler")
    return
}

func NewTransByDynamicContractHandlerResponse() (response *TransByDynamicContractHandlerResponse) {
    response = &TransByDynamicContractHandlerResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 版本升级
// 
// 根据动态部署的合约发送交易
func (c *Client) TransByDynamicContractHandler(request *TransByDynamicContractHandlerRequest) (response *TransByDynamicContractHandlerResponse, err error) {
    if request == nil {
        request = NewTransByDynamicContractHandlerRequest()
    }
    response = NewTransByDynamicContractHandlerResponse()
    err = c.Send(request, response)
    return
}
