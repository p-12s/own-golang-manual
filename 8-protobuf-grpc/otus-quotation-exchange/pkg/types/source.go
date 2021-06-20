package types

// Source - один из источников данных
type Source interface {
	GetL2OrderBook(symbol string, size int) (L2OrderBook, error)
}
