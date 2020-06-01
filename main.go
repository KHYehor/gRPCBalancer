package main

import (
	"container/ring"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	server "github.com/KHYehor/gRPCBalancer/src/server"
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
	s := server.NewLoadBalabcer()
	s.initServers(addresses)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
}
