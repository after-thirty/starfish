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

package main

import (
	"fmt"
	"os"
	"strconv"
)

import (
	gxnet "github.com/dubbogo/gost/net"

	"github.com/urfave/cli/v2"
)

import (
	"github.com/transaction-mesh/starfish/common/version"
	"github.com/transaction-mesh/starfish/pkg/base/common"
	_ "github.com/transaction-mesh/starfish/pkg/base/config_center/nacos"
	_ "github.com/transaction-mesh/starfish/pkg/base/registry/etcdv3"
	_ "github.com/transaction-mesh/starfish/pkg/base/registry/file"
	_ "github.com/transaction-mesh/starfish/pkg/base/registry/nacos"
	"github.com/transaction-mesh/starfish/pkg/tc/config"
	"github.com/transaction-mesh/starfish/pkg/tc/holder"
	"github.com/transaction-mesh/starfish/pkg/tc/lock"
	_ "github.com/transaction-mesh/starfish/pkg/tc/metrics"
	"github.com/transaction-mesh/starfish/pkg/tc/server"
	"github.com/transaction-mesh/starfish/pkg/util/log"
	"github.com/transaction-mesh/starfish/pkg/util/uuid"
)

var (
	appName = "starfish"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start starfish golang tc server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config, c",
						Usage: "Load configuration from `FILE`",
					},
					&cli.StringFlag{
						Name:  "serverNode, n",
						Value: "1",
						Usage: "server node id, such as 1, 2, 3. default is 1",
					},
				},
				Action: func(c *cli.Context) error {
					configPath := c.String("config")
					serverNode := c.Int64("serverNode")

					conf, err := config.InitConf(configPath)
					if err != nil {
						log.Fatal(err)
					}

					ip, _ := gxnet.GetLocalIP()
					port, err := strconv.Atoi(conf.Port)
					if err != nil {
						log.Fatal(err)
					}

					common.Init(ip, port)
					uuid.Init(serverNode)
					lock.Init()
					holder.Init()

					srv := server.NewServer()
					srv.Start(fmt.Sprintf(":%s", conf.Port))
					return nil
				},
			},
		},
		Version: version.Print(appName),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err)
	}
}
