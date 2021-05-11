package main

import (
	"bufio"
    "fmt"
    "log"
    "os"
)

func main() {

    file, err := os.Open("data.txt")
    if err != nil {
       log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if scanner.Err() != nil {
        log.Fatal(err)
    }

    err = file.Close()
    if err != nil {
        log.Fatal(err)
    }
}

