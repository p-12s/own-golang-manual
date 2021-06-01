package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"./models"
)

func getComDomains(filename string) map[string]uint32 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	comDomains := make(map[string]uint32)

	var line []byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Bytes()
		user := &models.User{}
		if err = user.UnmarshalJSON(line); err != nil {
			panic(err)
		}

		if strings.HasSuffix(user.Email, ".com") {
			domain := strings.SplitN(user.Email, "@", 2)[1]
			num := comDomains[domain]
			num++
			comDomains[domain] = num
		}
	}

	return comDomains
}

func main() {
	comDomains := getComDomains("data.dat")
	log.Printf("%v", comDomains)
}
