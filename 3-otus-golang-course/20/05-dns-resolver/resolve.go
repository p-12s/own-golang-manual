package main

import (
	"context"
	"log"
	"net"
)

func main() {
	resolver := net.Resolver{
		PreferGo:     true, //использовать гошный резолвер
		StrictErrors: false,
		Dial:         nil,
	}

	ctx := context.Background()

	addrs, err := resolver.LookupIPAddr(ctx, "yandex.ru")
	if err != nil {
		log.Fatalf("cannot lookup addr: %v", err)
	}

	//SHOW DNS BALANCING
	for _, addr := range addrs {
		log.Printf("addr of yandex.ru: %s", addr)
	}

	mxes, err := resolver.LookupMX(ctx, "yandex.ru")
	if err != nil {
		log.Fatalf("cannot lookup mx: %v", err)
	}

	//SHOW host command
	for _, mx := range mxes {
		log.Printf("mx of yandex.ru: %s (pref=%d)", mx.Host, mx.Pref)
	}

	//Show spf
	txts, err := resolver.LookupTXT(ctx, "yandex.ru")
	if err != nil {
		log.Fatalf("cannot lookup mx: %v", err)
	}

	for _, txt := range txts {
		log.Printf("txt of yandex.ru: %s", txt)
	}
}
