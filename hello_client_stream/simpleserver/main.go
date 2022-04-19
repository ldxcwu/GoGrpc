package main

import (
	"io"
	"log"
	"net"
	pb "streamclientservice"

	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(srv pb.Greeter_SayHelloServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			srv.SendAndClose(&pb.RpcResponse{Value: "Ok"})
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Recv: ", req.Name)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreeterServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
