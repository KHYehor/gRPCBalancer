package main

import (
	"container/ring"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	server "github.com/KHYehor/gRPCBalancer/src/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addresses = []string{"", "", ""}

func initClients() *ring.Ring {
	cliets2 := ring.New(len(addresses))
	for _, elemet := range addresses {
		conn, err := grpc.Dial(elemet)
		if err != nil {
			panic("error")
		}
		defer conn.Close()
		cliets2.Value = calculate.NewCalculateMatrixClient(conn)
		cliets2.Next()
	}
	return cliets2
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
