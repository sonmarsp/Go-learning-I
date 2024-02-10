package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"

	fmt.Println(str)
	ageCheck(1967)
}

func ageCheck(a int) {

	switch {
	case a > 1946 && a <= 1964:
		fmt.Println("Привет, бумер!")
	case a >= 1965 && a <= 1980:
		fmt.Println("Привет, представитель X!")
	case a >= 1981 && a <= 1996:
		fmt.Println("Привет, представитель X!")
	case a >= 1997 && a <= 2012:
		fmt.Println("Привет, представитель X!")
	case a >= 2013:
		fmt.Println("Привет, альфа!")

	}
}
