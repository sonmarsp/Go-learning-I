package main

import (
	"fmt"
	"log"
	"strconv"
)

//Объявляем тип Book, который удовлетворяет интерфейсу fmt.Stringer
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {

	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

//Объявляем тип Count, который удовлетворяет интерфейсу fmt.Stringer
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

//Объявляем функцию WriteLog(), которая берет любой объект,
// удовлетворяющий интерфейсу fmt.Stringer в виде параметра.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	//инициализируем объект Book и передаем в WriteLog().
	book := Book{"Alise in Wonderland", "Lewis Carrol"}

	WriteLog(book)

	count := Count(3)
	WriteLog(count)

}
