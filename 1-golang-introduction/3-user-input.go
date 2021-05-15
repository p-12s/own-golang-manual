package main

import (
	"bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    fmt.Println("How old are you:")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    check(err)

    input = strings.TrimSpace(input)

    years, err := strconv.ParseFloat(input, 64)
    check(err)

    if years >= 18 {
        fmt.Println("Willkommen!")
    } else {
        fmt.Println("Halt, yunge!")
    }
    fmt.Println(fmt.Sprintf("Du bist %f yahre alt", years))
}
