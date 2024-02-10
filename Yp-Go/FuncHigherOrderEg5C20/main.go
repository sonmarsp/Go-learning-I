package main

import (
	"fmt"
)

type figures int

const (
	square figures = iota
	circle
	triangle //равностороний тругольник
	nof
)

func area(f figures) (func(float64) float64, bool) {

	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
		//возращает функции первый параметр func(x float64) float64 { return 0.433 * x * x }, второй параметр true
	default:
		return nil, false
	}
}

func main() {

	var str string = `Функция высшего порядка, принимающая в качестве аргументов другие функции или
	возращающая другую функцию в качестве результата`
	fmt.Println(str)

	ar, ok := area(nof)
	if ok {
		fmt.Println(ar(2))
	} else {
		fmt.Println("Нет такой фигуры")
	}

}
