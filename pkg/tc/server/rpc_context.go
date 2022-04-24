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

package server

import (
	getty "github.com/apache/dubbo-getty"
)

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/base/model"
)

const IpPortSplitChar = ":"

type RpcContext struct {
	Version                 string
	TransactionServiceGroup string
	ClientRole              meta.TransactionRole
	ApplicationID           string
	ClientID                string
	ResourceSets            *model.Set
	Session                 getty.Session
}

type RpcContextOption func(ctx *RpcContext)

func WithRpcContextVersion(version string) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.Version = version
	}
}

func WithRpcContextTxServiceGroup(txServiceGroup string) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.TransactionServiceGroup = txServiceGroup
	}
}

func WithRpcContextClientRole(clientRole meta.TransactionRole) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.ClientRole = clientRole
	}
}

func WithRpcContextApplicationID(applicationID string) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.ApplicationID = applicationID
	}
}

func WithRpcContextClientID(clientID string) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.ClientID = clientID
	}
}

func WithRpcContextResourceSet(resourceSet *model.Set) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.ResourceSets = resourceSet
	}
}

func WithRpcContextSession(session getty.Session) RpcContextOption {
	return func(ctx *RpcContext) {
		ctx.Session = session
	}
}

func NewRpcContext(opts ...RpcContextOption) *RpcContext {
	ctx := &RpcContext{
		ResourceSets: model.NewSet(),
	}
	for _, o := range opts {
		o(ctx)
	}
	return ctx
}

func (context *RpcContext) AddResource(resource string) {
	if resource != "" {
		if context.ResourceSets == nil {
			context.ResourceSets = model.NewSet()
		}
		context.ResourceSets.Add(resource)
	}
}

func (context *RpcContext) AddResources(resources *model.Set) {
	if resources != nil {
		if context.ResourceSets == nil {
			context.ResourceSets = model.NewSet()
		}
		for _, resource := range resources.List() {
			context.ResourceSets.Add(resource)
		}
	}
}
