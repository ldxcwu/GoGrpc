package standarlize

import "net/rpc"

const HelloServiceName = "standarlize.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (hc *HelloServiceClient) Hello(request string, reply *string) error {
	return hc.Client.Call(HelloServiceName+".Hello", request, reply)
}
