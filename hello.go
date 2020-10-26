package main

import "fmt"

func main() {
	// default type: rune (int32)
	var myChar = 'c'
	fmt.Printf("Type of myChar: %T\n", myChar)

	//
	var myByteChar byte = 'a'
	var myRuneChar rune = 'a'
	fmt.Printf("%c = %d, %c = %U\n", myByteChar, myByteChar, myRuneChar, myRuneChar)

}
