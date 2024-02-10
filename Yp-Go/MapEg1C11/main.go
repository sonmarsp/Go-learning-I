package main

import (
	"fmt"
)

func main() {

	var str string = `Ассоциативные массивы`
	fmt.Println(str)

	m := make(map[string]string) //создаем мап с применением функции map для создания

	fmt.Println(m)
	m["foo"] = "bar"
	m["ping"] = "pong"

	fmt.Println(m)

	//В языке го мап ссылочный тип reference type, поэтому одного объявления
	// типа мап не достаточно
	//к примеру такой код скомпилируется но выдаст ощибку assignment to nil map (SA5000)
	//var me map[string]string
	//me["foo"] = "bar"
	//Чтобы объект ссылочного типа можно было использовать, его нужно сначала создать (инициализировать)
	// встроенной функцией make()

	type MyMap map[string]string

	var m1 MyMap = make(MyMap, 5)
	m1["foo"] = "bar"

	/* Встроенная функция make() универсальный конструктор объектов ссылочного типа, она
	применяется для создания объектов всех ссылочных типов
	*/

	//Сложный литерал (composite literal)

}
