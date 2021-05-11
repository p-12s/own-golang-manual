package main

import (
	"./datafile"
	"fmt"
	"log"
	"sort"
)

func main() {
	// Карты (под капотом - сегменты)
	myMap := map[string]float64 {
		"a": 1.3,
		"b": 3.23,
	}
	delete(myMap, "a")

	// Считываем из файла
	lines, err := datafile.GetStrings("data-names.txt")
	if err != nil {
		log.Fatal(err)
	}

	names := make(map[string]int)
	for index, line := range lines {
		names[line] = index
	}
	fmt.Println(names)

	// НЕупорядоченная карта
	langs := map[string]float64 {
		"go": 1000,
		"php": 256,
		"c#": 342,
	}
	for key, value := range langs {
		fmt.Println(key, " ", value)
	}

	// Создадим сегмент для сортировки
	var sortedNames []string
	for name := range langs {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)
	for _, value := range sortedNames {
		fmt.Println(value, langs[value])
	}

}