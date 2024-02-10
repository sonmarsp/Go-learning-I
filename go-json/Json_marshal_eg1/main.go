package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

/*
Дескриптор для color имеет дополнительный параметр, omitempty, который указывает
что если поле имеет нулевое значение своего типа здесь false или иным образом
оказывается пустым, никакого вывода JSON не должно быть.
*/
var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	{
		//!+Marshal
		/*Такие структуры данных отлично подходят для json и легко конвертируются в обоих
		направлениях. Преобразование структуры данных в Go наподобии movies в JSON
		называется маршалингом (marshaling). Маршалинг осуществляется с помощью
		json.Marshal
		*/

		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}

		fmt.Printf("%s\n", data)
	}

	{
		//!+MarshalIndent
		data, err := json.MarshalIndent(movies, "", "   ")

		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}

		fmt.Printf("%s\n", data)

		//!+Unmarshal
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
	}

}
