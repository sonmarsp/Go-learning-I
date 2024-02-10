package main

type Name string
type Fruit string

func main() {

	//var fruit Fruit
	//var name Name

	//fruit = "Apple"
	//name = Name(fruit)
	//ошибка компиляции

	var str string
	str = "Hello world"

	println(str[0])         //72
	println(string(str[0])) //H

}
