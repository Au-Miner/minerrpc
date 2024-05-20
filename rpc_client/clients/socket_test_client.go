package main

import (
	"fmt"
	apiServices "minerrpc/rpc_api/services"
	transportClient "minerrpc/rpc_core/transport/client"
)

func main() {
	client := transportClient.NewDefaultSocketClient()
	proxy := transportClient.NewRpcClientProxy(client)

	testService := proxy.NewProxyInstance(&apiServices.Test{}).(*apiServices.Test)
	res, _ := testService.Ping()
	fmt.Println("The result is: ", res)
	res, _ = testService.Hello()
	fmt.Println("The result is: ", res)
}
