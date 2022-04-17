package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	conn, err := lis.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}
