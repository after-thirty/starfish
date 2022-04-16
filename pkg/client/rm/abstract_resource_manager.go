/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rm

import (
	"strings"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/base/model"
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
	"github.com/transaction-mesh/starfish/pkg/client/config"
	"github.com/transaction-mesh/starfish/pkg/client/context"
	"github.com/transaction-mesh/starfish/pkg/client/rpc_client"
)

var (
	DBKEYS_SPLIT_CHAR = ","
)

type AbstractResourceManager struct {
	RpcClient     *rpc_client.RpcRemoteClient
	ResourceCache map[string]model.IResource
}

func NewAbstractResourceManager(client *rpc_client.RpcRemoteClient) AbstractResourceManager {
	resourceManager := AbstractResourceManager{
		RpcClient:     client,
		ResourceCache: make(map[string]model.IResource),
	}
	go resourceManager.handleRegisterRM()
	return resourceManager
}

func (resourceManager AbstractResourceManager) RegisterResource(resource model.IResource) {
	resourceManager.ResourceCache[resource.GetResourceID()] = resource
}

func (resourceManager AbstractResourceManager) UnregisterResource(resource model.IResource) {

}

func (resourceManager AbstractResourceManager) GetManagedResources() map[string]model.IResource {
	return resourceManager.ResourceCache
}

func (resourceManager AbstractResourceManager) BranchRegister(branchType meta.BranchType, resourceID string,
	clientID string, xid string, applicationData []byte, lockKeys string) (int64, error) {
	request := protocal.BranchRegisterRequest{
		XID:             xid,
		BranchType:      branchType,
		ResourceID:      resourceID,
		LockKey:         lockKeys,
		ApplicationData: applicationData,
	}
	resp, err := resourceManager.RpcClient.SendMsgWithResponse(request)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	response := resp.(protocal.BranchRegisterResponse)
	if response.ResultCode == protocal.ResultCodeSuccess {
		return response.BranchID, nil
	} else {
		return 0, response.GetError()
	}
}

func (resourceManager AbstractResourceManager) BranchReport(branchType meta.BranchType, xid string, branchID int64,
	status meta.BranchStatus, applicationData []byte) error {
	request := protocal.BranchReportRequest{
		XID:             xid,
		BranchID:        branchID,
		Status:          status,
		ApplicationData: applicationData,
	}
	resp, err := resourceManager.RpcClient.SendMsgWithResponse(request)
	if err != nil {
		return errors.WithStack(err)
	}
	response := resp.(protocal.BranchReportResponse)
	if response.ResultCode == protocal.ResultCodeFailed {
		return response.GetError()
	}
	return nil
}

func (resourceManager AbstractResourceManager) LockQuery(ctx *context.RootContext, branchType meta.BranchType, resourceID string, xid string,
	lockKeys string) (bool, error) {
	return false, nil
}

func (resourceManager AbstractResourceManager) handleRegisterRM() {
	for {
		serverAddress := <-resourceManager.RpcClient.GettySessionOnOpenChannel
		resourceManager.doRegisterResource(serverAddress)
	}
}

func (resourceManager AbstractResourceManager) doRegisterResource(serverAddress string) {
	if resourceManager.ResourceCache == nil || len(resourceManager.ResourceCache) == 0 {
		return
	}
	message := protocal.RegisterRMRequest{
		AbstractIdentifyRequest: protocal.AbstractIdentifyRequest{
			Version:                 config.GetClientConfig().StarfishVersion,
			ApplicationID:           config.GetClientConfig().ApplicationID,
			TransactionServiceGroup: config.GetClientConfig().TransactionServiceGroup,
		},
		ResourceIDs: resourceManager.getMergedResourceKeys(),
	}

	resourceManager.RpcClient.RegisterResource(serverAddress, message)
}

func (resourceManager AbstractResourceManager) getMergedResourceKeys() string {
	var builder strings.Builder
	if resourceManager.ResourceCache != nil && len(resourceManager.ResourceCache) > 0 {
		for key := range resourceManager.ResourceCache {
			builder.WriteString(key)
			builder.WriteString(DBKEYS_SPLIT_CHAR)
		}
		resourceKeys := builder.String()
		resourceKeys = resourceKeys[:len(resourceKeys)-1]
		return resourceKeys
	}
	return ""
}
