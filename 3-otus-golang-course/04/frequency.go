package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
Домашнее задание
Частотный анализ
Написать функцию, которая получает на вход текст и возвращает
10 самых часто встречающихся слов без учета словоформ
*/

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(fileName string) ([]string, error) {
	var words []string

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if len(line) >= MIN_BYTE_IN_WORD {
			words = append(words, line)
		}
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return words, nil
}

func Frequency(words []string) map[string]int {
	var frequencyDict = map[string]int{}

	for _, word := range words {
		_, ok := frequencyDict[word] // check if word (the key) is already in the map
		if ok == true { // if true add 1 to frequency (value of map)
			frequencyDict[word] += 1
		} else { 		// else start frequency at 1
			frequencyDict[word] = 1
		}
	}

	return frequencyDict
}

type Pair struct {
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }


func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

const MIN_BYTE_IN_WORD = 3
const TOP_WORD_COUNT = 10

func main() {
	words, err := ReadFile("path-to-file.txt")
	check(err)
	frequencyDict := Frequency(words)

	res := rankByWordCount(frequencyDict)
	fmt.Println(res[0:TOP_WORD_COUNT])
}