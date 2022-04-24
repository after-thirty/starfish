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
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/client/rm"
	"github.com/transaction-mesh/starfish/pkg/client/tm"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

type TransactionCoordinatorInbound interface {
	tm.TransactionManager
	rm.ResourceManagerOutbound
}

type TransactionCoordinatorOutbound interface {
	// Commit a branch transaction.
	branchCommit(globalSession *session.GlobalSession, branchSession *session.BranchSession) (meta.BranchStatus, error)

	// Rollback a branch transaction.
	branchRollback(globalSession *session.GlobalSession, branchSession *session.BranchSession) (meta.BranchStatus, error)
}

type TransactionCoordinator interface {
	TransactionCoordinatorInbound
	TransactionCoordinatorOutbound

	// Do global commit.
	doGlobalCommit(globalSession *session.GlobalSession, retrying bool) (bool, error)

	// Do global rollback.
	doGlobalRollback(globalSession *session.GlobalSession, retrying bool) (bool, error)

	// Do global report.
	doGlobalReport(globalSession *session.GlobalSession, xid string, param meta.GlobalStatus) error
}
