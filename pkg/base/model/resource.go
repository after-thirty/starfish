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
	"github.com/transaction-mesh/starfish/pkg/base/meta"
)

type IResource interface {
	/**
	 * Get the resource group id.
	 * e.g. master and slave data-source should be with the same resource group id.
	 *
	 * @return resource group id.
	 */
	GetResourceGroupID() string

	/**
	 * Get the resource id.
	 * e.g. url of a data-source could be the id of the db data-source resource.
	 *
	 * @return resource id.
	 */
	GetResourceID() string

	/**
	 * get resource type, BranchType_AT, BranchType_TCC, BranchType_SAGA and XA
	 *
	 * @return
	 */
	GetBranchType() meta.BranchType
}
