package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := pb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c pb.CalculatorServiceClient) {
	fmt.Println("start calculator client doUnary()")
	req := &pb.SumRequest{
		FirstNumber:  -2.22,
		SecondNumber: 1.1,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculator client doUnary() %v", err)
	}
	log.Printf("\nend calculator client doUnary() with result:\n%v\n", res.Result)
}

func doServerStreaming(c pb.CalculatorServiceClient) {
	fmt.Println("start calculator client doServerStreaming()")
	req := &pb.PrimeNumberDecompositionRequest{
		Number: 12,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculator client doServerStreaming() %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end data
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream in doServerStreaming %v", err)
		}
		log.Printf("%v\n", msg.GetPrimeFactor())
	}
	log.Printf("\nend doServerStreaming\n")
}
