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

package constant

const (
	// NacosDefaultGroup
	NacosDefaultGroup = "SEATA_GROUP"
	// NacosDefaultDataID
	NacosDefaultDataID = "starfish"
	// NacosKey
	NacosKey = "nacos"
	// FileKey
	FileKey = "file"

	Etcdv3Key                = "etcdv3"
	Etcdv3RegistryPrefix     = "etcdv3-starfish-" // according to starfish java version
	Etcdv3LeaseRenewInterval = 5               // according to starfish java version
	Etcdv3LeaseTtl           = 10              // according to starfish java version
	Etcdv3LeaseTtlCritical   = 6               // according to starfish java version
)
