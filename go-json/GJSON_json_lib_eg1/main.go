package main

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	// gjson

	//json := `{"name": {"first": "artur", "last": "karapetov"}, "age": 47}`
	//value := gjson.Get(json, "age")

	//fmt.Println(value.String())

	json := `{"name":{"first": "Tom", "last": "Andersen"},
	"age":37,
	"children" : ["Sara", "Alex", "Jack"],
	"fav.movie": "Deer Hunter",
	"friends" : [
		{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb", "tw"]},
        {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	]
	}`

	if !gjson.Valid(json) {
		panic("JSON IS NOT VALID")
	}

	//чтобы получить имя с точкой надо ее экранировать
	//value := gjson.Get(json, "fav\\.movie")

	//value := gjson.Get(json, "children.1")
	//можно использовать регулярные выражения
	//value := gjson.Get(json, "child*.2")

	//value := gjson.Get(json, "friends.1.last")

	//ищет первое совпадение
	//value := gjson.Get(json, `friends.#(last=="Murphy").first`)
	//чтобы искал все совпадения
	//value := gjson.Get(json, `friends.#(age==47)#.first`)

	//У бибилиотеки есть модификаторы
	gjson.AddModifier("case", func(json, arg string) string {

		if arg == "upper" {
			return strings.ToUpper(json)
		} else if arg == "lower" {
			return strings.ToLower(json)
		}

		return json
	})

	//можно передать модификатор для того чтобы изменить вывод json
	value := gjson.Get(json, `friends.#(age>=44)#.last|@case:upper`)

	fmt.Println(value.String())

	//gjson умеет парсить
	fmt.Println(gjson.Parse(json).Get("name"))

	//вывести в цикле
	for _, last := range value.Array() {

		fmt.Println(last.String())

	}

	//gjson можно unmarshal
	result, ok := gjson.Parse(json).Value().(map[string]interface{})
	if !ok {
		panic("NOT OKAY PARSING TO MAP")
	}
	fmt.Println(result)

}
