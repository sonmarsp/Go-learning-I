package main

import (
	"bufio"
	"fmt"
	"log"
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

func main() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("output:")
	for _, l := range lines {
		fmt.Println(l)
	}

	fmt.Println(lines)
}
