package main

import (
	"fmt"
	"json-rpc/protocs"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	s := rpc.NewServer()

	listener, err := net.Listen("tcp", ":")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
	defer listener.Close()

	fmt.Printf("listen Addr: %s",listener.Addr())

	err = s.Register(&HelloService{})
	if err != nil {
		log.Fatalf("Register error: %v", err)
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Fatalf("accept error: %v", err)
		}
	
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

func client(addr string) {
	client,err := net.DialTimeout("tcp", addr,3*time.Second)
	if err != nil {
		log.Fatalf("client dial error: %v", err)
	}

	defer client.Close()

	rpcClient := jsonrpc.NewClient(client)

	req := protocs.Request{
		Value: "mxy",
	}
	
	var resp protocs.Reply
	err = rpcClient.Call("HelloService.Say",&req,&resp)
	if err != nil {
		log.Fatalf("json rpc call error: %v", err)
	}
}

type HelloService struct {}

func (h *HelloService) Say(req *protocs.Request,rep *protocs.Reply) error {
	fmt.Printf("req: %s\n", req.Value)

	rep.Value = "hello "+ req.Value
	return nil
}
