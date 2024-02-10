package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	// sjson

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

	//value, _ := sjson.Set(json, "name.first", "Artur")
	//fmt.Println(value)

	//чтобы добавить последним элементом
	//value, _ := sjson.Set(json, "children.-1", "Ivanov")
	//fmt.Println(value)

	//добавить объект в json
	value, _ := sjson.Set(json, "new_obj", map[string]interface{}{"hello": "world"})
	fmt.Println(value)

	//есть функционал удаления
	//val, _ := sjson.Delete(json, "friends")
	val, _ := sjson.Delete(json, "friends.1")
	fmt.Println(val)

}
