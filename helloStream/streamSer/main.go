package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "streamservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
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
			log.Fatal("Recv error: ", err)
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
	//alts 认证
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	s := grpc.NewServer(grpc.Creds(altsTC))
	// s := grpc.NewServer()
	pb.RegisterCommunicationServer(s, &CommunicationServer{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Serve error: ", err)
	}
}
