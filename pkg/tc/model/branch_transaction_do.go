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

package model

import (
	"time"
)

// BranchTransactionDO for persist BranchTransaction.
type BranchTransactionDO struct {
	XID string `xorm:"xid"`

	TransactionID int64 `xorm:"transaction_id"`

	BranchID int64 `xorm:"branch_id"`

	ResourceGroupID string `xorm:"resource_group_id"`

	ResourceID string `xorm:"resource_id"`

	BranchType string `xorm:"branch_type"`

	Status int32 `xorm:"status"`

	ClientID string `xorm:"client_id"`

	ApplicationData []byte `xorm:"application_data"`

	GmtCreate time.Time `xorm:"gmt_create"`

	GmtModified time.Time `xorm:"gmt_modified"`
}
