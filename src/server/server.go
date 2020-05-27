package server

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc"
	"sync"
)

type Structures struct {
	server1 string "server1"
	server2 string "server2"
	server3 string "server3"
}

type Server struct {
	mutex sync.Mutex
	routingServer int
}

func (s *Server) MatrixSum(ctx context.Context, req *grpc.MatrixRequest) (*grpc.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
}

func (s *Server) MatrixMul(ctx context.Context, req *grpc.MatrixRequest) (*grpc.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
}
