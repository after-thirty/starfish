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

package registry

type Address struct {
	IP   string
	Port uint64
}

// Registry Extension - Registry
type Registry interface {
	//注册服务
	Register(addr *Address) error
	//取消注册
	UnRegister(addr *Address) error
	//查询服务地址
	Lookup() ([]string, error)
	//订阅
	Subscribe(EventListener) error
	//取消订阅
	UnSubscribe(EventListener) error

	Stop()
}

//订阅获取到的服务信息
type Service struct {
	EventType uint32 // EventType: 0 => PUT, 1 => DELETE
	IP        string
	Port      uint64
	Name      string
}

type EventListener interface {
	OnEvent(service []*Service) error
}
