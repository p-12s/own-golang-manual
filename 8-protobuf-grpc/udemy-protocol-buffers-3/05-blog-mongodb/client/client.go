package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/05-blog-mongodb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	fmt.Println("Blog client run")

	isTls := false
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

	c := pb.NewBlogServiceClient(cc)

	doUnary(c)
}

func doUnary(c pb.BlogServiceClient) {
	fmt.Println("send blog from client")
	blog := &pb.Blog{
		AuthorId: "Mister",
		Title:    "My title",
		Content:  "My content",
	}
	res, err := c.CreateBlog(context.Background(), &pb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("error while calling CreateBlog %v\n", err)
	}
	log.Printf("\nblog has been created:\n%v\n", res.GetBlog())
}
