package transportServer

import "jrpc/src/rpc-core/serializer"

const (
	DEFAULT_SERIALIZER = serializer.JSONSerializer
)

type RpcServer interface {
	Start()
	Register(iClass interface{})
}
