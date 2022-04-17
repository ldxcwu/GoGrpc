package main

import (
	"fmt"
	"log"
	"standarlize"
)

func main() {
	/* 	client, err := rpc.Dial("tcp", "localhost:8080")
	   	if err != nil {
	   		log.Fatal("dailing: ", err)
	   	}

	   	var reply string
	   	err = client.Call(standarlize.HelloServiceName+".Hello", "world", &reply)
	   	if err != nil {
	   		log.Fatal(err)
	   	}
	   	fmt.Println(reply) */
	client, err := standarlize.DialHelloService("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
