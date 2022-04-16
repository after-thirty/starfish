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
	"github.com/pkg/errors"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/tc/config"
	"github.com/transaction-mesh/starfish/pkg/tc/model"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

type DataBaseSessionManager struct {
	TaskName                string
	conf                    config.DBStoreConfig
	TransactionStoreManager TransactionStoreManager
}

func NewDataBaseSessionManager(taskName string, conf config.DBStoreConfig) SessionManager {
	logStore := &LogStoreDataBaseDAO{engine: conf.Engine}
	transactionStoreManager := &DBTransactionStoreManager{
		logQueryLimit: conf.LogQueryLimit,
		LogStore:      logStore,
	}
	sessionManager := &DataBaseSessionManager{
		TaskName:                taskName,
		conf:                    conf,
		TransactionStoreManager: transactionStoreManager,
	}
	return sessionManager
}

func (sessionManager *DataBaseSessionManager) AddGlobalSession(session *session.GlobalSession) error {
	if sessionManager.TaskName == "" {
		ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationGlobalAdd, session)
		if !ret {
			return errors.New("addGlobalSession failed.")
		}
	} else {
		ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationGlobalUpdate, session)
		if !ret {
			return errors.New("addGlobalSession failed.")
		}
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) FindGlobalSession(xid string) *session.GlobalSession {
	return sessionManager.FindGlobalSessionWithBranchSessions(xid, true)
}

func (sessionManager *DataBaseSessionManager) FindGlobalSessionWithBranchSessions(xid string, withBranchSessions bool) *session.GlobalSession {
	return sessionManager.TransactionStoreManager.ReadSessionWithBranchSessions(xid, withBranchSessions)
}

func (sessionManager *DataBaseSessionManager) UpdateGlobalSessionStatus(session *session.GlobalSession, status meta.GlobalStatus) error {
	if sessionManager.TaskName != "" {
		return nil
	}
	session.Status = status
	ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationGlobalUpdate, session)
	if !ret {
		return errors.New("updateGlobalSessionStatus failed.")
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) RemoveGlobalSession(session *session.GlobalSession) error {
	ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationGlobalRemove, session)
	if !ret {
		return errors.New("removeGlobalSession failed.")
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) AddBranchSession(globalSession *session.GlobalSession, session *session.BranchSession) error {
	if sessionManager.TaskName != "" {
		return nil
	}
	ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationBranchAdd, session)
	if !ret {
		return meta.NewTransactionException(nil,
			meta.WithTransactionExceptionCode(meta.TransactionExceptionCodeFailedToAddBranch),
			meta.WithMessage("addBranchSession failed."))
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) UpdateBranchSessionStatus(session *session.BranchSession, status meta.BranchStatus) error {
	if sessionManager.TaskName != "" {
		return nil
	}
	ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationBranchUpdate, session)
	if !ret {
		return errors.New("updateBranchSessionStatus failed.")
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) RemoveBranchSession(globalSession *session.GlobalSession, session *session.BranchSession) error {
	if sessionManager.TaskName != "" {
		return nil
	}
	ret := sessionManager.TransactionStoreManager.WriteSession(LogOperationBranchRemove, session)
	if !ret {
		return errors.New("addBranchSession failed.")
	}
	return nil
}

func (sessionManager *DataBaseSessionManager) AllSessions() []*session.GlobalSession {
	if sessionManager.TaskName == ASYNC_COMMITTING_SESSION_MANAGER_NAME {
		return sessionManager.FindGlobalSessions(model.SessionCondition{
			Statuses: []meta.GlobalStatus{meta.GlobalStatusAsyncCommitting},
		})
	} else if sessionManager.TaskName == RETRY_COMMITTING_SESSION_MANAGER_NAME {
		return sessionManager.FindGlobalSessions(model.SessionCondition{
			Statuses: []meta.GlobalStatus{meta.GlobalStatusCommitRetrying},
		})
	} else if sessionManager.TaskName == RETRY_ROLLBACKING_SESSION_MANAGER_NAME {
		ss := sessionManager.FindGlobalSessions(model.SessionCondition{
			Statuses: []meta.GlobalStatus{meta.GlobalStatusRollbackRetrying,
				meta.GlobalStatusRollingBack,
				meta.GlobalStatusTimeoutRollingBack,
				meta.GlobalStatusTimeoutRollbackRetrying,
			},
		})
		return ss
	} else {
		return sessionManager.FindGlobalSessions(model.SessionCondition{
			Statuses: []meta.GlobalStatus{meta.GlobalStatusUnknown, meta.GlobalStatusBegin,
				meta.GlobalStatusCommitting, meta.GlobalStatusCommitRetrying, meta.GlobalStatusRollingBack,
				meta.GlobalStatusRollbackRetrying, meta.GlobalStatusTimeoutRollingBack, meta.GlobalStatusTimeoutRollbackRetrying,
				meta.GlobalStatusAsyncCommitting,
			},
		})
	}
}

func (sessionManager *DataBaseSessionManager) FindGlobalSessions(condition model.SessionCondition) []*session.GlobalSession {
	return sessionManager.TransactionStoreManager.ReadSessionWithSessionCondition(condition)
}

func (sessionManager *DataBaseSessionManager) Reload() {

}
