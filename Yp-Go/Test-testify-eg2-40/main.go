package main

import (
	"fmt"
)

func EstimateValue(value int) string {

	switch {
	case value < 10:
		return "small"
	case value < 100:
		return "medium"
	default:
		return "big"

	}

}

func main() {

	fmt.Println("Start")

}
