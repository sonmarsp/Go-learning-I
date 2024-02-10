package main

import (
	"fmt"
)

func main() {

	var str string = `Разернуть строку таким образом не получится, так как строка — неизменяемый тип данных. 
	// Строку можно преобразовать к слайсу байт ([]byte), но это опасно, если строка содержит Unicode-символы
	// Лучше создать слайс рун из строки и развернуть его`
	fmt.Println(str)

	input := "The quick brown 狐 jumped over the lazy 犬"

	// Get Unicode code points.
	n := 0
	//создает слайс рун
	runes := make([]rune, len(input))
	//добавляем руны в слайс
	for _, r := range input {
		runes[n] = r
		n++
	}

	//убираем лишние нулевые руны
	runes = runes[0:n]
	//разворачиваем
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	//Преобразуем в строку UTF-8
	output := string(runes)
	fmt.Println(output)

}
