package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/04-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer cc.Close()

	c := pb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDirectionalStreaming(c)
	/*doErrorUnary(c, 122)
	doErrorUnary(c, 10)
	doErrorUnary(c, 0)
	doErrorUnary(c, -9)*/
	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 1*time.Second)
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
	fmt.Println("start calculator server streaming doServerStreaming()")
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

func doClientStreaming(c pb.CalculatorServiceClient) {
	fmt.Println("start calculator client streaming doClientStreaming()")

	resStream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling calculator client doServerStreaming() %v", err)
	}

	numbers := []int32{2, 3, 4}

	for _, number := range numbers {
		resStream.Send(&pb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := resStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while closing client stream doClientStreaming %v", err)
	}
	log.Printf("\nend doClientStreaming %v\n", res)
}

func doBiDirectionalStreaming(c pb.CalculatorServiceClient) {
	fmt.Println("start calculator client streaming doBiDirectionalStreaming()")

	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("error while calling calculator client doBiDirectionalStreaming() %v", err)
	}

	waitChan := make(chan struct{})

	// send
	go func() {
		numbers := []int32{320, 3, 43, 1000, 500, 21, 12, -3245}
		for _, req := range numbers {
			fmt.Printf("sending number: %v\n", req)
			stream.Send(&pb.FindMaxRequest{
				Number: req,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving stream in doBiDirectionalStreaming %v", err)
				break
			}
			max := res.GetMax()
			fmt.Printf("received a new max of: %v\n", max)
		}
		close(waitChan)
	}()

	<-waitChan
}

func doErrorUnary(c pb.CalculatorServiceClient, number int32) {
	fmt.Println("start calculator client doErrorUnary()")
	req := &pb.SquareRootRequest{
		Number: number,
	}
	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("we sent negative number")
				return
			}
		} else {
			log.Fatalf("error while calling calculator client doErrorUnary() %v", err)
			return
		}
	}
	log.Printf("\nnumber %d square is:\n%v\n", number, res.GetNumberRoot())
}

func doUnaryWithDeadline(c pb.CalculatorServiceClient, timeout time.Duration) {
	fmt.Println("start calculator client doUnaryWithDeadline()")
	req := &pb.DeadlineRequest{
		Number: 10,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.DeadlineExample(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Printf("timeout was hit: %v", err)
			} else {
				log.Printf("unexpected error doUnaryWithDeadline() %v", statusErr)
			}
		} else {
			log.Fatalf("error while calling calculator client doUnaryWithDeadline() %v", err)
		}
		return
	}
	log.Printf("\nend calculator client doUnaryWithDeadline() with result:\n%v\n", res.NumberRoot)
}
