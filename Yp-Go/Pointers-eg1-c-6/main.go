package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	i := 42
	p := &i

	//чтобы получить или изменить значение, хранящееся по указателю, применяют
	// оператор разыменовывания  (dereference) *
	fmt.Println(*p)

}
