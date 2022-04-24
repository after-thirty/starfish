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

package extension

import (
	"sync"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/config"
	"github.com/transaction-mesh/starfish/pkg/base/config_center"
)

var (
	configCentersMu sync.RWMutex
	configCenters   = make(map[string]func(conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error))
)

// SetConfigCenter set config center
func SetConfigCenter(name string, v func(conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error)) {
	configCentersMu.Lock()
	defer configCentersMu.Unlock()
	if v == nil {
		panic("configCenter: Register  configCenter is nil")
	}
	if _, dup := configCenters[name]; dup {
		panic("configCenter: Register called twice for configCenter " + name)
	}
	configCenters[name] = v
}

// GetConfigCenter get config center
func GetConfigCenter(name string, conf *config.ConfigCenterConfig) (config_center.DynamicConfigurationFactory, error) {
	configCentersMu.RLock()
	configCenter := configCenters[name]
	configCentersMu.RUnlock()
	if configCenter == nil {
		return nil, errors.Errorf("config center for " + name + " is not existing, make sure you have import the package.")
	}
	return configCenter(conf)
}
