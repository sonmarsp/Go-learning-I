package main

import (
	"fmt"
	"os"
)

/*
Чтобы создать аварийную ситуацию, нужно вызвать встроенную функцию panic
В аргументе при вызове panic можно передать значение любого типа
*/

func myFunc() {

	if _, err := os.ReadFile(`test.txt`); err != nil {
		panic(err)
	}
}
func main() {

	fmt.Println("Start")
	myFunc()
	fmt.Println("Finich")

}
