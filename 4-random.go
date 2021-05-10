package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

// Ввод числа в цикле
func main() {
    // Установка рандомного числа
    seconds := time.Now().Unix()
    rand.Seed(seconds)
    number := rand.Intn(10) + 1
    fmt.Println(number)

    isNumberGuessed := false
    for count := 0; count < 3; count++ {
        fmt.Println("Осталось ", 3 - count, " попыток")

        // Польз. ввод
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Ваше число: ")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        input = strings.TrimSpace(input)

        guess, err := strconv.Atoi(input)
        if err != nil {
            log.Fatal(err)
        }

        // Действие
        if guess > number {
            fmt.Println("Ваше число больше рандома")
        } else if guess < number {
            fmt.Println("Ваше число меньше рандома")
        } else {
            isNumberGuessed = true
            fmt.Println("Угадал!")
            break
        }
    }

    if !isNumberGuessed {
        fmt.Println("Попытки кончились... ", number)
    }

}

