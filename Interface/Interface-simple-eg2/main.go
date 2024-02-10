package main

import "fmt"

type strucHere struct {
	N1, N2 int
}

type InterfaceHere interface {
	/*В интерфейсе не может быть данных, только методы.
	методы отличаются от функций*/
	Sum() int
}

// создадим метод
func (s *strucHere) Sum() int {
	return s.N1 + s.N2
}

func main() {

	var a InterfaceHere
	sh := strucHere{N1: 1, N2: 2}
	//os := otherStruct{A: 2, B: 3}

	a = &sh
	fmt.Println(a.Sum())

}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
