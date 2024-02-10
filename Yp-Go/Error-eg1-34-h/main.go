package main

import (
	"fmt"
	"os"
)

// Обычно в Го используется паттерн ранний выход
//Если возникла ошибка, то, скорее всего продолжать нет смысла, и лучше
//завершить работу как можно раньше
func ReadTextFile() (string, error) {

	if data, err := os.ReadFile(`nothing.txt`); err != nil {
		//будет вызван метод Error() string , который преобразует ошибку в строку
		fmt.Println(err)
		return err
	}

	fmt.Println(string(data))
	return string(data), nil
}

func main() {

	var str string = "Проверка"
	fmt.Println(str)

	/**<Юлагодаря тому что функции могут возращать несколько значений, ошибки
	легко попадают в их ряды */

	if data, err := os.ReadFile(`nothing.txt`); err != nil {
		//будет вызван метод Error() string , который преобразует ошибку в строку
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	ReadTextFile()

}
