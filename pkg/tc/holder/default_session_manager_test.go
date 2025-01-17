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

package holder

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/common"
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

func TestDefaultSessionManager_AddGlobalSession_RemoveGlobalSession(t *testing.T) {
	gs := globalSessionProvider(t)

	sessionManager := NewDefaultSessionManager("default")
	sessionManager.AddGlobalSession(gs)
	sessionManager.RemoveGlobalSession(gs)
}

func TestDefaultSessionManager_FindGlobalSession(t *testing.T) {
	gs := globalSessionProvider(t)
	sessionManager := NewDefaultSessionManager("default")
	sessionManager.AddGlobalSession(gs)
	expected := sessionManager.FindGlobalSession(gs.XID)

	assert.NotNil(t, expected)
	assert.Equal(t, gs.TransactionID, expected.TransactionID)
	assert.Equal(t, gs.ApplicationID, expected.ApplicationID)
	assert.Equal(t, gs.TransactionServiceGroup, expected.TransactionServiceGroup)
	assert.Equal(t, gs.TransactionName, expected.TransactionName)
	assert.Equal(t, gs.Status, expected.Status)

	sessionManager.RemoveGlobalSession(gs)
}

func globalSessionsProvider() []*session.GlobalSession {
	common.Init("127.0.0.1", 9876)

	result := make([]*session.GlobalSession, 0)
	gs1 := session.NewGlobalSession(
		session.WithGsApplicationID("demo-cmd"),
		session.WithGsTransactionServiceGroup("my_test_tx_group"),
		session.WithGsTransactionName("test"),
		session.WithGsTimeout(6000),
	)

	gs2 := session.NewGlobalSession(
		session.WithGsApplicationID("demo-cmd"),
		session.WithGsTransactionServiceGroup("my_test_tx_group"),
		session.WithGsTransactionName("test"),
		session.WithGsTimeout(6000),
	)

	result = append(result, gs1)
	result = append(result, gs2)
	return result
}

func globalSessionProvider(t *testing.T) *session.GlobalSession {
	common.Init("127.0.0.1", 9876)

	gs := session.NewGlobalSession(
		session.WithGsApplicationID("demo-cmd"),
		session.WithGsTransactionServiceGroup("my_test_tx_group"),
		session.WithGsTransactionName("test"),
		session.WithGsTimeout(6000),
	)
	return gs
}

func branchSessionProvider(globalSession *session.GlobalSession) *session.BranchSession {
	bs := session.NewBranchSession(
		session.WithBsTransactionID(globalSession.TransactionID),
		session.WithBsBranchID(1),
		session.WithBsResourceGroupID("my_test_tx_group"),
		session.WithBsResourceID("tb_1"),
		session.WithBsLockKey("t_1"),
		session.WithBsBranchType(meta.BranchTypeAT),
		session.WithBsApplicationData([]byte("{\"data\":\"test\"}")),
	)

	return bs
}
