// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QuotesClient is the client API for Quotes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuotesClient interface {
	// GetL2OrderBook получаем запрос клиента и отправляем ответ в стриме
	// Server stream
	GetL2OrderBook(ctx context.Context, in *L2OrderBookRequest, opts ...grpc.CallOption) (Quotes_GetL2OrderBookClient, error)
}

type quotesClient struct {
	cc grpc.ClientConnInterface
}

func NewQuotesClient(cc grpc.ClientConnInterface) QuotesClient {
	return &quotesClient{cc}
}

func (c *quotesClient) GetL2OrderBook(ctx context.Context, in *L2OrderBookRequest, opts ...grpc.CallOption) (Quotes_GetL2OrderBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &Quotes_ServiceDesc.Streams[0], "/api.Quotes/GetL2OrderBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &quotesGetL2OrderBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Quotes_GetL2OrderBookClient interface {
	Recv() (*L2OrderBook, error)
	grpc.ClientStream
}

type quotesGetL2OrderBookClient struct {
	grpc.ClientStream
}

func (x *quotesGetL2OrderBookClient) Recv() (*L2OrderBook, error) {
	m := new(L2OrderBook)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QuotesServer is the server API for Quotes service.
// All implementations should embed UnimplementedQuotesServer
// for forward compatibility
type QuotesServer interface {
	// GetL2OrderBook получаем запрос клиента и отправляем ответ в стриме
	// Server stream
	GetL2OrderBook(*L2OrderBookRequest, Quotes_GetL2OrderBookServer) error
}

// UnimplementedQuotesServer should be embedded to have forward compatible implementations.
type UnimplementedQuotesServer struct {
}

func (UnimplementedQuotesServer) GetL2OrderBook(*L2OrderBookRequest, Quotes_GetL2OrderBookServer) error {
	return status.Errorf(codes.Unimplemented, "method GetL2OrderBook not implemented")
}

// UnsafeQuotesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuotesServer will
// result in compilation errors.
type UnsafeQuotesServer interface {
	mustEmbedUnimplementedQuotesServer()
}

func RegisterQuotesServer(s grpc.ServiceRegistrar, srv QuotesServer) {
	s.RegisterService(&Quotes_ServiceDesc, srv)
}

func _Quotes_GetL2OrderBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(L2OrderBookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuotesServer).GetL2OrderBook(m, &quotesGetL2OrderBookServer{stream})
}

type Quotes_GetL2OrderBookServer interface {
	Send(*L2OrderBook) error
	grpc.ServerStream
}

type quotesGetL2OrderBookServer struct {
	grpc.ServerStream
}

func (x *quotesGetL2OrderBookServer) Send(m *L2OrderBook) error {
	return x.ServerStream.SendMsg(m)
}

// Quotes_ServiceDesc is the grpc.ServiceDesc for Quotes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Quotes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Quotes",
	HandlerType: (*QuotesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetL2OrderBook",
			Handler:       _Quotes_GetL2OrderBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "quotes.proto",
}
