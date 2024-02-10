package main

import (
	"fmt"
)

/* Defer отложенных вызовов может быть несколько
тогда они выполняются в обратном порядке, тоесть с того, который был
отложен последним, так как вызовы складывались в стек

*/

func main() {

	fmt.Println("Hello")

	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("World")
}
