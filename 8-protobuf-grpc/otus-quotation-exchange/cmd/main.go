package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/api"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/pkg/onederx"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/pkg/rpc"
	"google.golang.org/grpc"
)

func main() {
	// создаем и подключаемся к источнику
	fmt.Println("create onederx")
	onederx := onederx.NewSource()
	fmt.Println("create onederx start")
	onederx.Start(context.Background())
	fmt.Println("create onederx end")

	// создаем сервис gRPC
	fmt.Println("create source middle")
	service := rpc.NewService()
	service.AddSource(onederx) //не реализ метод от интерф?
	fmt.Println("create source end")

	server := grpc.NewServer()
	api.RegisterQuotesServer(server, service)
	fmt.Println("register service on server end")

	reflection.Register(server)
	fmt.Println("evans reflection end")

	// запустим сервер
	lsn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server started OK")
}
