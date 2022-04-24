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

package config

import (
	"time"
)

import (
	config2 "github.com/transaction-mesh/starfish/pkg/base/config"
)

// GettyConfig
//Config holds supported types by the multiconfig package
type GettyConfig struct {
	ReconnectInterval int `default:"0" yaml:"reconnect_interval" json:"reconnect_interval,omitempty"`
	// getty_session pool
	ConnectionNum int `default:"16" yaml:"connection_number" json:"connection_number,omitempty"`

	// heartbeat
	HeartbeatPeriod time.Duration `default:"15s" yaml:"heartbeat_period" json:"heartbeat_period,omitempty"`

	// getty_session tcp parameters
	GettySessionParam config2.GettySessionParam `required:"true" yaml:"getty_session_param" json:"getty_session_param,omitempty"`
}

// GetDefaultGettyConfig ...
func GetDefaultGettyConfig() GettyConfig {
	return GettyConfig{
		ReconnectInterval: 0,
		ConnectionNum:     1,
		HeartbeatPeriod:   10 * time.Second,
		GettySessionParam: config2.GettySessionParam{
			CompressEncoding: false,
			TCPNoDelay:       true,
			TCPKeepAlive:     true,
			KeepAlivePeriod:  180 * time.Second,
			TCPRBufSize:      262144,
			TCPWBufSize:      65536,
			TCPReadTimeout:   time.Second,
			TCPWriteTimeout:  5 * time.Second,
			WaitTimeout:      time.Second,
			MaxMsgLen:        4096,
			SessionName:      "rpc_client",
		},
	}
}
