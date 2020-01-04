package handlers

import (
	"context"

	"grpc-poc/api"
)

// Server is the struct used for this grpc server
type Server struct{}

func (s *Server) Sum(ctx context.Context, in *api.SumRequest) (*api.SumResponse, error) {
	result := in.GetX() + in.GetY()

	return &api.SumResponse{Result: result}, nil
}
