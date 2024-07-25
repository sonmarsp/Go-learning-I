package main

import (
	"fmt"
)

func main() {

	//var age int
	//fmt.Println("Сколько тебе лет?")
	//fmt.Scanf("%d\n", &age)
	//fmt.Scan("%d\n", &age)

	//fmt.Printf("Привет, %s, твой возраст - %d\n", name, age)

	//var name string
	var age int
	fmt.Print("Введите имя: ")
	//fmt.Fscan(os.Stdin, &name)
	fmt.Scanf("%d\n", &age)

	//fmt.Print("Введите возраст: ")
	//fmt.Fscan(os.Stdin, &age)

	//fmt.Println(name, age)
	fmt.Println(age)

}
