package main

import (
	"fmt"
	"minerrpc/rpc_api/api_services"
	"minerrpc/rpc_core/transport/transport_client"
)

func main() {
	client := transport_client.NewDefaultSocketClient()
	proxy := transport_client.NewRpcClientProxy(client)

	testService := proxy.NewProxyInstance(&api_services.Test{}).(*api_services.Test)
	res, _ := testService.Ping()
	fmt.Println("The result is: ", res)
	res, _ = testService.Hello()
	fmt.Println("The result is: ", res)
}
