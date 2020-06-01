package main

import (
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	server "github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addresses = []string{"", "", ""}

func initClients() *[]calculate.CalculateMatrixClient {
	var clients []calculate.CalculateMatrixClient
	for i := 0; i < 3; i++ {
		conn, err := grpc.Dial(addresses[i])
		if err != nil {
			panic("error")
		}
		defer conn.Close()
		client := calculate.NewCalculateMatrixClient(conn)
		clients = append(clients, client)
	}
	return &clients
}

func main() {
	lis, err := net.Listen("tcp", ":32625")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	clients := initClients()
	grpcServer := grpc.NewServer()
	s := server.New(clients)
	calculate.RegisterCalculateMatrixServer(grpcServer, s)
}
