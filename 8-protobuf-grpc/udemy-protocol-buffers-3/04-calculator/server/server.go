package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.ResultResponse, error) {
	fmt.Printf("\nserver action Sum(): %v\n", req)
	return &pb.ResultResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil
}

func (*server) PrimeNumberDecomposition(req *pb.PrimeNumberDecompositionRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("\nserver action PrimeNumberDecomposition(): %v\n", req)
	number := req.GetNumber()
	var divisor int64 = 2

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("divisor++ = %v", divisor)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	fmt.Println("server serve OK")
}
