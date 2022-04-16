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

package file

import (
	"strings"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/constant"
	"github.com/transaction-mesh/starfish/pkg/base/extension"
	"github.com/transaction-mesh/starfish/pkg/base/registry"
	"github.com/transaction-mesh/starfish/pkg/client/config"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

func init() {
	extension.SetRegistry(constant.FileKey, newFileRegistry)
}

type fileRegistry struct {
}

func (r *fileRegistry) Register(addr *registry.Address) error {
	//文件不需要注册
	log.Info("file register")
	return nil
}

func (r *fileRegistry) UnRegister(addr *registry.Address) error {
	return nil
}
func (r *fileRegistry) Lookup() ([]string, error) {
	addressList := strings.Split(config.GetClientConfig().TransactionServiceGroup, ",")
	return addressList, nil
}
func (r *fileRegistry) Subscribe(notifyListener registry.EventListener) error {
	return nil
}

func (r *fileRegistry) UnSubscribe(notifyListener registry.EventListener) error {
	return nil
}

func (r *fileRegistry) Stop() {
	// TODO: Implement Stop interface
	return
}

// newNacosRegistry will create new instance
func newFileRegistry() (registry.Registry, error) {
	tmpRegistry := &fileRegistry{}
	return tmpRegistry, nil
}
