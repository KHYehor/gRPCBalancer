package server

import (
	"container/ring"
	"context"
	"fmt"
	"github.com/KHYehor/gRPCBalancer/src/grpc/calculate"
	"github.com/KHYehor/gRPCBalancer/src/grpc/health"
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

func (s *LoadBalancer) InitServers(ctx context.Context, addresses []string) {
	s.servers = ring.New(len(addresses))
	for _, address := range addresses {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err == nil {
			server := Servers {
				grpcServer: calculate.NewCalculateMatrixClient(conn),
				address: address,
			}
			s.servers.Value = &server
			s.servers.Next()
		} else {
			fmt.Println(err)
		}
	}
}

func (s *LoadBalancer) getServerFromPool(ctx context.Context) (calculate.CalculateMatrixClient) {
	s.mutex.Lock()
	value := s.servers.Value.(Servers).grpcServer
	s.servers.Next()
	s.mutex.Unlock()
	return value
}

func (s *LoadBalancer) MatrixSum(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	server := s.getServerFromPool(ctx)
	fmt.Println(req)
	fmt.Println(req.Matrix1)
	fmt.Println(req.Matrix2)
	request := &calculate.MatrixRequest{Matrix1: req.GetMatrix1(), Matrix2: req.GetMatrix2()}
	response, err := server.MatrixSum(ctx, request)
	return response, err
	//matrix := []*calculate.Array{}
	//matrix = append(matrix, &calculate.Array{Digit: []float64{1, 2, 3}})
	//matrix = append(matrix, &calculate.Array{Digit: []float64{1, 2, 3}})
	//matrix = append(matrix, &calculate.Array{Digit: []float64{1, 2, 3}})
	//return &calculate.MatrixResponse{Matrix: matrix}, nil
}

func (s *LoadBalancer) MatrixMul(ctx context.Context, req *calculate.MatrixRequest) (*calculate.MatrixResponse, error) {
	s.mutex.Lock()
	s.mutex.Unlock()
	return nil, nil
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
		/* response */_, err := healthChecker.Health(context.Background(), request)
		if err != nil {
			//fmt.Println(response)
		}
	}
}
