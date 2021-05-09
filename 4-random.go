package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    seconds := time.Now().Unix()
    rand.Seed(seconds)
    number := rand.Intn(100)

    fmt.Println(number)
}

