package main

import (
	"fmt"
	"time"
)

type Sched struct {
	n int
}

func (s *Sched) Start() {
	ticker := time.Tick(100 * time.Millisecond)

	for range ticker {
		s.n++
	}
}

func (s *Sched) Shutdown() {
	// завершение работающих горутин, чтобы взаимодействие с юзером было без ошибок
}

func main() {
	ticker := time.Tick(100 * time.Millisecond)
	fmt.Println(ticker)
}
