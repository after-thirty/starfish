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

package session

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/util/uuid"
)

func TestBranchSession_Encode_Decode(t *testing.T) {
	bs := branchSessionProvider()
	result, _ := bs.Encode()
	newBs := &BranchSession{}
	newBs.Decode(result)

	assert.Equal(t, bs.TransactionID, newBs.TransactionID)
	assert.Equal(t, bs.BranchID, newBs.BranchID)
	assert.Equal(t, bs.ResourceID, newBs.ResourceID)
	assert.Equal(t, bs.LockKey, newBs.LockKey)
	assert.Equal(t, bs.ClientID, newBs.ClientID)
	assert.Equal(t, bs.ApplicationData, newBs.ApplicationData)
}

func branchSessionProvider() *BranchSession {
	bs := NewBranchSession(
		WithBsTransactionID(uuid.NextID()),
		WithBsBranchID(1),
		WithBsResourceGroupID("my_test_tx_group"),
		WithBsResourceID("tb_1"),
		WithBsLockKey("t_1"),
		WithBsBranchType(meta.BranchTypeAT),
		WithBsStatus(meta.BranchStatusUnknown),
		WithBsClientID("c1"),
		WithBsApplicationData([]byte("{\"data\":\"test\"}")),
	)

	return bs
}
