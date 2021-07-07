package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println(net.IPv4len)
	fmt.Println(net.IPv6len)

	fmt.Println(net.ErrClosed)
	fmt.Println(net.ErrWriteToConnected)

	res, _ := net.LookupAddr("yandex.ru")
	fmt.Println("LookupMX", res)

	res, _ = net.LookupAddr("yandex.ru")
	fmt.Println("LookupAddr", res)

	res2, _ := net.LookupCNAME("yandex.ru")
	fmt.Println("LookupCNAME", res2)

	res3, _ := net.LookupHost("yandex.ru")
	fmt.Println("LookupHost", res3)

	res4, _ := net.LookupIP("yandex.ru")
	fmt.Println("LookupIP", res4)

	res5, _ := net.LookupMX("yandex.ru")
	fmt.Println("LookupMX", res5)

	ln, err := net.Listen("tcp", "0.0.0.0:2121")
	if err != nil {
		log.Fatalf("cannot listen %v", err.Error())
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("cannot accept conn: %v", err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) { // TODO пока ничего не возвращает. это сервер, надо стучать клиентом?
	defer conn.Close()
	// вытащим данные
	// обработаем и запишем ответ
	// вернем
	conn.Write([]byte(fmt.Sprintf("Hello net, LocalAddr: %s, RemoteAddr: %s", conn.LocalAddr(), conn.RemoteAddr())))

}