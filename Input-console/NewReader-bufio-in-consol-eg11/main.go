package main

import (
	"bufio"
	"fmt"

	//"log"
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
/* Такой вариант вроде берет их ввод наборы данных работает без лог*/

func main() {
	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		line, _ := reader.ReadString('\n')

		//if err != nil {
		//	log.Fatal(err)
		//}

		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println("output:")

	for k, l := range lines {
		fmt.Println("key")
		fmt.Println(k)
		fmt.Println(l)
	}

	fmt.Println(lines)
}

/*
[3
 4
 1 2 1 2
 3
 2 4 2
 5
 1 2 3 4 5
]
*/
