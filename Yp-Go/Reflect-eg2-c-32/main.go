package main

import (
	"fmt"
	"reflect"
)

/*
	nil может быть как значением самого интерфейса, так и значением величины на которую указывает
*/

type MyType struct{}

func NaiveIsNil(obj interface{}) bool {
	return obj == nil
}

func IsNil(obj interface{}) bool {

	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)
	//проверяем, что тип значения ссылочный, в принципе может быть равен nil
	if objValue.Kind() != reflect.Ptr {
		return false
	}

	//Проверяем что значение равно nil
	//важно, что IsNil() вызывает панику если value не является ссылочным типом. Поэтому всегда
	//проверяйте на Kind()
	if objValue.IsNil() {
		return true
	}

	return false

}

func main() {

	var str string = `Попробуем реализовать наивную реализацию сравнения.`
	fmt.Println(str)

	var t *MyType
	fmt.Printf("Проверка типа (%v) на nil: %v\n", reflect.TypeOf(t), NaiveIsNil(t))

}
