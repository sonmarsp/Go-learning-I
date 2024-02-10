package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {

	go hello()
	time.Sleep(1 * time.Second)
	//слеп нужен так как сразу переходит к следующей строке кода маин функции
	fmt.Println("main function")
}
