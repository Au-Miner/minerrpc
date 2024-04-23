package servicesDiscovery

import (
	"jrpc/src/rpc-common/utils"
	"jrpc/src/rpc-core/loadbalancer"
	"net"
)

type ZkServiceDiscovery struct {
	LoadBalancer loadbalancer.LoadBalancer
}

func NewZkServiceDiscovery(lb loadbalancer.LoadBalancer) *ZkServiceDiscovery {
	if lb == nil {
		return &ZkServiceDiscovery{LoadBalancer: &loadbalancer.RandomLoadBalancer{}}
	} else {
		return &ZkServiceDiscovery{LoadBalancer: lb}
	}
}

func (z *ZkServiceDiscovery) LookupService(serviceName string) (*net.TCPAddr, error) {
	instances, err := utils.GetAllInstances(serviceName)
	if err != nil {
		return nil, err
	}
	return z.LoadBalancer.Select(instances)
}
