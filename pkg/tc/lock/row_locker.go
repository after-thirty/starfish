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

package lock

import (
	"strings"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/common"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

const LOCK_SPLIT = "^^^"

type RowLock struct {
	XID string

	TransactionID int64

	BranchID int64

	ResourceID string

	TableName string

	Pk string

	RowKey string

	Feature string
}

func collectRowLocksByBranchSession(branchSession *session.BranchSession) []*RowLock {
	if branchSession == nil || branchSession.LockKey == "" {
		return nil
	}
	return collectRowLocks(branchSession.LockKey, branchSession.ResourceID, branchSession.XID, branchSession.TransactionID, branchSession.BranchID)
}

func collectRowLocksByLockKeyResourceIDAndXID(lockKey string,
	resourceID string,
	xid string) []*RowLock {

	return collectRowLocks(lockKey, resourceID, xid, common.GetTransactionID(xid), 0)
}

func collectRowLocks(lockKey string,
	resourceID string,
	xid string,
	transactionID int64,
	branchID int64) []*RowLock {
	var locks = make([]*RowLock, 0)
	tableGroupedLockKeys := strings.Split(lockKey, ";")
	for _, tableGroupedLockKey := range tableGroupedLockKeys {
		if tableGroupedLockKey != "" {
			idx := strings.Index(tableGroupedLockKey, ":")
			if idx < 0 {
				return nil
			}

			tableName := tableGroupedLockKey[0:idx]
			mergedPKs := tableGroupedLockKey[idx+1:]

			if mergedPKs == "" {
				return nil
			}

			pks := strings.Split(mergedPKs, ",")
			if len(pks) == 0 {
				return nil
			}

			for _, pk := range pks {
				if pk != "" {
					rowLock := &RowLock{
						XID:           xid,
						TransactionID: transactionID,
						BranchID:      branchID,
						ResourceID:    resourceID,
						TableName:     tableName,
						Pk:            pk,
					}
					locks = append(locks, rowLock)
				}
			}
		}
	}
	return locks
}
