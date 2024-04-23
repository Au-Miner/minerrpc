package loadbalancer

import "net"

type LoadBalancer interface {
	Select(services []*net.TCPAddr) (*net.TCPAddr, error)
}
