package main

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"net"
)

// List of all available servers for routing
var addresses = []string{
	"",
	"",
	"",
}

var healthAddresses = []string{
	"",
	"",
	"",
}

// start Load Balancer with routing
func startGrpcBalancer(host string, addresses []string) (error) {
	// Start tcp listening port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	// Create server service
	grpcServer := grpc.NewServer()
	s := &server.LoadBalancer{}
	// Init servers for routing
	s.InitServers(context.Background(), addresses)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
	// Attach listener to server
	grpcServer.Serve(lis)
	return nil
}

func main() {
	err := startGrpcBalancer("127.0.0.1:5000", addresses)
	if err != nil {
		panic(err)
	}
}
