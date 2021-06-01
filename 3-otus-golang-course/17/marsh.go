package bench_parallel

import (
	"encoding/json"
)

type A struct {
	I int
}

func (a *A) Reset() {
	*a = A{}
}

// bad
func Slow2() {
	for i := 0; i < 1000; i++ {
		a := &A{}
		json.Unmarshal([]byte("{\"i\":32}"), a)
	}
}

// OK
func Fast2() {
	a := &A{}
	for i := 0; i < 1000; i++ {
		a.Reset()
		json.Unmarshal([]byte("{\"i\":32}"), a)
	}
}
