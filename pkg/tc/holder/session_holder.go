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
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/tc/config"
	"github.com/transaction-mesh/starfish/pkg/tc/lock"
	"github.com/transaction-mesh/starfish/pkg/tc/session"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

var (
	ASYNC_COMMITTING_SESSION_MANAGER_NAME  = "async.commit.data"
	RETRY_COMMITTING_SESSION_MANAGER_NAME  = "retry.commit.data"
	RETRY_ROLLBACKING_SESSION_MANAGER_NAME = "retry.rollback.data"
)

type SessionHolder struct {
	RootSessionManager             SessionManager
	AsyncCommittingSessionManager  SessionManager
	RetryCommittingSessionManager  SessionManager
	RetryRollbackingSessionManager SessionManager
}

var sessionHolder SessionHolder

func Init() {
	if config.GetStoreConfig().StoreMode == "file" {
		sessionHolder = SessionHolder{
			RootSessionManager:             NewFileBasedSessionManager(config.GetStoreConfig().FileStoreConfig),
			AsyncCommittingSessionManager:  NewDefaultSessionManager(ASYNC_COMMITTING_SESSION_MANAGER_NAME),
			RetryCommittingSessionManager:  NewDefaultSessionManager(RETRY_COMMITTING_SESSION_MANAGER_NAME),
			RetryRollbackingSessionManager: NewDefaultSessionManager(RETRY_ROLLBACKING_SESSION_MANAGER_NAME),
		}
		sessionHolder.reload()
	}
	if config.GetStoreConfig().StoreMode == "db" {
		sessionHolder = SessionHolder{
			RootSessionManager:             NewDataBaseSessionManager("", config.GetStoreConfig().DBStoreConfig),
			AsyncCommittingSessionManager:  NewDataBaseSessionManager(ASYNC_COMMITTING_SESSION_MANAGER_NAME, config.GetStoreConfig().DBStoreConfig),
			RetryCommittingSessionManager:  NewDataBaseSessionManager(RETRY_COMMITTING_SESSION_MANAGER_NAME, config.GetStoreConfig().DBStoreConfig),
			RetryRollbackingSessionManager: NewDataBaseSessionManager(RETRY_ROLLBACKING_SESSION_MANAGER_NAME, config.GetStoreConfig().DBStoreConfig),
		}
		sessionHolder.reload()
	}
}

func GetSessionHolder() SessionHolder {
	return sessionHolder
}

func (sessionHolder SessionHolder) FindGlobalSession(xid string) *session.GlobalSession {
	return sessionHolder.FindGlobalSessionWithBranchSessions(xid, true)
}

func (sessionHolder SessionHolder) FindGlobalSessionWithBranchSessions(xid string, withBranchSessions bool) *session.GlobalSession {
	return sessionHolder.RootSessionManager.FindGlobalSessionWithBranchSessions(xid, withBranchSessions)
}

func (sessionHolder SessionHolder) reload() {
	sessionManager, reloadable := sessionHolder.RootSessionManager.(Reloadable)
	if reloadable {
		sessionManager.Reload()

		reloadedSessions := sessionHolder.RootSessionManager.AllSessions()
		if len(reloadedSessions) > 0 {
			for _, globalSession := range reloadedSessions {
				switch globalSession.Status {
				case meta.GlobalStatusUnknown, meta.GlobalStatusCommitted, meta.GlobalStatusCommitFailed, meta.GlobalStatusRolledBack,
					meta.GlobalStatusRollbackFailed, meta.GlobalStatusTimeoutRolledBack, meta.GlobalStatusTimeoutRollbackFailed,
					meta.GlobalStatusFinished:
					log.Errorf("Reloaded Session should NOT be %s", globalSession.Status.String())
				case meta.GlobalStatusAsyncCommitting:
					sessionHolder.AsyncCommittingSessionManager.AddGlobalSession(globalSession)
				default:
					branchSessions := globalSession.GetSortedBranches()
					for _, branchSession := range branchSessions {
						lock.GetLockManager().AcquireLock(branchSession)
					}
					switch globalSession.Status {
					case meta.GlobalStatusCommitting, meta.GlobalStatusCommitRetrying:
						if globalSession.Status == meta.GlobalStatusCommitting {
							globalSession.Status = meta.GlobalStatusCommitRetrying
						}
						sessionHolder.RetryCommittingSessionManager.AddGlobalSession(globalSession)
					case meta.GlobalStatusRollingBack, meta.GlobalStatusRollbackRetrying, meta.GlobalStatusTimeoutRollingBack,
						meta.GlobalStatusTimeoutRollbackRetrying:
						sessionHolder.RetryRollbackingSessionManager.AddGlobalSession(globalSession)
					case meta.GlobalStatusBegin:
						globalSession.Active = true
					default:
						log.Errorf("NOT properly handled %s", globalSession.Status)
					}
				}
			}
		}
	}
}
