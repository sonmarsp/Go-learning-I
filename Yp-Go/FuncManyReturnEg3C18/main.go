package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	var a rune = 'a'
	fmt.Println(Index("поиск", a))

}

func Index(st string, a rune) (index int, ok bool) {

	for i, c := range st {

		if c == a {
			return i, true
		}

	}

	return

}
