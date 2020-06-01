package server

import (
	"context"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"sync"
	"container/ring"
)

func New(servers *ring.Ring) *Server {
	return &Server{
		servers: servers,
	}
}

type Server struct {
	mutex sync.Mutex
	routingServer int
	servers *ring.Ring
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
