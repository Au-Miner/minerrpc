package main

import (
	"fmt"
	"minerrpc/rpc_core/transport/transport_server"
	"minerrpc/rpc_server/servers/servers_services"
	"os"
	"os/signal"
)

func main() {
	srv, err := transport_server.NewDefaultSocketServer("localhost:3212")
	if err != nil {
		panic(err)
	}
	ss := servers_services.TestImpl{}
	srv.Register(&ss)
	go srv.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	sig := <-stopChan
	fmt.Printf("Received %v, initiating shutdown...\n", sig)
}
