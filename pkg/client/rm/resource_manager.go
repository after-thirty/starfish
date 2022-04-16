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
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/base/model"
)

type ResourceManagerInbound interface {
	// Commit a branch transaction.
	BranchCommit(branchType meta.BranchType, xid string, branchID int64, resourceID string, applicationData []byte) (meta.BranchStatus, error)

	// Rollback a branch transaction.
	BranchRollback(branchType meta.BranchType, xid string, branchID int64, resourceID string, applicationData []byte) (meta.BranchStatus, error)
}

type ResourceManagerOutbound interface {
	// Branch register long.
	BranchRegister(branchType meta.BranchType, resourceID string, clientID string, xid string, applicationData []byte, lockKeys string) (int64, error)

	// Branch report.
	BranchReport(branchType meta.BranchType, xid string, branchID int64, status meta.BranchStatus, applicationData []byte) error

	// Lock query boolean.
	LockQuery(branchType meta.BranchType, resourceID string, xid string, lockKeys string) (bool, error)
}

type ResourceManager interface {
	ResourceManagerInbound
	ResourceManagerOutbound

	// Register a Resource to be managed by Resource Manager.
	RegisterResource(resource model.IResource)

	// Unregister a Resource from the Resource Manager.
	UnregisterResource(resource model.IResource)

	// Get all resources managed by this manager.
	GetManagedResources() map[string]model.IResource

	// Get the BranchType.
	GetBranchType() meta.BranchType
}
