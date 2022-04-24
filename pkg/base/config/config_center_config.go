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

// ConfigCenterConfig config center config
type ConfigCenterConfig struct {
	Mode        string            `yaml:"type" json:"type,omitempty"` //类型
	NacosConfig NacosConfigCenter `yaml:"nacos" json:"nacos,omitempty"`
	ETCDConfig  EtcdConfigCenter  `yaml:"etcdv3" json:"etcdv3,omitempty"`
}

// NacosConfigCenter nacos config center
type NacosConfigCenter struct {
	ServerAddr string `yaml:"server_addr" json:"server_addr,omitempty"`
	Group      string `default:"SEATA_GROUP" yaml:"group" json:"group,omitempty"`
	Namespace  string `yaml:"namespace" json:"namespace,omitempty"`
	Cluster    string `yaml:"cluster" json:"cluster,omitempty"`
	UserName   string `yaml:"username" json:"username,omitempty"`
	Password   string `yaml:"password" json:"password,omitempty"`
	DataID     string `default:"starfish" yaml:"data_id" json:"data_id,omitempty"`
}

type EtcdConfigCenter struct {
	Name       string        `default:"starfish-config-center" yaml:"name" json:"name"`
	ConfigKey  string        `default:"config-starfish" yaml:"config_key" json:"config_key,omitempty"`
	Endpoints  string        `yaml:"endpoints" json:"endpoints,omitempty"`
	Heartbeats int           `yaml:"heartbeats" json:"heartbeats"`
	Timeout    time.Duration `yaml:"timeout" json:"timeout"`
}
