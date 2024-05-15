package main

import (
	"fmt"
	transportServer "jrpc/src/rpc_core/transport/server"
	serversServices "jrpc/src/rpc_server/servers/services"
	"os"
	"os/signal"
)

func main() {
	srv, err := transportServer.NewDefaultSocketServer("localhost:3212")
	if err != nil {
		panic(err)
	}
	srv.Register(&serversServices.TestImpl{})
	go srv.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	sig := <-stopChan
	fmt.Printf("Received %v, initiating shutdown...\n", sig)
}
