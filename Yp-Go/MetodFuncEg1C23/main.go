package main

import (
	"fmt"
)

/* Метод представляет собой функцию, привязанную к конкретному типу. Методы позволяют
связывать поведение и данные типа в самом типе, обеспечивая инкапсуляцию.

*/
// объявление типа
type MyType int

/*
	объявление метода

Синатксис метода типа похож на синтаксис обычной функции, но добавляется получатель (receiver)

	после ключевого слова func. Можно сказать что получатель - это еще один аргумент функции
*/
func (m MyType) String() string {

	return fmt.Sprintf("MyType: %d", m)
}

func main() {

	fmt.Println("Hello")
	//вызов метода
	var m MyType = 5
	var s string = m.String()

	fmt.Println(s)
}