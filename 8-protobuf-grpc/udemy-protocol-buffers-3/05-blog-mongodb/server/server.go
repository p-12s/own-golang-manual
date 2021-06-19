package main

import (
	"context"
	"fmt"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/udemy-protocol-buffers-3/05-blog-mongodb/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct{}

func (*server) CreateBlog(ctx context.Context, req *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	data := blogItem{
		AuthorId: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error: %v", err),
		)
	}

	objId, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("can't convert to objId: %v", err),
		)
	}

	return &pb.CreateBlogResponse{
		Blog: &pb.Blog{
			Id:       objId.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

var collection *mongo.Collection

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	// if we crashed the code, we get the file name and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog service started")

	ctx := context.Background()

	fmt.Println("DB connect")
	// uri := "mongodb://mongos0.example.com:27017,mongos1.example.com:27017/"
	var uri = "mongodb://localhost:27017"
	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("mongodb connect error: %v", err)
	}
	defer func() { _ = client.Disconnect(ctx) }()

	collection = client.Database("mydb").Collection("blog")

	fmt.Println("OK")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	isTls := false
	opts := []grpc.ServerOption{}
	if isTls {
		creds, err := credentials.NewServerTLSFromFile("../ssl/server.crt", "../ssl/server.pem")
		if err != nil {
			log.Fatalf("can't read cert files: %v", err)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Server start")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve %v", err)
		}
	}()

	// Ctr+C waiting to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal is received
	<-ch
	fmt.Println("Server trying to stop")
	s.Stop()
	lis.Close()
	fmt.Println("DB trying to stop")
	client.Disconnect(context.TODO())
	fmt.Println("Server and DB stopped")
}
