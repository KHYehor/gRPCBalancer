package server

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"sync"
	"container/ring"
)

func New(servers *[]calculate.CalculateMatrixClient) *Server {
	return &Server{
		servers: servers,
	}
}

type Server struct {
	mutex sync.Mutex
	routingServer int
	serverss *[]*ring.Ring
	servers *[]calculate.CalculateMatrixClient
}

func (s *Server) MatrixSum(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
}

func (s *Server) MatrixMul(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
}
