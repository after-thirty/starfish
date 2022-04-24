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

var config *RegistryConfig

// RegistryConfig registry config
type RegistryConfig struct {
	Mode        string      `yaml:"type" json:"type,omitempty"` //类型
	NacosConfig NacosConfig `yaml:"nacos" json:"nacos,omitempty"`
	EtcdConfig  EtcdConfig  `yaml:"etcdv3" json:"etcdv3"`
}

// NacosConfig nacos config
type NacosConfig struct {
	Application string `yaml:"application" json:"application,omitempty"`
	ServerAddr  string `yaml:"server_addr" json:"server_addr,omitempty"`
	Group       string `default:"SEATA_GROUP" yaml:"group" json:"group,omitempty"`
	Namespace   string `yaml:"namespace" json:"namespace,omitempty"`
	Cluster     string `yaml:"cluster" json:"cluster,omitempty"`
	UserName    string `yaml:"username" json:"username,omitempty"`
	Password    string `yaml:"password" json:"password,omitempty"`
}

// InitRegistryConfig init registry config
func InitRegistryConfig(registryConfig *RegistryConfig) {
	config = registryConfig
}

// GetRegistryConfig get registry config
func GetRegistryConfig() *RegistryConfig {
	return config
}

type EtcdConfig struct {
	ClusterName string        `default:"starfish-etcdv3" yaml:"cluster_name" json:"cluster_name,omitempty"`
	Endpoints   string        `yaml:"endpoints" json:"endpoints,omitempty"`
	Heartbeats  int           `yaml:"heartbeats" json:"heartbeats"`
	Timeout     time.Duration `yaml:"timeout" json:"timeout"`
}
