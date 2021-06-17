package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/03-greet/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	return &pb.GreetResponse{
		Result: "Hello " + firstName + " " + lastName,
	}, nil
}

func (*server) GreetManyTimes(req *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		stream.Send(&pb.GreetManyTimesResponse{
			Result: "Hello " + firstName + ", this is #" + strconv.Itoa(i+1),
		})
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// finish
			return stream.SendAndClose(&pb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}

		result += req.GetGreeting().GetFirstName() + "! "
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	fmt.Println("server serve OK")
}
