package transportClient

import (
	"minerrpc/src/rpc_common/entities"
	"minerrpc/src/rpc_core/serializer"
)

const (
	DEFAULT_SERIALIZER = serializer.JSONSerializer
)

type RpcClient interface {
	SendRequest(req entities.RPCdata) (*entities.RPCdata, error)
}
