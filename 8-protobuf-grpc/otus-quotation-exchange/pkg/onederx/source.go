package onederx

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/pkg/types"
	"log"
	"sync"
	"time"
)

const (
	onederxWsUrl = "wss://api.onederx.com/v1/ws"

	defaultRetryInterval = time.Second
	defaultReadTimeout   = time.Second * 5
	defaultWriteTimeout  = time.Second * 5
)

// Source - конструктор
type Source struct {
	sync.RWMutex // обезопасим Source, потому что кроме получения-обновления данных к нему будут обращаться клиенты по gRPC
	l2BySymbol   map[string]*L2OrderBook
}

func NewSource() *Source {
	return &Source{
		l2BySymbol: make(map[string]*L2OrderBook),
	}
}

// Start запускает взаимодействие с биржей
func (s *Source) Start(ctx context.Context) {
	go func() {
		for {
			select {
			// при завершении приложения выходим из цикла
			case <-ctx.Done():
				log.Printf("source stopped: context cancelled")
				return
			// получим данные или запишем ошибку
			default:
				if err := s.receiveData(ctx); err != nil {
					log.Printf("receiving failed: %v\n", err)
				}
				log.Printf("sleep for %v\n", defaultRetryInterval)
				time.Sleep(defaultRetryInterval) // тут можно DDoS-ить ресурс. почему?
			}
		}
	}()
}

func (s *Source) GetL2OrderBook(symbol string, size int) (types.L2OrderBook, error) {
	s.RLock()
	defer s.RUnlock()

	l2, ok := s.l2BySymbol[symbol]
	if !ok {
		return types.L2OrderBook{}, fmt.Errorf("no data for symbol %s", symbol)
	}

	return types.L2OrderBook{
		Bid: l2.GetBid(size),
		Ask: l2.GetAsk(size),
	}, nil
}

func (s *Source) receiveData(ctx context.Context) error {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, onederxWsUrl, nil)
	if err != nil {
		return err
	}

	// у gorilla/websocket нет установленного по-умолчанию дедлайна, запрос может висеть долго
	if err := conn.SetWriteDeadline(time.Now().Add(defaultWriteTimeout)); err != nil {
		return err
	}
	// отправляем запрос - подпису на канал
	if err := conn.WriteJSON(GetWsL2SubscribeRequest()); err != nil {
		return err
	}

	// теперь вычитываем все, что нам приходит в ответ
	// Type - это snapshot, update, heartbeat
	var header struct{ Type string }
	for {
		select {
		// проверяем контекст, он всегда может "cancel-нуться"
		case <-ctx.Done():
			return nil
		// обрабатываем все, что приходит от биржи - heartbeat-ы, update-ы
		default:
			//
			if err := conn.SetReadDeadline(time.Now().Add(defaultReadTimeout)); err != nil {
				return err
			}

			mt, data, err := conn.ReadMessage()
			if err != nil {
				return err
			}

			if mt != websocket.TextMessage {
				return fmt.Errorf("unexpected message type %d", mt)
			}

			if err := json.Unmarshal(data, &header); err != nil {
				return err
			}

			switch header.Type {
			case "snapshot":
				err = s.onSnapshot(data)
			case "update":
				err = s.onUpdate(data)
			default: // иные типы сообщений - heartbeat и т.п. не обрабатываем
			}
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// onSnapshot - обработка первоначального снапшота (в теле может прийти и uppdate)
func (s *Source) onSnapshot(data []byte) error {
	var snapshot WsL2Snapshot
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return err
	}

	s.Lock() // лочим струкруту на время обновления
	defer s.Unlock()

	// получение снапшота происходит в начале подключения,
	// поэтому создаем новую структуру
	l2 := NewL2OrderBook()
	s.l2BySymbol[snapshot.Params.Symbol] = l2

	// при первом подключении вместе со снапшотом могут прийти и апдейты, поэтому
	// идем по пришедшим данным (snapshot, update) и применяем их
	for _, items := range [][]*WsL2Item{
		snapshot.Payload.Snapshot,
		snapshot.Payload.Updates,
	} {
		for _, item := range items {
			// применяем данные
			side := types.SideFromString(item.Side)
			tm := time.Unix(0, item.Timestamp)
			l2.Apply(item.Price, side, item.Volume, tm)
		}
	}

	log.Printf("snapshot applied: symbol=%s, bid=%d, ask=%d",
		snapshot.Params.Symbol, l2.bid.Len(), l2.ask.Len())

	return nil
}

// onUpdate - вторичная обработка обновлений
func (s *Source) onUpdate(data []byte) error {
	var update WsL2Update
	if err := json.Unmarshal(data, &update); err != nil {
		return err
	}

	s.Lock() // лочим струкруту на время обновления
	defer s.Unlock()

	l2, ok := s.l2BySymbol[update.Params.Symbol]
	if !ok {
		// почему-то ранее не пришел snapshot
		log.Printf("inconsistent update for symbol %s", update.Params.Symbol)
		return nil
	}

	side := types.SideFromString(update.Payload.Side)
	tm := time.Unix(0, update.Payload.Timestamp)
	l2.Apply(update.Payload.Price, side, update.Payload.Volume, tm)

	log.Printf("update applied: symbol=%s, bid=%d, ask=%d",
		update.Params.Symbol, l2.bid.Len(), l2.ask.Len())

	return nil
}
