package main

import (
	"fmt"
	"net"

	pb "test/a" // 引入编译生成的包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type HelloService struct {
	pb.UnimplementedHelloServer
}

// SayHello 实现Hello服务接口
func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("接收名字 %s.", in.Name)

	fmt.Println(in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, &HelloService{})

	grpclog.Info("Listen on " + Address)

	str := s.Serve(listen)
	fmt.Println(str)

}
