package main

import (
	"fmt"
)

func main() {

	//Для ключей используем строковый тип, а для значений
	// пустой интерфейсный тип interface{}.

	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	fmt.Printf("%+v", person)

}
