package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "streamclientservice"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatal("get stream client error: ", err)
	}
	for i := 0; i < 5; i++ {
		stream.Send(&pb.RpcRequest{Name: fmt.Sprintf("Name"+"%d", i)})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("failed to get response: ", err)
	}
	log.Print(res.Value)
}
