package main

import (
	"fmt"
	"os"
)

/*
3
4
1 2 1 2
3
2 4 2
5
1 2 3 4 5

1
5
1 2 3 4 5

*/
/* Взять набор данных для code force */

func main() {
	var w int

	//fmt.Print("Введите вес: ")
	//просто scan не срабатывает закрывается консоль
	fmt.Fscan(os.Stdin, &w)

	fmt.Println(w)

	sl := make([]int, w)
	for i := 0; i < w; i++ {
		var s int
		fmt.Fscan(os.Stdin, &s)
		sl[i] = s
	}
	fmt.Println(sl)
}
