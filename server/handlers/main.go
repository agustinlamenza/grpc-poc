package handlers

import (
	"context"
	"time"

	"grpc-poc/api"
)

// Server is the struct used for this grpc server
type Server struct{}

func (s *Server) Sum(ctx context.Context, req *api.SumRequest) (*api.SumResponse, error) {
	result := req.GetX() + req.GetY()

	return &api.SumResponse{Result: result}, nil
}

func (s *Server) Fibonacci(req *api.FibonacciRequest, stream api.CalculatorService_FibonacciServer) error {
	n := int64(req.GetNumber())
	for i, j := 0, 1; int64(i) < n; i, j = i+j, i {
		f := &api.FibonacciResponse{Number: int64(i)}
		stream.Send(f)
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}
