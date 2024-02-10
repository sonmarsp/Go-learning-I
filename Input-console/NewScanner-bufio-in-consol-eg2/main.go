package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sl := []int{}
	fmt.Println(sl)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	var t string
	if scanner.Scan() {
		t = scanner.Text()
		fmt.Printf("You wrote \"%s\"\n", scanner.Text())
	}

	fmt.Println(t)

	for i, s := range t {

		fmt.Println("index:", i)
		fmt.Println("text:", s)

		if s != 32 {
			xstr := string(s) //x is rune converted to string
			fmt.Printf("%c\n", s)
			fmt.Println(xstr)
			xint, _ := strconv.Atoi(xstr)
			fmt.Println(xint)

			sl = append(sl, xint)
		}
	}

	fmt.Println(sl)

}
