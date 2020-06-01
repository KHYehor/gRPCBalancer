package main

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"net"
)

var addresses = []string{
	"",
	"",
	"",
}

func startListening(address string) {
	lis, err := net.Listen("tcp", address)
	defer lis.Close()
	if err != nil {
		panic(err)
	}
}

func startLoadBalancer(addresses []string) {
	grpcServer := grpc.NewServer()
	s := &server.LoadBalancer{}
	s.InitServers(context.Background(), addresses)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
}

func main() {
	startListening("")
	startLoadBalancer(addresses)
}
