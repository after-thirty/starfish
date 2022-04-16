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

package meta

import (
	"fmt"
)

// BranchStatus
type BranchStatus byte

const (
	// The BranchStatus_Unknown.
	// description:BranchStatus_Unknown branch status.
	BranchStatusUnknown BranchStatus = iota

	// The BranchStatus_Registered.
	// description:BranchStatus_Registered to TC.
	BranchStatusRegistered

	// The Phase one done.
	// description:Branch logic is successfully done at phase one.
	BranchStatusPhaseOneDone

	// The Phase one failed.
	// description:Branch logic is failed at phase one.
	BranchStatusPhaseOneFailed

	// The Phase one timeout.
	// description:Branch logic is NOT reported for a timeout.
	BranchStatusPhaseOneTimeout

	// The Phase two committed.
	// description:Commit logic is successfully done at phase two.
	BranchStatusPhaseTwoCommitted

	// The Phase two commit failed retryable.
	// description:Commit logic is failed but retryable.
	BranchStatusPhaseTwoCommitFailedRetryable

	// The Phase two commit failed can not retry.
	// description:Commit logic is failed and NOT retryable.
	BranchStatusPhaseTwoCommitFailedCanNotRetry

	// The Phase two rolled back.
	// description:Rollback logic is successfully done at phase two.
	BranchStatusPhaseTwoRolledBack

	// The Phase two rollback failed retryable.
	// description:Rollback logic is failed but retryable.
	BranchStatusPhaseTwoRollbackFailedRetryable

	// The Phase two rollback failed can not retry.
	// description:Rollback logic is failed but NOT retryable.
	BranchStatusPhaseTwoRollbackFailedCanNotRetry
)

func (s BranchStatus) String() string {
	switch s {
	case BranchStatusUnknown:
		return "Unknown"
	case BranchStatusRegistered:
		return "Registered"
	case BranchStatusPhaseOneDone:
		return "PhaseOneDone"
	case BranchStatusPhaseOneFailed:
		return "PhaseOneFailed"
	case BranchStatusPhaseOneTimeout:
		return "PhaseOneTimeout"
	case BranchStatusPhaseTwoCommitted:
		return "PhaseTwoCommitted"
	case BranchStatusPhaseTwoCommitFailedRetryable:
		return "PhaseTwoCommitFailedRetryable"
	case BranchStatusPhaseTwoCommitFailedCanNotRetry:
		return "CommitFailedCanNotRetry"
	case BranchStatusPhaseTwoRolledBack:
		return "PhaseTwoRolledBack"
	case BranchStatusPhaseTwoRollbackFailedRetryable:
		return "RollbackFailedRetryable"
	case BranchStatusPhaseTwoRollbackFailedCanNotRetry:
		return "RollbackFailedCanNotRetry"
	default:
		return fmt.Sprintf("%d", s)
	}
}
