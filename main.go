package main

import (
	"google.golang.org/grpc"
	"github.com/KHYehor/gRPCBalancer/src/server"
)



func main() {
	conn, err := grpc.Dial("")
	if err != nil {

	}
	client := server.NewCa(conn)
}
