package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("How old are you:")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')

    input = strings.TrimSpace(input)

    years, err := strconv.ParseFloat(input, 64)

    if err != nil {
        log.Fatal(err)
    }

    if years >= 18 {
        fmt.Println("Willkommen!")
    } else {
        fmt.Println("Halt, yunge!")
    }
    fmt.Println(fmt.Sprintf("Du bist %f yahre alt", years))
}
