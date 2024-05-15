package transportClient

import (
	"fmt"
	"jrpc/src/rpc_common/entities"
	"jrpc/src/rpc_core/load_balancer"
	"jrpc/src/rpc_core/serializer"
	"jrpc/src/rpc_core/services/discovery"
	transportUtils "jrpc/src/rpc_core/transport/utils"
	"net"
)

type SocketClient struct {
	ServiceDiscovery servicesDiscovery.ServiceDiscovery
	Serializer       serializer.CommonSerializer
}

func NewDefaultSocketClient() *SocketClient {
	return NewSocketClient(DEFAULT_SERIALIZER, &load_balancer.RandomLoadBalancer{})
}

func NewSocketClient(serializerId int, loadBalancer load_balancer.LoadBalancer) *SocketClient {
	return &SocketClient{
		Serializer:       serializer.GetByCode(serializerId),
		ServiceDiscovery: servicesDiscovery.NewZkServiceDiscovery(loadBalancer),
	}
}

func (client *SocketClient) SendRequest(req entities.RPCdata) (*entities.RPCdata, error) {
	addr, err := client.ServiceDiscovery.LookupService(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup service: %w", err)
	}
	conn, err := net.Dial("tcp", addr.String())
	if err != nil {
		return nil, err
	}
	err = transportUtils.NewObjectWriter(conn).WriteObject(&req, client.Serializer)
	if err != nil {
		return nil, err
	}
	resp, err := transportUtils.NewObjectReader(conn).ReadObject()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
