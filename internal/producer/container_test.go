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

package producer_test

import (
	"database/sql"
	"os"
	"testing"
)

import (
	"github.com/transaction-mesh/starfish/pkg/test"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

var (
	utContainerMySQL *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testContainer := &test.MysqlContainer{
		Username: "root",
		Password: "123456",
		Database: "starfish",
	}
	ctx, container := test.SetupMysql(testContainer)
	utContainerMySQL, err = testContainer.OpenConnection(ctx, container)
	defer test.CloseConnection(ctx, container)
	if err != nil {
		log.Errorf("failed to setup MySQL container")
		panic(err)
	}
	os.Exit(m.Run())
}

func TestProducerSelect(t *testing.T) {
	query, err := utContainerMySQL.Query("select * from branch_table")
	if err != nil {
		t.Errorf("error on list branch_table : %s", err)
	}
	defer query.Close()
}
