package main

import (
	"google.golang.org/grpc"

	"grpc_article/pbfile"
	"grpc_article/etc"

	"net"
	"fmt"
)

var (
	serverAddrs = etc.My_Config.ListenOn.Host + ":" + fmt.Sprint(etc.My_Config.ListenOn.Port)
)

func main() {
	server := grpc.NewServer()

	pbfile.RegisterArticleServiceServer(server, &pbfile.ArticleService{})

	listener, err := net.Listen("tcp", serverAddrs)
	if err != nil {
		fmt.Printf("listen on %s error!\n", serverAddrs)
		panic(err)
	}

	fmt.Println("start listening on", serverAddrs)
	server.Serve(listener)
}