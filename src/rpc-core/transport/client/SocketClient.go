package transportClient

import (
	"fmt"
	"jrpc/src/rpc-common/constants"
	"jrpc/src/rpc-common/entities"
	"jrpc/src/rpc-core/loadbalancer"
	"jrpc/src/rpc-core/serializer"
	"jrpc/src/rpc-core/services/discovery"
	transportUtils "jrpc/src/rpc-core/transport/utils"
	"net"
)

type SocketClient struct {
	ServiceDiscovery servicesDiscovery.ServiceDiscovery
	Serializer       serializer.CommonSerializer
}

func NewDefaultSocketClient() *SocketClient {
	return NewSocketClient(DEFAULT_SERIALIZER, &loadbalancer.RandomLoadBalancer{})
}

func NewSocketClient(serializerId int, loadBalancer loadbalancer.LoadBalancer) *SocketClient {
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
	conn, err := net.DialTimeout("tcp", addr.IP.String(), constants.ZK_SESSION_TIMEOUT)
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
