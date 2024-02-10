package main

import (
	"fmt"
)

func main() {

	var str string = `Создайте слайс и заполните его числами от 1 до 100. Оставьте в 
	слайсе первые и последние 10 элементов и разверните слайс в обратном порядке.`
	fmt.Println(str)

	dim := 100
	s := make([]int, 0, dim)
	//заполняем слайс числами
	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}

	fmt.Println(s)

	//оставляем первые и последние элементы
	s = append(s[:10], s[dim-10:]...)
	dim = len(s)
	fmt.Println(dim)
	fmt.Println(s)

	//разворачиваем слайс
	for i := range s[:dim/2] {
		//меняем местами
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println(s)
}
