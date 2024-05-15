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
	AimIp            string
}

func NewDefaultSocketClient() *SocketClient {
	return NewSocketClient(DEFAULT_SERIALIZER, &load_balancer.RandomLoadBalancer{}, "")
}

func NewDefaultSocketClientWithAimIp(aimIp string) *SocketClient {
	return NewSocketClient(DEFAULT_SERIALIZER, &load_balancer.RandomLoadBalancer{}, aimIp)
}

func NewSocketClient(serializerId int, loadBalancer load_balancer.LoadBalancer, aimIp string) *SocketClient {
	return &SocketClient{
		Serializer:       serializer.GetByCode(serializerId),
		ServiceDiscovery: servicesDiscovery.NewZkServiceDiscovery(loadBalancer),
		AimIp:            aimIp,
	}
}

func (client *SocketClient) SendRequest(req entities.RPCdata) (*entities.RPCdata, error) {
	var addr *net.TCPAddr
	var err error
	if client.AimIp == "" {
		fmt.Println("走的是ServiceDiscovery")
		addr, err = client.ServiceDiscovery.LookupService(req.Name)
		fmt.Println("结果为addr", addr)
	} else {
		addr, err = net.ResolveTCPAddr("tcp", client.AimIp)
		fmt.Println("走的是ResolveTCPAddr")
	}
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
