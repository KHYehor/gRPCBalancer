package server

import (
	"container/ring"
	"context"
	"fmt"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/grpcBalancer/src/grpc/health"
	"google.golang.org/grpc"
	"sync"
	"time"
)

type Servers struct {
	address string
	grpcServer calculate.CalculateMatrixClient
}

type LoadBalancer struct {
	mutex sync.Mutex
	servers *ring.Ring
}

func (s *LoadBalancer) MatrixSum(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	s.mutex.Lock()
	server := s.servers.Value
	fmt.Println(server)
	s.mutex.Unlock()
	return nil, nil
}

func (s *LoadBalancer) MatrixMul(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
}

func (s *LoadBalancer) InitServers(ctx context.Context, addresses []string) {
	s.servers = ring.New(len(addresses))
	for _, address := range addresses {
		conn, err := grpc.Dial(address)
		if err == nil {
			server := Servers{
				grpcServer: calculate.NewCalculateMatrixClient(conn),
				address: address,
			}
			s.servers.Value = server
			s.servers.Next()
		}
	}
}

func (s *LoadBalancer) StartCheckHealth(ctx context.Context, addresses []string) {
	conn, err := grpc.Dial("")
	if err != nil {
		panic("error")
	}
	defer conn.Close()
	healthChecker := health.NewCheckHealthClient(conn)
	request := &health.HealthRequest{}
	for range time.Tick(time.Second * 1) {
		_, err := healthChecker.Health(context.Background(), request)
		if err != nil {
			// rebuild ring function
		}
	}
}
