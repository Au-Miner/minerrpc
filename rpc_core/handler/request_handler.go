package handler

import (
	"minerrpc/rpc_common/entities"
	"reflect"
)

type RequestHandler interface {
	Execute(req *entities.RPCdata, f reflect.Value) *entities.RPCdata
}
