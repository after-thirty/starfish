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

package server

import (
	getty "github.com/apache/dubbo-getty"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

func (coordinator *DefaultCoordinator) OnTrxMessage(rpcMessage protocal.RpcMessage, session getty.Session) {
	rpcContext := SessionManager.GetContextFromIdentified(session)
	log.Debugf("server received:%v,clientIp:%s,vgroup:%s", rpcMessage.Body, session.RemoteAddr(), rpcContext.TransactionServiceGroup)

	warpMessage, isWarpMessage := rpcMessage.Body.(protocal.MergedWarpMessage)
	if isWarpMessage {
		resultMessage := protocal.MergeResultMessage{Msgs: make([]protocal.MessageTypeAware, 0)}
		for _, msg := range warpMessage.Msgs {
			resp := coordinator.handleTrxMessage(msg, *rpcContext)
			resultMessage.Msgs = append(resultMessage.Msgs, resp)
		}
		coordinator.SendResponse(rpcMessage, session, resultMessage)
	} else {
		message := rpcMessage.Body.(protocal.MessageTypeAware)
		resp := coordinator.handleTrxMessage(message, *rpcContext)
		coordinator.SendResponse(rpcMessage, session, resp)
	}
}

func (coordinator *DefaultCoordinator) handleTrxMessage(msg protocal.MessageTypeAware, ctx RpcContext) protocal.MessageTypeAware {
	switch msg.GetTypeCode() {
	case protocal.TypeGlobalBegin:
		req := msg.(protocal.GlobalBeginRequest)
		resp := coordinator.doGlobalBegin(req, ctx)
		return resp
	case protocal.TypeGlobalStatus:
		req := msg.(protocal.GlobalStatusRequest)
		resp := coordinator.doGlobalStatus(req, ctx)
		return resp
	case protocal.TypeGlobalReport:
		req := msg.(protocal.GlobalReportRequest)
		resp := coordinator.doGlobalReport(req, ctx)
		return resp
	case protocal.TypeGlobalCommit:
		req := msg.(protocal.GlobalCommitRequest)
		resp := coordinator.doGlobalCommit(req, ctx)
		return resp
	case protocal.TypeGlobalRollback:
		req := msg.(protocal.GlobalRollbackRequest)
		resp := coordinator.doGlobalRollback(req, ctx)
		return resp
	case protocal.TypeBranchRegister:
		req := msg.(protocal.BranchRegisterRequest)
		resp := coordinator.doBranchRegister(req, ctx)
		return resp
	case protocal.TypeBranchStatusReport:
		req := msg.(protocal.BranchReportRequest)
		resp := coordinator.doBranchReport(req, ctx)
		return resp
	case protocal.TypeGlobalLockQuery:
		req := msg.(protocal.GlobalLockQueryRequest)
		resp := coordinator.doLockCheck(req, ctx)
		return resp
	default:
		return nil
	}
}

func (coordinator *DefaultCoordinator) OnRegRmMessage(rpcMessage protocal.RpcMessage, session getty.Session) {
	message := rpcMessage.Body.(protocal.RegisterRMRequest)

	//version things
	SessionManager.RegisterRmGettySession(message, session)
	log.Debugf("checkAuth for rpc_client:%s,vgroup:%s,applicationID:%s", session.RemoteAddr(), message.TransactionServiceGroup, message.ApplicationID)

	coordinator.SendResponse(rpcMessage, session, protocal.RegisterRMResponse{AbstractIdentifyResponse: protocal.AbstractIdentifyResponse{Identified: true}})
}

func (coordinator *DefaultCoordinator) OnRegTmMessage(rpcMessage protocal.RpcMessage, session getty.Session) {
	message := rpcMessage.Body.(protocal.RegisterTMRequest)

	//version things
	SessionManager.RegisterTmGettySession(message, session)
	log.Debugf("checkAuth for rpc_client:%s,vgroup:%s,applicationID:%s", session.RemoteAddr(), message.TransactionServiceGroup, message.ApplicationID)

	coordinator.SendResponse(rpcMessage, session, protocal.RegisterTMResponse{AbstractIdentifyResponse: protocal.AbstractIdentifyResponse{Identified: true}})
}

func (coordinator *DefaultCoordinator) OnCheckMessage(rpcMessage protocal.RpcMessage, session getty.Session) {
	coordinator.SendResponse(rpcMessage, session, protocal.HeartBeatMessagePong)
	log.Debugf("received PING from %s", session.RemoteAddr())
}
