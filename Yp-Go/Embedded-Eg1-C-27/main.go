package main

import (
	"fmt"
)

/**
Композиция это отношение, которое что-либо включает в себя.
Эмбеддинг, то есть встраивание - это реализация композиции в го
*/

// Person - структура описывающая человека
type Person struct {
	Name string
	Year int
}

// Самописный конструктор NewPerson возвращает новую структуру Person
func NewPerson(name string, year int) Person {

	return Person{
		Name: name,
		Year: year,
	}
}

// String возвращает информацию о человеке
func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print выводит информацию о человеке
func (p Person) Print() {
	fmt.Println(p.String())
}

/*
	Student описывает студента с использованием вложенной структуры Person.

То есть структура Student описывает.
*/
type Student struct {
	Person // вложенный объект Person
	Group  string
}

/**Конструктуры по соглашению называют с New*/
func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), //Явно создаем структуру Person
		Group:  group,
	}
}

// String возвращает информацию о студенте
func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func main() {

	var str string = `Например утка состоит из клюва, туловища, лап`
	fmt.Println(str)

	s := NewStudent("John Doe", 1980, "701")
	s.Print()
	fmt.Println(s)
	fmt.Println(s.Name, s.Year, s.Group)

}
