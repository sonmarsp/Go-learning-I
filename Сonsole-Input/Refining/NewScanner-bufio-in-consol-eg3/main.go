package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//fmt.Println("text")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	//как преобразовать в цифры
	//input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //только одну строку
	ai := strings.Split(input, "/n")

	fmt.Println("Вывод")
	fmt.Println(input)
	fmt.Println(ai)

	var sint []int
	for _, v := range ai {
		//fmt.Println(i)
		fmt.Println(v)
		//s = append(s, 0)
		t, _ := strconv.Atoi(v)
		sint = append(sint, t)
	}

	fmt.Println(sint)
}
