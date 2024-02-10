package main

import (
	"bufio"
	"fmt"
	"os"
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
*/
/* Такой вариант вроде берет их ввод наборы данных*/

func main() {

	reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Simple Shell")
	//fmt.Println("---------------------")

	var t []string
	for i := 0; i < 3; i++ {

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", " ", -1)
		fmt.Print(text)
		t = append(t, text)
	}

	fmt.Print(t)

}
