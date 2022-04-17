package main

import (
	"log"
	"net"
	"net/rpc"
	"standarlize"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	helloService := HelloService{}
	err := standarlize.RegisterHelloService(&helloService)
	if err != nil {
		log.Fatal("Register error: ", err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
