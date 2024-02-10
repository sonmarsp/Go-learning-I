package greetings

import (
	"errors"
	"fmt"
)

//Hello возвращает приветствие указанному человеку.
/**В Го функция, имя которой начинается с заглавной буквы, может быть вызвана из
другого пакета. Это известно в Го как экспортируемое имя.
*/

/*
func Hello(name string) string {

	//Вернуть приветствие, включающее имя в сообщение.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
*/

func Hello(name string) (string, error) {
	// Если имя не было указано, возвращаем ошибку с сообщением.

	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
