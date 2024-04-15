package main

import (
	"context"
	"log"
	"net"

	pb "jrpc/src/pb"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

// 对服务器的抽象，用来实现服务方法
type server struct {
	pb.UnimplementedProductInfoServer
}

// 存放商品，模拟业务逻辑
var productMap map[string]*pb.Product

// 实现 AddProduct 方法
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if productMap == nil {
		productMap = make(map[string]*pb.Product)
	}
	productMap[in.Id] = in
	log.Printf("Product %v : %v - Added.", in.Id, in.Name)
	return &pb.ProductID{Value: in.Id}, nil
}

// 实现 GetProduct 方法
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	product, exists := productMap[in.Value]
	if exists && product != nil {
		log.Printf("Product %v : %v - Retrieved.", product.Id, product.Name)
		return product, nil
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func main() {
	// 创建一个 tcp 监听器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个 gRPC 服务器实例
	s := grpc.NewServer()
	// 将服务注册到 gRPC 服务器上
	pb.RegisterProductInfoServer(s, &server{})
	// 绑定 gRPC 服务器到指定 tcp
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
