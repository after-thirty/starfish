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

package event

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
)

const (
	RoleTC = "tc"
	RoleTM = "tm"
	RoleRM = "rm"
)

type GlobalTransactionEvent struct {
	id        int64
	role      string
	name      string
	beginTime int64
	endTime   int64
	status    meta.GlobalStatus
}

func NewGlobalTransactionEvent(id int64, role string, name string, beginTime int64, endTime int64, status meta.GlobalStatus) GlobalTransactionEvent {
	return GlobalTransactionEvent{
		id,
		role,
		name,
		beginTime,
		endTime,
		status,
	}
}

func (event GlobalTransactionEvent) GetID() int64 { return event.id }

func (event GlobalTransactionEvent) GetRole() string { return event.role }

func (event GlobalTransactionEvent) GetName() string { return event.name }

func (event GlobalTransactionEvent) GetBeginTime() int64 { return event.beginTime }

func (event GlobalTransactionEvent) GetEndTime() int64 { return event.endTime }

func (event GlobalTransactionEvent) GetStatus() meta.GlobalStatus { return event.status }
