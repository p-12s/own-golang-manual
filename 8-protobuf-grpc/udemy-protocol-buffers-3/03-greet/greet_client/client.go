package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/03-greet/pb"
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

	c := pb.NewGreetServiceClient(cc)

	//doUnary(c)
	doServerStreaming(c)
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

func doServerStreaming(c pb.GreetServiceClient) {
	fmt.Println("start doServerStreaming")
	req := &pb.GreetManyTimesRequest{
		Greeting: &pb.Greeting{
			FirstName: "Petro",
			LastName:  "Malikov",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling doServerStreaming %v", err)
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
		log.Printf("%v\n", msg.GetResult())
	}
	log.Printf("\nend doServerStreaming\n")
}
