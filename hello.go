package main

import (
	"fmt"
)

func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}

func greeting(name string) string {
	return "hello" + name
}

func double(x int, y int) (a, b int) {
	a = x * 2
	b = y * 2
	return
}

func main() {
	i := 25

	switch {
	case i < 18:
		fmt.Printf("Children")
	case i < 50:
		fmt.Printf("Adult")
	default:
		fmt.Printf("Old person")
	}
}
