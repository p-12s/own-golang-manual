package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := pb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c pb.CalculatorServiceClient) {
	fmt.Println("start calculator client doUnary()")
	req := &pb.SumRequest{
		FirstNumber:  5.51,
		SecondNumber: 2.22,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculator client doUnary() %v", err)
	}
	log.Printf("\nend calculator client doUnary() with result:\n%v\n", res.Result)
}
