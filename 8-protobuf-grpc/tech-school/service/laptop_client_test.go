package service_test

import (
	"context"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/pb"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/sample"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/serializer"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore()) // -> LaptopServiceServer

	grpcServer := grpc.NewServer()
	//RegisterLaptopServiceServer(s grpc.ServiceRegistrar, srv LaptopServiceServer) {
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1 := serializer.ProtobufToJSON(laptop1)
	json2 := serializer.ProtobufToJSON(laptop2)

	require.Equal(t, json1, json2)
}
