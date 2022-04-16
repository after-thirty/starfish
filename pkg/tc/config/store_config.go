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
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
)

const (
	DefaultFileDir                      = "root.data"
	DefaultMaxBranchSessionSize         = 1024 * 16
	DefaultMaxGlobalSessionSize         = 512
	DefaultWriteBufferSize              = 1024 * 16
	DefaultServiceSessionReloadReadSize = 100
)

type FlushDiskMode int

const (
	/**
	 * sync flush disk
	 */
	FlushdiskModeSyncModel FlushDiskMode = iota

	/**
	 * async flush disk
	 */
	FlushdiskModeAsyncModel
)

type StoreConfig struct {
	MaxBranchSessionSize int             `default:"16384" yaml:"max_branch_session_size" json:"max_branch_session_size,omitempty"`
	MaxGlobalSessionSize int             `default:"512" yaml:"max_global_session_size" json:"max_global_session_size,omitempty"`
	StoreMode            string          `default:"file" yaml:"mode" json:"mode,omitempty"`
	FileStoreConfig      FileStoreConfig `yaml:"file" json:"file,omitempty"`
	DBStoreConfig        DBStoreConfig   `yaml:"db" json:"db,omitempty"`
}

type FileStoreConfig struct {
	FileDir                  string        `default:"root.data" yaml:"file_dir" json:"file_dir,omitempty"`
	FileWriteBufferCacheSize int           `default:"16384" yaml:"file_write_buffer_cache_size" json:"file_write_buffer_cache_size,omitempty"`
	FlushDiskMode            FlushDiskMode `default:"1" yaml:"flush_disk_mode" json:"flush_disk_mode,omitempty"`
	SessionReloadReadSize    int           `default:"100" yaml:"session_reload_read_size" json:"session_reload_read_size,omitempty"`
}

type DBStoreConfig struct {
	LogQueryLimit int    `default:"100" yaml:"log_query_limit" json:"log_query_limit"`
	DSN           string `yaml:"dsn" json:"dsn"`
	Engine        *xorm.Engine
}

func GetDefaultFileStoreConfig() FileStoreConfig {
	return FileStoreConfig{
		FileDir:                  DefaultFileDir,
		FileWriteBufferCacheSize: DefaultWriteBufferSize,
		FlushDiskMode:            0,
		SessionReloadReadSize:    DefaultServiceSessionReloadReadSize,
	}
}
