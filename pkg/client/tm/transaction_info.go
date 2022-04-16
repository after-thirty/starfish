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

package tm

import (
	"fmt"
)

type Propagation byte

const (
	/**
	 * The REQUIRED.
	 */
	REQUIRED Propagation = iota

	/**
	 * The REQUIRES_NEW.
	 */
	REQUIRES_NEW

	/**
	 * The NOT_SUPPORTED
	 */
	NOT_SUPPORTED

	/**
	 * The SUPPORTS
	 */
	SUPPORTS

	/**
	 * The NEVER
	 */
	NEVER

	/**
	 * The MANDATORY
	 */
	MANDATORY
)

func (t Propagation) String() string {
	switch t {
	case REQUIRED:
		return "REQUIRED"
	case REQUIRES_NEW:
		return "REQUIRES_NEW"
	case NOT_SUPPORTED:
		return "NOT_SUPPORTED"
	case SUPPORTS:
		return "SUPPORTS"
	case NEVER:
		return "NEVER"
	case MANDATORY:
		return "MANDATORY"
	default:
		return fmt.Sprintf("%d", t)
	}
}

type TransactionInfo struct {
	TimeOut     int32
	Name        string
	Propagation Propagation
}
