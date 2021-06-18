package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (*server) ComputeAverage(stream pb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("\nserver action ComputeAverage()\n")

	var sum int32
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func (*server) FindMax(stream pb.CalculatorService_FindMaxServer) error {
	fmt.Printf("\nserver action FindMax()\n")

	var max int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
			return err
		}
		number := req.GetNumber()
		if number > max {
			max = number
			sendErr := stream.Send(&pb.FindMaxResponse{
				Max: max,
			})
			if sendErr != nil {
				log.Fatalf("error while sending client stream %v", err)
				return err
			}
		}

	}
}

func (*server) SquareRoot(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	fmt.Printf("\nserver action SquareRoot()\n")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("received negative number: %d\n", number))
	}
	return &pb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func (*server) DeadlineExample(ctx context.Context, req *pb.DeadlineRequest) (*pb.DeadlineResponse, error) {
	fmt.Printf("\nserver action DeadlineExample()\n")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// client canceled
			fmt.Println("the client canceled the request")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("received negative number: %d\n", number))
	}
	return &pb.DeadlineResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	// register reflection service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	fmt.Println("server serve OK")
}
