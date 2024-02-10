package main

import (
	"fmt"
	"time"

	"practicum.local/randbyte"
)

func main() {

	var str string = `Создаем генератор случайных чисел`
	fmt.Println(str)

	//В качестве затравки передаем ему текущее врямя при каждом запуске оно будет разным
	generator := randbyte.New(time.Now().UnixNano())

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) //единственный доступный метов, но он нам и нужен
		fmt.Printf("Generate bytes: %v size(%d) \n", buf, n)
	}

}
