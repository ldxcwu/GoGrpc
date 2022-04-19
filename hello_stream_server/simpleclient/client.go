package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "streamserverservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	stream, err := client.SayHello(context.Background(), &pb.RpcRequest{Name: "Cc"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("res: %v\n", res)
	}
}
