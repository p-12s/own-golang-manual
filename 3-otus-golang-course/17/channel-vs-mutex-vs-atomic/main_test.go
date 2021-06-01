package benchmark

import (
	"context"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Скопируйте в любую (желательно пустую) директорию и запустите:
//	$ go test -v -bench=. .
// Пример вывода:
//	goos: linux
//	goarch: amd64
//	BenchmarkChannel-8   	 3000000	       429 ns/op
//	BenchmarkMutex-8     	10000000	       204 ns/op
//	BenchmarkAtomic-8    	10000000	       165 ns/op

// Значение, которое будет редактироваться из N горутин.
var sharedValue int64

var valueMutex sync.Mutex

func addAtomic(x int64) {
	atomic.AddInt64(&sharedValue, x)
}

// Функция, которая будет заниматься атомарным обновлением
// через операцию send+receive. Только одна горутина обновляет глобальную
// переменную - та, что держит канал на recveive.
func channelUpdater(ch <-chan int64) {
	for x := range ch {
		sharedValue += x
	}
}

func addWithLocking(x int64) {
	valueMutex.Lock()
	sharedValue += x
	valueMutex.Unlock()
}

func runBenchmark(b *testing.B, fn func()) {
	sharedValue = 0
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					fn()
				}
			}
		}()
	}

	for int(sharedValue) <= b.N {
		time.Sleep(100)
	}
	cancel()
	wg.Wait()
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int64)
	go channelUpdater(ch)

	var sendOnly chan<- int64 = ch
	runBenchmark(b, func() {
		sendOnly <- 1
	})

	close(ch)
}

func BenchmarkMutex(b *testing.B) {
	runBenchmark(b, func() {
		addWithLocking(1)
	})
}

func BenchmarkAtomic(b *testing.B) {
	runBenchmark(b, func() {
		addAtomic(1)
	})
}
