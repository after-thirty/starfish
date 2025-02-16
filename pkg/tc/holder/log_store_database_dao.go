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

package holder

import (
	"github.com/go-xorm/xorm"

	"xorm.io/builder"
)

import (
	"github.com/transaction-mesh/starfish/pkg/tc/model"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

const (
	QueryGlobalTransactionDOByXid = `select xid, transaction_id, status, application_id, transaction_service_group, transaction_name,
		timeout, begin_time, application_data, gmt_create, gmt_modified from global_table where xid = ?`
	QueryGlobalTransactionDOByTransactionID = `select xid, transaction_id, status, application_id, transaction_service_group, transaction_name,
		timeout, begin_time, application_data, gmt_create, gmt_modified from global_table where transaction_id = ?`
	InsertGlobalTransactionDO = `insert into global_table (xid, transaction_id, status, application_id, transaction_service_group,
        transaction_name, timeout, begin_time, application_data, gmt_create, gmt_modified) values(?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now())`
	UpdateGlobalTransactionDO     = "update global_table set status = ?, gmt_modified = now() where xid = ?"
	DeleteGlobalTransactionDO     = "delete from global_table where xid = ?"
	QueryBranchTransactionDOByXid = `select xid, branch_id, transaction_id, resource_group_id, resource_id, branch_type, status, client_id,
	    application_data, gmt_create, gmt_modified from branch_table where xid = ? order by gmt_create asc`
	InsertBranchTransactionDO = `insert into branch_table (xid, branch_id, transaction_id, resource_group_id, resource_id, branch_type,
        status, client_id, application_data, gmt_create, gmt_modified) values(?, ?, ?, ?, ?, ?, ?, ?, ?, now(6), now(6))`
	UpdateBranchTransactionDO = "update branch_table set status = ?, gmt_modified = now(6) where xid = ? and branch_id = ?"
	DeleteBranchTransactionDO = "delete from branch_table where xid = ? and branch_id = ?"
	QueryMaxTransactionID     = "select max(transaction_id) as maxTransactionID from global_table where transaction_id < ? and transaction_id > ?"
	QueryMaxBranchID          = "select max(branch_id) as maxBranchID from branch_table where branch_id < ? and branch_id > ?"
)

type LogStore interface {
	QueryGlobalTransactionDOByXID(xid string) *model.GlobalTransactionDO
	QueryGlobalTransactionDOByTransactionID(transactionID int64) *model.GlobalTransactionDO
	QueryGlobalTransactionDOByStatuses(statuses []int, limit int) []*model.GlobalTransactionDO
	InsertGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool
	UpdateGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool
	DeleteGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool
	QueryBranchTransactionDOByXID(xid string) []*model.BranchTransactionDO
	QueryBranchTransactionDOByXIDs(xids []string) []*model.BranchTransactionDO
	InsertBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool
	UpdateBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool
	DeleteBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool
	GetCurrentMaxSessionID(high int64, low int64) int64
}

type LogStoreDataBaseDAO struct {
	engine *xorm.Engine
}

func (dao *LogStoreDataBaseDAO) QueryGlobalTransactionDOByXID(xid string) *model.GlobalTransactionDO {
	var globalTransactionDO model.GlobalTransactionDO
	has, err := dao.engine.SQL(QueryGlobalTransactionDOByXid, xid).
		Get(&globalTransactionDO)
	if has {
		return &globalTransactionDO
	}
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (dao *LogStoreDataBaseDAO) QueryGlobalTransactionDOByTransactionID(transactionID int64) *model.GlobalTransactionDO {
	var globalTransactionDO model.GlobalTransactionDO
	has, err := dao.engine.SQL(QueryGlobalTransactionDOByTransactionID, transactionID).
		Get(&globalTransactionDO)
	if has {
		return &globalTransactionDO
	}
	if err != nil {
		log.Errorf(err.Error())
	}
	return nil
}

func (dao *LogStoreDataBaseDAO) QueryGlobalTransactionDOByStatuses(statuses []int, limit int) []*model.GlobalTransactionDO {
	var globalTransactionDOs []*model.GlobalTransactionDO
	err := dao.engine.Table("global_table").
		Where(builder.In("status", statuses)).
		OrderBy("gmt_modified").
		Limit(limit).
		Find(&globalTransactionDOs)

	if err != nil {
		log.Errorf(err.Error())
	}
	return globalTransactionDOs
}

func (dao *LogStoreDataBaseDAO) InsertGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool {
	_, err := dao.engine.Exec(InsertGlobalTransactionDO,
		globalTransaction.XID,
		globalTransaction.TransactionID,
		globalTransaction.Status,
		globalTransaction.ApplicationID,
		globalTransaction.TransactionServiceGroup,
		globalTransaction.TransactionName,
		globalTransaction.Timeout,
		globalTransaction.BeginTime,
		globalTransaction.ApplicationData)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) UpdateGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool {
	_, err := dao.engine.Exec(UpdateGlobalTransactionDO, globalTransaction.Status, globalTransaction.XID)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) DeleteGlobalTransactionDO(globalTransaction model.GlobalTransactionDO) bool {
	_, err := dao.engine.Exec(DeleteGlobalTransactionDO, globalTransaction.XID)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) QueryBranchTransactionDOByXID(xid string) []*model.BranchTransactionDO {
	var branchTransactionDos []*model.BranchTransactionDO
	err := dao.engine.SQL(QueryBranchTransactionDOByXid, xid).Find(&branchTransactionDos)
	if err != nil {
		log.Errorf(err.Error())
	}
	return branchTransactionDos
}

func (dao *LogStoreDataBaseDAO) QueryBranchTransactionDOByXIDs(xids []string) []*model.BranchTransactionDO {
	var branchTransactionDos []*model.BranchTransactionDO
	err := dao.engine.Table("branch_table").
		Where(builder.In("xid", xids)).
		OrderBy("gmt_create asc").
		Find(&branchTransactionDos)
	if err != nil {
		log.Errorf(err.Error())
	}
	return branchTransactionDos
}

func (dao *LogStoreDataBaseDAO) InsertBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool {
	_, err := dao.engine.Exec(InsertBranchTransactionDO,
		branchTransaction.XID,
		branchTransaction.BranchID,
		branchTransaction.TransactionID,
		branchTransaction.ResourceGroupID,
		branchTransaction.ResourceID,
		branchTransaction.BranchType,
		branchTransaction.Status,
		branchTransaction.ClientID,
		branchTransaction.ApplicationData)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) UpdateBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool {
	_, err := dao.engine.Exec(UpdateBranchTransactionDO,
		branchTransaction.Status,
		branchTransaction.XID,
		branchTransaction.BranchID)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) DeleteBranchTransactionDO(branchTransaction model.BranchTransactionDO) bool {
	_, err := dao.engine.Exec(DeleteBranchTransactionDO,
		branchTransaction.XID,
		branchTransaction.BranchID)

	return err == nil
}

func (dao *LogStoreDataBaseDAO) GetCurrentMaxSessionID(high int64, low int64) int64 {
	var maxTransactionID, maxBranchID int64
	_, err := dao.engine.SQL(QueryMaxTransactionID, high, low).
		Cols("maxTransactionID").
		Get(&maxTransactionID)
	if err != nil {
		log.Errorf(err.Error())
	}
	_, err = dao.engine.SQL(QueryMaxBranchID, high, low).
		Cols("maxBranchID").
		Get(&maxBranchID)
	if err != nil {
		log.Errorf(err.Error())
	}
	if maxTransactionID > maxBranchID {
		return maxTransactionID
	}
	return maxBranchID
}
