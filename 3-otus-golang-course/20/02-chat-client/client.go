package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func readRoutine(ctx context.Context, cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				cancel()
				break
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if !scanner.Scan() {
				break
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}

//TODO:

func main() {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*60*time.Second)

	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(ctx, cancel, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn)
		wg.Done()
	}()

	//time.Sleep(1 * time.Minute)
	//cancel()
	wg.Wait()
	conn.Close()
}
