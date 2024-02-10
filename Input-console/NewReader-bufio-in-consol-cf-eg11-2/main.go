package main

import (
	"bufio"
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
*/
/* Такой вариант вроде берет их ввод наборы данных*/
func inputConsole(reader *bufio.Reader) {

	fmt.Println("input text:")
	//reader := bufio.NewReader(os.Stdin)

	var lines []string

	line, _ := reader.ReadString('\n')

	//line = strings.Replace(line, "\n", " ", -1)

	lines = append(lines, line)

	//for k, l := range lines {
	//fmt.Println("key")
	//fmt.Println(k)
	//fmt.Println(l)
	//}

	fmt.Println(lines)

}

func main() {

	//fmt.Println(lines)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i <= 6; i++ {
		inputConsole(reader)
	}

}
