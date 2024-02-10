package main

import (
	"fmt"
)

func incr(p *int) int {
	*p++ //Увеличивает значение на которое указывает p

	return *p
}

func main() {

	x := 1
	p := &x

	fmt.Println(*p)

	*p = 2
	fmt.Println(x)

	v := 1
	incr(&v)
	fmt.Println(incr(&v))

}
