package types

import (
	"github.com/shopspring/decimal"
	"time"
)

type SideType int8

const (
	SideBid SideType = iota
	SideAsk
)

// SideFromString возвращает тип действия - покупка или продажа
func SideFromString(s string) SideType {
	if s == "buy" {
		return SideBid
	}
	return SideAsk
}

type L2OrderBookItem struct {
	Price  decimal.Decimal // не эффективно с точки зрения хранения данных, но нам нужна большая точность
	Volume uint64
	Time   time.Time
}

// L2OrderBook - структура для хранения Bid/Ask
// здесь проблема в том, что мы все время получаем апдейты,
// а значит нужно их вставлять в середину массива
// Это не производительно при вставке/удалении!
// Для решения проблемы создана другая структура на основе красно-черных деревьях:
// pkg/onederx/l2.go
type L2OrderBook struct {
	Bid, Ask []*L2OrderBookItem
}
