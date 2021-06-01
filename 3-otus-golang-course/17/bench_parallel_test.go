package bench_parallel

import "testing"

func BenchmarkParallelFastParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fast()
		}
	})
}

func BenchmarkParallelSlowParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Slow()
		}
	})
}
