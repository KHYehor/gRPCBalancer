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
	"127.0.0.1:5000",
}

var healthAddresses = []string{
	"127.0.0.1:6000",
}

// start Load Balancer with routing
func startGrpcBalancer(host string, addresses []string) (error) {
	// Start tcp listening port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	// Create server service
	grpc.WithInsecure()
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
	err := startGrpcBalancer("127.0.0.1:3000", addresses)
	if err != nil {
		panic(err)
	}
}
