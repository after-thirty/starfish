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

// BranchTrype
type BranchType byte

const (
	// The At.
	// BranchType_AT Branch
	BranchTypeAT BranchType = iota

	//The BranchType_TCC.
	BranchTypeTCC

	// The BranchType_SAGA.
	BranchTypeSAGA
)

// String string of branch type
func (t BranchType) String() string {
	switch t {
	case BranchTypeAT:
		return "AT"
	case BranchTypeTCC:
		return "TCC"
	case BranchTypeSAGA:
		return "SAGA"
	default:
		return fmt.Sprintf("%d", t)
	}
}

// ValueOfBranchType value of branch type
func ValueOfBranchType(branchType string) BranchType {
	switch branchType {
	case "AT":
		return BranchTypeAT
	case "TCC":
		return BranchTypeTCC
	case "SAGA":
		return BranchTypeSAGA
	default:
		return 0
	}
}
