package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/03-greet/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := pb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c pb.GreetServiceClient) {
	fmt.Println("start doUnary")
	req := &pb.GreetRequest{
		Greeting: &pb.Greeting{
			FirstName: "Stepan",
			LastName:  "Razin",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeting %v", err)
	}
	log.Printf("\nend doUnary with result:\n%v\n", res.Result)
}
