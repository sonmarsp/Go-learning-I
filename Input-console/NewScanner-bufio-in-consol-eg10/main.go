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

func inputConsole() []int {

	// мув получение строки
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	//for {
	scanner.Scan()
	line := scanner.Text()

	fmt.Println("line")
	fmt.Println(line)
	//if len(line) == 0 {
	//	break
	//}
	lines = append(lines, line)
	//}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

	//move 2 превращение в int
	var sint []int
	/*
		ai := strings.Split(line, " ")

		var sint []int
		for _, v := range lines {
			//fmt.Println(i)
			fmt.Println(v)
			t, _ := strconv.Atoi(v)

			fmt.Println("какой тип")
			fmt.Println(t)
			sint = append(sint, t)
		}

	*/
	fmt.Println("lines")
	fmt.Println(lines[0])

	return sint
}

func main() {

	t := inputConsole()
	b := inputConsole()
	fmt.Println(t)
	fmt.Println(b)
	

}
