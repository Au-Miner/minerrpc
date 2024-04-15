package main

import (
	"context"
	pb "jrpc/src/pb"
	"log"

	"google.golang.org/grpc"
)

const (
	// 服务端地址
	address = "localhost:50051"
)

func main() {
	// 创建 gRPC 连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建客户端 stub，利用它调用远程方法
	c := pb.NewProductInfoClient(conn)

	name := "XiaoMi 11"
	description := "XiaoMi 11 with MIUI 12.5"
	price := float32(3999.00)

	// 调用远程方法
	r, err := c.AddProduct(context.Background(), &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(context.Background(), &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %v", product.String())
}
