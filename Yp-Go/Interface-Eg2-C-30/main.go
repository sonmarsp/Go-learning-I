package main

import (
	"fmt"
	"strconv"
)

func PrintFC(v interface{}) {

	switch v2 := v.(type) {
	case int:
		fmt.Print("Это число " + strconv.FormatInt(int64(v2), 10))
	case string:
		fmt.Print("Это строка " + v2)
	case Stringer:
		fmt.Print("Это тип, реализующий Stringer, " + v2.String())
	default:
		fmt.Print("Неизвестный тип")
	}
}

func main() {

	var str string = `Создаем генератор случайных чисел`
	fmt.Println(str)

}
