package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/05-blog-mongodb/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
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

	createPost(c)

	/*readPost(c, "60ce3459872ab0ef155593a0")
	updatedData := &pb.Blog{
		Id:       "60ce3459872ab0ef155593a0",
		AuthorId: "Mister 2 NEW",
		Title:    "My title 2 NEW NEW",
		Content:  "My content 2 NEW NEW NEW",
	}
	updatePost(c, updatedData)
	readPost(c, "60ce3459872ab0ef155593a0")

	deletePost(c, "60ce0bfa0ee96f3c8e6e15df")*/
	getList(c)
}

func createPost(c pb.BlogServiceClient) {
	fmt.Println("create Post")
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
	log.Printf("\nPost has been created:\n%v\n", res.GetBlog())
}

func readPost(c pb.BlogServiceClient, blogId string) {
	fmt.Println("read Post")
	res, err := c.ReadBlog(context.Background(), &pb.ReadBlogRequest{
		BlogId: blogId,
	})
	if err != nil {
		log.Fatalf("error while read post %v\n", err)
	}
	fmt.Println(res)
}

func updatePost(c pb.BlogServiceClient, blog *pb.Blog) {
	fmt.Println("update Post")
	res, err := c.UpdateBlog(context.Background(), &pb.UpdateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("error while calling updatePost %v\n", err)
	}
	log.Printf("\nPost has been updated:\n%v\n", res.GetBlog())
}

func deletePost(c pb.BlogServiceClient, blogId string) {
	fmt.Println("delete Post")
	res, err := c.DeleteBlog(context.Background(), &pb.DeleteBlogRequest{
		BlogId: blogId,
	})
	if err != nil {
		log.Fatalf("error while delete post %v\n", err)
	}
	fmt.Println(res)
}

func getList(c pb.BlogServiceClient) {
	fmt.Println("get list of Posts")
	stream, err := c.ListBlog(context.Background(), &pb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error 1 while getting list of posts %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error 2 while getting list of posts: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
