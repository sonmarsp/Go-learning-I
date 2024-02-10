package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// Первоначальный пример
/* func main() {

	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
*/

func main() {
	// Устанавливаем свойства предопределенного Регистратора, включая
	// префикс записи журнала и флаг отключения печати
	// время, исходный файл и номер строки.
	// Get a greeting message and print it.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("ff")
	// Если вернулась ошибка, выводим ее в консоль и
	// выходим из программы.
	if err != nil {
		log.Fatal(err)
	}
	// Если ошибок не было, распечатываем возвращенное сообщение
	// в консоль.

	fmt.Println(message)
}
