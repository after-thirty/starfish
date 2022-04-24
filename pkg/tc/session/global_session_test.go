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

package session

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestGlobalSession_Encode_Decode(t *testing.T) {
	gs := globalSessionProvider()
	result, err := gs.Encode()
	assert.NoError(t, err, "Encode() should success")

	newGs := &GlobalSession{}
	newGs.Decode(result)

	assert.Equal(t, newGs.TransactionID, gs.TransactionID)
	assert.Equal(t, newGs.Timeout, gs.Timeout)
	assert.Equal(t, newGs.ApplicationID, gs.ApplicationID)
	assert.Equal(t, newGs.TransactionServiceGroup, gs.TransactionServiceGroup)
	assert.Equal(t, newGs.TransactionName, gs.TransactionName)
}

func globalSessionProvider() *GlobalSession {
	gs := NewGlobalSession(
		WithGsApplicationID("demo-cmd"),
		WithGsTransactionServiceGroup("my_test_tx_group"),
		WithGsTransactionName("test"),
		WithGsTimeout(6000),
		WithGsActive(true),
	)

	return gs
}
