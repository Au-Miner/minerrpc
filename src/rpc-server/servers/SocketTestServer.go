package main

import (
	transportServer "jrpc/src/rpc-core/transport/server"
	serversServices "jrpc/src/rpc-server/servers/services"
)

func main() {
	srv, err := transportServer.NewDefaultSocketServer("localhost:3212")
	if err != nil {
		panic(err)
	}
	srv.Register(&serversServices.TestImpl{})
	go srv.Start()
}
