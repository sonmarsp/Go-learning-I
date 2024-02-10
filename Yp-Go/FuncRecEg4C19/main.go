package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	fmt.Println(fact(1))

}

func fact(n int) int {

	if n == 0 { //терминальная ветка - то есть условие выхода из рекурсии
		return 1
	} else { //рекурсивная ветка
		return n * fact(n-1)
	}

}
