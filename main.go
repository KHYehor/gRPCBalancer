package main

import (
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addresses = []string{
	"",
	"",
	"",
}

func main() {
	lis, err := net.Listen("tcp", ":32625")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	grpcServer := grpc.NewServer()
	s := server.LoadBalancer{}
	s.initServers(addresses)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
}
