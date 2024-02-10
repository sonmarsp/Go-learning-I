package main

import (
	"fmt"
	"os"
)

func ReadTextFileByName(filename string) (string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		//вернем ошибку на русском
		return ``, fmt.Errorf(`не удалось получить файл (%s): %v`, filename, err)
	}

	return string(data), nil
}

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	e, err := ReadTextFileByName(`nothing.txt`)

	fmt.Println(e)
	fmt.Println(err)
}
