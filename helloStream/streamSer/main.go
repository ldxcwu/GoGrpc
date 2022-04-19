package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "streamservice"

	"google.golang.org/grpc"
)

type CommunicationServer struct {
	pb.UnimplementedCommunicationServer
}

func (s *CommunicationServer) Communicate(srv pb.Communication_CommunicateServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("req.Msg: %v\n", req.Msg)
		srv.Send(&pb.RpcResponse{Msg: req.Msg})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("lisnten error: ", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommunicationServer(s, &CommunicationServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Serve error: ", err)
	}
}
