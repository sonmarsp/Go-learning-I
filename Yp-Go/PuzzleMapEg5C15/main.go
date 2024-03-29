package main

import (
	"fmt"
)

func main() {

	var str string = `Дан массив целых чисел A и целое число k. Нужно найти и вывести индексы пары чисел,
	сумма которых равна k. Если таких чисел нет, то вернуть пустой слайс.
	Индексы можно вернуть в любом порядке.`
	fmt.Println(str)

}

func find(arr []int, k int) []int {

	//Создадим пустую мап
	m := make(map[int]int)
	//Будем складывать в нее индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {

		if j, ok := m[k-a]; ok {
			return []int{i, j}
		}
		//если искомого значения нет, то добавляем текущий индекс и значение в мап
		m[a] = i
	}
	//не нашли пары подходящих чисел
	return nil
}
