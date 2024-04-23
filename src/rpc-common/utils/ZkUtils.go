package utils

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"jrpc/src/rpc-common/constants"
	"net"
	"os"
	"os/signal"
)

var (
	conn       *zk.Conn
	serviceSet = make(map[string]bool)
)

func init() {
	var err error
	conn, _, err = zk.Connect([]string{constants.ZK_SERVER_ADDRESS}, constants.ZK_SESSION_TIMEOUT)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Zookeeper: %v", err))
	}
	fmt.Println("ZK client connected successfully.")
	cleanupHook()
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("ZK client is closing.")
			conn.Close()
		}
	}()
}

func GetAllInstances(serviceName string) ([]*net.TCPAddr, error) {
	servicePath := fmt.Sprintf("%s/%s", constants.ZK_BASE_PATH, serviceName)
	children, _, err := conn.Children(servicePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get children for %s: %v", servicePath, err)
	}

	var instances []*net.TCPAddr
	for _, child := range children {
		addr, err := net.ResolveTCPAddr("tcp", child)
		if err != nil {
			fmt.Printf("Invalid address %s: %v\n", child, err)
			continue
		}
		instances = append(instances, addr)
	}
	return instances, nil
}

func RegisterService(serviceName string, addr *net.TCPAddr) error {
	serviceSet[serviceName] = true
	servicePath := fmt.Sprintf("%s/%s/%s:%d", constants.ZK_BASE_PATH, serviceName, addr.IP.String(), addr.Port)
	flags := int32(zk.FlagEphemeral)

	_, err := conn.Create(servicePath, []byte{}, flags, zk.WorldACL(zk.PermAll))
	if err != nil {
		return fmt.Errorf("failed to create service node: %v", err)
	}
	return nil
}
