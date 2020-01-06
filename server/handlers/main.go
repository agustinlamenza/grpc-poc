package handlers

import (
	"context"
	"io"
	"log"
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
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func (s *Server) Average(stream api.CalculatorService_AverageServer) error {
	numbers := []int64{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var sum int64
			for _, v := range numbers {
				sum = sum + v
			}
			avr := float64(sum) / float64(len(numbers))
			res := &api.ArvResponse{Avr: avr}
			err := stream.SendAndClose(res)
			return err
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		numbers = append(numbers, req.GetNumber())
	}
}

func (s *Server) Max(stream api.CalculatorService_MaxServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v", err)
			return err
		}

		data := req.GetNumber()

		stream.Send(&api.MaxResponse{MaxNumber: data * data})
	}
}
