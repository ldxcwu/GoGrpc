# 客户端 流式消息 发送

## 1. 定义消息格式
```proto
message RpcRequest {
    string Name = 1;
}

message RpcResponse {
    string value = 1;
}

service Greeter {
    rpc SayHello(stream RpcRequest) returns (RpcResponse) {};
}
```
>  在rcp方法里请求参数前添加stream关键字即可    
>  生成 .go 文件
```shell
protoc --go_out=. --go_opt=paths=source_relative service.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative service.proto
```
## 2. 实现服务端
### 2.1 定义结构体，实现服务接口
> 注意包含生成的go文件中Umimplexxx结构，因为注册服务需要提供实现对应接口的对象，而该结构实现了服务接口
1. 监听端口
2. 创建grpc服务器
3. 注册服务
4. 对监听到的链接提供服务
```go
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
```
## 3. 实现客户端
1. grpc 拨号获得链接（注意defer关闭）
2. 根据链接创建客户端
3. 调用服务获得一个流
4. 使用流发送消息
5. 发送最后调用CloseAndRecv获得服务器反馈
```go
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

```