package main

import (
	"fmt"
	"reflect"
)

/*
	Каждое значение, вне зависимости от типа можно привести к универсальному типу reflect.Value

Делается это через вызов reflect.ValueOf(v interface()) Value
У самого value множество методов
Type и Kind
Type() возращает тип объекта
Kind возращает базовый тип объекта, то есть не пользовательский тип а один из встроенныех в язык Go:
структуру, канал, слайс, функцию, массив и другие.
*/
func main() {

	var str string = ``
	fmt.Println(str)

	var varBool *bool
	fmt.Println(reflect.ValueOf(varBool).Kind()) //ptr указатель
	fmt.Println(reflect.ValueOf(varBool).Type()) //*bool

	var varFloat float32
	fmt.Println(reflect.ValueOf(varFloat).Kind()) //float32
	fmt.Println(reflect.ValueOf(varFloat).Type()) //float32

	var varMap map[string]int
	fmt.Println(reflect.ValueOf(varMap).Kind()) //map
	fmt.Println(reflect.ValueOf(varMap).Type()) //map[string]int

	varStruct := struct{ Value int }{}
	fmt.Println(reflect.ValueOf(varStruct).Kind()) //struct
	fmt.Println(reflect.ValueOf(varStruct).Type()) // struct {Value int}
}
