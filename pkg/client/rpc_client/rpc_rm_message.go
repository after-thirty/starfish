package rpc_client

import (
	"github.com/gotrx/starfish/pkg/base/protocal"
)

type RpcRMMessage struct {
	RpcMessage    protocal.RpcMessage
	ServerAddress string
}
