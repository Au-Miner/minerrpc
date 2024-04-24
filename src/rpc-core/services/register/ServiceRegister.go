package servicesRegister

import "net"

type ServiceRegister interface {
	Register(serviceName string, addr *net.TCPAddr) error
}
