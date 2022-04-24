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

// GettySessionParam getty session param
type GettySessionParam struct {
	CompressEncoding bool          `default:"false" yaml:"compress_encoding" json:"compress_encoding,omitempty"`
	TCPNoDelay       bool          `default:"true" yaml:"tcp_no_delay" json:"tcp_no_delay,omitempty"`
	TCPKeepAlive     bool          `default:"true" yaml:"tcp_keep_alive" json:"tcp_keep_alive,omitempty"`
	KeepAlivePeriod  time.Duration `default:"180s" yaml:"keep_alive_period" json:"keep_alive_period,omitempty"`
	TCPRBufSize      int           `default:"262144" yaml:"tcp_r_buf_size" json:"tcp_r_buf_size,omitempty"`
	TCPWBufSize      int           `default:"65536" yaml:"tcp_w_buf_size" json:"tcp_w_buf_size,omitempty"`
	TCPReadTimeout   time.Duration `default:"1s" yaml:"tcp_read_timeout" json:"tcp_read_timeout,omitempty"`
	TCPWriteTimeout  time.Duration `default:"5s" yaml:"tcp_write_timeout" json:"tcp_write_timeout,omitempty"`
	WaitTimeout      time.Duration `default:"7s" yaml:"wait_timeout" json:"wait_timeout,omitempty"`
	MaxMsgLen        int           `default:"4096" yaml:"max_msg_len" json:"max_msg_len,omitempty"`
	SessionName      string        `default:"rpc" yaml:"session_name" json:"session_name,omitempty"`
}
