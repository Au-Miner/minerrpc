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

// func main() {
// 	conn, _ := net.Dial("tcp", "localhost:5000")
// 	defer conn.Close()
// 	req := &protofile.Request{
// 		Key: 200,
// 	}
// 	data, _ := proto.Marshal(req)
// 	conn.Write(data)
//
// 	buf := make([]byte, 1024)
// 	n, _ := conn.Read(buf)
// 	var resp protofile.Response
// 	proto.Unmarshal(buf[:n], &resp)
// 	fmt.Printf("Received response: %v\n", resp.Value)
// }
