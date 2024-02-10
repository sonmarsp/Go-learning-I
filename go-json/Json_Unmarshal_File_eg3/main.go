package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Чтобы задать имя нужно прописать тэг `json:"name"`
type HolydaysList struct {
	Holydays []string `json:"holydays"`
}

func main() {

	//прочитать открытый файл как массив байтов
	byteValue, err := os.ReadFile("holydays.json")
	if err != nil {
		fmt.Println(err)
	}

	var holydays HolydaysList
	//разбираем наши биты которые содержат jsonfile в структуру
	json.Unmarshal(byteValue, &holydays)

	//отложить закрытие нашего jsonFile, чтобы мы могли проанализировать его позже
	//defer jsonFile.Close()

	fmt.Println(holydays.Holydays)

}
