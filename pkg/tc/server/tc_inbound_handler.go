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
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
)

type TCInboundHandler interface {
	doGlobalBegin(request protocal.GlobalBeginRequest, ctx RpcContext) protocal.GlobalBeginResponse
	doGlobalStatus(request protocal.GlobalStatusRequest, ctx RpcContext) protocal.GlobalStatusResponse
	doGlobalReport(request protocal.GlobalReportRequest, ctx RpcContext) protocal.GlobalReportResponse
	doGlobalCommit(request protocal.GlobalCommitRequest, ctx RpcContext) protocal.GlobalCommitResponse
	doGlobalRollback(request protocal.GlobalRollbackRequest, ctx RpcContext) protocal.GlobalRollbackResponse
	doBranchRegister(request protocal.BranchRegisterRequest, ctx RpcContext) protocal.BranchRegisterResponse
	doBranchReport(request protocal.BranchReportRequest, ctx RpcContext) protocal.BranchReportResponse
	doLockCheck(request protocal.GlobalLockQueryRequest, ctx RpcContext) protocal.GlobalLockQueryResponse
}
