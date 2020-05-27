package main

import (
	calculate "github.com/KHYehor/gRPCBalancer/src/grpc"
	server "github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"sync"
)

var addresses = []string{"", "", ""}

func initClients() *[]calculate.CalculateMatrixClient {
	var clients []calculate.CalculateMatrixClient
	for i := 0; i < 3; i++ {
		conn, err := grpc.Dial(addresses[i])
		if err != nil {
			panic("error")
		}
		client := calculate.NewCalculateMatrixClient(conn)
		clients = append(clients, client)
	}
	return &clients
}

func main() {
	clients := initClients()
	server.Server{sync.Mutex, 0, clients}
}
