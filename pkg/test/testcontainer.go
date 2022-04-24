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

package test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

import (
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

type MysqlContainer struct {
	Username string `validate:"required" yaml:"username" json:"username"`
	Database string `validate:"required" yaml:"database" json:"database"`
	Password string `validate:"required" yaml:"password" json:"password"`
}

func SetupMysql(tester *MysqlContainer) (context.Context, testcontainers.Container) {
	log.Info("setup mysql container")
	ctx := context.Background()
	seedDataPath, err := os.Getwd()
	if err != nil {
		log.Errorf("Error get working directory: %s", err)
		panic(fmt.Sprintf("%v", err))
	}
	mountPath := seedDataPath + "/../../scripts/server/db"
	slashPath := filepath.ToSlash(mountPath)
	req := testcontainers.ContainerRequest{
		Image: "mysql:latest",
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": tester.Password,
			"MYSQL_DATABASE":      tester.Database,
		},
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		BindMounts: map[string]string{
			"/docker-entrypoint-initdb.d": slashPath,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		//.WithStartupTimeout(time.Minute * 2),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Errorf("Error Start MySQL container: %s", err)
		panic(fmt.Sprintf("%v", err))
	}
	return ctx, container
}

func (tester MysqlContainer) OpenConnection(ctx context.Context, container testcontainers.Container) (*sql.DB, error) {
	host, _ := container.Host(ctx)
	p, _ := container.MappedPort(ctx, "3306/tcp")
	port := p.Int()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify&parseTime=true&multiStatements=true",
		tester.Username, tester.Password, host, port, tester.Database)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Error("error connect to db: %+v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Errorf("error pinging db: %+v\n", err)
	}
	return db, err
}

func CloseConnection(ctx context.Context, container testcontainers.Container) {
	log.Info("Closing Container")
	err := container.Terminate(ctx)
	if err != nil {
		log.Errorf("error stop Container: %s", err)
		panic(fmt.Sprintf("%v", err))
	}
}
