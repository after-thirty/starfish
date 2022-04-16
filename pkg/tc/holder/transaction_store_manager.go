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
	"fmt"
)

import (
	"github.com/transaction-mesh/starfish/pkg/tc/model"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

type LogOperation byte

const (
	LogOperationGlobalAdd LogOperation = iota
	/**
	 * Global update log operation.
	 */
	LogOperationGlobalUpdate
	/**
	 * Global remove log operation.
	 */
	LogOperationGlobalRemove
	/**
	 * Branch add log operation.
	 */
	LogOperationBranchAdd
	/**
	 * Branch update log operation.
	 */
	LogOperationBranchUpdate
	/**
	 * Branch remove log operation.
	 */
	LogOperationBranchRemove
)

func (t LogOperation) String() string {
	switch t {
	case LogOperationGlobalAdd:
		return "GlobalAdd"
	case LogOperationGlobalUpdate:
		return "GlobalUpdate"
	case LogOperationGlobalRemove:
		return "GlobalRemove"
	case LogOperationBranchAdd:
		return "BranchAdd"
	case LogOperationBranchUpdate:
		return "BranchUpdate"
	case LogOperationBranchRemove:
		return "BranchRemove"
	default:
		return fmt.Sprintf("%d", t)
	}
}

type Reloadable interface {
	// Reload states.
	Reload()
}

type TransactionStoreManager interface {
	// Write session boolean.
	WriteSession(logOperation LogOperation, session session.SessionStorable) bool

	// Read global session global session.
	ReadSession(xid string) *session.GlobalSession

	// Read session global session.
	ReadSessionWithBranchSessions(xid string, withBranchSessions bool) *session.GlobalSession

	// Read session by status list.
	ReadSessionWithSessionCondition(sessionCondition model.SessionCondition) []*session.GlobalSession

	// Shutdown.
	Shutdown()
}

type AbstractTransactionStoreManager struct {
}

func (transactionStoreManager *AbstractTransactionStoreManager) WriteSession(logOperation LogOperation, session session.SessionStorable) bool {
	return true
}

func (transactionStoreManager *AbstractTransactionStoreManager) ReadSession(xid string) *session.GlobalSession {
	return nil
}

func (transactionStoreManager *AbstractTransactionStoreManager) ReadSessionWithBranchSessions(xid string, withBranchSessions bool) *session.GlobalSession {
	return nil
}

func (transactionStoreManager *AbstractTransactionStoreManager) ReadSessionWithSessionCondition(sessionCondition model.SessionCondition) []*session.GlobalSession {
	return nil
}

func (transactionStoreManager *AbstractTransactionStoreManager) Shutdown() {

}

func (transactionStoreManager *AbstractTransactionStoreManager) GetCurrentMaxSessionID() int64 {
	return 0
}
