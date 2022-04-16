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
	"github.com/transaction-mesh/starfish/pkg/base/registry"
)

var (
	registriesMu sync.RWMutex
	registries   = make(map[string]func() (registry.Registry, error))
)

// SetRegistry sets the registry extension with @name
func SetRegistry(name string, v func() (registry.Registry, error)) {
	//写加锁，参考sql.go
	registriesMu.Lock()
	defer registriesMu.Unlock()
	if v == nil {
		panic("registry: Register  v is nil")
	}
	if _, dup := registries[name]; dup {
		panic("registry: Register called twice for registry " + name)
	}
	registries[name] = v
}

// GetRegistry finds the registry extension with @name
func GetRegistry(name string) (registry.Registry, error) {
	registriesMu.RLock()
	registry := registries[name]
	registriesMu.RUnlock()
	if registry == nil {
		return nil, errors.Errorf("registry for " + name + " is not existing, make sure you have import the package.")
	}
	return registry()
}
