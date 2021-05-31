package service

// Batch - почка элементов
type Batch []Item

// Item - элемент, абстрактный тип
type Item struct{}

func CreateBatch(n uint64) Batch {
	var batch Batch
	batch = make([]Item, n)
	var i uint64 = 0
	for i = 0; i < n; i++ {
		batch[i] = Item{}
	}
	return batch
}
