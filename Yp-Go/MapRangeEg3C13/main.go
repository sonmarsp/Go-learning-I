package main

import (
	"fmt"
)

func main() {

	var str string = `Ассоциативные массивы map использование в циклах range`
	fmt.Println(str)

	m := make(map[string]string)
	m["foo"] = "bar"
	m["bazz"] = "yup"

	for k, v := range m {
		// k будет перебирать ключи,
		// v — соответствующие этим ключам значения
		fmt.Printf("Ключ %v, имеет значение %v \n", k, v)
	}

}
