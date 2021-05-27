package unpack

import (
	"log"
	"strconv"
)

/*
Распаковка строки
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)
*/

func UnpackStr(s string) string {
	var storage []byte
	var prevChar byte = '\n'

	for i := 0; i < len(s); i++ {

		// is digit
		if prevChar != '\n' && '0' <= s[i] && s[i] <= '9' {
			k := 1
			addingCount, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal(err)
			}
			for k <= addingCount {
				storage = append(storage, prevChar)
				k++
			}
			prevChar = '\n'
			// is char
		} else {
			if prevChar != '\n' {
				storage = append(storage, prevChar)
			}
			prevChar = s[i]
		}
		// last char in string if it not digit
		if i == (len(s)-1) && !('0' <= s[i] && s[i] <= '9') {
			storage = append(storage, prevChar)
		}
		// Check wrong string like "5c or 55"

		/*if prevChar == '\n' {
			log.Fatal("Строка не должна начинаться с цифры")
		}*/
	}

	return string(storage)
}
