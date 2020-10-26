package main

import (
	"fmt"
)

func sumDefer(x, y int) int {
	defer fmt.Printf("Added %d and %d\n", x, y)

	fmt.Printf("some stuff here ...\n")

	return x + y
}

func main() {
	var sum = sumDefer(2, 4)
	fmt.Printf("Sum = %d\n", sum)
}
