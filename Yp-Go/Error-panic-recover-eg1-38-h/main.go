package main

import (
	"fmt"
)

/*
Чтобы создать аварийную ситуацию, нужно вызвать встроенную функцию panic
В аргументе при вызове panic можно передать значение любого типа
*/

/*
Где вызывать функцию  recover, если программа остановила свое выполнение. Так как
при возникновении аварийной ситуации вызывается функция defer вставим вызов recover туда
*/

func mypanic() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(`Возникла паника: `, p)
		}
	}()
	panic(`Аварийная ситуация`)
}

func main() {

	fmt.Println("Start")
	mypanic()
	fmt.Println("Finich")

}
