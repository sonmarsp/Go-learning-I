package main

import (
	"encoding/json"
	"fmt"
)

// Чтобы задать имя нужно прописать тэг `json:"name"`
type User struct {
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
	byt := []byte(`{"name":"Paul","age":80,"is_blocked":true,"books":[{"name":"BN","year":1990},{"name":"BN2","year":2090}]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)
	//получаем через интерфейс первый интерфейс это неизвестно что за тип а потом массив уже интерфейсов.
	fmt.Println(dat["books"].([]interface{})[0].(map[string]interface{})["name"])

	fmt.Println("использование структуры")

	var datEg2 User

	if err := json.Unmarshal(byt, &datEg2); err != nil {
		panic(err)
	}

	fmt.Println(datEg2)
	fmt.Println(datEg2.Books[0].Name)
	fmt.Println(datEg2.Books[1].Name)

}
