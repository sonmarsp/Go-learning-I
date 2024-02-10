package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {

	arr := [2]int{30}      //30 0
	arr2 := [3]int{40, 50} //40 50 0

	//each int = 8 bytes unsafe.Sizeof(arr) показывает сколько занимает байт
	fmt.Println(arr, unsafe.Sizeof(arr))
	fmt.Println(arr2, unsafe.Sizeof(arr2))

	//copy array
	arrayToMerge := [len(arr) + len(arr2)]int{}
	fmt.Println(arrayToMerge, unsafe.Sizeof(arrayToMerge))

	counter := 0
	for i := 0; i < len(arr); i++ {
		arrayToMerge[i] = arr[i]
		counter++
	}

	for i := 0; i < len(arr2); i++ {
		arrayToMerge[counter+i] = arr2[i]
	}

	fmt.Println(arrayToMerge, unsafe.Sizeof(arrayToMerge))

	//посмотреть статистику по памяти
	stats := new(runtime.MemStats)
	runtime.ReadMemStats(stats)
	fmt.Printf("stats: %d\n", stats.HeapAlloc)
}
