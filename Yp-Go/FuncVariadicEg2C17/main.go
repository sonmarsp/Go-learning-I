package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	sum := Sum(2, 3, 5, 1, 2, 57)
	fmt.Println(sum)
}

func Sum(x ...int) (res int) {

	for _, v := range x {

		res += v
	}

	return
}
