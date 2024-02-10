package main

import (
	"fmt"
)

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	/* В языке Go есть метки (labels), которые позволяют перемещаться к разным частям кода.
	*  метку можно указать для операторов
	  break;
	  continue;
	  goto - безусловный оператор перехода, позволяет перейти в любое место кода
	*/

outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabel
		}
	}
	//Здесь break outerLoopLabel прерывает выполнение внешнего цикла.

	fmt.Println("End")

}
