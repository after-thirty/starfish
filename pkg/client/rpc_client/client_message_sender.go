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

package rpc_client

import (
	"time"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
)

type ClientMessageSender interface {

	// Send msg with response object.
	SendMsgWithResponse(msg interface{}) (interface{}, error)

	// Send msg with response object.
	SendMsgWithResponseAndTimeout(msg interface{}, timeout time.Duration) (interface{}, error)

	// Send response.
	SendResponse(request protocal.RpcMessage, serverAddress string, msg interface{})
}
