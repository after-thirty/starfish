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

package common

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	// IPAddress when started, initialized
	IPAddress string
	// Port when started, initialized
	Port int
)

// Init initialize ip address and port
func Init(ipAddress string, port int) {
	IPAddress = ipAddress
	Port = port
}

// GenerateXID generate xid
func GenerateXID(tranID int64) string {
	return fmt.Sprintf("%s:%d:%d", IPAddress, Port, tranID)
}

// GetTransactionID get transactionID by xid
func GetTransactionID(xid string) int64 {
	if xid == "" {
		return -1
	}

	idx := strings.LastIndex(xid, ":")
	if len(xid) == idx+1 {
		return -1
	}
	tranID, _ := strconv.ParseInt(xid[idx+1:], 10, 64)
	return tranID
}
