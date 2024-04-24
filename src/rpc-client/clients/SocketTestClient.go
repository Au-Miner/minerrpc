package main

import (
	"fmt"
	apiServices "jrpc/src/rpc-api/services"
	transportClient "jrpc/src/rpc-core/transport/client"
)

func main() {
	client := transportClient.NewDefaultSocketClient()
	proxy := transportClient.NewRpcClientProxy(client)

	testService := proxy.NewProxyInstance(&apiServices.Test{}).(*apiServices.Test)
	res, _ := testService.Ping()
	fmt.Println("结果是: ", res)
	res, _ = testService.Hello()
	fmt.Println("结果是: ", res)
}
