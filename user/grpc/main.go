package main

import (
	"google.golang.org/grpc"

	"myblogs/user/grpc/pbfile"

	"fmt"
	"net"
)

var (
	serverAddrs = "localhost:8000"
)

func main() {
	server := grpc.NewServer()

	pbfile.RegisterUserServiceServer(server, &pbfile.UserService{})
	listener, err := net.Listen("tcp", serverAddrs)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("service start working at ", serverAddrs)
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}