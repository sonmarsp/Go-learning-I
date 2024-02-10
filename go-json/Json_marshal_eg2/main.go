package main

import (
	"encoding/json"
	"fmt"
)

//Начинаться с большой буквы чтобы экспортировалось
//Если с маленькой то не передастся
type User struct {
	Name      string
	Age       int
	IsBlocked bool
}

//Чтобы задать имя нужно прописать тэг `json:"name"`
type UserTag struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsBlocked bool   `json:"is_blocked"`
	Books     []Book `json:"books"`
}

type Book struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func main() {
	sv := []string{"A", "B", "C"}

	boolVar, _ := json.Marshal(sv)
	fmt.Println(string(boolVar))

	eg2 := map[string]string{"field": "value", "field1": "123", "field2": "true"}
	resEg2, _ := json.Marshal(eg2)
	fmt.Println(string(resEg2))

	//Чтобы разные типы то интерфейс
	eg3 := map[string]interface{}{"field": "value", "field1": 123, "field2": true, "arr": []string{"a", "b", "c"}}
	resEg3, _ := json.Marshal(eg3)
	fmt.Println(string(resEg3))

	//вариант со структурой
	eg4 := User{
		Name:      "Paul",
		Age:       80,
		IsBlocked: true,
	}

	resEg4, _ := json.Marshal(eg4)
	fmt.Println(string(resEg4))

	//вариант с тегами для отображения имени
	var books []Book
	book1 := Book{
		Name: "BN",
		Year: 1990,
	}
	book2 := Book{
		Name: "BN",
		Year: 2090,
	}

	books = append(books, book1, book2)
	eg5 := UserTag{
		Name:      "Paul",
		Age:       80,
		IsBlocked: true,
		Books:     books,
	}

	resEg5, _ := json.Marshal(eg5)
	fmt.Println(string(resEg5))
}
