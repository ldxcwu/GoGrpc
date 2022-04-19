package main

import (
	"context"
	"fmt"
	"log"

	pb "streamservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	defer conn.Close()

	client := pb.NewCommunicationClient(conn)
	stream, err := client.Communicate(context.Background())
	if err != nil {
		log.Fatal("Call error: ", err)
	}
	for i := 0; i < 5; i++ {
		err = stream.Send(&pb.RpcRequest{Msg: fmt.Sprintf("Name"+"%d", i)})
		if err != nil {
			log.Fatal("Send error: ", err)
		}
		res, err := stream.Recv()
		if err != nil {
			log.Fatal("Recv error: ", err)
		}
		fmt.Printf("res.Msg: %v\n", res.Msg)
	}
	err = stream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
}
