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
	"time"
)

import (
	getty "github.com/apache/dubbo-getty"
)

import (
	getty2 "github.com/transaction-mesh/starfish/pkg/base/getty"
	"github.com/transaction-mesh/starfish/pkg/base/protocal"
	"github.com/transaction-mesh/starfish/pkg/util/log"
)

const (
	CronPeriod = 20e9
)

func (coordinator *DefaultCoordinator) OnOpen(session getty.Session) error {
	log.Infof("got getty_session:%s", session.Stat())
	return nil
}

func (coordinator *DefaultCoordinator) OnError(session getty.Session, err error) {
	SessionManager.ReleaseGettySession(session)
	session.Close()
	log.Errorf("getty_session{%s} got error{%v}, will be closed.", session.Stat(), err)
}

func (coordinator *DefaultCoordinator) OnClose(session getty.Session) {
	log.Info("getty_session{%s} is closing......", session.Stat())
}

func (coordinator *DefaultCoordinator) OnMessage(session getty.Session, pkg interface{}) {
	log.Debugf("received message:{%v}", pkg)
	rpcMessage, ok := pkg.(protocal.RpcMessage)
	if ok {
		_, isRegTM := rpcMessage.Body.(protocal.RegisterTMRequest)
		if isRegTM {
			coordinator.OnRegTmMessage(rpcMessage, session)
			return
		}

		heartBeat, isHeartBeat := rpcMessage.Body.(protocal.HeartBeatMessage)
		if isHeartBeat && heartBeat == protocal.HeartBeatMessagePing {
			coordinator.OnCheckMessage(rpcMessage, session)
			return
		}

		if rpcMessage.MessageType == protocal.MSGTypeRequest ||
			rpcMessage.MessageType == protocal.MSGTypeRequestOneway {
			log.Debugf("msgID:%s, body:%v", rpcMessage.ID, rpcMessage.Body)
			_, isRegRM := rpcMessage.Body.(protocal.RegisterRMRequest)
			if isRegRM {
				coordinator.OnRegRmMessage(rpcMessage, session)
			} else {
				if SessionManager.IsRegistered(session) {
					defer func() {
						if err := recover(); err != nil {
							log.Errorf("Catch Exception while do RPC, request: %v,err: %w", rpcMessage, err)
						}
					}()
					coordinator.OnTrxMessage(rpcMessage, session)
				} else {
					session.Close()
					log.Infof("close a unhandled connection! [%v]", session)
				}
			}
		} else {
			resp, loaded := coordinator.futures.Load(rpcMessage.ID)
			if loaded {
				response := resp.(*getty2.MessageFuture)
				response.Response = rpcMessage.Body
				response.Done <- true
				coordinator.futures.Delete(rpcMessage.ID)
			}
		}
	}
}

func (coordinator *DefaultCoordinator) OnCron(session getty.Session) {
	active := session.GetActive()
	if CronPeriod < time.Since(active).Nanoseconds() {
		log.Infof("OnCorn session{%s} timeout{%s}", session.Stat(), time.Since(active).String())
		session.Close()
	}
}
