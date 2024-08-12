package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// Возращаем слайс чисел из входного потока
func inputConsole(scanner *bufio.Scanner) []int {

	scanner.Scan()
	input := scanner.Text()
	ai := strings.Split(input, " ")

	var sint []int
	for _, v := range ai {

		t, _ := strconv.Atoi(v)
		sint = append(sint, t)
	}

	return sint
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := inputConsole(scanner)

	fmt.Println(t)

	i := 1

	for i <= t[0] {
		i++

		np := inputConsole(scanner)
		n := np[0]

		fmt.Println(n)

		x := inputConsole(scanner)
		fmt.Println(x)

	}

}
