package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type key int

const (
	userIdCtx key = 0
)

// ================= прерывание выполнения запросов

func processLongTask(ctx context.Context) { // context всегда передается 1-м параметром и с названием ctx context.Context

	select { // select тут для того, чтобы читать данные без блокировок из нескольких каналов
	case <-time.After(2 * time.Second):
		fmt.Println("the process ended on its own")
	case <-ctx.Done():
		fmt.Println("request canceled with user input")
	}
}

func userCancelProcess() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()

	processLongTask(ctx)
}

// ================= передача специфичных для контекста параметров

func processLongTask2(ctx context.Context) { // context всегда передается 1-м параметром и с названием ctx context.Context
	id := ctx.Value(userIdCtx)

	select { // select тут для того, чтобы читать данные без блокировок из нескольких каналов
	case <-time.After(2 * time.Second):
		fmt.Println("the process ended on its own", id)
	case <-ctx.Done():
		fmt.Println("request canceled with user input")
	}
}

func sentSpecificParamIntoContext() {
	ctx := context.WithValue(context.Background(), userIdCtx, 42)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()

	processLongTask2(ctx)
}

// =================

func handle(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User-Id")
	ctx := context.WithValue(r.Context(), userIdCtx, id)

	result := processLongTask3(ctx)
	w.Write([]byte(result))
}

func processLongTask3(ctx context.Context) string { // context всегда передается 1-м параметром и с названием ctx context.Context
	id := ctx.Value(userIdCtx)

	select { // select тут для того, чтобы читать данные без блокировок из нескольких каналов
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("the process ended on its own %i\n", id)
	case <-ctx.Done():
		log.Println("request canceled with user input")
		return ""
	}
}

func main() {
	// userCancelProcess() // прерывание выполнения запросов
	// sentSpecificParamIntoContext() // передача специфичных для контекста параметров

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


