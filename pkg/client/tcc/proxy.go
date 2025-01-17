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

package tcc

import (
	"encoding/json"
	"reflect"
	"strconv"
)

import (
	gxnet "github.com/dubbogo/gost/net"

	"github.com/pkg/errors"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/client/context"
	"github.com/transaction-mesh/starfish/pkg/client/proxy"
	"github.com/transaction-mesh/starfish/pkg/util/log"
	"github.com/transaction-mesh/starfish/pkg/util/time"
)

var (
	// TCCActionName
	TCCActionName = "TCCActionName"

	// TryMethod
	TryMethod = "Try"
	// ConfirmMethod
	ConfirmMethod = "Confirm"
	// CancelMethod
	CancelMethod = "Cancel"

	// ActionStartTime
	ActionStartTime = "action-start-time"
	// ActionName
	ActionName = "actionName"
	// PrepareMethod
	PrepareMethod = "sys::prepare"
	// CommitMethod
	CommitMethod = "sys::commit"
	// RollbackMethod
	RollbackMethod = "sys::rollback"
	// HostName
	HostName = "host-name"

	// TCCMethodArguments
	TCCMethodArguments = "arguments"
	// TCCMethodResult
	TCCMethodResult = "result"

	businessActionContextType = reflect.TypeOf(&context.BusinessActionContext{})
)

type TCCService interface {
	Try(ctx *context.BusinessActionContext) (bool, error)
	Confirm(ctx *context.BusinessActionContext) bool
	Cancel(ctx *context.BusinessActionContext) bool
}

type TCCServiceProxy interface {
	GetTCCService() TCCService
}

func makeCallProxy(methodDesc *proxy.MethodDescriptor, resource *TCCResource) func(in []reflect.Value) []reflect.Value {
	return func(in []reflect.Value) []reflect.Value {
		businessContextValue := in[0]
		businessActionContext := businessContextValue.Interface().(*context.BusinessActionContext)
		rootContext := businessActionContext.RootContext
		businessActionContext.XID = rootContext.GetXID()
		businessActionContext.ActionName = resource.ActionName
		if !rootContext.InGlobalTransaction() {
			args := make([]interface{}, 0)
			args = append(args, businessActionContext)
			return proxy.Invoke(methodDesc, nil, args)
		}

		returnValues, _ := proceed(methodDesc, businessActionContext, resource)
		return returnValues
	}
}

func ImplementTCC(v TCCServiceProxy) {
	valueOf := reflect.ValueOf(v)
	log.Debugf("[Implement] reflect.TypeOf: %s", valueOf.String())
	if valueOf.Kind() != reflect.Ptr {
		log.Errorf("%s must be a ptr", valueOf)
		return
	}

	valueOfElem := valueOf.Elem()
	typeOf := valueOfElem.Type()

	// check incoming interface, incoming interface's elem must be a struct.
	if typeOf.Kind() != reflect.Struct {
		log.Errorf("%s must be a struct ptr", valueOf.String())
		return
	}
	serviceProxy := v.GetTCCService()

	numField := valueOfElem.NumField()
	for i := 0; i < numField; i++ {
		t := typeOf.Field(i)
		methodName := t.Name
		f := valueOfElem.Field(i)
		if f.Kind() == reflect.Func && f.IsValid() && f.CanSet() && methodName == TryMethod {
			if t.Type.NumIn() != 1 && t.Type.In(0) != businessActionContextType {
				panic("prepare method argument is not BusinessActionContext")
			}

			actionName := t.Tag.Get(TCCActionName)
			if actionName == "" {
				panic("must tag TCCActionName")
			}

			commitMethodDesc := proxy.Register(serviceProxy, ConfirmMethod)
			cancelMethodDesc := proxy.Register(serviceProxy, CancelMethod)
			tryMethodDesc := proxy.Register(serviceProxy, TryMethod)

			tccResource := &TCCResource{
				ResourceGroupID:    "",
				AppName:            "",
				ActionName:         actionName,
				PrepareMethodName:  TryMethod,
				CommitMethodName:   CommitMethod,
				CommitMethod:       commitMethodDesc,
				RollbackMethodName: CancelMethod,
				RollbackMethod:     cancelMethodDesc,
			}

			tccResourceManager.RegisterResource(tccResource)

			// do method proxy here:
			f.Set(reflect.MakeFunc(f.Type(), makeCallProxy(tryMethodDesc, tccResource)))
			log.Debugf("set method [%s]", methodName)
		}
	}
}

func proceed(methodDesc *proxy.MethodDescriptor, ctx *context.BusinessActionContext, resource *TCCResource) ([]reflect.Value, error) {
	var (
		args = make([]interface{}, 0)
	)

	branchID, err := doTCCActionLogStore(ctx, resource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ctx.BranchID = strconv.FormatInt(branchID, 10)

	args = append(args, ctx)
	returnValues := proxy.Invoke(methodDesc, nil, args)
	errValue := returnValues[len(returnValues)-1]
	if errValue.IsValid() && !errValue.IsNil() {
		err := tccResourceManager.BranchReport(meta.BranchTypeTCC, ctx.XID, branchID, meta.BranchStatusPhaseOneFailed, nil)
		if err != nil {
			log.Errorf("branch report err: %v", err)
		}
	}

	return returnValues, nil
}

func doTCCActionLogStore(ctx *context.BusinessActionContext, resource *TCCResource) (int64, error) {
	ctx.ActionContext[ActionStartTime] = time.CurrentTimeMillis()
	ctx.ActionContext[PrepareMethod] = resource.PrepareMethodName
	ctx.ActionContext[CommitMethod] = resource.CommitMethodName
	ctx.ActionContext[RollbackMethod] = resource.RollbackMethodName
	ctx.ActionContext[ActionName] = ctx.ActionName
	ip, err := gxnet.GetLocalIP()
	if err == nil {
		ctx.ActionContext[HostName] = ip
	} else {
		log.Warn("getLocalIP error")
	}

	applicationContext := make(map[string]interface{})
	applicationContext[TCC_ACTION_CONTEXT] = ctx.ActionContext

	applicationData, err := json.Marshal(applicationContext)
	if err != nil {
		log.Errorf("marshal applicationContext failed:%v", applicationContext)
		return 0, err
	}

	branchID, err := tccResourceManager.BranchRegister(meta.BranchTypeTCC, ctx.ActionName, "", ctx.XID, applicationData, "")
	if err != nil {
		log.Errorf("TCC branch Register error, xid: %s", ctx.XID)
		return 0, errors.WithStack(err)
	}
	return branchID, nil
}
