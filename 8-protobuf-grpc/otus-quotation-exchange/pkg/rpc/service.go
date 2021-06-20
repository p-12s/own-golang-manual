package rpc

import (
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/api"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/pkg/types"
)

type Service struct {
	sources []types.Source
}

func NewService() *Service {
	return &Service{}
}

// AddSource - добавление новых ресурсов-бирж, кроме той, что в примере
func (s *Service) AddSource(source types.Source) {
	s.sources = append(s.sources, source)
}

// GetL2OrderBook реализация метода из proto-файла
func (s *Service) GetL2OrderBook(req *api.L2OrderBookRequest, stream api.Quotes_GetL2OrderBookServer) error {
	log.Printf("client connected")

	if req.Size <= 0 {
		return status.Error(codes.InvalidArgument, "invalid size")
	}
	if req.Interval <= 0 {
		return status.Error(codes.InvalidArgument, "invalid interval")
	}

	var (
		stop bool // флаг для выхода из цикла for
		l2   types.L2OrderBook
		err  error
	)
	for !stop {
		select {
		case <-time.After(time.Duration(req.Interval) * time.Millisecond):
			for _, source := range s.sources {
				l2, err = source.GetL2OrderBook(req.Symbol, int(req.Size))
				if err != nil {
					stop = true
					break // выйдет только из select, поэтому используем флаг stop
				}
				// отправляем клиенту данные
				l2Proto := ConvertToProtoL2(req.Symbol, l2)
				if err = stream.Send(l2Proto); err != nil {
					stop = true
				}
			}

		case <-stream.Context().Done():
			stop = true
		}
	}

	log.Printf("client disconnected")
	return err
}
