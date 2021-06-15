package main

import (
	"context"
	"flag"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/pb"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server:", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()

	// send empty Id (should create)
	//laptop.Id = ""

	// send already exists id
	//laptop.Id = "e9deb34c-f005-4343-991c-7ee38183af5b"

	// send invalid
	//laptop.Id = "invalid"

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout - cannot create laptop:rpc error: code = DeadlineExceeded desc = context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop:", err)
		}
		return
	}

	log.Printf("create laptop eith id: %s", res.Id)
}
