package server

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc"
	"sync"
)

type Server struct {
	mutex sync.Mutex
	routingServer int
	servers *[]grpc.CalculateMatrixClient
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
