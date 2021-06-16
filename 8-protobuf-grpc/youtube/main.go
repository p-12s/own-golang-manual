package main

import (
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

type GRPCServer struct{}

func (g *GRPCServer) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	return &AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func main() {
	s := grpc.NewServer()
	srv := &GRPCServer{}

	RegisterAdderServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := s.Server(l); err != nil {
		log.Fatalf(err.Error())
	}
}
