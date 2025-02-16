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

package tm

import (
	"github.com/pkg/errors"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
	"github.com/transaction-mesh/starfish/pkg/client/rpc_client"
)

type DefaultTransactionManager struct {
	rpcClient *rpc_client.RpcRemoteClient
}

func (manager DefaultTransactionManager) Begin(applicationID string, transactionServiceGroup string, name string, timeout int32) (string, error) {
	request := protocal.GlobalBeginRequest{
		Timeout:         timeout,
		TransactionName: name,
	}
	resp, err := manager.syncCall(request)
	if err != nil {
		return "", errors.WithStack(err)
	}
	response := resp.(protocal.GlobalBeginResponse)
	return response.Xid, nil
}

func (manager DefaultTransactionManager) Commit(xid string) (meta.GlobalStatus, error) {
	globalCommit := protocal.GlobalCommitRequest{AbstractGlobalEndRequest: protocal.AbstractGlobalEndRequest{XID: xid}}
	resp, err := manager.syncCall(globalCommit)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	response := resp.(protocal.GlobalCommitResponse)
	return response.GlobalStatus, nil
}

func (manager DefaultTransactionManager) Rollback(xid string) (meta.GlobalStatus, error) {
	globalRollback := protocal.GlobalRollbackRequest{AbstractGlobalEndRequest: protocal.AbstractGlobalEndRequest{XID: xid}}
	resp, err := manager.syncCall(globalRollback)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	response := resp.(protocal.GlobalRollbackResponse)
	return response.GlobalStatus, nil
}

func (manager DefaultTransactionManager) GetStatus(xid string) (meta.GlobalStatus, error) {
	queryGlobalStatus := protocal.GlobalStatusRequest{AbstractGlobalEndRequest: protocal.AbstractGlobalEndRequest{XID: xid}}
	resp, err := manager.syncCall(queryGlobalStatus)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	response := resp.(protocal.GlobalStatusResponse)
	return response.GlobalStatus, nil
}

func (manager DefaultTransactionManager) GlobalReport(xid string, globalStatus meta.GlobalStatus) (meta.GlobalStatus, error) {
	globalReport := protocal.GlobalReportRequest{
		AbstractGlobalEndRequest: protocal.AbstractGlobalEndRequest{XID: xid},
		GlobalStatus:             globalStatus,
	}
	resp, err := manager.syncCall(globalReport)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	response := resp.(protocal.GlobalReportResponse)
	return response.GlobalStatus, nil
}

func (manager DefaultTransactionManager) syncCall(request interface{}) (interface{}, error) {
	return manager.rpcClient.SendMsgWithResponse(request)
}
