package main

import (
	"context"
	"fmt"
	"github.com/zhashkevych/scheduler"
	"time"
)

func main() {
	s := scheduler.NewScheduler()
	s.Add(context.Background(), func(ctx context.Context) {
		fmt.Printf("Time is: %s\n", time.Now())
	}, time.Second*1)
	time.Sleep(time.Minute)
}
