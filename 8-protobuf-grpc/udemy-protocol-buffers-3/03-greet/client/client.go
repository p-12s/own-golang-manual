package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/03-greet/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"time"
)

func main() {

	isTls := true
	opts := grpc.WithInsecure()
	if isTls {
		creds, err := credentials.NewClientTLSFromFile("../ssl/ca.crt", "")
		if err != nil {
			log.Fatalf("could not read cert files: %v", err)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := pb.NewGreetServiceClient(cc)

	doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
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

func doBiDirectionalStreaming(c pb.GreetServiceClient) {
	fmt.Println("start doBiDirectionalStreaming")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream %v", err)
	}

	request := []*pb.GreetEveryoneRequest{
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "Ilay",
				LastName:  "Smith",
			},
		},
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "John",
				LastName:  "Pack",
			},
		},
		&pb.GreetEveryoneRequest{
			Greeting: &pb.Greeting{
				FirstName: "Mike",
				LastName:  "Longer",
			},
		},
	}

	waitChan := make(chan struct{})

	// send a bunch of messages
	go func() {
		for _, req := range request {
			stream.Send(req)
			fmt.Printf("send mess %v\n", req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive a bunch of messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving stream %v", err)
				break
			}
			fmt.Printf("recieved %v\n", res.GetResult())
		}
		close(waitChan)
	}()

	<-waitChan
}

func doClientStreaming(c pb.GreetServiceClient) {
	fmt.Println("start doClientStreaming")
	request := []*pb.LongGreetRequest{
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Ilay",
				LastName:  "Smith",
			},
		},
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "John",
				LastName:  "Pack",
			},
		},
		&pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: "Mike",
				LastName:  "Longer",
			},
		},
	}

	resStream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling doClientStreaming %v", err)
	}

	for i, req := range request {
		fmt.Printf("send req: %d\n", i)
		resStream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := resStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while closing client stream doClientStreaming %v", err)
	}
	log.Printf("\nend doClientStreaming %v\n", res)
}
