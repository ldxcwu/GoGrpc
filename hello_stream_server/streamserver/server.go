package main

import (
	"log"
	"net"
	pb "streamserverservice"

	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(req *pb.RpcRequest, srv pb.Greeter_SayHelloServer) error {
	for i := 0; i < 5; i++ {
		err := srv.Send(&pb.RpcResponse{
			StreamValue: "Hello " + req.Name,
		})
		if err != nil {
			log.Fatal("Send error: ", err)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
