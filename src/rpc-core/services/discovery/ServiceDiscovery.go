package servicesDiscovery

import "net"

type ServiceDiscovery interface {
	LookupService(serviceName string) (*net.TCPAddr, error)
}
