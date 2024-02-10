package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3}

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	change(s1)

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	changePoint(&s1)

	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

}

//Слайс передается по ссылке по умолчанию а cap и len нет
func change(s2 []int) {

	s2[2] = 5
	//s2 = append(s2, 4) //ничего не добавит в исходный слайс при этом изменить можно

}

// Нужно разименовывание указателя
func changePoint(s2 *[]int) {

	*s2 = append(*s2, 4)
}
