package main

import (
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/05-blog-mongodb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type server struct{}

func main() {
	fmt.Println("Blog service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	isTls := false
	opts := []grpc.ServerOption{}
	if isTls {
		creds, err := credentials.NewServerTLSFromFile("../ssl/server.crt", "../ssl/server.pem")
		if err != nil {
			log.Fatalf("can't read cert files: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
