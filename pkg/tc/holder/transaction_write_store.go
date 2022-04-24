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
	"github.com/transaction-mesh/starfish/pkg/tc/session"
)

type TransactionWriteStore struct {
	SessionRequest session.SessionStorable
	LogOperation   LogOperation
}

func (transactionWriteStore *TransactionWriteStore) Encode() ([]byte, error) {
	bySessionRequest, err := transactionWriteStore.SessionRequest.Encode()
	if err != nil {
		return nil, err
	}
	byOpCode := transactionWriteStore.LogOperation

	var result = make([]byte, 0)
	result = append(result, bySessionRequest...)
	result = append(result, byte(byOpCode))
	return result, nil
}

func (transactionWriteStore *TransactionWriteStore) Decode(src []byte) {
	bySessionRequest := src[:len(src)-1]
	byOpCode := src[len(src)-1:]

	transactionWriteStore.LogOperation = LogOperation(byOpCode[0])
	sessionRequest, _ := transactionWriteStore.getSessionInstanceByOperation()
	sessionRequest.Decode(bySessionRequest)
	transactionWriteStore.SessionRequest = sessionRequest
}

func (transactionWriteStore *TransactionWriteStore) getSessionInstanceByOperation() (session.SessionStorable, error) {
	var sessionStorable session.SessionStorable
	switch transactionWriteStore.LogOperation {
	case LogOperationGlobalAdd, LogOperationGlobalUpdate, LogOperationGlobalRemove:
		sessionStorable = session.NewGlobalSession()
	case LogOperationBranchAdd, LogOperationBranchUpdate, LogOperationBranchRemove:
		sessionStorable = session.NewBranchSession()
	default:
		return nil, errors.New("incorrect logOperation.")
	}
	return sessionStorable, nil
}
