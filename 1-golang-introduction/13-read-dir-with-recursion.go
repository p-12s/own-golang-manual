package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println("Dir: ", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err) // все прекратить и раскрутить стек ошибки
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println("File: ", file.Name())
		}
	}
}

func main() {

	// AHTUNG !!!
	// разработчики языка предпочитают работать с return error + if err != nil
	// и используют panic - recover только в крайнем случае!

	defer reportPanic() // defer (отложить) - вызовет метод после выполнения scanDirectory()
	scanDirectory("./my-dir")
}
