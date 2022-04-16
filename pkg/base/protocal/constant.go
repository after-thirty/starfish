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

package protocal

var MAGIC_CODE_BYTES = [2]byte{0xda, 0xda}

const (
	// version
	VERSION = 1

	// MaxFrameLength max frame length
	MaxFrameLength = 8 * 1024 * 1024

	// V1HeadLength v1 head length
	V1HeadLength = 16

	// MSGTypeRequest request message type
	MSGTypeRequest = 0

	// MSGTypeResponse response message type
	MSGTypeResponse = 1

	// MSGTypeRequestOneway request one way
	MSGTypeRequestOneway = 2

	// MSGTypeHeartbeatRequest heart beat request
	MSGTypeHeartbeatRequest = 3

	// MSGTypeHeartbeatResponse heart beat response
	MSGTypeHeartbeatResponse = 4
)
