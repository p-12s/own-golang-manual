package bench_parallel

import "sync"

var mu = sync.Mutex{}

func FastParallel() int {
	acc := 0
	for i := 0; i < 1000; i++ {
		acc++
	}
	return acc
}

func SlowParallel() int {
	mu.Lock()
	defer mu.Unlock()

	acc := 0
	for i := 0; i < 1000; i++ {
		acc++
	}
	return acc
}
