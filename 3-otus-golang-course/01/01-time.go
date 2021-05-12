package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		log.Fatal(err)
	}
	const layout = "3:04:05 PM (MST) on Monday, January _2, 2006"
	fmt.Println("Время:")
	fmt.Println(time.Local().Format(layout))
}
