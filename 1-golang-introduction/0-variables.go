package main

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
)

type Counters struct {
	sync.Mutex
	m map[string]int
}

func main() {
	//var mu sync.Mutex
	var counters = map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int) {
			for j := 5; j < 5; j++ {
				//mu.Lock()
				counters[th*10+j]++
				//mu.Unlock()
			}
		}(counters, i)
	}
	fmt.Scanln()
	fmt.Println("counters result", counters)

	var s string
	s = "Hello"
	fmt.Printf("%v %s %v %p %d\n", s, s, &s, s, len(s))

	s = strings.Replace(s, "e", "🚳", 1)
	fmt.Printf("%v %s %v %p %d\n", s, s, &s, s, len(s))

	s = "new str"
	fmt.Printf("%v %s %v %p %d\n", s, s, &s, s, len(s))

	var a unsafe.Pointer
	fmt.Printf("%v %s %v %p\n", a, a, &a, a)

	fmt.Println("========")
	var p string
	p = "NEW"
	fmt.Printf("%v %s | %v %v | %p %d\n", p, p, &p, *&p, p, len(p))

	p = strings.Replace(p, "N", "🚳", 1)
	fmt.Printf("%v %s | %v %v | %p %d\n", p, p, &p, *&p, p, len(p))

	// Types
	/*var num2, num3 float64 = 2.31, 42.33223
	num4 := 42
	fmt.Println(num2, num3, num4)

	// Math
	fmt.Println(math.Floor(2.15))
	fmt.Println(strings.Title("Заголовок"))
	fmt.Println('A') // тип данных руны - символ в юникоде

	// Reflect
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.43))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("This is str"))*/

	// int32 Range: -2147483648 through 2147483647.
	/*for i := -1; i <= 100; i++ {
		fmt.Print(string(rune(i)))
	}*/

	/*var p *int
	i := 42
	p = &i
	fmt.Print(*p)
	r := &i
	fmt.Print(*r)*/

}
