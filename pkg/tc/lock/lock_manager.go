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
	"sync"
)

import (
	"github.com/transaction-mesh/starfish/pkg/tc/config"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

var lockManager LockManager

type LockManager interface {
	// AcquireLock Acquire lock boolean.
	AcquireLock(branchSession *session.BranchSession) bool

	// ReleaseLock Unlock boolean.
	ReleaseLock(branchSession *session.BranchSession) bool

	// GlobalSession 是没有锁的，所有的锁都在 BranchSession 上，因为
	// BranchSession 才持有资源，释放 GlobalSession 锁是指释放它所有
	// 的 BranchSession 上的锁.
	// ReleaseGlobalSessionLock Unlock boolean.
	ReleaseGlobalSessionLock(globalSession *session.GlobalSession) bool

	// IsLockable Is lockable boolean.
	IsLockable(xid string, resourceID string, lockKey string) bool

	// CleanAllLocks Clean all locks.
	CleanAllLocks()

	GetLockKeyCount() int64
}

func Init() {
	if config.GetStoreConfig().StoreMode == "db" {
		lockStore := &LockStoreDataBaseDao{engine: config.GetStoreConfig().DBStoreConfig.Engine}
		lockManager = &DataBaseLocker{LockStore: lockStore}
	} else {
		lockManager = &MemoryLocker{
			LockMap:      &sync.Map{},
			BucketHolder: &sync.Map{},
		}
	}
}

func GetLockManager() LockManager {
	return lockManager
}
