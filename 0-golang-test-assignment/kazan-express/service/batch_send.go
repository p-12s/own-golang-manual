package service

// Batch - почка элементов
type Batch []Item

// Item - элемент, абстрактный тип
type Item struct{}

func CreateBatch(n uint64) Batch {
	index := int(n)
	var batch Batch = make([]Item, index)
	for i := 0; i < index; i++ {
		batch[i] = Item{}
	}
	return batch
}
