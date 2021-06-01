package bench_parallel

import (
	"testing"
)

func BenchmarkSlow2(b *testing.B) {

	// тут можно сделать работу по-инициализации контекста, время кторого не должно учитываться в бенчмарке
	b.ResetTimer() // сбросили таймер

	for i := 0; i < b.N; i++ { // вызываем бенчмарк в цикле, чтобы была статистика
		Slow2()
	}
}

func BenchmarkFast2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast2()
	}
}
