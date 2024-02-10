package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Нужно обработать один ввод с переносом строк
3
4
1 2 1 2
3
2 4 2
5
1 2 3 4 5

в vscode powershel перенос строк "строка1 `n  строка `n"
*/

func inputConsole() []int {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println(input)
	ai := strings.Split(input, "\n")

	fmt.Println("inputConsole")
	//fmt.Println("Вывод")

	fmt.Println(ai)

	var sint []int
	for _, v := range ai {
		//fmt.Println(i)
		//fmt.Println(v)
		t, _ := strconv.Atoi(v)
		sint = append(sint, t)
	}

	return sint
}

func main() {

	t := inputConsole()
	fmt.Println(t)

}
