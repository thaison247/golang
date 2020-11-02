package main

import "fmt"

func square(c chan int) {
	fmt.Println("squares reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("cube reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("Main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3
	fmt.Println("Send testNum to square()")

	squareChan <- testNum

	fmt.Println("Main() resuming")
	fmt.Println("Send testNum to cube()")

	cubeChan <- testNum

	fmt.Println("Main() resuming")
	fmt.Println("Main() reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Printf("[main] sum of square and cube of %d is %d\n", testNum, sum)
	fmt.Println("Main() stopped")
}
