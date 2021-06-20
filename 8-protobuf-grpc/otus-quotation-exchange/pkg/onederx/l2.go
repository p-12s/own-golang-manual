package onederx

import (
	"time"

	"github.com/HuKeping/rbtree"
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/otus-quotation-exchange/pkg/types"
	"github.com/shopspring/decimal"
)

type L2OrderBookItem types.L2OrderBookItem

// Less - у используемой библиотеки красно-черного дерева
// узел должен реализовывать метод Less
func (item L2OrderBookItem) Less(than rbtree.Item) bool {
	return item.Price.LessThan(than.(*L2OrderBookItem).Price)
}

// L2OrderBook - структура на основе красно-черного дерева для быстрого
// удаления/вставки update
type L2OrderBook struct {
	bid, ask *rbtree.Rbtree
}

func NewL2OrderBook() *L2OrderBook {
	return &L2OrderBook{
		bid: rbtree.New(),
		ask: rbtree.New(),
	}
}

// Apply - получение данных из биржи
func (ob *L2OrderBook) Apply(price decimal.Decimal, side types.SideType, volume uint64, tm time.Time) {
	// сначала определяем, какую сторону стакана обрабатываем
	obs := ob.bid
	if side == types.SideAsk {
		obs = ob.ask
	}
	// вставляем в дерево
	item := obs.InsertOrGet(&L2OrderBookItem{Price: price})

	// для цены объем пришел = 0, значит его съели. Удаляем его, чтоб не хранить
	if volume == 0 {
		obs.Delete(item)
		return
	}

	// применяем обновления
	item.(*L2OrderBookItem).Volume = volume
	item.(*L2OrderBookItem).Time = tm
}

// GetBid - получаем срез "покупателей"
// от максимальной цены вниз, пока не заполним требуемый размер
func (ob *L2OrderBook) GetBid(size int) []*types.L2OrderBookItem {
	ret := make([]*types.L2OrderBookItem, 0, size)

	ob.bid.Descend(ob.bid.Max(), func(item rbtree.Item) bool {
		itemCopy := types.L2OrderBookItem(*item.(*L2OrderBookItem))
		ret = append(ret, &itemCopy)
		size--
		return size != 0
	})

	return ret
}

// GetAsk - получаем срез "продавцов"
// от минимальной цены вверх
func (ob *L2OrderBook) GetAsk(size int) []*types.L2OrderBookItem {
	ret := make([]*types.L2OrderBookItem, 0, size)

	ob.ask.Ascend(ob.ask.Min(), func(item rbtree.Item) bool {
		itemCopy := types.L2OrderBookItem(*item.(*L2OrderBookItem))
		ret = append(ret, &itemCopy)
		size--
		return size != 0
	})

	return ret
}
