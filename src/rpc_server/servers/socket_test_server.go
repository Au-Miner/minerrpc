package main

import (
	"fmt"
	transportServer "minerrpc/src/rpc_core/transport/server"
	serversServices "minerrpc/src/rpc_server/servers/services"
	"os"
	"os/signal"
)

func main() {
	srv, err := transportServer.NewDefaultSocketServer("localhost:3212")
	if err != nil {
		panic(err)
	}
	ss := serversServices.TestImpl{}
	srv.Register(&ss)
	go srv.Start()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	sig := <-stopChan
	fmt.Printf("Received %v, initiating shutdown...\n", sig)
}
