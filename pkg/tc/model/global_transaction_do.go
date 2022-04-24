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

// GlobalTransactionDO for persist GlobalTransaction.
type GlobalTransactionDO struct {
	XID string `xorm:"xid"`

	TransactionID int64 `xorm:"transaction_id"`

	Status int32 `xorm:"status"`

	ApplicationID string `xorm:"application_id"`

	TransactionServiceGroup string `xorm:"transaction_service_group"`

	TransactionName string `xorm:"transaction_name"`

	Timeout int32 `xorm:"timeout"`

	BeginTime int64 `xorm:"begin_time"`

	ApplicationData []byte `xorm:"application_data"`

	GmtCreate time.Time `xorm:"gmt_create"`

	GmtModified time.Time `xorm:"gmt_modified"`
}
